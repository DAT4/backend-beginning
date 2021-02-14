package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"hello-dtu/dao"

	"github.com/gorilla/mux"
)

type Message struct {
	Content string   `json:"content"`
	User    dao.User `json:"user"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/user", addUser).Methods("POST")
	r.HandleFunc("/user", listUsers).Methods("get")
	http.ListenAndServe(":5080", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, `{"message":"hello"}`)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(dao.GetUsers())
}

func addUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user dao.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, `{"error":"error"}`)
		return
	}
	fmt.Println(user)
	dao.CreateUser(user)

	msg := Message{
		Content: "You have now been logged!",
		User:    user,
	}

	res, err := json.Marshal(&msg)
	if err != nil {
		io.WriteString(w, `{"error":"error"}`)
		return
	}

	fmt.Println(string(res))

	w.Write(res)
}
