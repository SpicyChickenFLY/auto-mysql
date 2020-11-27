package utils

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/romberli/log"
)

var stdout bytes.Buffer
var stderr bytes.Buffer

// ExecCommand is a func for go to call shell command
// TODO: let user decide if he/she need sudoer premission
func ExecCommand(cmdStr string) error {
	// fmt.Println(cmdStr)
	// sudo mv srcFile dstFile
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdout = &stdout
	if stdout.String() != "" {
		fmt.Printf("cmd-[%s]%s\n", cmdStr, stdout.String())
		log.Infof("cmd-[%s]%s\n", cmdStr, stdout.String())
	}
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Warnf("cmd-[%s]%s:%s\n",
			cmdStr, err, stderr.String())
		fmt.Printf("cmd-[%s]%s:%s\n",
			cmdStr, err, stderr.String())
		return err
	}
	return nil
}

// Mv is a shell command for moving file
func Mv(srcFile, dstFile string) error {
	// sudo mv srcFile dstFile
	return ExecCommand(
		fmt.Sprintf("sudo mv %s %s", srcFile, dstFile))
}

// Cp is a shell command for copying file
func Cp(srcFile, dstFile string) error {
	// sudo cp -r srcFile dstFile
	return ExecCommand(
		fmt.Sprintf("sudo cp -rf %s %s", srcFile, dstFile))
}

// Rm is a shell command for deleteng file
func Rm(dstFile string) error {
	// sudo rm -rf dstFile
	return ExecCommand(
		fmt.Sprintf("sudo rm -rf %s", dstFile))
}

// Chown is a shell command for changing owner of dir/file
func Chown(dirPath, userName, groupName string) error {
	// sudo chown -R userName.groupName dirPath
	return ExecCommand(
		fmt.Sprintf("sudo chown -R %s.%s %s", userName, groupName, dirPath))

}

// Chmod is a shell command for changing
// access premission for file/dir
func Chmod(dirPath string, fileMode uint32) error {
	// sudo chmod -R 755 dirPath
	return ExecCommand(
		fmt.Sprintf("sudo chmod -R %d %s", fileMode, dirPath))

}

// Mkdir is a shell command for making dir
func Mkdir(dirPath string) error {
	// sudo mkdir -p dirPath
	return ExecCommand(
		fmt.Sprintf("sudo mkdir -p %s", dirPath))
}

// Tar is a shell command for tar/untar file
func Tar(srcFile string, dstPath string) error {
	// tar -zxvf srcFile -C dstPath
	return ExecCommand(
		fmt.Sprintf("sudo tar -zxf %s -C %s ", srcFile, dstPath))
}

// Useradd is a shell command for adding user in linux
func Useradd(userName string) error {
	return ExecCommand(
		fmt.Sprintf("sudo useradd -M %s", userName))
}

// Groupadd is a shell command for adding group in linux
func Groupadd(groupName string) error {
	return ExecCommand(
		fmt.Sprintf("sudo groupadd %s", groupName))

}

// UseraddWithGroup is a function to create user in specified group
func UseraddWithGroup(groupName, userName string) error {
	return ExecCommand(
		fmt.Sprintf("sudo useradd -M -g %s %s", groupName, userName))
}
