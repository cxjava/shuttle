package api

import (
	"github.com/cxjava/shuttle"
	"github.com/gin-gonic/gin"
)

func DNSCacheList(ctx *gin.Context) {
	ctx.JSON(200, &Response{
		Data: shuttle.DNSCacheList(),
	})
}
func ClearDNSCache(ctx *gin.Context) {
	shuttle.ClearDNSCache()
	ctx.JSON(200, &Response{})
}
