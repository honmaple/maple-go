/*********************************************************************************
Copyright © 2018 jianglin
File Name: router.go
Author: jianglin
Email: mail@honmaple.com
Created: 2018-02-01 10:52:02 (CST)
Last Update: Wednesday 2018-09-12 14:09:09 (CST)
		 By:
Description:
*********************************************************************************/
package cgin

// cgin mean custom gin

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

// RouterType ..
type RouterType interface {
	GET(c *gin.Context)
	GETITEM(c *gin.Context)
	GETHTML(c *gin.Context)
	POST(c *gin.Context)
	PUT(c *gin.Context)
	DELETE(c *gin.Context)
}

// Init ..
func Init(r *gin.Engine, router RouterType, prefix string) {
	group := r.Group(prefix)
	ins := reflect.TypeOf(router)
	{
		if _, ok := ins.MethodByName("GETHTML"); ok {
			group.GET(".html", router.GETHTML)
		}
		if _, ok := ins.MethodByName("GET"); ok {
			group.GET("", router.GET)
		}
		if _, ok := ins.MethodByName("POST"); ok {
			group.POST("", router.POST)
		}
		if _, ok := ins.MethodByName("GETITEM"); ok {
			group.GET("/:pk", router.GETITEM)
		}
		if _, ok := ins.MethodByName("PUT"); ok {
			group.PUT("/:pk", router.PUT)
		}
		if _, ok := ins.MethodByName("DELETE"); ok {
			group.DELETE("/:pk", router.DELETE)
		}
	}
}
