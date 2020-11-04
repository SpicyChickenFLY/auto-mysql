package installer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// === RUN   TestUnTarWithGzipGo
// 1               2/2Byte
// 2/              0
// 2/1             2/2Byte
// 3/              0
// 3/1             2/2Byte
// 3/3             3/3Byte
// --- PASS: TestUnTarWithGzipGo (0.00s)
// === RUN   TestMoveFile
// --- PASS: TestMoveFile (0.00s)
// === RUN   TestCreateDir
// --- PASS: TestCreateDir (0.00s)
// === RUN   TestModifyDir
// --- PASS: TestModifyDir (0.09s)

func TestUnTarWithGzipShell(t *testing.T) {
	type testCaseTemplate struct {
		srcFile string
		dstPath string
	}
	testCases := []testCaseTemplate{
		{
			srcFile: "testsrc/1.tar.gz",
			dstPath: "testdst",
		},
		{
			srcFile: "testsrc/2.tar.gz",
			dstPath: "testdst",
		},
		{
			srcFile: "testsrc/3.tar.gz",
			dstPath: "testdst",
		},
	}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		// Feasibility test
		err := unTarWithGzipShell(
			testCase.srcFile,
			testCase.dstPath,
		)
		asst.Nil(err, fmt.Sprintf(
			"unTar failed - testCase%2d:%v", testCaseIndex, err))
	}
}

func TestCreateDirWithDetail(t *testing.T) {
	testCases := []string{
		"testdst/4/7",
		"testdst/4/7/8",
	}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		// Feasibility test
		err := createDirWithDetail(testCase, USER_NAME, GROUP_NAME, FILE_MODE)
		asst.Nil(err, fmt.Sprintf(
			"createDir failed - testCase%2d:%v", testCaseIndex, err))
	}
}

func TestModify(t *testing.T) {
	testCases := []string{
		"testdst/4",
		"testdst/4/5",
		"testdst/4/6",
	}
	asst := assert.New(t)
	for testCaseIndex, testCase := range testCases {
		// Feasibility test
		err := modify(testCase, USER_NAME, GROUP_NAME, FILE_MODE)
		asst.Nil(err, fmt.Sprintf(
			"createDir failed - testCase%2d:%v", testCaseIndex, err))
	}
}
