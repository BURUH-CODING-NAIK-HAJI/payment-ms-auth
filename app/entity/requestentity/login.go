package requestentity

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *Login) Validate() error {
	payload := *login
	return validation.ValidateStruct(&payload,
		validation.Field(&payload.Username, validation.Required),
		validation.Field(&payload.Password, validation.Required),
	)
}
