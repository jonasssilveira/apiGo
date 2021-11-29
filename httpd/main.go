package main

import (
	"goAPi/httpd/handler"
	"github.com/gin-gonic/gin"
	"goAPi/platform/newsfeed"
)

func main(){

	feed := newsfeed.New();

	r := gin.Default()
	r.GET("/getall", handler.NewsfeedGet(feed))
	r.POST("/addy", handler.NewsfeedPost(feed))
	r.GET("/ping", handler.PingGet())
	r.Run()


}
