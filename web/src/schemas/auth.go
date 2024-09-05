package schemas

type Login struct {
	Username string `json:"username"`
	Password string `json:"id"`
}

type CreateAccount struct {
	Id                   int    `json:"id"`
	Username             string `json:"username"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}
