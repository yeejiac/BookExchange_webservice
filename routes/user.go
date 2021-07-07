package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yeejiac/BookExchange_webservice/database"
	"github.com/yeejiac/BookExchange_webservice/models"
)

func Create_User(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("parse body error")
		panic(err)
	}
	fmt.Println(string(body))
	var t models.User
	res := json.Unmarshal(body, &t)
	if res != nil {
		fmt.Println("decode body error")
		return
	}
}

func Get_User(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.User
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	res := database.RedisGet(t.Name, conn)
	u, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(u)
}

func Delete_User(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.User
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	key := t.Name
	database.RedisDelete(key, conn)
}

func Modify_User(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.User
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	if database.RedisCheckKey(t.Name, conn) {
		key := t.Name
		value := string(body)
		database.RedisSet(key, value, conn)
	} else {
		return
	}
}
