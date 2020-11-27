package utils

import (
	"fmt"
	"io"
	"os/exec"
	"os/user"

	"github.com/romberli/log"
)

// Warning: Do NOT create passwd for user "mysql"!!!

// CreateUserWithGroup is a func for installer to create group&user
func CreateUserWithGroup(userName, groupName string) error {
	log.Info(fmt.Sprintf("CreateUserWithGroup(%s,%s)",
		userName, groupName))
	if _, ok := findGroup(groupName); ok {
		log.Warn("Group already exists!")
	} else {
		if err := Groupadd(groupName); err != nil {
			return err
		}
	}

	if _, ok := findUser(userName); ok {
		log.Warn("User already exists!")
	} else {
		if err := UseraddWithGroup(groupName, userName); err != nil {
			return err
		}
	}

	return nil
}

// FIXME: This function is not used for now
func ModifyPwdForUser(userName string, usePwd string) error {
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

// findGroup is a func to find if group is existed
func findGroup(groupName string) (*user.Group, bool) {
	g, err := user.LookupGroup(groupName)
	if err != nil {
		return nil, false
	}
	return g, true
}

// findUser is a func to find if user is existed
func findUser(userName string) (*user.User, bool) {
	u, err := user.Lookup(userName)
	if err != nil {
		return nil, false
	}
	return u, true
}
