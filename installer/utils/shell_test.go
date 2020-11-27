package utils

import "testing"

func TestExecCommand(t *testing.T) {
	type args struct {
		cmdStr string
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
			if err := ExecCommand(tt.args.cmdStr); (err != nil) != tt.wantErr {
				t.Errorf("ExecCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMv(t *testing.T) {
	type args struct {
		srcFile string
		dstFile string
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
			if err := Mv(tt.args.srcFile, tt.args.dstFile); (err != nil) != tt.wantErr {
				t.Errorf("Mv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCp(t *testing.T) {
	type args struct {
		srcFile string
		dstFile string
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
			if err := Cp(tt.args.srcFile, tt.args.dstFile); (err != nil) != tt.wantErr {
				t.Errorf("Cp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRm(t *testing.T) {
	type args struct {
		dstFile string
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
			if err := Rm(tt.args.dstFile); (err != nil) != tt.wantErr {
				t.Errorf("Rm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChown(t *testing.T) {
	type args struct {
		dirPath   string
		userName  string
		groupName string
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
			if err := Chown(tt.args.dirPath, tt.args.userName, tt.args.groupName); (err != nil) != tt.wantErr {
				t.Errorf("Chown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChmod(t *testing.T) {
	type args struct {
		dirPath  string
		fileMode uint32
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
			if err := Chmod(tt.args.dirPath, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("Chmod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMkdir(t *testing.T) {
	type args struct {
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
			if err := Mkdir(tt.args.dirPath); (err != nil) != tt.wantErr {
				t.Errorf("Mkdir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTar(t *testing.T) {
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
			if err := Tar(tt.args.srcFile, tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("Tar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseradd(t *testing.T) {
	type args struct {
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
			if err := Useradd(tt.args.userName); (err != nil) != tt.wantErr {
				t.Errorf("Useradd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGroupadd(t *testing.T) {
	type args struct {
		groupName string
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
			if err := Groupadd(tt.args.groupName); (err != nil) != tt.wantErr {
				t.Errorf("Groupadd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseraddWithGroup(t *testing.T) {
	type args struct {
		groupName string
		userName  string
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
			if err := UseraddWithGroup(tt.args.groupName, tt.args.userName); (err != nil) != tt.wantErr {
				t.Errorf("UseraddWithGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
