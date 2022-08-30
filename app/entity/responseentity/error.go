package responseentity

type Error struct {
	Id      string `json:"id"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
