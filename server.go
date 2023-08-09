package main
import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)
type Post struct{
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}
func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/post/",handleRequest)
	server.ListenAndServe()
}
func handleRequest(w http.ResponseWriter,r *http.Request){
	var err error
	switch r.Method{
	case "GET":
		err=handleGet(w,r)
	case "POST":
		err=handlePost(w,r)
	case "PUT":
		err=handlePut(w,r)
	case "DELETE":
		err=handleDelete(w,r)
	}
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}
func handleGet(w http.ResponseWriter,r *http.Request)(err error){
	id,err:=
}
func handlePost(w http.ResponseWriter,r *http.Request)(err error){

}
func handlePut(w http.ResponseWriter,r *http.Request)(err error){

}
func handledelete(w http.ResponseWriter,r *http.Request)(err error){

}