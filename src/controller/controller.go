package controller

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"sample/src/models"
)

var (
	ProductCatalogue = map[string]*models.Product{
		"product1": {ID: "product1", Availability: true, Price: 10, Category: "Regular"},
		"product2": {ID: "product2", Availability: false, Price: 20, Category: "Premium"},
		"product3": {ID: "product3", Availability: true, Price: 5, Category: "Budget"},
	}

	Orders      = map[string]*models.Order{}
	OrdersMutex sync.Mutex
)

func GetCatalogue(c *gin.Context) {
	productList := []*models.Product{}
	for _, product := range ProductCatalogue {
		productList = append(productList, product)
	}

	c.JSON(http.StatusOK, productList)
}

func PlaceOrder(c *gin.Context) {
	var order models.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, ok := ProductCatalogue[order.ProductID]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if !product.Availability {
		c.JSON(http.StatusForbidden, gin.H{"error": "Product not available"})
		return
	}

	if order.Quantity > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity limit exceeded"})
		return
	}

	orderValue := product.Price * float64(order.Quantity)
	if product.Category == "Premium" {
		orderValue *= 0.9
	}

	order.OrderValue = orderValue
	order.OrderStatus = "Placed"
	order.Product = product

	product.Availability = false

	OrdersMutex.Lock()
	Orders[order.ID] = &order
	OrdersMutex.Unlock()

	c.JSON(http.StatusCreated, order)
}

func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	OrdersMutex.Lock()
	order, ok := Orders[orderID]
	if !ok {
		OrdersMutex.Unlock()
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var updateData models.UpdateData

	err := c.ShouldBindJSON(&updateData)
	if err != nil {
		OrdersMutex.Unlock()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateData.OrderStatus == "Dispatched" {
		order.DispatchDate = updateData.DispatchDate
	}

	order.OrderStatus = updateData.OrderStatus
	OrdersMutex.Unlock()

	c.JSON(http.StatusOK, order)
}
