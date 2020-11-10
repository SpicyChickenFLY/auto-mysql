package installer

import "testing"

func TestInstall(t *testing.T) {
	type args struct {
		srcSqlFile string
		dstSqlPath string
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
			if err := Install(tt.args.srcSqlFile, tt.args.dstSqlPath, tt.args.srcCnfFile, tt.args.dstCnfFile); (err != nil) != tt.wantErr {
				t.Errorf("Install() error = %v, wantErr %v", err, tt.wantErr)
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
