package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/romberli/log"
)

const ( // MYSQL CONFIG
	MYSQL_DRIVER_NAME = "mysql"
	MYSQL_SERVER_HOST = "localhost"
	MYSQL_SERVER_PORT = "3306"
	MYSQL_USER_NAME   = "root"
	// MYSQL_USER_PWD = "123" // defined by user
	MYSQL_DB_NAME    = "mysql" // defined by user
	MYSQL_DB_CHARSET = "utf8"
)

const ( // Test Data
	TEST_CREATE_DB    = "CREATE DATABASE spdb"
	TEST_CREATE_TABLE = `CREATE TABLE spdb.test(
		id int(10) not null auto_increment,
		name varchar(100),
		primary key(id)
	)engine=innodb default charset=utf8mb4`
	TEST_CREATE_INIT_DATA = `
	INSERT INTO spdb.test(name) VALUES ('a'), ('b'), ('c')`
	TEST_CREATE_NEW_DATA = `
	INSERT INTO spdb.test(name) VALUES ('d')`
	TEST_RETRIVE_NEW_DATA = `
	SELECT name FROM spdb.test
		WHERE name='d'`
)

// PrepareMysqlUbuntu is a func to install all dependencies for
// MySQL in Ubuntu(Linux OS)
// dependencies list: libaio, libncurese
func PrepareMysqlUbuntu() error {
	// install libaio with apt
	err := ExecCommand("sudo apt-cache search libaio")
	if err != nil {
		return err
	}
	err = ExecCommand("sudo apt install libaio1")
	if err != nil {
		return err
	}
	// install libncurses with apt
	err = ExecCommand("sudo apt install libncurses5")
	return err
}

// InitMysql is a func to initialize the mysql instance without password
func InitMysql(
	dstSQLPath, userName, dataDir string) error {
	log.Info(fmt.Sprintf("InitMysql(%s,%s)", dstSQLPath, userName))
	// first insecure, change pwd after initialized
	return ExecCommand(
		fmt.Sprintf(
			`sudo runuser -l mysql -c '%s/bin/mysqld \
			--initialize-insecure --user=%s --basedir=%s --datadir=%s'`,
			dstSQLPath, userName, dstSQLPath, dataDir))

}

// StartSingleInst is a func to start mysql Instance automaticaly
func StartSingleInst(dstDaemonPath string) error {
	log.Info(fmt.Sprintf("StartSingleInst(%s)", dstDaemonPath))
	// start mysql.server
	return ExecCommand(
		fmt.Sprintf("sudo %s start", dstDaemonPath))
}

// StartMultiInst is a func to start mysql Instance automaticaly
func StartMultiInst(dstSQLPath string, ports []int) error {
	log.Info(fmt.Sprintf("StartMultiInst(%s)", dstSQLPath))
	// start mysql.server
	for _, port := range ports {
		if err := ExecCommand(
			fmt.Sprintf(
				`sudo runuser -l mysql -c '%s/bin/mysqld_multi start %d'`,
				dstSQLPath, port)); err != nil {
			return err
		}
	}
	// show the instances status
	// if err := ExecCommand(
	// 	fmt.Sprintf(
	// 		`sudo runuser -l mysql -c '%s/bin/mysqld_multi report'`,
	// 		dstSQLPath)); err != nil {
	// 	return err
	// }
	// sleep 5 second for starting instance
	time.Sleep(time.Duration(10) * time.Second)
	return nil
}

// StopMultiInst is a func to start mysql Instance automaticaly
func StopMultiInst(dstDaemonPath string, ports []int) error {
	log.Info(fmt.Sprintf("StopMultiInst(%s)", dstDaemonPath))
	// start mysql.server
	for _, port := range ports {
		if err := ExecCommand(
			fmt.Sprintf(
				`sudo runuser -l mysql -c '%s/bin/mysqld_multi stop %d'`,
				dstDaemonPath, port)); err != nil {
			return err
		}
	}
	return nil
}

// CreateConn is a func to create a connection pool
func CreateConn(port int, passwd string) (*sql.DB, error) {
	var serverURL string
	serverURL = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		MYSQL_USER_NAME, passwd,
		MYSQL_SERVER_HOST, port,
		MYSQL_DB_NAME, MYSQL_DB_CHARSET)
	// connect the Mysql instance and select specified db
	db, err := sql.Open(
		MYSQL_DRIVER_NAME,
		serverURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Error("failed to connect database")
		return nil, err
	}
	return db, nil
}

// CreateConnBySock is a func to create a connection pool
func CreateConnBySock(sockFile, passwd string) (*sql.DB, error) {
	var serverURL string
	serverURL = fmt.Sprintf(
		"%s:%s@unix(%s)/%s?charset=%s&timeout=10s",
		MYSQL_USER_NAME, passwd,
		sockFile,
		MYSQL_DB_NAME, MYSQL_DB_CHARSET)
	// connect the Mysql instance and select specified db
	db, err := sql.Open(
		MYSQL_DRIVER_NAME,
		serverURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Error("failed to connect database")
		return nil, err
	}
	return db, nil
}

// CloseConn is a func to close the connection pooll
func CloseConn(db *sql.DB) error {
	db.Close()
	return nil
}

// ModifyMysqlPwd is a func to change passwd for user
func ModifyMysqlPwd(db *sql.DB, passwd string) error {
	// change password for user root in mysql
	sqlStmt := fmt.Sprintf(
		`UPDATE mysql.user SET
		 authentication_string=password('%s')
		 WHERE user='root';`, passwd)
	if _, err := db.Exec(sqlStmt); err != nil {
		log.Error("execute sql successfully")
		return err
	}
	return nil
}

// GrantUser is a func to grant permission
func GrantUser(
	db *sql.DB,
	permission, scope,
	user, passwd string) error {
	// change password for user root in mysql
	sqlStmt := fmt.Sprintf(
		`GRANT %s ON %s
			TO  %s IDENTIFIED BY '%s';`,
		permission, scope, user, passwd)
	if _, err := db.Exec(sqlStmt); err != nil {
		return err
	}
	log.Info("execute sql successfully")
	return nil
}

// CreateTestEnv is a func to create a db/table/data
func CreateTestEnv(db *sql.DB) error {
	if _, err := db.Exec(TEST_CREATE_DB); err != nil {
		return err
	}
	if _, err := db.Exec(TEST_CREATE_TABLE); err != nil {
		return err
	}
	if _, err := db.Exec(TEST_CREATE_INIT_DATA); err != nil {
		return err
	}
	return nil
}

// CreateTestData is a func to create test data
func CreateTestData(db *sql.DB) error {
	_, err := db.Exec(TEST_CREATE_NEW_DATA)
	return err
}

// RetriveTestData is a func to retrive test data
func RetriveTestData(db *sql.DB) error {
	rows, err := db.Query(TEST_RETRIVE_NEW_DATA)
	if err != nil {
		return err
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
	}
	return nil
}

// ChangeMaster is a func to change master of slave instance
func ChangeMaster(
	db *sql.DB,
	masterHost string, masterPort int,
	masterUser, masterPwd string) error {
	sqlStmt := fmt.Sprintf(
		`CHANGE MASTER TO 
			master_host='%s', 
			master_port=%d,
			master_user='%s',
			master_password='%s',
			master_auto_position=1;`,
		masterHost, masterPort, masterUser, masterPwd)
	_, err := db.Exec(sqlStmt)
	return err
}

// StartSlave is a func to start replication
func StartSlave(db *sql.DB) error {
	_, err := db.Exec("start slave;")
	return err
}
