package mysql

import (
	"os"
	"path"

	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/linux"
	"github.com/lingdor/stackerror"
)

const (
	allFile         = "*"
	dstCnfPathDef   = "/etc"
	dstCnfFileDef   = "/etc/my.cnf"
	autoCnfFileName = "auto.cnf"
	tmpSQLPath      = "/tmp/mysql"
	tmpSQLFile      = "mysql.tar.gz"
	usrBinPath      = "/usr/bin"
)

// modify is a func for changing mode and owner for file/dir
func modifyDataDir(s *linux.ServerInfo, dirPath string, mode uint32) error {
	if _, err := linux.ExecuteCommand(
		s, linux.Chown(dirPath, userName, groupName)); err != nil {
		return err
	}
	if _, err := linux.ExecuteCommand(
		s, linux.Chmod(dirPath, mode)); err != nil {
		return err
	}
	return nil
}

// createFile is a func for creating a specified file
func createFile(
	s *linux.ServerInfo,
	filePath string) error {
	if _, err := linux.ExecuteCommand(
		s, linux.Touch(filePath)); err != nil {
		return err
	}
	return modifyDataDir(s, filePath, sqlFileMode)
}

// createDir is a func for creating a specified dir
func createDir(
	s *linux.ServerInfo,
	dirPath string) error {
	if _, err := linux.ExecuteCommand(
		s, linux.Rm(dirPath)); err != nil {
		return err
	}
	if _, err := linux.ExecuteCommand(
		s, linux.Mkdir(dirPath)); err != nil {
		return err
	}
	return modifyDataDir(s, dirPath, sqlFileMode)
}

// unTarWithGzip extract a .tar.gz file by name write in SHELL
func unTarWithGzip(
	s *linux.ServerInfo,
	srcFile string, dstPath string) error {
	if _, err := linux.ExecuteCommand(
		s, linux.Mkdir(dstPath)); err != nil {
		return err
	}
	if _, err := linux.ExecuteCommand(
		s, linux.Tar(srcFile, dstPath)); err != nil {
		return err
	}
	return nil
}

// ExtractSoftware is a func to extract mysql archive to specfied dir
// and modify its mode and owner
func ExtractSoftware(
	s *linux.ServerInfo,
	srcSQLFile, dstSQLPath string) error {

	// Judge if software exist
	if _, err := os.Stat(srcSQLFile); err != nil {
		return stackerror.New("Software not exists")
	}
	// Move mysql file to dstSQLPath
	if err := linux.CopyDirOrFileBetweenServers(
		linux.LocalHost, s,
		srcSQLFile, path.Join(dstSQLPath, tmpSQLFile)); err != nil {
		return err
	}
	if err := unTarWithGzip(
		s,
		path.Join(dstSQLPath, tmpSQLFile), dstSQLPath); err != nil {
		return err
	}
	if err := linux.MoveDirOrFileOnServer(
		s, path.Join(dstSQLPath, "mysql-*", allFile), dstSQLPath); err != nil {
		return err
	}
	if err := linux.CopyDirOrFileBetweenServers(
		s, s, path.Join(dstSQLPath, daemonPathRel, allFile), usrBinPath); err != nil {
		return err
	}

	// Modify Directory/File
	return modifyDataDir(s, dstSQLPath, sqlFileMode)
}

// MoveCnfFile is a func to move custom my.cnf configure to specified dir
func MoveCnfFile(
	s *linux.ServerInfo,
	srcCnfFile string) error {
	if err := linux.CopyDirOrFileBetweenServers(
		linux.LocalHost, s, srcCnfFile, dstCnfPathDef); err != nil {
		return err
	}
	return modifyDataDir(s, dstCnfFileDef, cnfFileMode)
}

// MoveDaemonFile is a func to move mysql.server/mysqld_multi.server
func MoveDaemonFile(
	servInstInfo *ServerInstanceInfo,
	srcDaemonFileRel, dstDaemonFile string) error {
	// move mysql.Daemon to /etc/init.d/mysqld for auto start
	srcDaemonFile := path.Join(servInstInfo.BaseDir, srcDaemonFileRel)
	if err := linux.CopyDirOrFileBetweenServers(
		servInstInfo.ServerInfo, servInstInfo.ServerInfo,
		srcDaemonFile, dstDaemonFile); err != nil {
		return err
	}
	return modifyDataDir(servInstInfo.ServerInfo, dstDaemonFile, sqlFileMode)
}

// CopyDataDir is a func to copy files from datadir for master to slaves
func CopyDataDir(
	masterServInstInfo *ServerInstanceInfo,
	allServInstInfos []*ServerInstanceInfo) error {
	// for _, servInstInfo := range allServInstInfos {

	// }
	// if _, err := linux.ExecuteCommand(
	// 	dstServer, linux.Rm(dstDirPath)); err != nil {
	// 	return err
	// }
	// if err := linux.CopyDirOrFileBetweenServers(
	// 	srcServer, dstServer, srcDirPath, dstDirPath); err != nil {
	// 	return err
	// }
	// if err := modifyDataDir(dstServer, dstDirPath, sqlFileMode); err != nil {
	// 	return err
	// }
	// autoCnfFile := path.Join(dstDirPath, autoCnfFileName)
	// if _, err := linux.ExecuteCommand(
	// 	dstServer, linux.Rm(autoCnfFile)); err != nil {
	// 	return err
	// }
	return nil
}
