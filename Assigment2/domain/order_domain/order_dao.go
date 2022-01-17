package order_domain

import (
	"ass2/db"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	queryCreateOrder = `INSERT INTO orders (customer_name, ordered_at) VALUES($1, $2) RETURNING order_id, customer_name`
	queryCreateItem  = `INSERT INTO items (item_code, description, quantity, order_id) VALUES($1, $2, $3, $4) RETURNING item_id, item_code, description, quantity, order_id`
	queryGetorder    = `SELECT * from orders `
	queryGetitem     = `Select * from items where order_id=$1`
	queryOneOrder    = `Select order_id, customer_name, ordered_at from orders where order_id=$1`
	queryUpdateOrder = `Update Orders set customer_name = $1, ordered_at = $2 where order_id=$3 returning order_id`
	queryUpdateitem  = `Update items set item_code = $1, description=  $2, quantity = $3, order_id=$4 where item_id=$5  returning item_code, description, quantity, order_id,item_id`
	queryDeleteOrder = `Delete from orders where order_id = $1`
	queryDeleteItem  = `Delete from items where order_id = $1`
)

var OrderDomain orderDomain = &orderRepo{}

type orderDomain interface {
	CreateOrder(*Order) *Order
	DeleteOrder(int) string
	UpdateOrder(*Order) *Order
	GetOrder() *[]Order
}
type orderRepo struct{}

func (m *orderRepo) CreateOrder(orderReq *Order) *Order {
	db := db.GetDB()
	row := db.QueryRow(queryCreateOrder, orderReq.CustomerName, orderReq.OrderedAt)
	var orderID = 0
	var customer_name = ""
	err := row.Scan(&orderID, &customer_name)
	if err != nil {
		log.Fatal(err)
	}

	items := []Item{}
	for _, itemReq := range orderReq.Items {
		row = db.QueryRow(queryCreateItem, itemReq.ItemCode, itemReq.Description, itemReq.Quantity, orderID)
		var itemRes Item
		err = row.Scan(&itemRes.ItemId, &itemRes.ItemCode, &itemRes.Description, &itemRes.Quantity, &itemRes.OrderId)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, itemRes)
	}

	orderReq.Items = items
	orderReq.OrderId = orderID
	orderReq.CustomerName = customer_name

	return orderReq

}

func (m *orderRepo) GetOrder() *[]Order {
	db := db.GetDB()
	row, err := db.Query(queryGetorder)
	if err != nil {
		log.Fatal("fail to get data")
	}
	var dataOrder []Order

	for row.Next() {
		var ord Order
		if err := row.Scan(&ord.OrderId, &ord.CustomerName, &ord.OrderedAt); err != nil {
			log.Fatal("err")
		}

		row2, err := db.Query(queryGetitem, &ord.OrderId)
		if err != nil {
			log.Fatal("Gagal query item")
		}
		var dataItem []Item
		for row2.Next() {
			var itm Item
			if err := row2.Scan(&itm.ItemId, &itm.ItemCode, &itm.Description, &itm.Quantity, &itm.OrderId); err != nil {
				log.Fatal(err)
			}
			dataItem = append(dataItem, itm)

		}
		ord.Items = append(ord.Items, dataItem...)

		dataOrder = append(dataOrder, ord)

	}
	if err = row.Err(); err != nil {
		log.Fatal("Error ")
	}

	return &dataOrder
}

func (m *orderRepo) UpdateOrder(orderReq *Order) *Order {
	db := db.GetDB()
	fmt.Println(orderReq)
	row := db.QueryRow(queryUpdateOrder, orderReq.CustomerName, orderReq.OrderedAt, orderReq.OrderId)

	var orderdata Order
	err := row.Scan(&orderdata.OrderId)
	if err != nil {
		log.Fatal(err)
	}

	var resdata Order
	row2 := db.QueryRow(queryOneOrder, orderReq.OrderId)

	if err := row2.Scan(&resdata.OrderId, &resdata.CustomerName, &resdata.OrderedAt); err != nil {
		log.Fatal(err)
	}

	items := []Item{}
	for _, itemReq := range orderReq.Items {
		row3 := db.QueryRow(queryUpdateitem, itemReq.ItemCode, itemReq.Description, itemReq.Quantity, orderReq.OrderId, itemReq.ItemId)
		var itemRes Item
		err = row3.Scan(&itemRes.ItemCode, &itemRes.Description, &itemRes.Quantity, &itemRes.OrderId, &itemRes.ItemId)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, itemRes)
	}
	resdata.Items = append(resdata.Items, items...)
	return &resdata
}

func (m *orderRepo) DeleteOrder(data int) string {
	db := db.GetDB()
	var datastring string
	row, err := db.Exec(queryDeleteOrder, data)

	if err == nil {

		count, err := row.RowsAffected()
		if err == nil {
			datastring = fmt.Sprintf("%d, row order affected", count)
		}

		row2, err := db.Exec(queryDeleteItem, data)
		if err == nil {
			count2, err := row2.RowsAffected()
			if err == nil {
				datastring = datastring + fmt.Sprintf("%d, row item affected", count2)
			}

		}

	}
	return datastring

}
