package user

const (
	Table string = "user"
	PK    string = "id"
)

const (
	ID        string = "id"
	Acc       string = "account"
	Pwd       string = "password"
	Name      string = "name"
	RoleCode  string = "role_code"
	Status    string = "status"
	CreatedAt string = "created_at"
	UpdatedAt string = "updated_at"
)

type Entity struct {
	ID        string `json:"id,omitempty"`
	Acc       string `json:"account,omitempty"`
	Pwd       string `json:"password,omitempty"`
	Name      string `json:"name,omitempty"`
	RoleCode  string `json:"role_code,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
