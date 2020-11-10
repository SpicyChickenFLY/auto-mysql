package installer

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_prepareMysqlUbuntu(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := prepareMysqlUbuntu(); (err != nil) != tt.wantErr {
				t.Errorf("prepareMysqlUbuntu() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_initMysql(t *testing.T) {
	type args struct {
		dstPath  string
		userName string
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
			if err := initMysql(tt.args.dstPath, tt.args.userName); (err != nil) != tt.wantErr {
				t.Errorf("initMysql() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_startMysql(t *testing.T) {
	type args struct {
		dstPath string
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
			if err := startMysql(tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("startMysql() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_modifyMysqlPwd(t *testing.T) {
	type args struct {
		userName string
		userPwd  string
		sockPath string
		dbName   string
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
			if err := modifyMysqlPwd(tt.args.userName, tt.args.userPwd, tt.args.sockPath, tt.args.dbName); (err != nil) != tt.wantErr {
				t.Errorf("modifyMysqlPwd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
