package api

import (
	"ark/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllRole(c *gin.Context) {
	objs, err := model.AllRole()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    objs,
		"success": true,
	})
}

func SaveRole(c *gin.Context) {
	var obj model.Role
	err := c.BindJSON(&obj)
	err = obj.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    obj,
		"success": true,
	})
}

func DeleteRole(c *gin.Context) {
	var obj model.Role
	err := c.BindJSON(&obj)
	if obj.Id != 0 {
		err = obj.Delete()
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func FindRole(c *gin.Context) {
	var obj model.Role
	err := c.BindJSON(&obj)
	Objs, err := obj.Find()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Objs,
	})
}

func QueryRole(c *gin.Context) {
	var data QueryValue
	err := c.BindJSON(&data)
	objs, err := model.QueryRole(data.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    objs,
	})
}
