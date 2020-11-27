package utils

import "testing"

func Test_compareFile(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareFile(); got != tt.want {
				t.Errorf("compareFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_modify(t *testing.T) {
	type args struct {
		dirPath   string
		userName  string
		groupName string
		fileMode  uint32
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
			if err := modify(tt.args.dirPath, tt.args.userName, tt.args.groupName, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("modify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createDirWithDetail(t *testing.T) {
	type args struct {
		dirPath   string
		userName  string
		groupName string
		fileMode  uint32
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
			if err := createDirWithDetail(tt.args.dirPath, tt.args.userName, tt.args.groupName, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("createDirWithDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unTarWithGzip(t *testing.T) {
	type args struct {
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
			if err := unTarWithGzip(tt.args.srcFile, tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("unTarWithGzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExtractSQLFile(t *testing.T) {
	type args struct {
		srcSQLFile string
		dstSQLPath string
		userName   string
		groupName  string
		fileMode   uint32
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
			if err := ExtractSQLFile(tt.args.srcSQLFile, tt.args.dstSQLPath, tt.args.userName, tt.args.groupName, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("ExtractSQLFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMoveCnfFile(t *testing.T) {
	type args struct {
		srcCnfFile string
		dstCnfFile string
		userName   string
		groupName  string
		fileMode   uint32
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
			if err := MoveCnfFile(tt.args.srcCnfFile, tt.args.dstCnfFile, tt.args.userName, tt.args.groupName, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("MoveCnfFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMoveDaemonFile(t *testing.T) {
	type args struct {
		dstSQLPath    string
		srcDaemonFile string
		dstDaemonFile string
		fileMode      uint32
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
			if err := MoveDaemonFile(tt.args.dstSQLPath, tt.args.srcDaemonFile, tt.args.dstDaemonFile, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("MoveDaemonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyDataDir(t *testing.T) {
	type args struct {
		userName    string
		groupName   string
		srcDirPath  string
		dstDirPaths []string
		autoCnfName string
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
			if err := CopyDataDir(tt.args.userName, tt.args.groupName, tt.args.srcDirPath, tt.args.dstDirPaths, tt.args.autoCnfName); (err != nil) != tt.wantErr {
				t.Errorf("CopyDataDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
