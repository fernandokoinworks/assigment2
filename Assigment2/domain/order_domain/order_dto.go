package order_domain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ItemId      int    `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"orderId"`
}

type Order struct {
	OrderId      int       `json:"orderId"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items"`
}

func (m *Order) GetOrderParamId(c *gin.Context) int64 {
	paramId := c.Param("id")
	fmt.Println("Ini adalah order id =>", paramId)
	ID, err := strconv.Atoi(paramId)

	if err != nil {
		return 0
	}

	return int64(ID)
}
