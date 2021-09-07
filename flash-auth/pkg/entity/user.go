package entity

type User struct {
	Id         string   `json:"id"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Email      string   `json:"email"`
	Name       string   `json:"name"`
	AccessList []Access `json:"accessList"`
}

type Access struct {
	Service  string `json:"service"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
