package db

import (
	"database/sql"
	"fmt"

	"github.com/romberli/log"
)

const ( // MYSQL CONFIG
	MySQLAllHost   = `%`
	MySQLAllScope  = "*.*"
	MySQLLocalHost = mysqlLocalHost
)

// ModifyPwdForFirstTime is a func to change passwd for user
func ModifyPwdForFirstTime(db *sql.DB, passwd string) error {
	// change password for user root in mysql
	sqlStmt := fmt.Sprintf(
		`ALTER user user() IDENTIFIED BY '%s';`, passwd)
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
		`GRANT %s ON %s TO  %s IDENTIFIED BY '%s';`,
		permission, scope, user, passwd)
	if _, err := db.Exec(sqlStmt); err != nil {
		return err
	}
	log.Info("execute sql successfully")
	return nil
}
