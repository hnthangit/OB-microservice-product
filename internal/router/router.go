package router

import (
	"ob-product/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {

	var controller = di.GetInstance().Injec()

	client := route.Group("/api")
	{
		client.GET("/products", controller.GetProductsByName)
		client.GET("/products/", controller.GetProducts)
		client.GET("/products/:id", controller.GetProductsById)
	}

}
