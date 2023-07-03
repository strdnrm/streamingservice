package routes

import (
	"streamingservice/pkg/controllers"
	"streamingservice/pkg/store"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(store store.Store) *gin.Engine {
	router := gin.Default()

	controllers.SetUpStore(store)

	api := router.Group("/api")
	{
		api.GET("/order/list", controllers.HandleGetOrderList())
		api.GET("/order/:uid", controllers.HandleGetOrder())
	}

	return router
}
