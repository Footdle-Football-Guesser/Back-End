package main

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Score string `json:"score"`
}

func NewUser(id int, name string, score string) *User {
	user := new(User)
	user.Id = id
	user.Name = name
	user.Score = score
	return user
}
