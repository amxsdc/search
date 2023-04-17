package routes

import (
	"github.com/gin-gonic/gin"
	"search/src/search"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	searchGroup := r.Group("/api")
	{
		searchGroup.GET("/suggest", search.Search)
	}

	return r

}
