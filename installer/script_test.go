package installer

import (
	"reflect"
	"testing"
)

func TestInstall(t *testing.T) {
	type args struct {
		srcSQLFile string
		dstSQLPath string
		srcCnfFile string
		dstCnfFile string
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
			if err := InstallSingleInstance(tt.args.srcSQLFile, tt.args.dstSQLPath, tt.args.srcCnfFile, tt.args.dstCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("InstallSingleInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Remove()
		})
	}
}

func Test_checkErr(t *testing.T) {
	type args struct {
		err  error
		info string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkErr(tt.args.err, tt.args.info)
		})
	}
}

func TestInstallSingleInstance(t *testing.T) {
	type args struct {
		srcSQLFile string
		dstSQLPath string
		srcCnfFile string
		dstCnfFile string
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
			if err := InstallSingleInstance(tt.args.srcSQLFile, tt.args.dstSQLPath, tt.args.srcCnfFile, tt.args.dstCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("InstallSingleInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstallMultiInstance(t *testing.T) {
	type args struct {
		srcSQLFile string
		dstSQLPath string
		srcCnfFile string
		dstCnfFile string
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
			if err := InstallMultiInstance(tt.args.srcSQLFile, tt.args.dstSQLPath, tt.args.srcCnfFile, tt.args.dstCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("InstallMultiInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_prepare(t *testing.T) {
	type args struct {
		srcSQLFile string
		dstSQLPath string
		srcCnfFile string
		dstCnfFile string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []string
		want2 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := prepare(tt.args.srcSQLFile, tt.args.dstSQLPath, tt.args.srcCnfFile, tt.args.dstCnfFile)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepare() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("prepare() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("prepare() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_setupMasterInst(t *testing.T) {
	type args struct {
		sockFileMaster string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMasterInst(tt.args.sockFileMaster)
		})
	}
}

func Test_setupSlaveInst(t *testing.T) {
	type args struct {
		sockFileSlaves []string
		portMaster     int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupSlaveInst(tt.args.sockFileSlaves, tt.args.portMaster)
		})
	}
}

func Test_testReplication(t *testing.T) {
	type args struct {
		sockFileMaster string
		sockFileSlaves []string
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
			if err := testReplication(tt.args.sockFileMaster, tt.args.sockFileSlaves); (err != nil) != tt.wantErr {
				t.Errorf("testReplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
