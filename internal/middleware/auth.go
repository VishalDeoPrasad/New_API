package middleware

import (
	"context"
	"errors"
	"job-application-api/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (m *Mid) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := c.Request.Context()
		traceId, ok := ctx.Value(TraceIdkey).(string)
		if !ok {
			log.Error().Msg("trace id not present in the context")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
		authHeader := c.Request.Header.Get("Authorization")

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("provide, Bearer token:")
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}
		ctx = context.WithValue(ctx, auth.Ctxkey, claims)
		c.Request = c.Request.WithContext(ctx)

		next(c)
	}
}
