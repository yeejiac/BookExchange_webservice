package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/yeejiac/BookExchange_webservice/database"
	"github.com/yeejiac/BookExchange_webservice/models"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	auth := session.Values["auth"]
	if auth != nil {
		isAuth, ok := auth.(bool)
		if ok && isAuth {
			t, _ := template.ParseFiles("./views/index.gtpl")
			log.Println(t.Execute(w, nil))
		} else {
			http.Error(w, "unauthorizeed", http.StatusUnauthorized)
			return
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login method:", r.Method) //取得請求的方法
	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// usr := strings.Join(r.Form["Username"], " ")
		// password := strings.Join(r.Form["Password"], " ")
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

		status := models.Status{}
		if LoginVerification(t.Username, t.Password) { // login request pass
			session, err := store.Get(r, "session_token")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			session.Options.MaxAge = 600
			session.Values["auth"] = true
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// http.Redirect(w, r, "/index", http.StatusSeeOther)
			fmt.Println("Login success")
			status.Status = "True"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(status)
			return
		} else {
			fmt.Println("Login failed")
			// http.Redirect(w, r, "/login", http.StatusSeeOther)
			status.Status = "False"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(status)
			return
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/register.gtpl")
		log.Println(t.Execute(w, nil))
	}
}

func LoginVerification(username string, password string) bool {
	res := database.RedisGet(username, conn)
	if res == "" {
		return false
	}
	fmt.Println(res)
	data := []byte(res)
	var t models.User
	err := json.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}

	if t.Password == password {
		fmt.Println(t.Name + " Login success")
		return true
	}
	fmt.Println(t.Name + " Login failed")
	return false
}

func GenerateSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = 600
	session.Values["auth"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
