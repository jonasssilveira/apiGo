package handler

import (
	"github.com/gin-gonic/gin"
	"goAPi/platform/newsfeed"
	"net/http"
)

func NewsfeedGet(feed newsfeed.GetAll) gin.HandlerFunc {

	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}

}
