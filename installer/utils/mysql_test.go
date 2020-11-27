package utils

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestPrepareMysqlUbuntu(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrepareMysqlUbuntu(); (err != nil) != tt.wantErr {
				t.Errorf("PrepareMysqlUbuntu() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInitMysql(t *testing.T) {
	type args struct {
		dstSQLPath string
		userName   string
		dataDir    string
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
			if err := InitMysql(tt.args.dstSQLPath, tt.args.userName, tt.args.dataDir); (err != nil) != tt.wantErr {
				t.Errorf("InitMysql() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartSingleInst(t *testing.T) {
	type args struct {
		dstDaemonPath string
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
			if err := StartSingleInst(tt.args.dstDaemonPath); (err != nil) != tt.wantErr {
				t.Errorf("StartSingleInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartMultiInst(t *testing.T) {
	type args struct {
		dstSQLPath string
		ports      []int
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
			if err := StartMultiInst(tt.args.dstSQLPath, tt.args.ports); (err != nil) != tt.wantErr {
				t.Errorf("StartMultiInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStopMultiInst(t *testing.T) {
	type args struct {
		dstDaemonPath string
		ports         []int
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
			if err := StopMultiInst(tt.args.dstDaemonPath, tt.args.ports); (err != nil) != tt.wantErr {
				t.Errorf("StopMultiInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateConn(t *testing.T) {
	type args struct {
		port   int
		passwd string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateConn(tt.args.port, tt.args.passwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateConnBySock(t *testing.T) {
	type args struct {
		sockFile string
		passwd   string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateConnBySock(tt.args.sockFile, tt.args.passwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConnBySock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConnBySock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseConn(t *testing.T) {
	type args struct {
		db *sql.DB
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
			if err := CloseConn(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("CloseConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModifyMysqlPwd(t *testing.T) {
	type args struct {
		db     *sql.DB
		passwd string
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
			if err := ModifyMysqlPwd(tt.args.db, tt.args.passwd); (err != nil) != tt.wantErr {
				t.Errorf("ModifyMysqlPwd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGrantUser(t *testing.T) {
	type args struct {
		db         *sql.DB
		permission string
		scope      string
		user       string
		passwd     string
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
			if err := GrantUser(tt.args.db, tt.args.permission, tt.args.scope, tt.args.user, tt.args.passwd); (err != nil) != tt.wantErr {
				t.Errorf("GrantUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTestEnv(t *testing.T) {
	type args struct {
		db *sql.DB
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
			if err := CreateTestEnv(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("CreateTestEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTestData(t *testing.T) {
	type args struct {
		db *sql.DB
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
			if err := CreateTestData(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("CreateTestData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRetriveTestData(t *testing.T) {
	type args struct {
		db *sql.DB
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
			if err := RetriveTestData(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("RetriveTestData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChangeMaster(t *testing.T) {
	type args struct {
		db         *sql.DB
		masterHost string
		masterPort int
		masterUser string
		masterPwd  string
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
			if err := ChangeMaster(tt.args.db, tt.args.masterHost, tt.args.masterPort, tt.args.masterUser, tt.args.masterPwd); (err != nil) != tt.wantErr {
				t.Errorf("ChangeMaster() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartSlave(t *testing.T) {
	type args struct {
		db *sql.DB
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
			if err := StartSlave(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("StartSlave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
