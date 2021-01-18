package installer

import (
	"testing"

	"github.com/SpicyChickenFLY/auto-mysql/installer/mysql"
)

func Test_prepare(t *testing.T) {
	type args struct {
		servInstInfo *mysql.ServerInstanceInfo
		srcSQLFile   string
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
			if err := prepare(tt.args.servInstInfo, tt.args.srcSQLFile, tt.args.srcCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("prepare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstallCustomInstance(t *testing.T) {
	type args struct {
		servInstInfo *mysql.ServerInstanceInfo
		srcSQLFile   string
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
			if err := InstallCustomInstance(tt.args.servInstInfo, tt.args.srcSQLFile, tt.args.srcCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("InstallCustomInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstallStandardMultiInstanceOnMultiServer(t *testing.T) {
	type args struct {
		srcSQLFile string
		infoStr    string
		mysqlPwd   string
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
			if err := InstallStandardMultiInstanceOnMultiServer(tt.args.srcSQLFile, tt.args.infoStr, tt.args.mysqlPwd); (err != nil) != tt.wantErr {
				t.Errorf("InstallStandardMultiInstanceOnMultiServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
