package main

import (
	"encoding/json"
	"log"
	"os"
)

// Estrutura DAO para gerenciar a lista de jogadores
type BrasileiraoPlayerDAO struct {
	brasileiraoPlayers []*BrasileiraoPlayer
}

// Função para inicializar a DAO carregando os dados do arquivo JSON
func NewBrasileiraoPlayerDAO(filePath string) *BrasileiraoPlayerDAO {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	var players []*BrasileiraoPlayer
	err = json.Unmarshal(data, &players)
	if err != nil {
		log.Fatalf("Erro ao desserializar JSON: %v", err)
	}

	return &BrasileiraoPlayerDAO{brasileiraoPlayers: players}
}

// Método para obter todos os jogadores
func (dao *BrasileiraoPlayerDAO) GetAll() []*BrasileiraoPlayer {
	return dao.brasileiraoPlayers
}
