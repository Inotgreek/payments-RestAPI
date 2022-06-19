package repository

import (
	"github.com/Inotgreek/constanta2/pkg/models"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Authorization
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Payment: NewPaymentPostgres(db),
	}
}
