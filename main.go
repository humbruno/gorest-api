package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	id    int
}

var data []User = []User{
	{Name: "bruno", Email: "hsbruno1@gmail.com", id: 1},
	{Name: "patricia", Email: "pattccs@gmail.com", id: 2},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler)
	mux.HandleFunc("GET /{id}", handleById)
	mux.HandleFunc("POST /", handleAdd)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("Couldnt start server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	rand := rand.Intn(100)
	fmt.Fprintf(w, "Here is the response! %d", rand)
}

func handleById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "Invalid id, could not be converted to a whole number")
		return
	}

	var name string

	for _, v := range data {
		if v.id == idInt {
			name = v.Name
			break
		}
	}

	if name == "" {
		fmt.Fprintf(w, "Couldnt find user with this id")
		return
	}

	fmt.Fprintf(w, "User with id %s has the name of %s", id, name)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	newData := append(data, User{Name: newUser.Name, Email: newUser.Email, id: len(data) + 1})
	data = newData

	fmt.Fprintf(w, "User with name %s and email %s added successfully", newUser.Name, newUser.Email)
}
