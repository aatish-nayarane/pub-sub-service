package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InjectLogger(
	logger *zap.Logger,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		reqID :=
			c.GetHeader(
				"X-Request-ID",
			)

		userID :=
			c.GetHeader(
				"X-UserID",
			)

		childLogger :=
			logger.With(
				zap.String(
					"reqid",
					reqID,
				),

				zap.String(
					"userID",
					userID,
				),
			)

		c.Set(
			"logger",
			childLogger,
		)

		c.Next()
	}
}