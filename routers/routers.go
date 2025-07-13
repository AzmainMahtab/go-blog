// Package routers
package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/AzmainMahtab/go-blog/controller"
)

func Routes() *gin.Engine {
	r := gin.Default()

	V1 := r.Group("/api/v1")
	{
		V1.POST("/posts", controller.CreatePost)
		V1.GET("/posts", controller.GetPosts)
		// V1.GET("/posts/:id", GetPostByID)
		// V1.PUT("/posts/:id", UpdatePost)
		// V1.DELETE("/posts/:id", DeletePost)
		// V1.PATCH("/posts/:id", PatchPost)
	}
	return r
}
