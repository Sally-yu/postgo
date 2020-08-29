package api

import (
	"ark/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllUser(c *gin.Context) {
	users, err := model.AllUser()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"success": true,
	})
}

func FindUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	Objs, err := user.Find()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    Objs,
		"success": true,
	})
}

func SaveUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"success": true,
	})
}

func DeleteUser(c *gin.Context) {
	var user model.User
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

type QueryValue struct {
	Value string `json:"value"`
}

func QueryUser(c *gin.Context) {
	var data QueryValue
	err := c.BindJSON(&data)
	objs, err := model.QueryUser(data.Value)
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

func Login(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	has, Obj, _ := user.Has()
	if !(has && Obj.Code == user.Code && Obj.Pwd == user.Pwd) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	user.Pwd = ""
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Obj,
	})
}
