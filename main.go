package main

import (
	"fmt"
	"os"

	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"github.com/jahanarun/hyperv-vm-shut/hyperv"
)

func main() {
	if len(os.Args) <=1 {
		fmt.Println("No virtual machine specified!")
		fmt.Println("Exiting")
		return
	}
	vmNames := os.Args[1:]
	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/status", func(c *gin.Context) {
		virtualMachines := make([]hyperv.VirtualMachine, len(vmNames))
		fmt.Println(virtualMachines)
		for i, vm := range vmNames {
			virtualMachines[i], _ = hyperv.GetVM(vm)
		}

		fmt.Println(virtualMachines)

		var statusCode int = 503

		if linq.From(virtualMachines).All(func(vm interface{}) bool {
			virtualMachine := vm.(hyperv.VirtualMachine)
			return virtualMachine.State == "off" || virtualMachine.State == "saved" 
		}) {
			statusCode = 200
		}


		c.JSON(statusCode, virtualMachines)
	})

	r.POST("/shutdown", func(c *gin.Context) {
		for _, vm := range vmNames {
			hyperv.SaveVM(vm)
		}
		c.JSON(200, gin.H{
			"message": "shutting down",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}