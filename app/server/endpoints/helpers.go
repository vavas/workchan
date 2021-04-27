package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/vavas/workchan/app/server/middleware"
	"github.com/vavas/workchan/pkg/db"
)

// DB extracts the db from the gin context.
func DB(c *gin.Context) *db.DB {
	return middleware.GetDB(c)
}
