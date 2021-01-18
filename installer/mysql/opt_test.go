package mysql

import (
	"reflect"
	"testing"
)

func TestNewServInstInfo(t *testing.T) {
	type args struct {
		infoStr string
	}
	tests := []struct {
		name             string
		args             args
		wantServInstInfo *ServerInstanceInfo
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotServInstInfo, err := NewServInstInfo(tt.args.infoStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServInstInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotServInstInfo, tt.wantServInstInfo) {
				t.Errorf("NewServInstInfo() = %v, want %v", gotServInstInfo, tt.wantServInstInfo)
			}
		})
	}
}

func TestServerInstanceInfo_findInstByPort(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name    string
		s       *ServerInstanceInfo
		args    args
		want    *InstanceInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.findInstByPort(tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerInstanceInfo.findInstByPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerInstanceInfo.findInstByPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseServerStr(t *testing.T) {
	type args struct {
		infoStr string
	}
	tests := []struct {
		name                 string
		args                 args
		wantAllServInstInfos []*ServerInstanceInfo
		wantErr              bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAllServInstInfos, err := ParseServerStr(tt.args.infoStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseServerStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAllServInstInfos, tt.wantAllServInstInfos) {
				t.Errorf("ParseServerStr() = %v, want %v", gotAllServInstInfos, tt.wantAllServInstInfos)
			}
		})
	}
}

func TestCreateMysqlUserWithGroup(t *testing.T) {
	type args struct {
		servInstInfo *ServerInstanceInfo
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
			if err := CreateMysqlUserWithGroup(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("CreateMysqlUserWithGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShowErrorLog(t *testing.T) {
	type args struct {
		servInstInfo *ServerInstanceInfo
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
			if err := ShowErrorLog(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("ShowErrorLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKillMysqlProcess(t *testing.T) {
	type args struct {
		servInstInfo *ServerInstanceInfo
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
			if err := KillMysqlProcess(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("KillMysqlProcess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
