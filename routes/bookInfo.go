package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yeejiac/BookExchange_webservice/internal"
	"github.com/yeejiac/BookExchange_webservice/models"
)

func Create_BookInfo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("parse body error")
		panic(err)
	}
	fmt.Println(string(body))
	var t models.BookInfo
	res := json.Unmarshal(body, &t)
	if res != nil {
		fmt.Println("decode body error")
		return
	}

}

func Get_BookInfo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.BookInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	res := internal.RedisGet(t.ISBN, conn)
	u, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(u)
}

func Delete_BookInfo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.BookInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	key := t.ISBN
	internal.RedisDelete(key, conn)
}

func Modify_BookInfo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.BookInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	if internal.RedisCheckKey(t.ISBN, conn) {
		key := t.ISBN
		value := string(body)
		internal.RedisSet(key, value, conn)
	} else {
		return
	}
}