package controller

import (
	"RestAPI_GO/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context){
	req := CreateOrderRequest{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		orderId, err := createOrderAndItems(ctx, req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success":  true,
			"order_id": orderId,
		})
}



type CreateOrderRequest struct {
	OrderedAt    string             `json:"orderedAt"`
	CustomerName string             `json:"customerName"`
	Items        []ItemOrderRequest `json:"items"`
}

type ItemOrderRequest struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}



type Item struct {
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
}

type Order struct {
	OrderId      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
}

func createOrderAndItems(ctx context.Context, req CreateOrderRequest) (int, error) {
	orderedAt, _ := time.Parse(time.RFC3339, req.OrderedAt)

	order := Order{
		CustomerName: req.CustomerName,
		OrderedAt:    orderedAt,
	}

	tx, err := database.Get().Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var orderID int64
	err = tx.QueryRowContext(ctx, "insert into orders(customer_name, ordered_at) values($1, $2) returning order_id", order.CustomerName, order.OrderedAt).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	var items []Item
	for _, item := range req.Items {
		items = append(items, Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderId:     int(orderID),
		})
	}

	for _, item := range items {
		_, err := tx.ExecContext(ctx, "INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4)", item.ItemCode, item.Description, item.Quantity, item.OrderId)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(orderID), nil
}



