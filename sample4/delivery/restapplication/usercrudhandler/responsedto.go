package usercrudhandler

type UserGetRespDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	CreatedOn uint64 `json:"createdOn"`
}

type UserGetListRespDTO struct {
	Users []UserGetRespDTO `json:"users"`
	Count int              `json:"count"`
}

type UserCreateRespDTO struct {
	ID string `json:"id"`
}
