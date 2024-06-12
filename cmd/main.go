package main

import (
	"net/http"

	"github.com/Anandhu4456/rms/pkg/controller"
	"github.com/gorilla/mux"
)


func main(){

	r:=mux.NewRouter()
	r.HandleFunc("/signup",controller.Signup).Methods("POST")


	http.ListenAndServe(":8080",r)
}