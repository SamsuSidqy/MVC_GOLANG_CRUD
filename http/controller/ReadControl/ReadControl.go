package ReadControl

import (
	"fmt"
	"net/http"
	"html/template"	
	"crudMvc/db/models/Person"
)


type D map[string]interface{}

func ControllerRead(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method)	
	if r.Method == "GET"{	
		result := Person.GetallData()
		if result != nil {
			fmt.Printf("%T\n",result[0])
		}
		var data = D {
			"title":"Halaman Read",
			"result":Person.GetallData(),	
		}
		t,err := template.ParseFiles("template/read.html")
		if err!= nil { fmt.Println("Ada Error =>",err)}
		t.ExecuteTemplate(w,"read.html",data)
	}else if r.Method == "POST"{
		fmt.Println(r.Method)
	}else{

	}
}