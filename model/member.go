package model

type Member struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Test struct {
	Id      int `json:"id"`
	Balance int `json:"balance"`
}
