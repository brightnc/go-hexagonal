package dto

type CustomerResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Status   int    `json:"status"`
}
