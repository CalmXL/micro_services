package main

import (
	"fmt"
	"micro_services/user_service/utils"
)

func main() {
	finalPwd := utils.GeneratePassWord("12345678")

	isPass := utils.VerifyPassword("12345678", finalPwd)

	fmt.Println(isPass)
}
