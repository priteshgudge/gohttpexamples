package domain

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Age       int    `json:"age"`
	CreatedOn uint64 `json:"createdOn"`
}
