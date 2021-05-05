package main

import (
	"ob-product/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":3000")
}
