package utils

import (
	"fmt"

	"github.com/romberli/log"
)

// compareFile is a func to check if two file is the same
// TODO: complete this function to make module_test feasible
func compareFile() bool {
	return true
}

// modify is a func for changing mode and owner for file/dir
func modify(dirPath, userName, groupName string, fileMode uint32) error {
	if err := Chown(dirPath, userName, groupName); err != nil {
		log.Error("encount error.")
		return err
	}
	if err := Chmod(dirPath, fileMode); err != nil {
		log.Error("encount error.")
		return err
	}
	return nil
}

// createDirWithDetail is a func for creating a specified dir
func createDirWithDetail(
	dirPath, userName, groupName string, fileMode uint32) error {
	if err := Mkdir(dirPath); err != nil {
		log.Error("encount error.")
		return err
	}
	if err := modify(
		dirPath, userName, groupName, fileMode); err != nil {
		log.Error("encount error.")
		return err
	}
	return nil
}

// unTarWithGzip extract a .tar.gz file by name write in SHELL
func unTarWithGzip(srcFile string, dstPath string) error {
	if err := Mkdir(dstPath); err != nil {
		return err
	}
	if err := Tar(srcFile, dstPath); err != nil {
		return err
	}

	return nil
}

// ExtractSQLFile is a func to extract mysql archive to specfied dir
// and modify its mode and owner
func ExtractSQLFile(srcSQLFile, dstSQLPath, userName, groupName string, fileMode uint32) error {
	log.Info(fmt.Sprintf("ExtractSQLFile(%s,%s)",
		srcSQLFile, dstSQLPath))
	// Extract File from *.tar.gz
	if err := unTarWithGzip(srcSQLFile, dstSQLPath); err != nil {
		return err
	}
	// Move mysql file to parent dir
	if err := Mv(dstSQLPath+"/mysql-*/*", dstSQLPath); err != nil {
		return err
	}
	// Modify Directory/File
	return modify(dstSQLPath, userName, groupName, fileMode)
}

// MoveCnfFile is a func to move custom my.cnf configure to specified dir
func MoveCnfFile(
	srcCnfFile, dstCnfFile, userName, groupName string,
	fileMode uint32) error {
	log.Info(fmt.Sprintf("MoveCnfFile(%s,%s)",
		srcCnfFile, dstCnfFile))
	if err := Cp(srcCnfFile, dstCnfFile); err != nil {
		return err
	}
	return modify(dstCnfFile, userName, groupName, fileMode)
}

// MoveDaemonFile is a func to move mysql.server/mysqld_multi.server
func MoveDaemonFile(
	dstSQLPath, srcDaemonFile, dstDaemonFile string,
	fileMode uint32) error {
	// move mysql.Daemon to /etc/init.d/mysqld for auto start
	if err := Cp(
		dstSQLPath+"/"+srcDaemonFile, dstDaemonFile); err != nil {
		return err
	}
	return Chmod(dstDaemonFile, fileMode)
}

// CopyDataDir is a func to copy files from datadir for master to slaves
func CopyDataDir(
	userName, groupName string,
	srcDirPath string, dstDirPaths []string,
	autoCnfName string) error {
	for _, dstDirPath := range dstDirPaths {
		if err := Rm(
			fmt.Sprintf(
				"%s", dstDirPath)); err != nil {
			return err
		}
		if err := Cp(srcDirPath, dstDirPath); err != nil {
			return err
		}
		if err := Chown(dstDirPath, userName, groupName); err != nil {
			return err
		}
		if err := Rm(
			fmt.Sprintf(
				"%s/%s", dstDirPath, autoCnfName)); err != nil {
			return err
		}
	}
	return nil
}
