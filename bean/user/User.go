package user

import (
	"encoding/json"

	"github.com/lhs960906/myprojects/bean/identity"
)

type User struct {
	Id       string
	Email    string
	Password string
	Identity identity.Identity
}

func (user *User) GetIdentity() []byte {
	jsonBytes, _ := json.Marshal(user.Identity)
	return jsonBytes
}

