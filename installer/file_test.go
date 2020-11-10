package installer

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

func Test_extractSqlFile(t *testing.T) {
	type args struct {
		srcSqlFile string
		dstSqlPath string
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
			if err := extractSqlFile(tt.args.srcSqlFile, tt.args.dstSqlPath); (err != nil) != tt.wantErr {
				t.Errorf("extractSqlFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_moveCnfFile(t *testing.T) {
	type args struct {
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
			if err := moveCnfFile(tt.args.srcCnfFile, tt.args.dstCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("moveCnfFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
