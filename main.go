package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

	Router(db)
}

func Router(db *sql.DB){
	router := gin.Default()
	router.POST("/orders", func(ctx *gin.Context) {
		req := CreateOrderRequest{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success" : false,

			})
			return
		}

		orderId, err := createOrderAndItems(db , ctx , req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success" : false,
				"error" : err.Error(),

			})
			return
		}
	
		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success" : true,
			"order_id" : orderId,
		})
	})

	router.GET("/orders", func(ctx *gin.Context) {

		orders,err := getOrders(db, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success" : false,
				"error" : err.Error(),

			})
			return
		}
		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success" : true,
			"data" : orders,
		})
	})

	router.PUT("/orders/:orderId", func(ctx *gin.Context) {
		req := UpdateOrderRequest{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success" : false,

			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"success" : true,
			"order" : req,
		})
	})

	router.DELETE("/orders/:orderId", func(ctx *gin.Context) {
		orderIdStr := ctx.Param("orderId")
    
   		orderId, err := strconv.Atoi(orderIdStr)
    	if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "invalid orderId",
        })
        return
    }
	err = deleteOrder(db , ctx , orderId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success" : false,
				"error" : err.Error(),

			})
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"success" : true,
		})
	})

	

	if err := router.Run(); err != nil{
		panic(err)
	}
}

type CreateOrderRequest struct {
	OrderedAt string `json:"orderedAt"`
	CustomerName string `json:"customerName"`
	Items []ItemOrderRequest `json:"items"`
}

type ItemOrderRequest struct {
	ItemCode string `json:"itemCode"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}

type UpdateOrderRequest struct {
	OrderedAt string `json:"orderedAt"`
	CustomerName string `json:"customerName"`
	Items []UpdateItemOrderRequest `json:"items"`
}

type UpdateItemOrderRequest struct {
	LineItemId int `json:"lineItemId"`
	ItemCode string `json:"itemCode"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}

type Item struct{
	ItemId int 
	ItemCode string
	Description string
	Quantity int
	OrderId int
}

type Order struct {
	OrderId int `json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt time.Time `json:"ordered_at"`
}


func createOrderAndItems(db *sql.DB, ctx context.Context,req CreateOrderRequest)(int,error){
	orderedAt, _ := time.Parse(time.RFC3339,req.OrderedAt)
		
	
		order := Order{
			CustomerName: req.CustomerName,
			OrderedAt: orderedAt,

		}

		tx, err:= db.Begin()
		if err != nil {
			return 0,err
		}

		defer func ()  {
			if err != nil {
				tx.Rollback()
			}
		}()

		var orderID int64
		err = tx.QueryRowContext(ctx, "insert into orders(customer_name, ordered_at) values($1, $2) returning order_id",order.CustomerName, order.OrderedAt).Scan(&orderID)
		if err != nil {
			return 0,err
		}
		
		
			var items[]Item
			for _, item := range req.Items{
				items = append(items, Item{
					ItemCode: item.ItemCode,
					Description: item.Description,
					Quantity: item.Quantity,
					OrderId: int(orderID),
				})
			}

			for _, item := range items {
				_, err := tx.ExecContext(ctx, "INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4)", item.ItemCode, item.Description, item.Quantity, item.OrderId)
				if err != nil {
					return 0,err
				}
			}

			err = tx.Commit()
			if err != nil {
				return 0,err
			}

			return int(orderID),nil
}

func deleteOrder(db *sql.DB, ctx context.Context,orderID int)(error){
	tx, err:= db.Begin()
		if err != nil {
			return err
		}

		defer func ()  {
			if err != nil {
				tx.Rollback()
			}
		}()

		_,err = tx.ExecContext(ctx, "Delete from items where order_id = $1", orderID)
		if err != nil {
			return err
		}

		res,err := tx.ExecContext(ctx, "Delete from orders where order_id = $1", orderID)
		if err != nil {
			return err
		}

		if rowsAffected, err := res.RowsAffected(); err != nil {
			return err
			} else if rowsAffected < 1{
			return fmt.Errorf("Not Found")
		}
		
		return tx.Commit()
}

func getOrders(db *sql.DB, ctx context.Context)([]Order,error){
	var orders []Order
	rows,err := db.QueryContext(ctx, "select order_id, customer_name, ordered_at from orders")
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}