package service

import (
	"github.com/Inotgreek/constanta2/pkg/models"
	"github.com/Inotgreek/constanta2/pkg/repository"
)

type Authorization interface {
}

type Payment interface {
	Create(payment models.Payment) error
	ChangeStatus(paymentId int, payment models.Payment) (string, error)
	GetStatusById(paymentId int) (string, error)
	GetPaymentsByUserID(userId int) ([]models.Payment, error)
	GetPaymentsByEmail(userEmail string) ([]models.Payment, error)
	CancelPaymentByID(paymentId int) error
}

type Service struct {
	Authorization
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Payment: NewPaymentService(repos.Payment),
	}
}
