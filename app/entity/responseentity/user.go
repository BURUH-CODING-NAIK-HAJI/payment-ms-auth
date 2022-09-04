package responseentity

type User struct {
	Id       string      `json:"id"`
	Username string      `json:"username"`
	Password interface{} `json:"password,omitempty"`
	Profile  Profile     `json:"profile"`
}
