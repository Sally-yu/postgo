package api

import (
	"ark/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllCompanys(c *gin.Context) {
	objs, err := model.AllCompany()
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

func SaveCompany(c *gin.Context) {
	var obj model.Company
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

func DeleteCompany(c *gin.Context) {
	var obj model.Company
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

func FindCompany(c *gin.Context) {
	var obj model.Company
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

func FindParentCompany(c *gin.Context) {
	var obj model.Company
	err := c.BindJSON(&obj)
	var company model.Company
	if obj.Parent > 0 {
		company.Id = obj.Parent
	} else {
		temp, err := obj.Find()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "没有符合条件的记录",
			})
			return
		}
		if len(temp) >= 1 {
			company.Id = temp[0].Parent
			if company.Id <= 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "没有符合条件的记录",
				})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "没有符合条件的记录",
			})
			return
		}
	}
	Objs, err := company.Find()
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

func QueryCompany(c *gin.Context) {
	var data QueryValue
	err := c.BindJSON(&data)
	objs, err := model.QueryCompany(data.Value)
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

func FindChildrenCompany(c *gin.Context) {
	var obj model.Company
	err := c.BindJSON(&obj)
	list, err := obj.Find()
	if len(list) < 1 || list[0].Id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "没有符合条件的记录",
		})
		return
	}
	objs, err := list[0].FindChildren()
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
