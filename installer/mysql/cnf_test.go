package mysql

import (
	"testing"

	"github.com/SpicyChickenFLY/ini"
)

func TestCheckRquireDirExists(t *testing.T) {
	type args struct {
		servInstInfo *ServerInstanceInfo
		srcCnfFile   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckRquireDirExists(tt.args.servInstInfo, tt.args.srcCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("CheckRquireDirExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateStdCnf(t *testing.T) {
	type args struct {
		servInstInfo *ServerInstanceInfo
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateStdCnf(tt.args.servInstInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateStdCnf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateStdCnf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceAllValueInSection(t *testing.T) {
	type args struct {
		sec         *ini.Section
		placeHolder string
		replaceStr  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceAllValueInSection(tt.args.sec, tt.args.placeHolder, tt.args.replaceStr)
		})
	}
}
