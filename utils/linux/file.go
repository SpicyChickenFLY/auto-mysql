package linux

import (
	"fmt"

	"github.com/romberli/go-util/linux"
)

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

func Unarchive(
	dstServ *ServerInfo,
	srcFile, dstPath string,
	stripComponents int) error {
	_, err := ExecuteCommand(dstServ,
		fmt.Sprintf("tar -zxf %s -C %s --strip-components=%d ",
			srcFile, dstPath, stripComponents))
	return err
}

// func Archive(
// 	dstServ *ServerInfo,
// 	srcPath, dstPath string) error {

// }
