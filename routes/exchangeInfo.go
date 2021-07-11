package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/yeejiac/BookExchange_webservice/database"
	"github.com/yeejiac/BookExchange_webservice/models"
)

func Send_BookExchangeInfo(w http.ResponseWriter, r *http.Request) { //Parse the info from Group's attender
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t models.ExchangeInfo
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

func RandomlyAssigned(groupsize int) int {
	if groupsize == 2 {
		return 1
	} else {
		rand.Seed(time.Now().Unix())
		randNum := rand.Intn(groupsize - 1)
		return randNum
	}
}
