package installer

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open(
		"mysql",
		`root:@unix(/tmp/mysql.sock)/mysql`)
	if err != nil {
		fmt.Printf("[Error] Connection:%s\n", err)
	} else {
		fmt.Printf("[Info] connect mysql successfully\n")
		// fmt.Printf("[Info] driver:%v\n", db.Driver())
		fmt.Printf("[Info] alive: %v\n", nil == db.Ping())
	}

	_, err = db.Exec(
		`DROP TABLE IF EXISTS student`)

	if err != nil {
		fmt.Printf("[Error] Execution:%s\n", err)
	} else {
		fmt.Printf("[Info] execute sql successfully\n")
	}

	_, err = db.Exec( //
		`update mysql.user set
		 authentication_string=password('123')
		 where user='root';`)

	if err != nil {
		fmt.Printf("[Error] Execution:%s\n", err)
	} else {
		fmt.Printf("[Info] execute sql successfully\n")
	}
}
