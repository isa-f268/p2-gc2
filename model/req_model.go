package model

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoanReq struct {
	BookID int `json:"bookId"`
}
