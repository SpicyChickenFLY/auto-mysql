package linux

import (
	"strings"
)

// Warning: Do NOT create passwd for user "mysql"!!!

// CreateUserWithGroup is a func for installer to create group&user
func CreateUserWithGroup(s *ServerInfo, userName, groupName string) error {
	if output, err := ExecuteCommand(
		s, Groupadd(groupName)); err != nil {
		if !strings.Contains(output, "already exists") {
			return err
		}
	}
	if output, err := ExecuteCommand(
		s, UseraddWithGroup(groupName, userName)); err != nil {
		if !strings.Contains(output, "already exists") {
			return err
		}
	}
	return nil
}
