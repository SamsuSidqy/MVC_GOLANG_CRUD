package Person

import (
	"fmt"
	"crudMvc/db/conn"
	"html/template"
)


type Orang struct {
	Id int
	Nama string
	Alamat string
}

func GetallData() []Orang {
	db,err := conn.Connect()

	if err != nil {
		fmt.Println("Error Koneksi", err)
		return nil
	}

	rows,errquer := db.Query("SELECT * FROM person")

	if errquer != nil {
		fmt.Println("Query Gagal",errquer)
		return nil
	}

	defer rows.Close()
	defer db.Close()
	var results []Orang

	for rows.Next(){
		var each Orang

		var erreach = rows.Scan(&each.Id,&each.Nama,&each.Alamat)

		if erreach != nil{
			fmt.Println(erreach.Error())
			return nil
		}

		results = append(results,each)

	}

	return results
}



func InsertData(nama string, alamat string) bool {
	db,err := conn.Connect()
	if err != nil {
		fmt.Println("Gagal Koneksi ===>",err)
		return false
	}
	

	getNama := template.HTMLEscapeString(nama)
	getAlamat := template.HTMLEscapeString(alamat)
	
	rows,errquer := db.Exec("INSERT INTO person VALUES('',?,?)",getNama,getAlamat)

	if errquer != nil{
		fmt.Println("Query Berhasil Di Insert",rows)
		return true
	}else{
		fmt.Println(errquer)

		return false
	}

	
}

