package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/yeejiac/BookExchange_webservice/database"
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
	r.ParseForm()
	in, header, err := r.FormFile("image")
	if err != nil {
		fmt.Println("parse image error")
	}
	defer in.Close()
	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)

	tmpfile, err := os.Create("./tempfile/" + header.Filename)
	defer tmpfile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully Uploaded File\n")
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
	database.RedisDelete(key, conn)
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

	if database.RedisCheckKey(t.ISBN, conn) {
		key := t.ISBN
		value := string(body)
		database.RedisSet(key, value, conn)
	} else {
		return
	}
}
