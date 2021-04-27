package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.Writer.Write([]byte("Server is up"))
	c.Status(http.StatusOK)
}

func DBCheck(c *gin.Context) {

	err := DB(c).Master.Ping()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.Writer.Write([]byte("DB alive!"))
	c.Status(http.StatusOK)
}
