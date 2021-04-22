package mysql

import (
	"database/sql"
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
