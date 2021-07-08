package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func GoogleLoginCallBack(w http.ResponseWriter, r *http.Request) {
	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 1         // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store
	go GenerateSession(w, r)
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, r)
		return
	}
	t, _ := template.ParseFiles("/views/index.gtpl")
	log.Println(t.Execute(w, user))
}

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}
