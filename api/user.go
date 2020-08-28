package api

import (
	. "ark/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserList(c *gin.Context) {
	users, err := AllUser()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list":    users,
		"success": true,
	})
}

func SaveUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"success": true,
	})
}

func DeleteUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if user.Id != 0 {
		err = user.Delete()
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

func OneUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	err = user.One()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    user,
	})
}
