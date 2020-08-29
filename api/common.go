package api

import (
	"ark/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SyncDB(c *gin.Context) {
	err := model.SyncDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "同步数据库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
