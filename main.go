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

	// GET
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

	//POST
	router.HandleFunc("/newBrasileiraoPlayer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		var player BrasileiraoPlayer
		err := json.NewDecoder(r.Body).Decode(&player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = brasileiraoPlayerDAO.Insert(&player)

		if err != nil {
			http.Error(w, "Erro ao deletar o player: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Resposta de sucesso
		response := map[string]string{
			"message": "Player cadastrado com sucesso!",
		}
		w.WriteHeader(http.StatusCreated) // 201 - Created
		json.NewEncoder(w).Encode(response)
	})

	// DELETE
	router.HandleFunc("/deleteBrasileiraoPlayer/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		// Configurando o cabeçalho da resposta
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		// Obtendo o ID dos parâmetros da URL
		id := r.PathValue("id")

		// Chama o método Delete no DAO
		brasileiraoPlayerDAO.Delete(id)

		// Retorna uma mensagem de sucesso
		response := map[string]string{
			"message": fmt.Sprintf("Player com ID %s deletado com sucesso!", id),
		}
		w.WriteHeader(http.StatusOK) // 200 - OK
		json.NewEncoder(w).Encode(response)
	})

	//PUT
	// PUT
	router.HandleFunc("/updateBrasileiraoPlayer/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		id := r.PathValue("id")

		var updatedPlayer BrasileiraoPlayer
		err := json.NewDecoder(r.Body).Decode(&updatedPlayer)
		if err != nil {
			http.Error(w, "Erro ao decodificar os dados: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = brasileiraoPlayerDAO.Update(id, &updatedPlayer)
		if err != nil {
			http.Error(w, "Erro ao atualizar o player: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Retornando mensagem de sucesso
		response := map[string]string{
			"message": fmt.Sprintf("Player com ID %s atualizado com sucesso!", id),
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	c := cors.AllowAll()
	r := c.Handler(router)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Ouvindo na porta 8080...")
}
