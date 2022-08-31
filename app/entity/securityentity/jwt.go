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
	UserData    UserData    `json:"userData"`
	TokenSchema TokenSchema `json:"token"`
}
