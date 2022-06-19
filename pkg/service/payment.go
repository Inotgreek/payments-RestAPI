package service

import (
	"github.com/Inotgreek/constanta2/pkg/models"
	"github.com/Inotgreek/constanta2/pkg/repository"
)

type PaymentService struct {
	repo repository.Payment
}

func NewPaymentService(repo repository.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(payment models.Payment) error {
	return s.repo.Create(payment)
}
func (s *PaymentService) ChangeStatus(paymentId int, payment models.Payment) (string, error) {
	return s.repo.ChangeStatus(paymentId, payment)
}

func (s *PaymentService) GetStatusById(paymentId int) (string, error) {
	return s.repo.GetStatusById(paymentId)
}
func (s *PaymentService) GetPaymentsByUserID(userId int) ([]models.Payment, error) {
	return s.repo.GetPaymentsByUserID(userId)
}

func (s *PaymentService) GetPaymentsByEmail(userEmail string) ([]models.Payment, error) {
	return s.repo.GetPaymentsByEmail(userEmail)
}

func (s *PaymentService) CancelPaymentByID(paymentId int) error {
	return s.repo.CancelPaymentByID(paymentId)
}
