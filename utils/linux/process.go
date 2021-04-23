package linux

import "fmt"

// KillMysqlProcess kill all process of mysql
func KillProcess(
	s *ServerInfo, processName string) error {
	_, err := ExecuteCommand(s,
		fmt.Sprintf(
			"ps -ef | grep %s | grep -v grep | cut -c 9-15 | xargs kill -s 9",
			processName))
	return err
}
