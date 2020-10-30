package installer

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

// Warning: Do NOT create passwd for user "mysql"!!!

func createUser(userName string) error {
	cmdAddUser := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo useradd -M %s", userName))
	// if err := cmdAddUser.Run(); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdAddUser.Stdout = &out
	cmdAddUser.Stderr = &stderr
	err := cmdAddUser.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	return nil
}

func createUserWithGroup(userName string, groupName string) error {
	cmdAddGroup := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo groupadd %s", groupName))
	if err := cmdAddGroup.Run(); err != nil {
		return err
	}
	cmdAddUser := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf(
			"sudo useradd -M -s /sbin/nologin -g %s %s",
			groupName, userName))
	if err := cmdAddUser.Run(); err != nil {
		return err
	}
	return nil
}

// This function is not used
func modifyPwdForUser(userName string, usePwd string) error {
	cmdEcho := exec.Command("echo", usePwd)
	cmdPasswd := exec.Command("passwd", "--stdin", userName)

	r, w := io.Pipe() // create a Pipe
	defer r.Close()
	defer w.Close()
	cmdEcho.Stdout = w
	cmdPasswd.Stdin = r

	err := cmdEcho.Start()
	if err != nil {
		return nil
	}
	err = cmdPasswd.Start()
	if err != nil {
		return nil
	}
	cmdEcho.Wait()
	w.Close()
	return nil
}
