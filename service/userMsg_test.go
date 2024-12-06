package service

import (
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := repository.Init0(); err != nil {
		os.Exit(1)
	}
	if err := util.InitLogger(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestCreateUserMsg(t *testing.T) {
	//  带有参数的结构体
	type args struct {
		userId, name, phone, avatar, address string
	}
	// 测试用例
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "添加用户信息",
			args: args{
				userId:  "12345",
				name:    "John Doe",
				phone:   "13361641234",
				avatar:  "https://example.com/avatar.jpg",
				address: "123 Main St",
			},
			wantErr: false,
		},
	}
	// 测试发布函数功能
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateUserMsg(tt.args.userId, tt.args.name, tt.args.phone, tt.args.avatar, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("CREATEuserMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateUserMsg(t *testing.T) {
	//  带有参数的结构体
	type args struct {
		userId, name, phone, avatar, address string
	}
	// 测试用例
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "更新用户信息",
			args: args{
				userId:  "12345",
				name:    "haha",
				phone:   "13361641234",
				avatar:  "https://example.com/avatar.jpg",
				address: "123 Main St",
			},
			wantErr: false,
		},
	}
	// 测试发布函数功能
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateUserMsg(tt.args.userId, tt.args.name, tt.args.phone, tt.args.avatar, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("UPDATEUSERMSG() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQueryUserMsg(t *testing.T) {
	//  带有参数的结构体
	type args struct {
		userId string
	}
	// 测试用例
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "查询用户信息",
			args: args{
				userId: "12345",
			},
			wantErr: false,
		},
	}
	// 测试发布函数功能
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := QueryUserMsg(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QUERYUSERMSG() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
