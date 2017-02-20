package db


import (
	//"database/sql"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"

)

func GetDBConnection() (*sqlx.DB,error) {
	//db,err:=sql.Open("postgres","dbname=mydb  sslmode=disable")
	db,err:=sqlx.Connect("postgres","dbname=mydb  sslmode=disable")
	return db,err
}
