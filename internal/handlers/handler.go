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
	r.POST("/api/register", h.SignUp)
	r.POST("/api/login", h.Login)
	r.GET("/check", m.Authenticate((check)))
	
	
	r.POST("/api/companies", m.Authenticate(h.AddCompany))
	r.GET("/api/companies", m.Authenticate(h.ViewAllCompanies))
	r.GET("/api/companies/:id", m.Authenticate(h.ViewCompany))
	

	r.POST("/api/companies/:cid/jobs", m.Authenticate(h.AddJob))
	//r.GET("/api/companies/:cid/jobs", m.Authenticate(h.ViewJobByCompanyId))
	r.GET("/api/jobs", m.Authenticate(h.ViewAllJobs))
	r.GET("/api/jobs/:id", m.Authenticate(h.ViewJobById))
	r.POST("/api/process", m.Authenticate(h.ProcessApplication))

	return r
}
func check(c *gin.Context) {
	c.JSON(http.StatusOK, "Msg :ok")

}
