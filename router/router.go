package router

import (
	"net/http"

	"codesignal.com/example/gin/models"
	"github.com/gin-gonic/gin"
)

var products = []models.Product{
    {ID: 1, Name: "MacBook Pro", Category: "Laptop", Price: 2499.99},
    {ID: 2, Name: "Dell XPS", Category: "Laptop", Price: 1999.99},
    {ID: 3, Name: "HP Spectre", Category: "Laptop", Price: 1899.95},
    {ID: 4, Name: "Lenovo ThinkPad", Category: "Laptop", Price: 1499.85},
    {ID: 5, Name: "Samsung Galaxy", Category: "Smartphone", Price: 999.99},
    {ID: 6, Name: "iPhone 13", Category: "Smartphone", Price: 1099.99},
    {ID: 7, Name: "Google Pixel", Category: "Smartphone", Price: 899.99},
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	var smartphones = []models.Product{}
	var laptops = []models.Product{}

	for i := 0; i < len(products); i++{
		if products[i].Category == "Laptop"{
			laptops = append(laptops, products[i])
		}else{
			smartphones = append(smartphones, products[i])
		}
	}

	r.GET("/products/laptops", func(c *gin.Context){
		c.JSON(http.StatusOK, laptops)
	})

	r.GET("/products/smartphones", func(c *gin.Context){
		c.JSON(http.StatusOK, smartphones)
	})

	r.GET("/products", func(c *gin.Context){
		c.JSON(http.StatusOK, products)
	})

	return r
}
