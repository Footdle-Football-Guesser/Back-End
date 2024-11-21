package main

import "fmt"

// Estrutura DAO para gerenciar a lista de jogadores
type BrasileiraoPlayerDAO struct {
	db                 *Database
	brasileiraoPlayers []*BrasileiraoPlayer
}

func NewBrasileiraoPlayerDAO(db *Database) *BrasileiraoPlayerDAO {
	brasileiraoPlayerDAO := new(BrasileiraoPlayerDAO)
	brasileiraoPlayerDAO.db = db
	brasileiraoPlayerDAO.brasileiraoPlayers = make([]*BrasileiraoPlayer, 0)
	return brasileiraoPlayerDAO
}

// Método para obter todos os jogadores
func (dao *BrasileiraoPlayerDAO) GetAll() []*BrasileiraoPlayer {
	db := dao.db.Get()
	rows, err := db.Query(`SELECT * FROM "brasileiraoPlayers"`)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var id int
		var name string
		var position string
		var nationality string
		var age int
		var shirtNumber int
		var team string
		rows.Scan(&id, &name, &position, &nationality, &age, &shirtNumber, &team)
		dao.brasileiraoPlayers = append(dao.brasileiraoPlayers, NewBrasileiraoPlayer(id, name, position, nationality, age, shirtNumber, team))
	}
	return dao.brasileiraoPlayers
}

func (playerDAO *BrasileiraoPlayerDAO) Insert(player *BrasileiraoPlayer) error {
	db := playerDAO.db.Get()
	query := `INSERT INTO "brasileiraoPlayers" ("name", "position", nationality, "shirtNumber", age, team) values ($1, $2, $3, $4, $5, $6)`
	_, err := db.Query(query, player.Name, player.Position, player.Nationality, player.ShirtNumber, player.Age, player.Team)
	if err != nil {
		return err
	}
	return nil
}

func (playerDAO *BrasileiraoPlayerDAO) Delete(id string) error {
	db := playerDAO.db.Get()
	query := `DELETE FROM "brasileiraoPlayers" WHERE id=$1`
	_, err := db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (playerDAO *BrasileiraoPlayerDAO) Update(id string, player *BrasileiraoPlayer) error {
	db := playerDAO.db.Get()
	query := `
		UPDATE "brasileiraoPlayers"
		SET "name"=$1, "position"=$2, nationality=$3, "shirtNumber"=$4, age=$5, team=$6
		WHERE id=$7
	`
	_, err := db.Exec(query, player.Name, player.Position, player.Nationality, player.ShirtNumber, player.Age, player.Team, id)
	if err != nil {
		return err
	}
	return nil
}
