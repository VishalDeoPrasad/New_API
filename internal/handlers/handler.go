package handlers

import (
	"job-application-api/internal/auth"
	"job-application-api/internal/middleware"
	"job-application-api/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API(a auth.TokenAuth, sc service.UserService) *gin.Engine {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		log.Panic("middleware not setup")
		return nil
	}
	h, err := NewHandlerFunc(sc)
	if err != nil {
		log.Panic("handler not setup")
		return nil
	}
	r.Use(middleware.Log(), gin.Recovery())
	r.POST("signup", h.SignUp)
	r.POST("login", h.Login)
	r.GET("/check", m.Authenticate((check)))

	r.POST("/addcompanies", m.Authenticate(h.AddCompany))
	r.GET("/getcompanies", m.Authenticate(h.ViewAllCompanies))
	r.GET("/fetchcompanies/:id", m.Authenticate(h.ViewCompany))

	r.POST("/companies/addjob/:cid", m.Authenticate(h.AddJob))
	r.GET("/fetchAlljobs", m.Authenticate(h.ViewAllJobs))
	r.GET("/job/:id", m.Authenticate(h.ViewJobById))
	r.POST("/apply", m.Authenticate(h.ApplyingProcess))

	return r
}
func check(c *gin.Context) {
	c.JSON(http.StatusOK, "Message :Working... Status Ok")

}
