package utils

import (
	"reflect"
	"testing"
)

func TestCheckCnfDir(t *testing.T) {
	type args struct {
		srcCnfFile string
		userName   string
		groupName  string
		fileMode   uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		want1   []string
		want2   []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := CheckCnfDir(tt.args.srcCnfFile, tt.args.userName, tt.args.groupName, tt.args.fileMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckCnfDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckCnfDir() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CheckCnfDir() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("CheckCnfDir() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
