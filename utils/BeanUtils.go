package utils

import (
	"fmt"

	"github.com/lhs960906/myprojects/bean/user"
)

func PrintUserIdentity(u user.User) {
	jsonBytes := u.GetIdentity()
	fmt.Println(string(jsonBytes))
}

