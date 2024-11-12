package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := http.NewServeMux()
	userDAO := NewUserDAO()

	router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userDAO.Insert(NewUser(user.Id, user.Name, user.Score))
	})

	c := cors.AllowAll()
	r := c.Handler(router)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}
