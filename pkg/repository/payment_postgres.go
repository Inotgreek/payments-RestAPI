package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Inotgreek/constanta2/pkg/models"
	"github.com/jmoiron/sqlx"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}
func (r *PaymentPostgres) Create(payment models.Payment) error {
	time := time.Now()
	status := "new"
	if payment.Sum > 0 {
		status = "НОВЫЙ"
	} else {
		status = "ОШИБКА"
	}
	createPaymentQuery := "INSERT INTO payments(user_id, email, sum, value, create_date, last_change, status) VALUES ($1, $2, $3, $4,$5,$6,$7)"
	_, err := r.db.Exec(createPaymentQuery, payment.UserId, payment.Email, payment.Sum, payment.Value, time, time, status)
	if err != nil {
		return err
	}
	return nil
}

func (r *PaymentPostgres) ChangeStatus(paymentId int, payment models.Payment) (string, error) {
	status, err := r.GetStatusById(paymentId)
	if err != nil {
		return "", err
	}
	var ChangeStatusQuery string
	if status == "НОВЫЙ" {
		ChangeStatusQuery = "UPDATE payments SET status = $1, last_change = $2 WHERE id = $3"
	} else {
		return "", err
	}
	_, err = r.db.Exec(ChangeStatusQuery, payment.Status, time.Now(), paymentId)
	if err != nil {
		return "", err
	}
	return payment.Status, nil
}

func (r *PaymentPostgres) GetStatusById(paymentId int) (string, error) {
	var status string
	getStatusByIdQuery := "SELECT status FROM payments WHERE id=$1"
	err := r.db.Get(&status, getStatusByIdQuery, paymentId)
	if err != nil {
		return "", err
	}
	return status, nil
}

func (r *PaymentPostgres) GetPaymentsByEmail(userEmail string) ([]models.Payment, error) {
	var payments []models.Payment
	getPaymentsByEmailQuery := fmt.Sprintf("SELECT id, user_id, email, sum, value, create_date, last_change, status FROM payments WHERE email=$1")
	err := r.db.Select(&payments, getPaymentsByEmailQuery, userEmail)
	if err != nil {
		return []models.Payment{}, err
	}
	return payments, nil
}

func (r *PaymentPostgres) GetPaymentsByUserID(userId int) ([]models.Payment, error) {
	var payments []models.Payment
	getPaymentsByUserIDQuery := fmt.Sprintf("SELECT id, user_id, user_email, sum, value FROM payments WHERE user_id=$1")
	err := r.db.Select(&payments, getPaymentsByUserIDQuery, userId)
	if err != nil {
		return []models.Payment{}, err
	}
	return payments, nil
}

func (r *PaymentPostgres) CancelPaymentByID(paymentId int) error {
	status, err := r.GetStatusById(paymentId)
	if err != nil {
		err = errors.New("статус платежа не позволяет его отменить")
		return err
	}
	var CancelPaymentByIDQuery string
	if status == "НОВЫЙ" || status == "ОШИБКА" {
		CancelPaymentByIDQuery = "DELETE FROM payments WHERE id=$1"
	} else {
		err = errors.New("статус платежа не позволяет его отменить")
		return err
	}

	_, err = r.db.Exec(CancelPaymentByIDQuery, paymentId)
	if err != nil {
		return err
	}
	return nil
}
