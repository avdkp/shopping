package domain

type User struct {
	Id             int    `json:"id"`
	UserName       string `json:"username"`
	HashedPassword string `json:"password"`
	Role           string `json:"role"`
}
