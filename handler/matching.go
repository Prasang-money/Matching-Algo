package handler

import (
	"container/heap"
	"project/MatchingAlgo/Heap"

	"github.com/gin-gonic/gin"
)

func minimum(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func MatchingAlgo(c *gin.Context) {
	l1 := Queue.Buy.Len()
	l2 := Queue.Sell.Len()

	for minimum(l1, l2) > 0 {
		bord := heap.Pop(&Queue.Buy).(*Heap.Order)
		sord := heap.Pop(&Queue.Buy).(*Heap.Order)
		 
		

	}

}
