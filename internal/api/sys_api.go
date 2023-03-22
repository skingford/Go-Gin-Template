/*
 * @Author: kingford
 * @Date: 2023-03-21 23:24:35
 * @LastEditTime: 2023-03-22 23:31:29
 */
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
)

type SysApi struct {
	api.Api
}

func (e *SysApi) Find(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Find Hello, world!",
	})
}

func (e *SysApi) First(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "First Hello, world!",
	})
}

func (e *SysApi) Create(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}

	c.BindJSON(&request)

	c.JSON(200, gin.H{
		"message": "Hello, " + request.Name + "!",
	})
}

func (e *SysApi) Update(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}

	c.BindJSON(&request)

	c.JSON(200, gin.H{
		"message": "Hello, " + request.Name + "!",
	})
}
