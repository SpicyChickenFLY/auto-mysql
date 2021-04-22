package mysql

import (
	"fmt"
	"path"
	"time"

	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/db"
	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/linux"
)

const (
	instWaitTimeout      = 100 // ms
	instWaitTimeoutRetry = 1
)

// InitInstance is a func to initialize the mysql instance without password
func InitInstance(servInstInfo *ServerInstanceInfo) error {
	for _, instInfo := range servInstInfo.InstInfos {
		if _, err := linux.ExecuteCommand(
			servInstInfo.ServerInfo,
			fmt.Sprintf(
				`runuser -l mysql -c '%s --initialize-insecure --user=%s --basedir=%s --datadir=%s'`,
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
		fmt.Sprintf(`runuser -l mysql -c '%s start'`,
			path.Join(servInstInfo.BaseDir, singleServerFileRel))); err != nil {
		return err
	}
	// if err := ShowErrorLog(servInstInfo); err != nil {
	// 	return err
	// }
	return waitInstanceStartStop(servInstInfo, true)
}

// StopSingleInst is a func to stop mysql Instance automaticaly
func StopSingleInst(servInstInfo *ServerInstanceInfo) error {
	if _, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo,
		fmt.Sprintf(`runuser -l mysql -c '%s stop'`,
			path.Join(servInstInfo.BaseDir, singleServerFileRel))); err != nil {
		return err
	}
	return waitInstanceStartStop(servInstInfo, true)
}

// StartMultiInst is a func to start mysql Instance automaticaly
func StartMultiInst(servInstInfo *ServerInstanceInfo) error {
	for _, instInfo := range servInstInfo.InstInfos {
		// fmt.Println(servInstInfo.ServerInfo)
		if _, err := linux.ExecuteCommand(
			servInstInfo.ServerInfo,
			fmt.Sprintf(`runuser -l mysql -c '%s start %d'`,
				path.Join(servInstInfo.BaseDir, multiServerFileRel),
				instInfo.Port)); err != nil {
			return err
		}
	}
	return waitInstanceStartStop(servInstInfo, true)
}

// StopMultiInst is a func to start mysql Instance automaticaly
func StopMultiInst(servInstInfo *ServerInstanceInfo) error {
	for _, instInfo := range servInstInfo.InstInfos {
		if _, err := linux.ExecuteCommand(
			servInstInfo.ServerInfo,
			fmt.Sprintf(`runuser -l mysql -c '%s stop %d'`,
				path.Join(servInstInfo.BaseDir, multiServerFileRel),
				instInfo.Port)); err != nil {
			return err
		}
	}
	return waitInstanceStartStop(servInstInfo, false)
}

// ModifyPwdForAllInstOfServer modify pwd for all instances in one server
func ModifyPwdForAllInstOfServer(
	servInstInfo *ServerInstanceInfo, port []int, prevPwd, newPwd string) error {
	for _, instInfo := range servInstInfo.InstInfos {
		conn, err := db.CreateConn(servInstInfo.ServerInfo.Host, instInfo.Port, prevPwd)
		if err != nil {
			return err
		}
		defer conn.Close()
		if err := db.ModifyMysqlPwd(conn, newPwd); err != nil {
			return err
		}
	}
	return nil
}

func waitInstanceStartStop(
	servInstInfo *ServerInstanceInfo, startStop bool) error {
	time.Sleep(time.Duration(5) * time.Second)
	return nil
}
