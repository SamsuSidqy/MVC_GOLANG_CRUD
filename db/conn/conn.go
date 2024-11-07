package conn


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:@/mvc_crud")

	if err != nil {
		panic(err)
		return nil,err
	}

	return db,nil
}