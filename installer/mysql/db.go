package mysql

import (
	"database/sql"
	"fmt"

	"github.com/romberli/log"
)

const ( // MYSQL CONFIG
	mysqlDriverName = "mysql"
	mysqlLocalHost  = "localhost"
	mysqlUserName   = "root"
	mysqlDBName     = "mysql" // defined by user
	mysqlDBCharset  = "utf8"

	mysqlAllHost  = `%`
	mysqlAllScope = "*.*"

	mysqlAdminUserName   = `admin`
	mysqlAdminPermission = "SHUTDOWN"
	mysqlAdminScope      = mysqlAllScope

	mysqlReplicaHost       = mysqlLocalHost
	mysqlReplicaUserName   = "replication"
	mysqlReplicaPermission = "replication slave, replication client"
	mysqlReplicaScope      = mysqlAllScope
)

const ( // Test Data
	testCreateDB    = "CREATE DATABASE spdb"
	testCreateTable = `CREATE TABLE spdb.test(
		id int(10) not null auto_increment,
		name varchar(100),
		primary key(id)
	)engine=innodb default charset=utf8mb4`
	testCreateInitData = `
	INSERT INTO spdb.test(name) VALUES ('a'), ('b'), ('c')`
	testCreateNewData = `
	INSERT INTO spdb.test(name) VALUES ('d')`
	testRetriveNewData = `
	SELECT name FROM spdb.test
		WHERE name='d'`
)

// CreateConn is a func to create a connection pool
func CreateConn(servHost string, servPort int, userPwd string) (*sql.DB, error) {
	var serverURL string
	serverURL = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUserName, userPwd,
		servHost, servPort,
		mysqlDBName, mysqlDBCharset)
	// connect the Mysql instance and select specified db
	db, err := sql.Open(
		mysqlDriverName,
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
		mysqlUserName, passwd,
		sockFile,
		mysqlDBName, mysqlDBCharset)
	// connect the Mysql instance and select specified db
	db, err := sql.Open(
		mysqlDriverName,
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

// CreateAdminUser create a mysql user: admin
func CreateAdminUser(
	db *sql.DB,
	passwd string) error {
	mysqlAdminUser := fmt.Sprintf(
		"'%s'@'%s'", mysqlAdminUserName, mysqlLocalHost)
	return GrantUser(db,
		mysqlAdminPermission, mysqlAdminScope, mysqlAdminUser, passwd)
}

// CreateReplicaUser create a mysql user: replication
func CreateReplicaUser(
	db *sql.DB,
	passwd string) error {
	mysqlReplicaUser := fmt.Sprintf(
		"'%s'@'%s'", mysqlReplicaUserName, mysqlAllHost)
	return GrantUser(db,
		mysqlReplicaPermission, mysqlReplicaScope, mysqlReplicaUser, passwd)
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
	if _, err := db.Exec(testCreateDB); err != nil {
		return err
	}
	if _, err := db.Exec(testCreateTable); err != nil {
		return err
	}
	if _, err := db.Exec(testCreateInitData); err != nil {
		return err
	}
	return nil
}

// CreateTestData is a func to create test data
func CreateTestData(db *sql.DB) error {
	_, err := db.Exec(testCreateNewData)
	return err
}

// RetriveTestData is a func to retrive test data
func RetriveTestData(db *sql.DB) error {
	rows, err := db.Query(testRetriveNewData)
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
	masterHost string, masterPort int, masterPwd string) error {
	sqlStmt := fmt.Sprintf(
		`CHANGE MASTER TO 
			master_host='%s', 
			master_port=%d,
			master_user='%s',
			master_password='%s',
			master_auto_position=1;`,
		mysqlReplicaHost, masterPort, mysqlReplicaUserName, masterPwd)
	_, err := db.Exec(sqlStmt)
	return err
}

// StartSlave is a func to start replication
func StartSlave(db *sql.DB) error {
	_, err := db.Exec("start slave;")
	return err
}
