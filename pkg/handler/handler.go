package handler

import (
	"github.com/Inotgreek/constanta2/pkg/service"
	gin "github.com/gin-gonic/gin"
)

const (
	paymentURL      = "/payment"
	paymentIDURL    = "/payment/:id"
	userIDURL       = "/id/:id"
	paymentEmailURL = "/email/:email"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		payments := api.Group("/payments")
		{
			payments.POST(paymentURL, h.CreatePayment)
			payments.GET(paymentIDURL, h.GetStatusByID)
			payments.PATCH(paymentIDURL, h.ChangeStatus)
			payments.DELETE(paymentIDURL, h.CancelPaymentByID)
		}
		users := api.Group("/users")
		{
			users.GET(userIDURL, h.GetPaymentsByUserID)
			users.GET(paymentEmailURL, h.GetPaymentsByEmail)
		}
	}

	return router
}
