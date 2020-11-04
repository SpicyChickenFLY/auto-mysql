package installer

import (
	"fmt"
)

func initMysql(dstPath, userName string) error {
	// first insecure, change pwd after initialized
	return execCommand(
		fmt.Sprintf(
			"sudo %s/bin/mysqld --initialize-insecure --user=%s",
			dstPath, userName))

}

func startMysql(dstPath string) error {
	// move mysql.server to /etc/init.d/mysqld for auto start
	if err := mv(
		dstPath+"/"+SERVER_FILE_REL, DST_SERVER_FILE); err != nil {
		return err
	}
	if err := chmod(
		DST_SERVER_FILE, FILE_MODE); err != nil {
		return err
	}
	// start mysql.server
	return execCommand(
		fmt.Sprintf("sudo %s start", DST_SERVER_FILE))
}
