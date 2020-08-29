package api

import (
	"ark/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllDeparts(c *gin.Context) {
	objs, err := model.AllDepart()
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

func SaveDepart(c *gin.Context) {
	var obj model.Department
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

func DeleteDepart(c *gin.Context) {
	var obj model.Department
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

func FindDepart(c *gin.Context) {
	var obj model.Department
	err := c.BindJSON(&obj)
	Obj, err := obj.Find()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Obj,
	})
}

func QueryDepart(c *gin.Context) {
	var data QueryValue
	err := c.BindJSON(&data)
	objs, err := model.QueryDepart(data.Value)
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

func FindParentDepartment(c *gin.Context) {
	var obj model.Department
	err := c.BindJSON(&obj)
	var department model.Department
	if obj.Parent > 0 {
		department.Id = obj.Parent
	} else {
		temp, err := obj.Find()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "没有符合条件的记录",
			})
			return
		}
		if len(temp) >= 1 {
			department.Id = temp[0].Parent
			if department.Id <= 0 {
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
	Objs, err := department.Find()
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

func FindChildrenDepartment(c *gin.Context) {
	var obj model.Department
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
