package Person

import (
	"fmt"
	"crudMvc/db/conn"
)


type Orang struct {
	id int
	nama string
	alamat string
}

func GetallData() []Orang {
	db,err := conn.Connect()

	if err != nil {
		fmt.Println("Error Koneksi", err)
		return nil
	}

	rows,errquer := db.Query("SELECT * FROM person")

	if errquer != nil {
		fmt.Println("Query Gagal",err)
		return nil
	}

	defer rows.Close()

	var results []Orang

	for rows.Next(){
		var each = Orang{}

		var erreach = rows.Scan(&each.id,&each.nama,&each.alamat)

		if erreach != nil{
			fmt.Println(erreach.Error())
			return nil
		}
		results = append(results,each)
	}
	return results


}

