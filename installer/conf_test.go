package installer

import "testing"

func Test_checkCnfDir(t *testing.T) {
	type args struct {
		srcCnfFile string
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
			if err := checkCnfDir(tt.args.srcCnfFile, tt.args.userName, tt.args.groupName, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("checkCnfDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
