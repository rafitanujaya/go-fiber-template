package response

type UserRegister struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
