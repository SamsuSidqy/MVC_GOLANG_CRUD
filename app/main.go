package main

import (
	"fmt"
	"net/http"
	"log"
	"crudMvc/http/controller/ReadControl"
)

func main(){
	fmt.Println("Listening Port 8000")
	http.HandleFunc("/",ReadControl.ControllerRead)
	log.Fatal(http.ListenAndServe(":8000",nil))
}