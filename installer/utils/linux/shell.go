package linux

import (
	"fmt"

	"github.com/lingdor/stackerror"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/log"
)

const tempPath = "/tmp/temp_data"

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

// Touch is a shell command for create file
func Touch(dstFile string) string {
	return fmt.Sprintf("touch %s", dstFile)
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

// Mkdir is a shell command for making dir
func Mkdir(dirPath string) string {
	return fmt.Sprintf("mkdir -p %s", dirPath)
}

// Tar is a shell command for tar/untar file
func Tar(srcFile string, dstPath string) string {
	return fmt.Sprintf("tar -zxf %s -C %s ", srcFile, dstPath)
}

// Cat is a shell command for display file content
func Cat(dstFile string) string {
	return fmt.Sprintf("cat %s", dstFile)
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

// Useradd is a shell command for adding user in linux
func Useradd(userName string) string {
	return fmt.Sprintf("useradd -M -s /bin/false %s", userName)
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
