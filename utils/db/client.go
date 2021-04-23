package db

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
