package handlers

import (
	"encoding/json"
	"errors"
	"job-application-api/internal/middleware"
	"job-application-api/internal/models"
	"job-application-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type handler struct {
	service service.UserService
}

func NewHandlerFunc(s service.UserService) (NewHandler, error) {
	if s == nil {
		return nil, errors.New("the service cannot be nil")
	}
	return &handler{service: s}, nil
}

type NewHandler interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	ViewJobById(c *gin.Context)
	ViewAllJobs(c *gin.Context)
	ViewJobByCompanyId(c *gin.Context)
	AddJob(c *gin.Context)
	ViewAllCompanies(c *gin.Context)
	AddCompany(c *gin.Context)
	ViewCompany(c *gin.Context)
	ProcessApplication(c *gin.Context)
}

func (h *handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdkey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	var userData models.UserLogin

	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid email and password",
		})
		return
	}
	token, err := h.service.UserLogin(ctx, userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (h *handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	traceid, ok := ctx.Value(middleware.TraceIdkey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	var userData models.UserSignup

	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid username, email and password",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid username, email and password",
		})
		return
	}
	userDetails, err := h.service.UserSignup(ctx, userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userDetails)

}
