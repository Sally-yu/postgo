package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	."ark/model"
)

func UserList(c *gin.Context)  {
	var users []User
	users=append(users, User{"1", "yu", "0001", 20, 0, 15665870081, "yu@qq.com", "aaaaaa", time.Now().UnixNano() / 1e6, time.Now().UnixNano() / 1e6})
	users=append(users, User{"1", "yu", "0001", 20, 0, 15665870081, "yu@qq.com", "aaaaaa", time.Now().UnixNano() / 1e6, time.Now().UnixNano() / 1e6})
	users=append(users, User{"1", "yu", "0001", 20, 0, 15665870081, "yu@qq.com", "aaaaaa", time.Now().UnixNano() / 1e6, time.Now().UnixNano() / 1e6})
	users=append(users, User{"1", "yu", "0001", 20, 0, 15665870081, "yu@qq.com", "aaaaaa", time.Now().UnixNano() / 1e6, time.Now().UnixNano() / 1e6})
	users=append(users, User{"1", "yu", "0001", 20, 0, 15665870081, "yu@qq.com", "aaaaaa", time.Now().UnixNano() / 1e6, time.Now().UnixNano() / 1e6})
	c.JSON(http.StatusOK,gin.H{
		"list":users,
	})
}