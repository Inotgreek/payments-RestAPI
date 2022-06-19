package handler

import (
	"github.com/Inotgreek/constanta2/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	var input models.Payment
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Payment.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "payment created",
	})
}

func (h *Handler) GetStatusByID(c *gin.Context) {

	paymentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.services.Payment.GetStatusById(paymentId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status is": status,
	})
}
func (h *Handler) GetPaymentsByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	payments, err := h.services.Payment.GetPaymentsByUserID(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"payments list": payments,
	})
}
func (h *Handler) GetPaymentsByEmail(c *gin.Context) {
	userEmail := c.Param("email")
	payments, err := h.services.Payment.GetPaymentsByEmail(userEmail)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"payments list": payments,
	})
}
func (h *Handler) CancelPaymentByID(c *gin.Context) {
	paymentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Payment.CancelPaymentByID(paymentId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "платеж отменен",
	})

}
func (h *Handler) ChangeStatus(c *gin.Context) {
	paymentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.Payment
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.services.Payment.ChangeStatus(paymentId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if status == "" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Статус платежа не позволяет его изменить",
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"Статус платежа": status,
		})
	}
}
