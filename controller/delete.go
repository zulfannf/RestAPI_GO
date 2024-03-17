package controller

import (
	"RestAPI_GO/database"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context){
	orderIdStr := ctx.Param("orderId")

		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "invalid orderId",
			})
			return
		}
		err = deleteOrder(ctx, orderId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
		})
}

func deleteOrder(ctx context.Context, orderID int) error {
	tx, err := database.Get().Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "Delete from items where order_id = $1", orderID)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, "Delete from orders where order_id = $1", orderID)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil {
		return err
	} else if rowsAffected < 1 {
		return fmt.Errorf("Not Found")
	}

	return tx.Commit()
}

