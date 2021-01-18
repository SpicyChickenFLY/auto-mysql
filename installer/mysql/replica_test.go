package mysql

import (
	"reflect"
	"testing"
)

func TestCreateMasterSlaveRelation(t *testing.T) {
	type args struct {
		allServInstInfos []*ServerInstanceInfo
		newPwd           string
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
			if err := CreateMasterSlaveRelation(tt.args.allServInstInfos, tt.args.newPwd); (err != nil) != tt.wantErr {
				t.Errorf("CreateMasterSlaveRelation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_testReplication(t *testing.T) {
	type args struct {
		masterServInstInfo *ServerInstanceInfo
		slaveServInstInfos []*ServerInstanceInfo
		newPwd             string
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
			if err := testReplication(tt.args.masterServInstInfo, tt.args.slaveServInstInfos, tt.args.newPwd); (err != nil) != tt.wantErr {
				t.Errorf("testReplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setupMasterInst(t *testing.T) {
	type args struct {
		masterServInst *ServerInstanceInfo
		generalPwd     string
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
			if err := setupMasterInst(tt.args.masterServInst, tt.args.generalPwd); (err != nil) != tt.wantErr {
				t.Errorf("setupMasterInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setupSlaveInst(t *testing.T) {
	type args struct {
		masterServInst *ServerInstanceInfo
		slaveServInsts []*ServerInstanceInfo
		generalPwd     string
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
			if err := setupSlaveInst(tt.args.masterServInst, tt.args.slaveServInsts, tt.args.generalPwd); (err != nil) != tt.wantErr {
				t.Errorf("setupSlaveInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sperateMasterSlaveInstance(t *testing.T) {
	type args struct {
		all []*ServerInstanceInfo
	}
	tests := []struct {
		name               string
		args               args
		wantMasterServInst *ServerInstanceInfo
		wantSlaveServInsts []*ServerInstanceInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMasterServInst, gotSlaveServInsts := sperateMasterSlaveInstance(tt.args.all)
			if !reflect.DeepEqual(gotMasterServInst, tt.wantMasterServInst) {
				t.Errorf("sperateMasterSlaveInstance() gotMasterServInst = %v, want %v", gotMasterServInst, tt.wantMasterServInst)
			}
			if !reflect.DeepEqual(gotSlaveServInsts, tt.wantSlaveServInsts) {
				t.Errorf("sperateMasterSlaveInstance() gotSlaveServInsts = %v, want %v", gotSlaveServInsts, tt.wantSlaveServInsts)
			}
		})
	}
}
