package model

type User struct {
	GormModel
	ID           int32  `json:"id" gorm:"primaryKey"`
	MobileNumber string `json:"mobile_number" gorm:"index_mobile_number;unique;type:varchar(11);not null"`
	NickName     string `json:"nick_name" gorm:"type:varchar(15);unique"`
	Password     string `json:"password" gorm:"type:varchar(100);not null"`
	Gender       int32  `json:"gender" gorm:"default:0;type:tinyint(1) comment '0 for male, 1 for female'"`
	Role         int32  `json:"role" gorm:"default:1;type:tinyint(1) comment '0 for banner-user, 1 for common-user, 2 for administrator'"`
}
