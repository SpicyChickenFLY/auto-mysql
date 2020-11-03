package installer

import (
	"fmt"
	"os/exec"

	"github.com/romberli/log"
)

func initMysql(dstPath, userName string) error {
	// first insecure, change pwd after initialized
	cmdMysqlInit := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf(
			"sudo %s/bin/mysqld --initialize-insecure --user=%s",
			dstPath, userName))
	cmdMysqlInit.Stdout = &out
	cmdMysqlInit.Stderr = &stderr
	if err := cmdMysqlInit.Run(); err != nil {
		log.Warnf("cmdMysqlInit:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdMysqlInit: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}

func startMysql(dstPath string) error {
	// move mysql.server to /etc/init.d/mysqld for auto start
	if err := moveFile(
		dstPath+"/"+SERVER_FILE_REL, DST_SERVER_FILE); err != nil {
		return err
	}
	if err := modifyMode(
		DST_SERVER_FILE, FILE_MODE); err != nil {
		return err
	}
	// start mysql.server
	cmdMysqlStart := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf(
			"sudo %s start",
			DST_SERVER_FILE))
	cmdMysqlStart.Stdout = &out
	cmdMysqlStart.Stderr = &stderr
	if err := cmdMysqlStart.Run(); err != nil {
		log.Warnf("cmdMysqlStart:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdMysqlStart: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}
