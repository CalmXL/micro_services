package handler

import (
	"context"
	"micro_services/user_service/config"
	"micro_services/user_service/model"
	"micro_services/user_service/proto"
	"micro_services/user_service/utils"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
	proto.UnimplementedUserServer
}

func UserModelToResponse(user model.User) *proto.UserInfo {
	return &proto.UserInfo{
		Id:           user.ID,
		MobileNumber: user.MobileNumber,
		Nickname:     user.NickName,
		Password:     user.Password,
		Gender:       user.Gender,
		Role:         user.Role,
	}
}

func (u *User) GetUserList(c context.Context, r *proto.PageInfo) (*proto.UserList, error) {

	var userList []model.User
	result := u.DB.Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}

	resp := &proto.UserList{}
	resp.UserCount = int32(result.RowsAffected)

	u.DB.Scopes(utils.Paginate(r.PageNumber, r.PageSize)).Find(&userList)

	for _, user := range userList {
		userInfo := UserModelToResponse(user)
		resp.Users = append(resp.Users, userInfo)
	}

	return resp, nil
}
func (u *User) GetUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	/**
	id,
	mobileNumber
	*/

	var user *model.User
	var err error

	if r.Id > 0 {
		user, err = getUserByWhatever(u.DB, r, "ID")
	}

	if r.MobileNumber != "" {
		user, err = getUserByWhatever(u.DB, r, "MobileNumber")
	}

	if err != nil {
		return nil, err
	}

	return UserModelToResponse(*user), nil
}
func (u *User) CreateUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	/**
	1. MobileNumber (unique), password 必传参数
	2. Nickname 可传
	*/
	var user model.User

	result := u.DB.Where(&model.User{
		MobileNumber: r.MobileNumber,
	}).First(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "%s", result.Error.Error())
	}

	if result.RowsAffected != 0 {
		return nil, status.Errorf(codes.AlreadyExists, "user already exists")
	}

	user.MobileNumber = r.MobileNumber
	user.NickName = r.Nickname
	user.Password = utils.GeneratePassWord(r.Password)

	result = u.DB.Create(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "%s", result.Error.Error())
	}

	resp := UserModelToResponse(user)

	return resp, nil
}
func (u *User) UpdateUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {

	user, err := getUserByWhatever(u.DB, r, "ID")

	if err != nil {
		return nil, err
	}

	user.NickName = r.Nickname
	user.Gender = r.Gender

	result := u.DB.Save(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "%s", result.Error.Error())
	}

	return UserModelToResponse(*user), nil
}
func (u *User) VerifyPassword(c context.Context, r *proto.PasswordVerify) (*proto.PasswordVerifyPass, error) {

	var user model.User

	result := u.DB.First(r.Id)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "%s", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	encodedPassword := user.Password

	isPass := utils.VerifyPassword(r.RawPassword, encodedPassword)

	return &proto.PasswordVerifyPass{
		IsPass: isPass,
	}, nil
}

func (u *User) UpdateMobileNumber(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	user, error := getUserByWhatever(u.DB, r, "MobileNumber")

	if error != nil {
		return nil, error
	}

	matched, err := regexp.MatchString(config.REGPHONENUMBER, r.MobileNumber)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if !matched {
		return nil, status.Errorf(codes.InvalidArgument, "Mobile number is valid")
	}

	user.MobileNumber = r.MobileNumber

	result := u.DB.Save(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	resp := UserModelToResponse(*user)

	return resp, nil
}

func (u *User) UpdatePassword(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	user, error := getUserByWhatever(u.DB, r, "ID")

	if error != nil {
		return nil, error
	}

	result := u.DB.First(&user, r.Id)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "%s", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	user.Password = utils.GeneratePassWord(r.Password)

	return nil, nil
}

func getUserByWhatever(db *gorm.DB, r *proto.UserInfo, field string) (*model.User, error) {
	var user model.User

	switch field {
	case "ID":
		user = model.User{
			ID: r.Id,
		}
	case "MobileNumber":
		user = model.User{
			MobileNumber: r.MobileNumber,
		}
	default:
		user = model.User{
			ID: r.Id,
		}
	}

	result := db.Where(&user).First(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "%s", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "User Not Found")
	}

	return &user, nil
}
