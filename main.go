package main

import (
	. "ark/api"
	"ark/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	gin.SetMode(gin.ReleaseMode)
	wg.Add(1)
	db.ConnectDB()
	go Server()
	wg.Wait()
}

func Server() {
	router := gin.Default()
	router.Use(Cors())
	router = Handle(router)
	log.Printf("server start,listening...")
	router.Run(":2020")
	wg.Done()
}

//路由
func Handle(e *gin.Engine) *gin.Engine {

	e.GET("syncdb", SyncDB)
	api := e.Group("api")
	api.GET("/alluser", AllUser)
	api.POST("/saveuser", SaveUser)
	api.POST("/deleteuser", DeleteUser)
	api.POST("/finduser", FindUser)
	api.POST("/queryuser", QueryUser)

	api.POST("/login", Login)

	api.GET("/allrole", AllRole)
	api.POST("/saverole", SaveRole)
	api.POST("/findrole", FindRole)
	api.POST("/queryrole", QueryRole)
	api.POST("/deleterole", DeleteRole)

	api.GET("/allcompany", AllCompanys)
	api.POST("/savecompany", SaveCompany)
	api.POST("/findcompany", FindCompany)
	api.POST("/querycompany", QueryCompany)
	api.POST("/findparentcompany", FindParentCompany)
	api.POST("/findchildrencompany", FindChildrenCompany)
	api.POST("/deletecompany", DeleteCompany)

	api.GET("/alldepartment", AllDeparts)
	api.POST("/savedepartment", SaveDepart)
	api.POST("/finddepartment", FindDepart)
	api.POST("/querydepartment", QueryDepart)
	api.POST("/findparentdepartment", FindParentDepartment)
	api.POST("/findchildrendepartment", FindChildrenDepartment)
	api.POST("/deletedepartment", DeleteDepart)

	return e
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, User,Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
