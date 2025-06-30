package utils

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

var opts = &password.Options{16, 100, 32, md5.New}

func GeneratePassWord(rawPassword string) string {
	salt, encodedPwd := password.Encode(rawPassword, opts)
	encodedPassword := fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return encodedPassword
}

func VerifyPassword(rawPassword string, encodedPassword string) bool {
	passwordInfo := strings.Split(encodedPassword, "$")
	salt := passwordInfo[1]
	encodedPwd := passwordInfo[2]

	isPass := password.Verify(rawPassword, salt, encodedPwd, opts)

	return isPass
}
