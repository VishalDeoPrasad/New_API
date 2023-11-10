package handlers

import (
	"encoding/json"
	"job-application-api/internal/auth"
	"job-application-api/internal/middleware"
	"job-application-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (h *handler) ApplyingProcess(c *gin.Context) {

	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdkey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	var jobApply []models.RespondJApplicant

	err := json.NewDecoder(c.Request.Body).Decode(&jobApply)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "some field is missing, please provide all fields",
		})
		return
	}
	appicants, err := h.service.SelectApplications(ctx, jobApply)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "no filter records",
		})
		return
	}

	c.JSON(http.StatusOK, appicants)
}
