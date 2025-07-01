package main

import (
	"log"
	"micro_services/user_service/config"
	"micro_services/user_service/proto"
	"micro_services/user_service/utils"
)

// func main() {
// 	finalPwd := utils.GeneratePassWord("12345678")

// 	isPass := utils.VerifyPassword("12345678", finalPwd)

// 	fmt.Println(isPass)
// }

func main() {
	conn, err := utils.GrpcDial(*config.IP, *config.PORT)

	if err != nil {
		log.Fatalf("Failed to connect gRPC: %v", err)
		return
	}

	client := proto.NewUserClient(conn)

	// 获取 用户列表
	// resp, _ := client.GetUserList(context.Background(), &proto.PageInfo{
	// 	PageNumber: 3,
	// 	PageSize:   5,
	// })

	// for _, user := range resp.Users {
	// 	fmt.Println(user)
	// }

	// 获取用户
	// resp, _ := client.GetUser(context.Background(), &proto.UserInfo{
	// 	// Id: 3,
	// 	MobileNumber: "13800000021",
	// })

	// fmt.Println(resp)

	// CreateUser
	// resp, error := client.CreateUser(context.Background(), &proto.UserInfo{
	// 	MobileNumber: "13800000033",
	// 	Password:     "qwer1234",
	// 	Nickname:     "xl",
	// })

	// if error != nil {
	// 	log.Fatalf("%v", error)
	// }

	// fmt.Println(resp)

	// 更新用户
	// resp, err := client.UpdateUser(context.Background(), &proto.UserInfo{
	// 	Id:           12,
	// 	Nickname:     "lisi",
	// 	MobileNumber: "1388886666",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(resp)

	// 验证密码
	// resp, _ := client.VerifyPassword(context.Background(), &proto.PasswordVerify{
	// 	Id:          30,
	// 	RawPassword: "qwer1234",
	// })

	// fmt.Println(resp.IsPass)

	// 更新 MobileNumber
	// resp, err := client.UpdateMobileNumber(context.Background(), &proto.UserInfo{
	// 	Id:           1,
	// 	MobileNumber: "13111313333",
	// })

	// fmt.Println(err)
	// fmt.Println(resp)

	// 更新密码
	// resp, err := client.UpdatePassword(context.Background(), &proto.UserInfo{
	// 	Id:       30,
	// 	Password: "qwer1234",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(resp)
}
