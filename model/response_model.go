package model

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type RespRegister struct {
	User_id int    `json:"user_id"`
	Name    string `json:"name"`
}
