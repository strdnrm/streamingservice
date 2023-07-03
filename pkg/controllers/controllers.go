package controllers

import (
	"context"
	"fmt"
	"net/http"
	"streamingservice/pkg/store"

	"github.com/gin-gonic/gin"
)

var (
	st store.Store
)

func SetUpStore(s store.Store) {
	st = s
}

func HandleGetOrderList() gin.HandlerFunc {
	return func(c *gin.Context) {

		orders, err := st.Order().GetAll(context.Background())
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(orders)
		c.JSON(http.StatusOK, orders)
	}
}

func HandleGetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		order, err := st.Order().GetById(context.Background(), c.Param("uid"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}
