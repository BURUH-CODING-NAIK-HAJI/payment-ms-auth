package securityentity

type UserData struct {
	Id   string
	Name string
}

type TokenSchema struct {
	Bearer  string `json:"bearer"`
	Refresh string `json:"refresh"`
}

type GeneratedResponseJwt struct {
	UserData    UserData    `json:"userData"`
	TokenSchema TokenSchema `json:"token"`
}
