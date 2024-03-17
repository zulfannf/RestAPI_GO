package controller

import (
	"RestAPI_GO/database"
	"context"
	"net/http"
	

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context){
	orders, err := getOrders(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, map[string]interface{}{
			"success": true,
			"data":    orders,
		})
}

func getOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	rows, err := database.Get().QueryContext(ctx, "select order_id, customer_name, ordered_at from orders")
	if err != nil {
		return nil, err
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