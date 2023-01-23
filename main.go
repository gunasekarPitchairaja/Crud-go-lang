package main

import (
	"log"
	"makecrud/controllers"
	"net/http"
	"makecrud/config"
	"github.com/gorilla/mux"
	
)



 

func main() {
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter()

	uc:= controllers.NewBookController(config.DB())
	s.HandleFunc("/book/{id}",uc.GetBook).Methods("GET")
	s.HandleFunc("/book",uc.CreateBook).Methods("POST")
	s.HandleFunc("/books",uc.Books).Methods("GET")
	s.HandleFunc("/book/search/{title}",uc.SearchBook).Methods("GET")
	s.HandleFunc("/book/{id}",uc.DeleteBook).Methods("DELETE")
	s.HandleFunc("/book",uc.UpdateBook).Methods("PUT")
	log.Fatal(http.ListenAndServe("localhost:4300",s))
}

