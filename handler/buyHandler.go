package handler

import (
	"container/heap"
	"fmt"
	"net/http"
	"project/MatchingAlgo/Heap"

	"github.com/gin-gonic/gin"
)

func (token *Orders) AddBuyOrder(ord Heap.Order) {

	heap.Push(&token.Buy, &ord)
}
func RegisterBuy(c *gin.Context) {
	var ord Heap.Order

	err := c.BindJSON(&ord)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to unmarshal request data",
		})
		return
	}
	Queue.AddBuyOrder(ord)
	fmt.Println(Queue.Buy[0])
	c.JSON(http.StatusOK, gin.H{
		"message": "Your request for buying  share is added into queue",
	})
}
