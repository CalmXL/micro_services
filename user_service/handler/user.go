package handler

import (
	"context"
	"micro_services/user_service/model"
	"micro_services/user_service/proto"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
	proto.UnimplementedUserServer
}

func (u *User) GetUserList(c context.Context, r *proto.PageInfo) (*proto.UserList, error) {

	var userList []model.User
	result := u.DB.Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}

	resp := &proto.UserList{}
	resp.UserCount = int32(result.RowsAffected)

	return nil, nil
}
func (u *User) GetUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	return nil, nil
}
func (u *User) CreateUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	return nil, nil
}
func (u *User) UpdateUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	return nil, nil
}
func (u *User) VerifyPassword(c context.Context, r *proto.PasswordVerify) (*proto.PasswordVerifyPass, error) {
	return nil, nil
}
