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
	

	rows,errquer := db.Exec("INSERT INTO person (nama,alamat) VALUES(?,?)",getNama,getAlamat)

	
	defer db.Close()

	if errquer != nil{
		fmt.Println("Gagal Insert",errquer)
		return false
	}else{
		fmt.Println(rows)
		return true
	}

	
}

func CekData(id string) bool {
	db,err := conn.Connect()

	if err != nil {
		fmt.Println("Error Koneksi", err)
		return false
	}
	removeChar := template.HTMLEscapeString(id)	
	fmt.Println(removeChar)
	var result Orang
	defer db.Close()
	rows := db.QueryRow("SELECT * FROM person WHERE id=?",removeChar).Scan(&result.Id,&result.Nama,&result.Alamat)
	if rows != nil{
		fmt.Println("Failed Query",rows)
		return false
	}else{
		return true
	}
}

func DeleteData(id string) bool {
	db,err := conn.Connect()

	if err != nil {
		fmt.Println("Error Koneksi", err)
		return false
	}
	removeChar := template.HTMLEscapeString(id)	
	fmt.Println(removeChar)
	defer db.Close()

	result, err := db.Exec("DELETE FROM person WHERE id = ?", removeChar)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
    	fmt.Println("Failed Query :",err)
    	return false
	}else{
		fmt.Println("Success Query :",rowsAffected)
		return true
	}

}

func UpdateData(id string,nama string,alamat string) bool {
	db,err := conn.Connect()

	if err != nil {
		fmt.Println("Error Koneksi", err)
		return false
	}
	id_s := template.HTMLEscapeString(id)
	nama_s := template.HTMLEscapeString(nama)	
	alamat_s := template.HTMLEscapeString(alamat)
	
	defer db.Close()

	result, err := db.Exec("UPDATE person SET nama = ?, alamat = ? WHERE id = ?", nama_s,alamat_s,id_s)
	rowsAffected, err := result.RowsAffected()

	if err != nil {
    	fmt.Println("Failed Query :",err)
    	return false
	}else{
		fmt.Println("Success Query :",rowsAffected)
		return true
	}
}

