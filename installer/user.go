package installer

import (
	"fmt"
	"io"
	"os/exec"
	"os/user"

	"github.com/romberli/log"
)

// Warning: Do NOT create passwd for user "mysql"!!!

func createUser(userName string) error {
	if _, ok := findUser(userName); ok {
		return nil
	}
	cmdAddUser := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo useradd -M %s", userName))
	cmdAddUser.Stdout = &out
	cmdAddUser.Stderr = &stderr
	if err := cmdAddUser.Run(); err != nil {
		log.Warnf("cmdAddUser:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdAddUser: %s:%s\n",
			err, stderr.String())
		return err
	}

	return nil
}

func createUserWithGroup(userName string, groupName string) error {
	if _, ok := findGroup(groupName); ok {
		log.Warn("Group already exists!")
		return nil
	}
	cmdAddGroup := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo groupadd %s", groupName))
	cmdAddGroup.Stdout = &out
	cmdAddGroup.Stderr = &stderr
	if err := cmdAddGroup.Run(); err != nil {
		log.Warnf("cmdAddGroup:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdAddGroup: %s:%s\n",
			err, stderr.String())
		return err
	}

	if _, ok := findUser(userName); ok {
		log.Warn("User already exists!")
		return nil
	}
	cmdAddUser := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf(
			"sudo useradd -M -s /sbin/nologin -g %s %s",
			groupName, userName))
	cmdAddUser.Stdout = &out
	cmdAddUser.Stderr = &stderr
	if err := cmdAddUser.Run(); err != nil {
		log.Warnf("cmdAddUser:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdAddUser: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}

// This function is not used
func modifyPwdForUser(userName string, usePwd string) error {
	if _, ok := findUser(userName); ok {
		return nil
	}
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

func findGroup(groupName string) (*user.Group, bool) {
	g, err := user.LookupGroup(groupName)
	if err != nil {
		return nil, false
	}
	return g, true
}

func findUser(userName string) (*user.User, bool) {
	u, err := user.Lookup(userName)
	if err != nil {
		return nil, false
	}
	return u, true
}
