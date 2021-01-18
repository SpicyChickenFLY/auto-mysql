package mysql

import "testing"

func TestInitInstance(t *testing.T) {
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
			if err := InitInstance(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("InitInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartSingleInst(t *testing.T) {
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
			if err := StartSingleInst(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("StartSingleInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStopSingleInst(t *testing.T) {
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
			if err := StopSingleInst(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("StopSingleInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartMultiInst(t *testing.T) {
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
			if err := StartMultiInst(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("StartMultiInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStopMultiInst(t *testing.T) {
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
			if err := StopMultiInst(tt.args.servInstInfo); (err != nil) != tt.wantErr {
				t.Errorf("StopMultiInst() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModifyPwdForAllInstOfServer(t *testing.T) {
	type args struct {
		servInstInfo *ServerInstanceInfo
		port         []int
		prevPwd      string
		newPwd       string
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
			if err := ModifyPwdForAllInstOfServer(tt.args.servInstInfo, tt.args.port, tt.args.prevPwd, tt.args.newPwd); (err != nil) != tt.wantErr {
				t.Errorf("ModifyPwdForAllInstOfServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_waitInstanceStartStop(t *testing.T) {
	type args struct {
		servInstInfo *ServerInstanceInfo
		startStop    bool
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
			if err := waitInstanceStartStop(tt.args.servInstInfo, tt.args.startStop); (err != nil) != tt.wantErr {
				t.Errorf("waitInstanceStartStop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
