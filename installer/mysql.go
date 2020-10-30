package installer

import (
	"fmt"
	"os/exec"
)

func initMysql(dstPath string) error {
	// first insecure, change pwd after initialized
	cmdMysqlInit := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf(
			"%s/mysqld --initialize-insecure --user=%s",
			dstPath, USER_NAME))
	if err := cmdMysqlInit.Run(); err != nil {
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
	if err := modifyDirOnMode(
		DST_SERVER_FILE, FILE_MODE); err != nil {
		return err
	}
	// start mysql.server
	cmdMysqlStart := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf(
			"sudo %s start",
			DST_SERVER_FILE))
	if err := cmdMysqlStart.Run(); err != nil {
		return err
	}
	return nil
}
