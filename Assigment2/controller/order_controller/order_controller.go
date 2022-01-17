package order_controller

import (
	"ass2/domain/order_domain"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errMsg": err.Error(),
		})
		return
	}

	fmt.Println(data)
	var orderedAtStr string
	_ = orderedAtStr
	if val, ok := data["orderedAt"].(string); ok {
		orderedAtStr = val
	}

	y, err := time.Parse("2006-01-02", orderedAtStr)
	if err != nil {
		fmt.Println(err)
	}

	items := []order_domain.Item{}
	for _, v := range data["items"].([]interface{}) {
		item := v.(map[string]interface{})
		itemReq := []order_domain.Item{
			{
				ItemCode:    item["itemCode"].(string),
				Description: item["description"].(string),
				Quantity:    int(item["quantity"].(float64)),
			},
		}
		items = append(items, itemReq...)
	}
	req := order_domain.Order{
		CustomerName: data["customerName"].(string),
		OrderedAt:    y,
		Items:        items,
	}
	_ = req

	res := order_domain.OrderDomain.CreateOrder(&req)
	c.JSON(201, res)

}

func GetOrder(c *gin.Context) {
	res := order_domain.OrderDomain.GetOrder()
	c.JSON(201, res)
}

func UpdateOrder(c *gin.Context) {

	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errMsg": err.Error(),
		})
		return
	}
	fmt.Println(data)

	var orderedAtStr string
	_ = orderedAtStr
	if val, ok := data["orderedAt"].(string); ok {
		orderedAtStr = val
	}

	y, err := time.Parse("2006-01-02", orderedAtStr)
	if err != nil {
		fmt.Println(err)
	}

	items := []order_domain.Item{}
	for _, v := range data["items"].([]interface{}) {
		item := v.(map[string]interface{})
		itemReq := []order_domain.Item{
			{
				ItemId:      int(item["itemId"].(float64)),
				ItemCode:    item["itemCode"].(string),
				Description: item["description"].(string),
				Quantity:    int(item["quantity"].(float64)),
			},
		}
		items = append(items, itemReq...)
	}
	req := order_domain.Order{
		CustomerName: data["customerName"].(string),
		OrderedAt:    y,
		Items:        items,
	}
	_ = req

	orderid := req.GetOrderParamId(c)

	req.OrderId = int(orderid)

	fmt.Println("item : ", req.Items)
	res := order_domain.OrderDomain.UpdateOrder(&req)
	c.JSON(201, res)

}

func DeleteOrder(c *gin.Context) {
	var order order_domain.Order

	orderid := order.GetOrderParamId(c)
	fmt.Println(orderid)
	res := order_domain.OrderDomain.DeleteOrder(int(orderid))
	fmt.Println(res)
	c.String(201, res)

}
