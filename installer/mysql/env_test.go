package mysql

import "testing"

func TestPrepareMysqlUbuntu(t *testing.T) {
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
			if err := PrepareMysqlUbuntu(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("PrepareMysqlUbuntu() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrepareMysqlCentos(t *testing.T) {
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
			if err := PrepareMysqlCentos(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("PrepareMysqlCentos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
