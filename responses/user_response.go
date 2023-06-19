package responses

// TODO: is this necessary?
type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
