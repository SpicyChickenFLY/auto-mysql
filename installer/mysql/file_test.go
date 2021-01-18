package mysql

import (
	"testing"

	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/linux"
)

func Test_modifyDataDir(t *testing.T) {
	type args struct {
		s       *linux.ServerInfo
		dirPath string
		mode    uint32
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
			if err := modifyDataDir(tt.args.s, tt.args.dirPath, tt.args.mode); (err != nil) != tt.wantErr {
				t.Errorf("modifyDataDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createFile(t *testing.T) {
	type args struct {
		s        *linux.ServerInfo
		filePath string
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
			if err := createFile(tt.args.s, tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("createFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createDir(t *testing.T) {
	type args struct {
		s       *linux.ServerInfo
		dirPath string
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
			if err := createDir(tt.args.s, tt.args.dirPath); (err != nil) != tt.wantErr {
				t.Errorf("createDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unTarWithGzip(t *testing.T) {
	type args struct {
		s       *linux.ServerInfo
		srcFile string
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
			if err := unTarWithGzip(tt.args.s, tt.args.srcFile, tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("unTarWithGzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExtractSoftware(t *testing.T) {
	type args struct {
		s          *linux.ServerInfo
		srcSQLFile string
		dstSQLPath string
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
			if err := ExtractSoftware(tt.args.s, tt.args.srcSQLFile, tt.args.dstSQLPath); (err != nil) != tt.wantErr {
				t.Errorf("ExtractSoftware() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMoveCnfFile(t *testing.T) {
	type args struct {
		s          *linux.ServerInfo
		srcCnfFile string
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
			if err := MoveCnfFile(tt.args.s, tt.args.srcCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("MoveCnfFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMoveDaemonFile(t *testing.T) {
	type args struct {
		servInstInfo     *ServerInstanceInfo
		srcDaemonFileRel string
		dstDaemonFile    string
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
			if err := MoveDaemonFile(tt.args.servInstInfo, tt.args.srcDaemonFileRel, tt.args.dstDaemonFile); (err != nil) != tt.wantErr {
				t.Errorf("MoveDaemonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyDataDir(t *testing.T) {
	type args struct {
		masterServInstInfo *ServerInstanceInfo
		allServInstInfos   []*ServerInstanceInfo
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
			if err := CopyDataDir(tt.args.masterServInstInfo, tt.args.allServInstInfos); (err != nil) != tt.wantErr {
				t.Errorf("CopyDataDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
