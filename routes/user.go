package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"

	"github.com/yeejiac/BookExchange_webservice/database"
	"github.com/yeejiac/BookExchange_webservice/internal"
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
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println("decode body error")
		return
	}

	if !CheckUserDataApply(t) {
		var status models.Status
		status.Status = "email wrong"
		b, err := json.Marshal(status)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	var status models.Status
	if database.RedisCheckKey(t.Name, conn) {
		fmt.Println("Already Exist")
		status.Status = "Failed"
	} else {
		go internal.SendRegisterMail(t)
		status.Status = "Success"
		key := t.Name
		value := string(body)
		database.RedisSet(key, value, conn)
	}
	b, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
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

func CheckUserDataApply(userinfo models.User) bool {
	_, err := mail.ParseAddress(userinfo.Email)
	return err == nil
}
