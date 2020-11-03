package installer

import (
	"bytes"
	"fmt"

	"github.com/romberli/log"
)

var out bytes.Buffer
var stderr bytes.Buffer

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
//  6. Start server
//  7. TODO: change the password for root in mysql
func Install(
	srcSqlFile, dstSqlPath,
	srcCnfFile, dstCnfFile string) error {
	// Create group/user
	log.Info(fmt.Sprintf("createUserWithGroup(%s,%s):", USER_NAME, GROUP_NAME))
	if err := createUserWithGroup(USER_NAME, GROUP_NAME); err != nil {
		fmt.Printf("[ %s ] Create group/user\n",
			RenderStr("FAIL", "highlight", "black", "red"))
		return err
	}
	fmt.Printf("[  %s  ] Create group/user\n",
		RenderStr("OK", "highlight", "black", "green"))

	// Decompress the archive
	log.Info(fmt.Sprintf("extactFile(%s,%s):", srcSqlFile, dstSqlPath))
	if err := extractFile(srcSqlFile, dstSqlPath); err != nil {
		fmt.Printf("[ %s ] Decompress the archive\n",
			RenderStr("FAIL", "highlight", "black", "red"))
		return err
	}
	fmt.Printf("[  %s  ] Decompress the archive\n",
		RenderStr("OK", "highlight", "black", "green"))

	// Move configure file
	log.Info(fmt.Sprintf("extactFile(%s,%s):", srcSqlFile, dstSqlPath))
	if err := moveCnfFile(srcCnfFile, dstCnfFile); err != nil {
		fmt.Printf("[ %s ] Move configure file\n",
			RenderStr("FAIL", "highlight", "black", "red"))
		return err
	}
	fmt.Printf("[  %s  ] Move configure file\n",
		RenderStr("OK", "highlight", "black", "green"))

	// Create data directory
	log.Info(fmt.Sprintf("checkCnfDir(%s,%s,%s,%4d)",
		dstCnfFile, USER_NAME, GROUP_NAME, FILE_MODE))
	if err := checkCnfDir(
		dstCnfFile, USER_NAME, GROUP_NAME, FILE_MODE); err != nil {
		fmt.Printf("[ %s ] Create data directory\n",
			RenderStr("FAIL", "highlight", "black", "red"))
		return err
	}
	fmt.Printf("[  %s  ] Create data directory\n",
		RenderStr("OK", "highlight", "black", "green"))

	// Initialize MySQL(without password)
	log.Info(fmt.Sprintf("initMysql(%s,%s)", dstSqlPath, USER_NAME))
	if err := initMysql(dstSqlPath, USER_NAME); err != nil {
		fmt.Printf("[ %s ] Initialize MySQL(without password)\n",
			RenderStr("FAIL", "highlight", "black", "red"))
		return err
	}
	fmt.Printf("[  %s  ] Initialize MySQL(without password)\n",
		RenderStr("OK", "highlight", "black", "green"))

	fmt.Println((RenderStr("Install Compelete\n", "highlight", "black", "red")))
	log.Info("Install Compelete")
	return nil
}

func TestInstall(
	startPos int,
	srcSqlFile, dstSqlPath,
	srcCnfFile, dstCnfFile string) error {
	switch startPos {
	case 0:
		if err := createUserWithGroup(USER_NAME, GROUP_NAME); err != nil {
			return err
		}
		fallthrough
	case 1:
		if err := unTarWithGzipShell(srcSqlFile, dstSqlPath); err != nil {
			return err
		}
		fallthrough
	case 2:
		if err := modifyDir(dstSqlPath, USER_NAME, GROUP_NAME, FILE_MODE); err != nil {
			return err
		}
		fallthrough
	case 3:
		if err := moveFile(srcCnfFile, dstCnfFile); err != nil {
			return err
		}
		fallthrough
	case 4:
		if err := modifyDir(dstCnfFile, USER_NAME, GROUP_NAME, CNF_FILE_MODE); err != nil {
			return err
		}
		fallthrough
	case 5:
		if err := checkCnfDir(
			dstCnfFile, USER_NAME, GROUP_NAME, FILE_MODE); err != nil {
			return err
		}
		fallthrough
	case 6:
		if err := initMysql(dstSqlPath, USER_NAME); err != nil {
			return err
		}
	}
	return nil
}

// TODO: feature: auto uninstall
func Remove() {

}

func extractFile(srcSqlFile, dstSqlPath string) error {
	// Extract File from *.tar.gz
	if err := unTarWithGzipShell(srcSqlFile, dstSqlPath); err != nil {
		return err
	}
	// Move mysql file to parent dir
	if err := moveFileShell(dstSqlPath+"/mysql-*/*", dstSqlPath); err != nil {
		return err
	}
	// Modify Directory/File
	if err := modifyDir(dstSqlPath, USER_NAME, GROUP_NAME, FILE_MODE); err != nil {
		return err
	}
	return nil
}

func moveCnfFile(srcCnfFile, dstCnfFile string) error {
	if err := copyFileShell(srcCnfFile, dstCnfFile); err != nil {
		return err
	}
	if err := modifyDir(dstCnfFile, USER_NAME, GROUP_NAME, CNF_FILE_MODE); err != nil {
		return err
	}
	return nil
}
