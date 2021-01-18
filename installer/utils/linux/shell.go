package linux

import (
	"fmt"

	"github.com/lingdor/stackerror"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/log"
)

const tempPath = "/tmp/temp_data"

// MoveDirOrFileOnServer move file on one server
func MoveDirOrFileOnServer(
	s *ServerInfo,
	srcPath, dstPath string) error {
	_, err := ExecuteCommand(
		s, Mv(srcPath, dstPath))
	return err
}

// CopyDirOrFileBetweenServers copy file between any servers
func CopyDirOrFileBetweenServers(
	srcServ, dstServ *ServerInfo,
	srcPath, dstPath string) error {
	if srcServ.Host == dstServ.Host {
		_, err := ExecuteCommand(
			srcServ, Cp(srcPath, dstPath))
		return err
	}
	if srcServ.Host == "localhost" {
		sshConn, err := linux.NewMySSHConn(
			dstServ.Host, dstServ.Port, dstServ.UserName, dstServ.UserPwd)
		if err != nil {
			return err
		}
		return sshConn.CopyToRemote(srcPath, dstPath)
	} else if dstServ.Host == "localhost" {
		sshConn, err := linux.NewMySSHConn(
			srcServ.Host, srcServ.Port, srcServ.UserName, srcServ.UserPwd)
		if err != nil {
			return err
		}
		return sshConn.CopyFromRemote(srcPath, dstPath)
	}
	localServer := ServerInfo{Host: "localhost"}
	// 清空临时目录中的文件
	if _, err := ExecuteCommand(&localServer, Rm(tempPath)); err != nil {
		return err
	}
	if _, err := ExecuteCommand(&localServer, Mkdir(tempPath)); err != nil {
		return err
	}
	// 并将远端源服务器中的目录文件传到该目录下
	sshConnSrcServ, err := linux.NewMySSHConn(
		srcServ.Host, srcServ.Port, srcServ.UserName, srcServ.UserPwd)
	if err != nil {
		return err
	}
	if err := sshConnSrcServ.CopyFromRemote(srcPath, tempPath); err != nil {
		return err
	}
	// 获取临时目录中所有文件，传至远端目标服务器
	sshConnDstServ, err := linux.NewMySSHConn(
		dstServ.Host, dstServ.Port, dstServ.UserName, dstServ.UserPwd)
	if err != nil {
		return err
	}
	return sshConnDstServ.CopyToRemote(tempPath, dstPath)

}

// ExecuteCommand exec cmd on any server
func ExecuteCommand(s *ServerInfo, cmdStr string) (output string, err error) {
	// fmt.Println(s, cmdStr)
	log.Info(fmt.Sprintln(s, cmdStr))
	if s == nil || s.Host == "localhost" {
		output, err = linux.ExecuteCommand(cmdStr)
	} else {
		sshConn, err := linux.NewMySSHConn(s.Host, s.Port, s.UserName, s.UserPwd)
		if err != nil {
			return output, err
		}
		_, output, err = sshConn.ExecuteCommand(cmdStr)
	}
	// fmt.Println(output)
	log.Info(fmt.Sprintln(output))
	if err != nil {
		// fmt.Println(err)
		err = stackerror.New(err.Error() + output)
	}
	return output, err
}

// Mv is a shell command for moving file
func Mv(srcFile, dstFile string) string {
	return fmt.Sprintf("mv %s %s", srcFile, dstFile)
}

// Cp is a shell command for copying file
func Cp(srcFile, dstFile string) string {
	return fmt.Sprintf("cp -rf %s %s", srcFile, dstFile)
}

// Rm is a shell command for deleteng file
func Rm(dstFile string) string {
	return fmt.Sprintf("rm -rf %s", dstFile)
}

// Chown is a shell command for changing owner of dir/file
func Chown(dirPath, userName, groupName string) string {
	return fmt.Sprintf("chown -R %s.%s %s", userName, groupName, dirPath)
}

// Chmod is a shell command for changing
// access premission for file/dir
func Chmod(dirPath string, fileMode uint32) string {
	// chmod -R 755 dirPath
	return fmt.Sprintf("chmod -R %d %s", fileMode, dirPath)
}

// Mkdir is a shell command for making dir
func Mkdir(dirPath string) string {
	return fmt.Sprintf("mkdir -p %s", dirPath)
}

// Tar is a shell command for tar/untar file
func Tar(srcFile string, dstPath string) string {
	return fmt.Sprintf("tar -zxf %s -C %s ", srcFile, dstPath)
}

// Useradd is a shell command for adding user in linux
func Useradd(userName string) string {
	return fmt.Sprintf("useradd -M -s /sbin/nologin %s", userName)
}

// Userdel is a shell command for deleting user in linux
func Userdel(userName string) string {
	return fmt.Sprintf("userdel %s", userName)
}

// Groupadd is a shell command for adding group in linux
func Groupadd(groupName string) string {
	return fmt.Sprintf("groupadd %s", groupName)
}

// Groupdel is a shell command for deleting group in linux
func Groupdel(groupName string) string {
	return fmt.Sprintf("groupdel %s", groupName)
}

// UseraddWithGroup is a function to create user in specified group
func UseraddWithGroup(groupName, userName string) string {
	return fmt.Sprintf("useradd -M -g %s %s", groupName, userName)
}

// Cat is a shell command for display file content
func Cat(dstFile string) string {
	return fmt.Sprintf("cat %s", dstFile)
}

// Touch is a shell command for create file
func Touch(dstFile string) string {
	return fmt.Sprintf("touch %s", dstFile)
}
