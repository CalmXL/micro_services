package main

import (
	"log"
	"micro_services/user_service/config"
	"micro_services/user_service/utils"
)

func main() {
	db, errConn := utils.DBConnect(config.DBName)

	if errConn != nil {
		log.Fatalf("Failed to start connect DB: %v", errConn)
		return
	}

	// if errMigrate := db.AutoMigrate(
	// 	&model.User{},
	// ); errMigrate != nil {
	// 	log.Fatalf("Failed to migrate the data model: %v", errMigrate)
	// 	return
	// }

	generateUsers(db)

	defer utils.DBClose(db)
}
