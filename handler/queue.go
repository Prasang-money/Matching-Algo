package handler

import (
	"container/heap"
	"fmt"
	"net/http"
	"project/MatchingAlgo/Heap"

	"github.com/gin-gonic/gin"
)

type Orders struct {
	Buy  Heap.MaxHeap
	Sell Heap.MinHeap
}

// type Request struct {
// 	Pps int `json:"pps"`
// 	Qty int `json:"qty"`
// }

var Queue Orders

func GetOrder(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"OrderList": Queue,
	})
}

func Getdata(c *gin.Context) {

	for Queue.Buy.Len() > 0 {
		item := heap.Pop(&Queue.Buy).(*Heap.Order)
		fmt.Println(item)

		c.JSON(http.StatusAccepted, gin.H{
			"Price":    item.Pps,
			"Quantity": item.Qty,
		})

	}

	c.String(http.StatusAccepted, "end of Buy queue")
}
