package main

import "fmt"

// Estrutura DAO para gerenciar a lista de jogadores
type BrasileiraoPlayerDAO struct {
	db                 *Database
	brasileiraoPlayers []*BrasileiraoPlayer
}

// Função para inicializar a DAO carregando os dados do arquivo JSON
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
