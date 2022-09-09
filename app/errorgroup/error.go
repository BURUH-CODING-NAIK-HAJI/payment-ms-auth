package errorgroup

type Error struct {
	Id      interface{} `json:"id"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func (error Error) Error() string {
	return error.Message
}
