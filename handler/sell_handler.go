package handler

import (
	"container/heap"
	"fmt"
	"net/http"
	"project/MatchingAlgo/Heap"

	"github.com/gin-gonic/gin"
)

func (token *Orders) AddSellorder(ord Heap.Order) {

	heap.Push(&token.Sell, &ord)
}
func RegisterSell(c *gin.Context) {
	var ord Heap.Order

	err := c.BindJSON(&ord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to unmarshal request data",
		})
		return
	}
	Queue.AddSellorder(ord)
	fmt.Println(Queue.Sell[0])
	c.JSON(http.StatusOK, gin.H{
		"message": "Your request for selling share is added into queue",
	})

}
