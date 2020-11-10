package shell

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
		{
			name: "OK",
			args: args{
				cmdStr: "echo OK",
			},
			wantErr: false,
		},
		{
			name: "ERROR",
			args: args{
				cmdStr: "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExecCommand(tt.args.cmdStr); (err != nil) != tt.wantErr {
				t.Errorf("ExecCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
