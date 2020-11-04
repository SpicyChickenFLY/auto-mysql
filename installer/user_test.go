package installer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// === RUN   TestCreateUser
// --- PASS: TestCreateUser (0.09s)
// === RUN   TestCreateUserWithGroup
// --- PASS: TestCreateUserWithGroup (0.09s)

func TestCreateUser(t *testing.T) {
	testCases := []string{
		"test_user_1",
	}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		err := useradd(testCase)
		asst.Nil(err, fmt.Sprintf(
			"createUser failed - testCase%2d:%v", testCaseIndex, err))
	}
}

func TestCreateUserWithGroup(t *testing.T) {
	type testCaseTemplate struct {
		userName  string
		groupName string
	}
	testCases := []testCaseTemplate{
		{
			userName:  "test_user_2",
			groupName: "test_group_2",
		},
	}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		err := createUserWithGroup(testCase.userName, testCase.groupName)
		asst.Nil(err, fmt.Sprintf(
			"createUserWithGroup failed - testCase%2d:%v", testCaseIndex, err))
	}
}
