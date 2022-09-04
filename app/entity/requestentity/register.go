package requestentity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Register struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (register *Register) Validate() error {
	payload := *register
	return validation.ValidateStruct(
		&payload,
		validation.Field(&payload.Name, validation.Required),
		validation.Field(&payload.Username, validation.Required),
		validation.Field(&payload.Password, validation.Length(8, 32)),
	)
}
