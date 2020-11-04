package installer

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/romberli/log"
)

var out bytes.Buffer
var stderr bytes.Buffer

func execCommand(cmdStr string) error {
	// sudo mv srcFile dstFile
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdout = &out
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

func mv(srcFile, dstFile string) error {
	// sudo mv srcFile dstFile
	return execCommand(
		fmt.Sprintf("sudo mv %s %s", srcFile, dstFile))
}

func cp(srcFile, dstFile string) error {
	// sudo mv srcFile dstFile
	return execCommand(
		fmt.Sprintf("sudo cp -r %s %s", srcFile, dstFile))
}

func chown(dirPath, userName, groupName string) error {
	// sudo chown -R userName:groupName dirPath
	return execCommand(
		fmt.Sprintf("sudo chown -R %s:%s %s", userName, groupName, dirPath))

}

func chmod(dirPath string, fileMode uint32) error {
	// sudo chmod -R 755 dirPath
	return execCommand(
		fmt.Sprintf("sudo chmod -R %d %s", fileMode, dirPath))

}

func mkdir(dirPath string) error {
	// sudo mkdir -p dirPath
	return execCommand(
		fmt.Sprintf("sudo mkdir -p %s", dirPath))
}

func tar(srcFile string, dstPath string) error {
	// tar -zxvf srcFile -C dstPath
	return execCommand(
		fmt.Sprintf("sudo tar -zxvf %s -C %s ", srcFile, dstPath))
}

func useradd(userName string) error {
	return execCommand(
		fmt.Sprintf("sudo useradd -M %s", userName))
}

func groupadd(groupName string) error {
	return execCommand(
		fmt.Sprintf("sudo groupadd %s", groupName))

}
