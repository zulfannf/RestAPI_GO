package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
    	panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("Successfully Connected to database")

	Router()
}

func Router(){
	router := gin.Default()
	router.POST("/orders", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success" : true,
		})
	})

	router.GET("/orders", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success" : true,
		})
	})

	router.PUT("/orders/:orderId", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success" : true,
		})
	})

	router.DELETE("/orders/:orderId", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success" : true,
		})
	})

	

	if err := router.Run(); err != nil{
		panic(err)
	}
}