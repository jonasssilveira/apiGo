package handler

import (
	"goAPi/platform/newsfeed"
	"net/http"
	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json "title"`
	Post  string `json "post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc{
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)
		item := newsfeed.Item{"Jonas", "Silveira"}

		feed.Add(&item)

		c.Status(http.StatusNoContent)

	}
}
