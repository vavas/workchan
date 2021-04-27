package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vavas/workchan/pkg/db"
)

const dbCtxKey = "dbConns"

// InjectDBConnections creates a gin middleware that adds the
// database connections to all requests contexts
func InjectDBConnections(dbConns *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(dbCtxKey, dbConns)
		c.Next()
	}
}

// GetDB extracts the database connections from the request context.
// if it's not present, it returns nil
func GetDB(c *gin.Context) *db.DB {
	val, exists := c.Get(dbCtxKey)
	if !exists {
		return nil
	}
	dbConns, ok := val.(*db.DB)
	if !ok {
		return nil
	}
	return dbConns
}
