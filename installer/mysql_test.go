package installer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitMysql(t *testing.T) {
	testCases := []string{}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		// Feasibility test
		err := initMysql(testCase)
		asst.Nil(err, fmt.Sprintf(
			"initMysql failed - testCase%2d:%v", testCaseIndex, err))
	}
}

func TestStartMysql(t *testing.T) {
	testCases := []string{}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		// Feasibility test
		err := startMysql(testCase)
		asst.Nil(err, fmt.Sprintf(
			"startMysql failed - testCase%2d:%v", testCaseIndex, err))
	}
}
