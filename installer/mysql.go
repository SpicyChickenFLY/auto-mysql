package installer

import (
	"database/sql"
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/installer/shell"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_DRIVER_NAME = "mysql"
)

// prepareMysqlUbuntu is a func to install all dependencies for
// MySQL in Ubuntu(Linux OS)
// dependencies list: libaio, libncurese
func prepareMysqlUbuntu() error {
	// install libaio with apt
	err := shell.ExecCommand("sudo apt-cache search libaio")
	if err != nil {
		return err
	}
	err = shell.ExecCommand("sudo apt install libaio1")
	if err != nil {
		return err
	}
	// install libncurses with apt
	err = shell.ExecCommand("sudo apt install libncurses5")
	return err
}

// initMysql is a func to initialize the mysql instance without password
func initMysql(dstPath, userName string) error {
	// first insecure, change pwd after initialized
	return shell.ExecCommand(
		fmt.Sprintf(
			"sudo %s/bin/mysqld --initialize-insecure --user=%s",
			dstPath, userName))

}

// startMysql is a func to start mysql instance automaticaly
func startMysql(dstPath string) error {
	// move mysql.server to /etc/init.d/mysqld for auto start
	if err := shell.Cp(
		dstPath+"/"+SERVER_FILE_REL, DST_SERVER_FILE); err != nil {
		return err
	}
	if err := shell.Chmod(
		DST_SERVER_FILE, FILE_MODE); err != nil {
		return err
	}
	// start mysql.server
	return shell.ExecCommand(
		fmt.Sprintf("sudo %s start", DST_SERVER_FILE))
}

// modifyMysqlPwd is a func to connect insecure db and change passwd for user
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
