package main

type BrasileiraoPlayer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Nationality string `json:"nationality"`
	ShirtNumber int    `json:"shirtNumber"`
	Age         int    `json:"age"`
	Team        string `json:"team"`
}

// Método construtor de brasileiraoPlayer
func NewBrasileiraoPlayer(id int,
	name string,
	position string,
	nationality string,
	shirtNumber int,
	age int,
	team string) *BrasileiraoPlayer {
	player := new(BrasileiraoPlayer)
	player.Id = id
	player.Name = name
	player.Age = age
	player.Nationality = nationality
	player.Team = team
	player.Position = position
	player.ShirtNumber = shirtNumber
	return player
}
