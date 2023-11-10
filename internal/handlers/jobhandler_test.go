package handlers

import (
	"context"
	"job-application-api/internal/auth"
	"job-application-api/internal/middleware"
	"job-application-api/internal/service"
	"job-application-api/internal/service/mockmodels"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func Test_handler_AddJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "error in validating json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
					"jobname" : "Developer",
					"minNoticePeriod" : 3,
					"maxNoticePeriod" : 30,
					"location" : [
						1,
						2
						],
					"technologyStack":[
						1,
						2
					],
					"description":"A developer ",
					"minExperience":3,
					"maxExperience":6,
					"qualifications":[
						1,
						2
					],
					"shifts":[
						1,
						2
					],
					"jobtype":"part-time
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := mockmodels.NewMockUserService(mc)

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"provide valid details"}`,
		},
		// {
		// 	name: "Failure",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
		// 			"jobname" : "Developer",
		// 			"minNoticePeriod" : 3,
		// 			"maxNoticePeriod" : 30,
		// 			"location" : [
		// 				1,
		// 				2
		// 				],
		// 			"technologyStack":[
		// 				1,
		// 				2
		// 			],
		// 			"description":"A developer ",
		// 			"minExperience":3,
		// 			"maxExperience":6,
		// 			"qualifications":[
		// 				1,
		// 				2
		// 			],
		// 			"shifts":[
		// 				1,
		// 				2
		// 			],
		// 			"jobtype":"part-time"
		// 		}`))

		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
		// 		ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

		// 		c.Request = httpRequest

		// 		mc := gomock.NewController(t)
		// 		ms := mockmodels.NewMockUserService(mc)

		// 		ms.EXPECT().AddJobDetails(c.Request.Context(), gomock.Any()).Return(models.ResponseJob{}, errors.New("test error"))

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `{"error":"test error"}`,
		// },
		// {
		// 	name: "success",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
		// 			"jobname" : "Developer",
		// 			"minNoticePeriod" : 3,
		// 			"maxNoticePeriod" : 30,
		// 			"location" : [
		// 				1,
		// 				2
		// 				],
		// 			"technologyStack":[
		// 				1,
		// 				2
		// 			],
		// 			"description":"A developer ",
		// 			"minExperience":3,
		// 			"maxExperience":6,
		// 			"qualifications":[
		// 				1,
		// 				2
		// 			],
		// 			"shifts":[
		// 				1,
		// 				2
		// 			],
		// 			"jobtype":"part-time"
		// 		}`))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
		// 		ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
		// 		c.Request = httpRequest

		// 		mc := gomock.NewController(t)
		// 		ms := mockmodels.NewMockUserService(mc)
		// 		ms.EXPECT().AddJobDetails(gomock.Any(), gomock.Any()).Return(models.ResponseJob{
		// 			Id: 1,
		// 		}, nil)

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   `{"id":1}`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := handler{
				service: ms,
			}
			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
