/*********************************************************************************
Copyright © 2018 jianglin
File Name: router.go
Author: jianglin
Email: mail@honmaple.com
Created: 2018-02-01 10:52:02 (CST)
Last Update: Wednesday 2018-09-12 13:59:19 (CST)
		 By:
Description:
*********************************************************************************/
package cgin

// cgin mean custom gin

import (
	"github.com/gin-gonic/gin"
)

var err error

// RouterType ..
type RouterType interface {
	GET(c *gin.Context)
	GETITEM(c *gin.Context)
	POST(c *gin.Context)
	PUT(c *gin.Context)
	DELETE(c *gin.Context)
}

// RouterAPI ..
type RouterAPI struct {
	Router RouterType
	Prefix string
}

// Router ..
type Router struct{}

// GET ..
func (r *Router) GET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// POST ..
func (r *Router) POST(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// GETITEM ..
func (r *Router) GETITEM(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// PUT ..
func (r *Router) PUT(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// DELETE ..
func (r *Router) DELETE(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Init ..
func Init(r *gin.Engine, apis []RouterAPI) {
	for _, api := range apis {
		group := r.Group(api.Prefix)
		{
			group.GET("", api.Router.GET)
			group.POST("", api.Router.POST)
			group.PUT("", api.Router.PUT)
			group.DELETE("", api.Router.DELETE)
			group.GET("/:pk", api.Router.GETITEM)
		}
	}
}
