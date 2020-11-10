package installer

import (
	"github.com/SpicyChickenFLY/auto-mysql/installer/shell"
	"github.com/romberli/log"
)

const (
	DEFAULT_FILE_MODE = 755
)

// compareFile is a func to check if two file is the same
// TODO: complete this function to make module_test feasible
func compareFile() bool {
	return true
}

// modify is a func for changing mode and owner for file/dir
func modify(dirPath, userName, groupName string, fileMode uint32) error {
	if err := shell.Chown(dirPath, userName, groupName); err != nil {
		log.Error("encount error.")
		return err
	}
	if err := shell.Chmod(dirPath, fileMode); err != nil {
		log.Error("encount error.")
		return err
	}
	return nil
}

// createDirWithDetail is a func for creating a specified dir
func createDirWithDetail(
	dirPath, userName, groupName string, fileMode uint32) error {
	if err := shell.Mkdir(dirPath); err != nil {
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
	if err := shell.Mkdir(dstPath); err != nil {
		return err
	}
	if err := shell.Tar(srcFile, dstPath); err != nil {
		return err
	}

	return nil
}

// extractSqlFile is a func to extract mysql archive to specfied dir
// and modify its mode and owner
func extractSqlFile(srcSqlFile, dstSqlPath string) error {
	// Extract File from *.tar.gz
	if err := unTarWithGzip(srcSqlFile, dstSqlPath); err != nil {
		return err
	}
	// Move mysql file to parent dir
	if err := shell.Mv(dstSqlPath+"/mysql-*/*", dstSqlPath); err != nil {
		return err
	}
	// Modify Directory/File
	return modify(dstSqlPath, USER_NAME, GROUP_NAME, FILE_MODE)
}

// moveCnfFile is a func to move custom my.cnf configure to specified dir
func moveCnfFile(srcCnfFile, dstCnfFile string) error {
	if err := shell.Cp(srcCnfFile, dstCnfFile); err != nil {
		return err
	}
	return modify(dstCnfFile, USER_NAME, GROUP_NAME, CNF_FILE_MODE)
}
