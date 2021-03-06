package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cxjava/shuttle/assets"
	"github.com/cxjava/shuttle/controller/api"
	"github.com/gin-gonic/gin"
)

func staticHandler(urlPrefix string, fs http.FileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		fileserver.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func index(ctx *gin.Context) {
	b, err := assets.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	ctx.Data(200, "text/html; charset=utf-8", b)
}

func StartController(inter, port string, shutdownSignal chan bool, reloadConfigSignal chan bool, upgradeSignal chan string, level string) {
	if level == "info" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	e := gin.Default()
	e.Use(Cors())
	api.APIRoute(e.Group("/api"), shutdownSignal, reloadConfigSignal, upgradeSignal)

	e.GET("/", index)
	e.GET("/records", index)
	e.GET("/dns-cache", index)
	e.GET("/servers", index)
	e.GET("/mitm", index)
	e.GET("/general", index)
	e.Use(staticHandler("/", assets.HTTP))

	e.Run(inter + ":" + port)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//下面的都是乱添加的-_-~
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
