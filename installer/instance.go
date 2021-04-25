package installer

import (
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/SpicyChickenFLY/auto-mysql/utils/db"
	"github.com/SpicyChickenFLY/auto-mysql/utils/linux"
)

const (
	daemonFileName      = "mysqld"
	daemonFileRel       = "bin/mysqld"
	singleServerFileRel = "support-files/mysql.server"
	multiServerFileRel  = "bin/mysqld_multi"
)

const (
	instWaitTimeout      = 100 // ms
	instWaitTimeoutRetry = 5
)

// InitInstance is a func to initialize the mysql instance without password
func InitInstance(servInstInfo *ServerInstanceInfo) error {
	for _, instInfo := range servInstInfo.InstInfos {
		if _, err := linux.ExecuteCommand(
			servInstInfo.ServerInfo,
			fmt.Sprintf(
				`%s --initialize-insecure --user=%s --basedir=%s --datadir=%s`,
				path.Join(servInstInfo.BaseDir, daemonFileRel),
				linuxUserMysql,
				servInstInfo.BaseDir,
				instInfo.DataDir)); err != nil {
			return err
		}

	}
	return nil
}

// StartSingleInst is a func to start mysql Instance automaticaly
func StartSingleInst(servInstInfo *ServerInstanceInfo) error {
	if _, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo,
		fmt.Sprintf(`%s start`,
			path.Join(servInstInfo.BaseDir, singleServerFileRel))); err != nil {
		return err
	}
	// wait for instance start
	for i := 0; i < instWaitTimeoutRetry; i++ {
		output, err := linux.FindProcess(
			servInstInfo.ServerInfo, daemonFileName)
		if err != nil {
			return err
		}
		if len(output) > 0 {
			return nil
		}
		time.Sleep(time.Duration(instWaitTimeout) * time.Millisecond)
	}
	return errors.New("wait too long to start single instance ")
}

// StopSingleInst is a func to stop mysql Instance automaticaly
func StopSingleInst(servInstInfo *ServerInstanceInfo) error {
	if _, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo,
		fmt.Sprintf(`%s stop`,
			path.Join(servInstInfo.BaseDir, singleServerFileRel))); err != nil {
		return err
	}
	// wait for instance stop
	for i := 0; i < instWaitTimeoutRetry; i++ {
		output, err := linux.FindProcess(
			servInstInfo.ServerInfo, daemonFileName)
		if err != nil {
			return err
		}
		if len(output) == 0 {
			return nil
		}
		time.Sleep(time.Duration(instWaitTimeout) * time.Millisecond)
	}
	return errors.New("wait too long to stop single instance ")
}

// StartMultiInst is a func to start mysql Instance automaticaly
func StartMultiInst(servInstInfo *ServerInstanceInfo) error {
	for _, instInfo := range servInstInfo.InstInfos {
		// fmt.Println(servInstInfo.ServerInfo)
		if _, err := linux.ExecuteCommand(
			servInstInfo.ServerInfo,
			fmt.Sprintf(`%s start %d`,
				path.Join(servInstInfo.BaseDir, multiServerFileRel),
				instInfo.Port)); err != nil {
			return err
		}
		// wait for multi instance start
		i := 0
		for i < instWaitTimeoutRetry {
			output, err := linux.FindProcess(
				servInstInfo.ServerInfo, daemonFileName, fmt.Sprint(instInfo.Port))
			if err != nil {
				return err
			}
			if len(output) == 0 {
				break
			}
			time.Sleep(time.Duration(instWaitTimeout) * time.Millisecond)
			i++
		}
		if i == instWaitTimeoutRetry {
			return errors.New("wait too long to start multi instance ")
		}
	}
	return nil
}

// StopMultiInst is a func to start mysql Instance automaticaly
func StopMultiInst(servInstInfo *ServerInstanceInfo) error {
	for _, instInfo := range servInstInfo.InstInfos {
		if _, err := linux.ExecuteCommand(
			servInstInfo.ServerInfo,
			fmt.Sprintf(`%s stop %d`,
				path.Join(servInstInfo.BaseDir, multiServerFileRel),
				instInfo.Port)); err != nil {
			return err
		}
		// wait for multi instance stop
		i := 0
		for i < instWaitTimeoutRetry {
			output, err := linux.FindProcess(
				servInstInfo.ServerInfo, daemonFileName, fmt.Sprint(instInfo.Port))
			if err != nil {
				return err
			}
			if len(output) == 0 {
				break
			}
			time.Sleep(time.Duration(instWaitTimeout) * time.Millisecond)
			i++
		}
		if i == instWaitTimeoutRetry {
			return errors.New("wait too long to stop multi instance ")
		}
	}
	return nil
}

// ModifyPwdForAllInstForFirstTime modify pwd for all instances in one server
func ModifyPwdForAllInstForFirstTime(
	servInstInfo *ServerInstanceInfo, port []int, prevPwd, newPwd string) error {
	for _, instInfo := range servInstInfo.InstInfos {
		conn, err := db.CreateConn(servInstInfo.ServerInfo.Host, instInfo.Port, prevPwd)
		if err != nil {
			return err
		}
		defer conn.Close()
		if err := db.ModifyPwdForFirstTime(conn, newPwd); err != nil {
			return err
		}
	}
	return nil
}
