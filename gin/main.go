package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/net/Interfaces", func(c *gin.Context) {
		ifaces, err := net.Interfaces()
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
		} else {
			fmt.Print(ifaces)
			c.JSON(200, ifaces)
		}
	})

	r.GET("/net/Interface/:index/Addrs", func(c *gin.Context) {
		indexParam := c.Param("index")
		i64, err := strconv.ParseInt(indexParam, 10, 64)
		index := int(i64)
		iface, err := net.InterfaceByIndex(index)
		addrs, err := iface.Addrs()
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
		} else {
			fmt.Print(addrs)
			c.JSON(200, addrs)
		}
	})

	r.GET("/net/InterfaceAddrs", func(c *gin.Context) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
		} else {
			fmt.Print(addrs)
			c.JSON(200, addrs)
		}
	})

	r.Run(":9080") // listen and serve on 0.0.0.0:9080
}
