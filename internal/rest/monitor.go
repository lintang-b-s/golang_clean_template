package rest

import (
	"context"
	"net/http"

	"lintangbs.org/lintang/template/domain"

	"github.com/gin-gonic/gin"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type MonitorService interface {
	TesDoang(ctx context.Context) (domain.Monitor, error)
}

type MonitorHandler struct {
	Service MonitorService
}

func NewMonitorHandler(rg *gin.RouterGroup, svc MonitorService) {
	handler := &MonitorHandler{
		Service: svc,
	}
	h := rg.Group("/monitors")
	//h.Use(ginkeycloak.Auth(ginkeycloak.AuthCheck(), sbbEndpoint))
	{
		h.GET("/tes", handler.TesDoang)
	}

}

type tesResponse struct {
	Message string `json:"message" `
}

func (m *MonitorHandler) TesDoang(c *gin.Context) {
	tesResult, err := m.Service.TesDoang(c)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.JSON(http.StatusOK, tesResponse{tesResult.Message})
}
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	// logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
