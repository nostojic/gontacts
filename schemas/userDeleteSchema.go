package schemas

type UserDeleteSchema struct {
	Password string `json:"password,omitempty"`
	Email string `json:"email,omitempty"`
}