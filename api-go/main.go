package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ello/sandbox/api-go/ello"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/users", usersHandler).Methods("GET")
	r.HandleFunc("/user/{username}", userHandler).Methods("GET")
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func usersHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Responding to /users request")
	log.Println(req.UserAgent())

	users, _ := ello.Users()

	outgoingJSON, error := json.Marshal(users)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(res, string(outgoingJSON))
}

func userHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Responding to /user request")
	log.Println(req.UserAgent())

	username := mux.Vars(req)["username"]
	if username == "" {
		log.Fatal("Username must be provided")
		http.Error(res, "Username must be provided", http.StatusNotAcceptable)
	}
	user, _ := ello.UserForUsername(username)

	outgoingJSON, error := json.Marshal(user)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(res, string(outgoingJSON))
}
