package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	router := http.NewServeMux()
	db := NewDatabase()
	defer db.Close()
	brasileiraoPlayerDAO := NewBrasileiraoPlayerDAO(db)

	// Nova rota para obter todos os jogadores do Brasileirão
	router.HandleFunc("/brasileiraoPlayers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		players := brasileiraoPlayerDAO.GetAll()
		err := json.NewEncoder(w).Encode(players)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	c := cors.AllowAll()
	r := c.Handler(router)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Ouvindo na porta 8080...")
}
