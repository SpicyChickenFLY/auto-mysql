package installer

import (
	"github.com/romberli/log"
)

const (
	DEFAULT_FILE_MODE = 755
)

// TODO: complete this function to make module_test feasible
func compareFile() bool {
	return true
}

func modify(dirPath, userName, groupName string, fileMode uint32) error {
	if err := chown(dirPath, userName, groupName); err != nil {
		log.Error("encount error.")
		return err
	}
	if err := chmod(dirPath, fileMode); err != nil {
		log.Error("encount error.")
		return err
	}
	return nil
}

func createDirWithDetail(
	dirPath, userName, groupName string, fileMode uint32) error {
	if err := mkdir(dirPath); err != nil {
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
func unTarWithGzipShell(srcFile string, dstPath string) error {
	if err := mkdir(dstPath); err != nil {
		return err
	}
	if err := tar(srcFile, dstPath); err != nil {
		return err
	}

	return nil
}

func extractSqlFile(srcSqlFile, dstSqlPath string) error {
	// Extract File from *.tar.gz
	if err := unTarWithGzipShell(srcSqlFile, dstSqlPath); err != nil {
		return err
	}
	// Move mysql file to parent dir
	if err := mv(dstSqlPath+"/mysql-*/*", dstSqlPath); err != nil {
		return err
	}
	// Modify Directory/File
	return modify(dstSqlPath, USER_NAME, GROUP_NAME, FILE_MODE)
}

func moveCnfFile(srcCnfFile, dstCnfFile string) error {
	if err := cp(srcCnfFile, dstCnfFile); err != nil {
		return err
	}
	return modify(dstCnfFile, USER_NAME, GROUP_NAME, CNF_FILE_MODE)
}
