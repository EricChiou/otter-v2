package model

type User struct {
	ID          string `pk:"id"`
	Acc         string
	Pwd         string
	Name        string
	RoleCode    string
	Status      string
	CreatedDate string
	UpdatedDate string
}
