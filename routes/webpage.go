package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/yeejiac/BookExchange_webservice/database"
	"github.com/yeejiac/BookExchange_webservice/models"
)

// custom claims
type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

// jwt secret key
var jwtSecret = []byte("secret")

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
		usr := strings.Join(r.Form["Username"], " ")
		password := strings.Join(r.Form["Password"], " ")
		if LoginVerification(usr, password) { // login request pass
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
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			fmt.Println("success")
		} else {
			fmt.Println("Login failed")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
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
		generateToken(username)
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

func generateToken(username string) string {
	now := time.Now()
	jwtId := username + strconv.FormatInt(now.Unix(), 10)
	role := "Member"

	// set claims and sign
	claims := Claims{
		Account: username,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			Audience:  username,
			ExpiresAt: now.Add(20 * time.Second).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "yeejiac",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   username,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString(jwtSecret)
	return token
}
