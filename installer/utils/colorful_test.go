package utils

import (
	"fmt"
	"testing"
)

func TestRenderStr(t *testing.T) {
	type args struct {
		str       string
		modeStr   string
		bColorStr string
		fColorStr string
	}
	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testOK",
			args: args{
				str:       "OK",
				modeStr:   "default",
				bColorStr: "black",
				fColorStr: "green",
			},
			want: fmt.Sprintf("%c[%d;%d;%dm%s%c[0m",
				MARK, ModeDefault, BackBlack, FrontGreen, "OK", MARK),
		},
		{
			name: "testFAIL",
			args: args{
				str:       "FAIL",
				modeStr:   "line",
				bColorStr: "white",
				fColorStr: "red",
			},
			want: fmt.Sprintf("%c[%d;%d;%dm%s%c[0m",
				MARK, ModeLine, BackWhite, FrontRed, "FAIL", MARK),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := RenderStr(
				testCase.args.str,
				testCase.args.modeStr,
				testCase.args.bColorStr,
				testCase.args.fColorStr); got != testCase.want {
				t.Errorf("RenderStr() = %v, want %v", got, testCase.want)
			}
		})
	}
}
