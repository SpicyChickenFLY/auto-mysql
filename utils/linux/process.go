package linux

import "fmt"

// KillProcess kill all process by param
func KillProcess(
	s *ServerInfo, processParam string) error {
	_, err := ExecuteCommand(s,
		fmt.Sprintf(
			"ps -ef | grep %s | grep -v grep | cut -c 9-15 | xargs kill -s 9",
			processParam))
	return err
}

// FindProcess find process by param one by one
func FindProcess(
	s *ServerInfo, processParams ...string) (string, error) {
	cmdStr := "ps -ef"
	for _, param := range processParams {
		cmdStr += fmt.Sprintf("| grep %s | grep -v grep", param)
	}
	output, err := ExecuteCommand(s, cmdStr)
	return output, err
}
