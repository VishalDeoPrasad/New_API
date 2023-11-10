package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type key string

const TraceIdkey key = "1"

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.NewString()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, TraceIdkey, traceId)
		c.Request = c.Request.WithContext(ctx)
		log.Info().Str("traceId", traceId).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).Msg("request started")
		defer log.Info().Str("Trace Id", traceId).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).
			Int("status Code", c.Writer.Status()).Msg("Request processing completed")
		c.Next()
	}
}
