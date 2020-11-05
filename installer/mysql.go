package installer

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

const (
	MYSQL_DRIVER_NAME = "mysql"
)

func modifyMysqlPwd(userName, userPwd, sockPath, dbName string) error {
	// connect the Mysql instance and select specified db
	db, err := sql.Open(
		MYSQL_DRIVER_NAME,
		fmt.Sprintf(
			"%s:%s@unix(%s)/%s",
			userName, userPwd, sockPath, dbName))
	if err != nil {
		fmt.Printf("[Error] Connection:%s\n", err)
		return err
	}
	fmt.Printf("[Info] connect to db, alive: %v\n", nil == db.Ping())
	defer db.Close()

	// change password for user root in mysql
	if _, err = db.Exec(
		`update mysql.user set
		 authentication_string=password('123')
		 where user='root';`); err != nil {
		fmt.Printf("[Error] Execution:%s\n", err)
		return err
	} else {
		fmt.Printf("[Info] execute sql successfully\n")
	}

	return nil
}
