package installer

import (
	"os/user"
	"reflect"
	"testing"
)

func Test_createUserWithGroup(t *testing.T) {
	type args struct {
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
			if err := createUserWithGroup(tt.args.userName, tt.args.groupName); (err != nil) != tt.wantErr {
				t.Errorf("createUserWithGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_modifyPwdForUser(t *testing.T) {
	type args struct {
		userName string
		usePwd   string
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
			if err := modifyPwdForUser(tt.args.userName, tt.args.usePwd); (err != nil) != tt.wantErr {
				t.Errorf("modifyPwdForUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_findGroup(t *testing.T) {
	type args struct {
		groupName string
	}
	tests := []struct {
		name  string
		args  args
		want  *user.Group
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findGroup(tt.args.groupName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findGroup() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findGroup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findUser(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name  string
		args  args
		want  *user.User
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findUser(tt.args.userName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
