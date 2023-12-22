package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"time"
)

func Context(c *gin.Context) {
	traceId := c.GetHeader("traceId")
	if len(traceId) == 0 {
		traceId, _ = uuid.GenerateUUID()
	}
	c.Set("traceId", traceId)
	c.Set("startTime", time.Now().UnixNano())
}
