package installer

import "flag"

const (
	USER_NAME     = "mysql"
	GROUP_NAME    = "mysql"
	FILE_MODE     = 755
	CNF_FILE_MODE = 644

	SRC_SQL_FILE = "./src/mysql/mysql.tar.gz"
	DST_SQL_PATH = "/usr/local/mysql"
	SRC_CNF_FILE = "./src/conf/my.cnf"
	DST_CNF_FILE = "/etc/my.cnf"

	SERVER_FILE_REL = "support-files/mysql.server"
	DST_SERVER_FILE = "/etc/init.d/mysqld"
	CLIENT_FILE_REL = "bin/mysql"
)

/* Install Procedure
 * 1. Create group/user
 * 2. Decompress the archive
 * 3. Move configure file
 * 4. Create data directory
 * 5. Initialize MySQL(without password)
 * 6. Start server
 * 7. TODO: change the password for root in mysql
 */
func Install() error {
	// Custom parameters
	srcSqlFile := *flag.String(
		"s", SRC_SQL_FILE, "postion of mysql-binary file")
	dstSqlPath := *flag.String(
		"d", DST_SQL_PATH, "position for installation")
	srcCnfFile := *flag.String(
		"c", SRC_CNF_FILE, "postion of you configure file")
	// Fixed parameters
	dstCnfFile := DST_CNF_FILE

	if err := createUserWithGroup(USER_NAME, GROUP_NAME); err != nil {
		return err
	}
	if err := unTarWithGzipGo(srcSqlFile, dstSqlPath); err != nil {
		return err
	}
	if err := modifyDir(dstSqlPath, USER_NAME, GROUP_NAME, FILE_MODE); err != nil {
		return err
	}
	if err := moveFile(srcCnfFile, dstCnfFile); err != nil {
		return err
	}
	if err := modifyDir(dstCnfFile, USER_NAME, GROUP_NAME, CNF_FILE_MODE); err != nil {
		return err
	}
	if err := checkCnfDir(
		dstCnfFile, USER_NAME, GROUP_NAME, FILE_MODE); err != nil {
		return err
	}
	if err := initMysql(dstSqlPath); err != nil {
		return err
	}

	return nil
}

// TODO: feature: auto uninstall
func Remove() {

}
