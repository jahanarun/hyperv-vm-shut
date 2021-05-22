package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jahanarun/hyperv-vm-shut/hyperv"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// stdOut, stdErr, err := posh.Execute("Write-Host \"test\"")
	vm := hyperv.GetVM("dex-adfs")
	// fmt.Printf("\nEnableHyperV:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s", strings.TrimSpace(stdOut), stdErr, err)
	fmt.Println(vm)
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}