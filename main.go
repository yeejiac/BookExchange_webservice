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
	r.HandleFunc("/index", routes.Home).Methods("GET")
	r.HandleFunc("/login", routes.Login).Methods("GET")
	r.HandleFunc("/register", routes.Register).Methods("GET")

	r.HandleFunc("/api/bookgroup", routes.Check_BookGroup).Methods("GET")
	r.HandleFunc("/api/bookgroup", routes.Create_BookGroup).Methods("POST")
	r.HandleFunc("/api/bookgroup", routes.Modify_BookGroup).Methods("PUT")
	r.HandleFunc("/api/bookgroup", routes.Delete_BookGroup).Methods("DELETE")

	r.HandleFunc("/api/user", routes.Get_User).Methods("GET")
	r.HandleFunc("/api/user", routes.Create_User).Methods("POST")
	r.HandleFunc("/api/user", routes.Modify_User).Methods("PUT")
	r.HandleFunc("/api/user", routes.Delete_User).Methods("DELETE")

	r.HandleFunc("/api/BookInfo", routes.Get_BookInfo).Methods("GET")
	r.HandleFunc("/api/BookInfo", routes.Create_BookInfo).Methods("POST")
	r.HandleFunc("/api/BookInfo", routes.Modify_BookInfo).Methods("PUT")
	r.HandleFunc("/api/BookInfo", routes.Delete_BookInfo).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
