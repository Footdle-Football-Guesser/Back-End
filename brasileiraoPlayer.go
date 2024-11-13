package main

import (
	"encoding/json"
	"os"
)

type BrasileiraoPlayer struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	Nationality string `json:"nationality"`
	Age         int32  `json:"age"`
	ShirtNumber int32  `json:"shirtNumber"`
	Team        string `json:"team"`
}

func getAllPlayers() ([]BrasileiraoPlayer, error) {
	data, err := os.ReadFile("brasileiraoPlayersBD.json")
	if err != nil {
		return nil, err
	}

	var players []BrasileiraoPlayer
	err = json.Unmarshal(data, &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}
