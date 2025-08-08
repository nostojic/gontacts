package schemas

import "strings"

type UserCreateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func (u *UserCreateInput) TrimWhitespace() {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
}