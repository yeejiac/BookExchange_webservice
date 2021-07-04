package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yeejiac/BookExchange_webservice/internal"
	"github.com/yeejiac/BookExchange_webservice/routes"
)

func main() {
	f, err := os.OpenFile("./log/testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	rc := internal.RedisConnection()
	defer rc.Close()
	routes.SetConnectionObject(rc)

	r := mux.NewRouter()
	// r.HandleFunc("/index", routes.Home).Methods("GET")
	// r.HandleFunc("/login", routes.LoginHandle).Methods("GET")
	// r.HandleFunc("/login", routes.LoginHandle).Methods("POST")
	// r.HandleFunc("/register", routes.ShowRegisterPage).Methods("GET")
	// r.HandleFunc("/validation", routes.Account_Validation).Methods("GET") //確認帳戶已註冊
	r.HandleFunc("/api/bookgroup", routes.Check_BookGroup).Methods("GET")
	r.HandleFunc("/api/bookgroup", routes.Create_BookGroup).Methods("POST")
	r.HandleFunc("/api/bookgroup", routes.Modify_BookGroup).Methods("PUT")
	r.HandleFunc("/api/bookgroup", routes.Delete_BookGroup).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
