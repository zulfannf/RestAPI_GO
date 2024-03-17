package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host	= "localhost"
	port	= 5432
	user	= "postgres"
	password = "akiyama23"
	dbname = "latihan2"

)

var (
	db *sql.DB
	err error
)

func init(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)


	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
    	panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}
}

func Get()*sql.DB{
	return db
}