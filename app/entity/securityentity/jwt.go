package securityentity

type UserData struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type TokenSchema struct {
	Bearer  string `json:"bearer"`
	Refresh string `json:"refresh"`
}

type GeneratedResponseJwt struct {
	User  UserData    `json:"User"`
	Token TokenSchema `json:"token"`
}
