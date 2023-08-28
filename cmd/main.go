package main

import (
	"github.com/airstack/erc/client"
	"github.com/airstack/erc/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	httpClient := http.Client{}
	rpcClient := client.NewClient(httpClient)

	handler := controller.NewController(rpcClient)

	r := gin.Default()
	r.GET("/api/v1/:txhash/check", handler.GetTransactionReceiptsByHash)
	r.Run(":8080")
}
