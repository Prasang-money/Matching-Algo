package main

import (
	"project/MatchingAlgo/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// h := Heap.MaxHeap{}
	// heap.Init(&h)
	// item := Heap.Order{Pps: 25, Qty: 6}
	// heap.Push(&h, item)
	// fmt.Printf("minimum: %d\n", h[0])

	// for h.Len() > 0 {
	// 	fmt.Printf("%d\n ", heap.Pop(&h))
	// }
	// fmt.Print(h.Len())
	router := gin.Default()
	//queue := Orders{}
	//queue.Buy.Push(Heap.Order{Pps: 2, Qty: 5})
	router.GET("/", handler.Getdata)
	router.GET("getorders", handler.GetOrder)
	router.POST("/request/buy", handler.RegisterBuy)
	router.POST("/request/sell", handler.RegisterSell)

	router.Run()

}
