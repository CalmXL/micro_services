package main

import (
	"micro_services/user_service/model"
	"micro_services/user_service/utils"
	"strconv"

	"gorm.io/gorm"
)

func generateUsers(db *gorm.DB) {
	for i := 0; i < 30; i++ {
		suffix := ""

		if i < 10 {
			suffix = "0" + strconv.Itoa(i)
		} else {
			suffix = strconv.Itoa(i)
		}

		user := model.User{
			MobileNumber: "138000000" + suffix,
			Password:     utils.GeneratePassWord("a12345678"),
			NickName:     "user-" + suffix,
		}

		db.Save(&user)
	}
}
