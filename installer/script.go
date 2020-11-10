package installer

import (
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/installer/colorful"
	"github.com/romberli/log"
)

const (
	USER_NAME     = "mysql"
	GROUP_NAME    = "mysql"
	FILE_MODE     = 775
	CNF_FILE_MODE = 644

	SERVER_FILE_REL = "support-files/mysql.server"
	DST_SERVER_FILE = "/etc/init.d/mysqld"
	CLIENT_FILE_REL = "bin/mysql"
)

// Install Procedure
//  1. Create group/user
//  2. Decompress the archive
//  3. Move configure file
//  4. Create data directory
//  5. Initialize MySQL(without password)
//  6. Start MySQL instance
//  7. TODO: change the password for root in mysql
func Install(
	srcSqlFile, dstSqlPath,
	srcCnfFile, dstCnfFile string) error {
	// Create group/user
	log.Info(fmt.Sprintf(
		"createUserWithGroup(%s,%s):",
		USER_NAME, GROUP_NAME))
	checkErr(
		createUserWithGroup(USER_NAME, GROUP_NAME),
		"Create group/user")

	// Decompress the archive
	log.Info(fmt.Sprintf(
		"extactFile(%s,%s):", srcSqlFile, dstSqlPath))
	checkErr(
		extractSqlFile(srcSqlFile, dstSqlPath),
		"Decompress the archive")

	// Move configure file
	log.Info(fmt.Sprintf(
		"moveCnfFile(%s,%s):", srcCnfFile, dstCnfFile))
	checkErr(
		moveCnfFile(srcCnfFile, dstCnfFile),
		"Move configure file")

	// Create data directory
	log.Info(fmt.Sprintf("checkCnfDir(%s,%s,%s,%4d)",
		dstCnfFile, USER_NAME, GROUP_NAME, FILE_MODE))
	checkErr(
		checkCnfDir(dstCnfFile, USER_NAME, GROUP_NAME, FILE_MODE),
		"Create data directory")

	// Initialize MySQL(without password)
	log.Info(fmt.Sprintf("initMysql(%s,%s)", dstSqlPath, USER_NAME))
	checkErr(
		initMysql(dstSqlPath, USER_NAME),
		"Initialize MySQL(without password)")

	// Start MySQL instance
	log.Info(fmt.Sprintf("startMysql(%s)", dstSqlPath))
	checkErr(
		startMysql(dstSqlPath),
		"Start MySQL instance")

	fmt.Println((colorful.RenderStr("Install Compelete\n", "highlight", "black", "green")))
	log.Info("Install Compelete")
	return nil
}

// Remove all file you installed and restore your configure
// TODO: feature: auto uninstall
func Remove() {

}

// checkErr is a built-in func for checking error and output its result
func checkErr(err error, info string) {
	if err != nil {
		fmt.Printf("[ %s ] %s\n",
			colorful.RenderStr("FAIL", "highlight", "black", "red"), info)
		fmt.Println(err)
		fmt.Println((colorful.RenderStr("Install Failed", "highlight", "black", "red")))
		log.Info("Install Failed")
		panic("exit")
	} else {
		fmt.Printf("[  %s  ] %s\n",
			colorful.RenderStr("OK", "highlight", "black", "green"), info)
	}
}
