package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/yeejiac/BookExchange_webservice/database"
	"github.com/yeejiac/BookExchange_webservice/internal"
	"github.com/yeejiac/BookExchange_webservice/models"
)

var conn redis.Conn

func SetConnectionObject(rc redis.Conn) {
	conn = rc
}

func Create_BookGroup(w http.ResponseWriter, r *http.Request) { //post
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("parse body error")
		panic(err)
	}
	fmt.Println(string(body))
	var t models.ExchangeGroup
	res := json.Unmarshal(body, &t)
	if res != nil {
		fmt.Println("decode body error")
		return
	}
	t.GroupID = internal.GenerateGroupid(6)
	t.EstablishTime = time.Now()
	t.ExpireTime = t.EstablishTime.Add(time.Hour * time.Duration(24))
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	key := t.GroupID
	value := string(b)
	database.RedisSetTimeout(key, value, 86400, conn) //expired after one day
}

func Modify_BookGroup(w http.ResponseWriter, r *http.Request) { //put
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.ExchangeGroup
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	if database.RedisCheckKey(t.GroupID, conn) {
		key := t.GroupID
		value := string(body)
		database.RedisSet(key, value, conn)
	} else {
		return
	}
}

func Delete_BookGroup(w http.ResponseWriter, r *http.Request) { //delete
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.ExchangeGroup
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	key := t.GroupID
	database.RedisDelete(key, conn)
}

func Check_BookGroup(w http.ResponseWriter, r *http.Request) { //get
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.ExchangeGroup
	// var status models.Status
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	res := database.RedisGet(t.GroupID, conn)
	u, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(u)
}
