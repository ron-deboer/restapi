package routes

import (
	item "RESTapi/controllers/item"
	user "RESTapi/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// start router
func StartGin() {
	router := gin.Default()
	apiUser := router.Group("/api/users")
	{
		apiUser.GET("", user.GetAll)
		apiUser.POST("", user.Create)
		apiUser.GET("/:id", user.Get)
		apiUser.PUT("/:id", user.Update)
		apiUser.DELETE("/:id", user.Delete)
	}
	apiItem := router.Group("/api/items")
	{
		apiItem.GET("", item.GetAll)
		apiItem.POST("", item.Create)
		apiItem.GET("/:id", item.Get)
		apiItem.PUT("/:id", item.Update)
		apiItem.DELETE("/:id", item.Delete)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8001")
}
