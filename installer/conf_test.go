package installer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckCnfDir(t *testing.T) {
	type testCaseTemplate struct {
		cnfDir    string
		userName  string
		groupName string
		fileMode  uint32
	}
	testCases := []testCaseTemplate{
		{
			cnfDir:    "testdst/my.cnf",
			userName:  USER_NAME,
			groupName: GROUP_NAME,
			fileMode:  FILE_MODE,
		},
	}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		err := checkCnfDir(
			testCase.cnfDir, testCase.userName, testCase.groupName, testCase.fileMode)
		asst.Nil(err, fmt.Sprintf(
			"checkCnfDir failed - testCase%2d:%v", testCaseIndex, err))
	}
}
