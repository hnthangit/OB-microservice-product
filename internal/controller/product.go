package controller

import (
	"ob-product/internal/entity"
	"ob-product/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.IProductService
}

func (ctl *ProductController) GetProducts(c *gin.Context) {
	var products []*entity.Product = ctl.ProductService.ListProducts()
	c.JSON(200, products)
}

func (ctl *ProductController) GetProductsById(c *gin.Context) {
	id := c.Param("id")
	var products *entity.Product = ctl.ProductService.GetProduct(id)
	c.JSON(200, products)
}

func (ctl *ProductController) GetProductsByName(c *gin.Context) {
	name := c.Query("name")
	var products []*entity.Product = ctl.ProductService.SearchProduct(name)
	c.JSON(200, products)
}
