package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Player struct {
	gorm.Model
	Name string
}

var player Player
var players []Player

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/players", fetchPlayers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func fetchPlayers(w http.ResponseWriter, r *http.Request) {
	db := GetDBConn()

	db.Find(&players)
	fmt.Println(players)
	profJson, _ := json.Marshal(players)
	fmt.Fprintf(w, string(profJson))
}

func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := ""
	DBNAME := "gorm-example"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	handleRequests()
}
