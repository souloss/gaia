package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_node(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "get node successful!"})
}

func get_all_node(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "get all node successful!"})
}

func delete_node(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "delete node successful!"})
}

func post_node(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "post node successful!"})
}

func SetRouterGroup(instance *gin.Engine) {
	v1 := instance.Group("/v1")
	{
		v1.GET("/node", get_all_node)
		v1.GET("/node/:id", get_node)
		v1.POST("/node", post_node)
		v1.DELETE("/node/:id", delete_node)
	}
}
