package main

import (
	"fmt"
	"net/http"
	"log"
	"crudMvc/http/controller/ReadControl"
	"crudMvc/http/controller/DeleteControl"
)

func main(){
	fmt.Println("Listening Port 8000")
	http.HandleFunc("/",ReadControl.ControllerRead)
	http.HandleFunc("/delete/",DeleteControl.ControllerDelete)
	log.Fatal(http.ListenAndServe(":8000",nil))
}