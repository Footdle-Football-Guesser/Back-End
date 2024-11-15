package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database class
type Database struct {
	connection *sql.DB
	dbtype     string
	info       string
	err        error
}

// NewDatabase Constructor
func NewDatabase() *Database {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Ler as variáveis de ambiente
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Montar a string de conexão
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpassword, dbname)

	// Criar a instância do Database
	db := &Database{
		dbtype: "postgres",
		info:   dbInfo,
	}

	return db
}

// Get retorna a conexão com o banco
func (db *Database) Get() *sql.DB {
	if db.connection == nil {
		db.connection, db.err = sql.Open(db.dbtype, db.info)
		if db.err != nil {
			log.Fatalf("Erro ao conectar ao banco de dados: %v", db.err)
		}
	}
	return db.connection
}

// Close fecha a conexão com o banco
func (db *Database) Close() {
	db.connection.Close()
}
