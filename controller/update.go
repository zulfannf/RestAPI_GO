package controller

import (
	"RestAPI_GO/database"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateOrderRequest struct {
	OrderedAt    string                   `json:"orderedAt"`
	CustomerName string                   `json:"customerName"`
	Items        []UpdateItemOrderRequest `json:"items"`
}

type UpdateItemOrderRequest struct {
	LineItemId  int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func Update(ctx *gin.Context){
	req := UpdateOrderRequest{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		orderIdStr := ctx.Param("orderId")

		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "invalid orderId",
			})
			return
		}
		err = UpdateOrderAndItems(ctx, req, orderId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"order":   req,
		})
}

func UpdateOrderAndItems(ctx context.Context, req UpdateOrderRequest, orderid int) error {
	tx, err := database.Get().Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	res, err := tx.ExecContext(ctx, "update orders set customer_name = $1, ordered_at = $2 where order_id = $3", req.CustomerName, req.OrderedAt, orderid)
	if err != nil {
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return err
	} else if rowsAffected < 1 {
		return fmt.Errorf("Not Found")
	}

	for _, item := range req.Items {
		res, err := tx.ExecContext(ctx,
			"update items set item_code = $1, description = $2, quantity = $3 where item_id = $4 and order_id = $5",
			item.ItemCode, item.Description, item.Quantity, item.LineItemId, orderid)
		if err != nil {
			return err
		}
		if rowsAffected, err := res.RowsAffected(); err != nil {
			return err
		} else if rowsAffected < 1 {
			return fmt.Errorf("Not Found")
		}
	}
	return tx.Commit()
}