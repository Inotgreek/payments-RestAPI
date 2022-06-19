package models

type Payment struct {
	ID         int64   `json:"id" db:"id"`
	UserId     int64   `json:"userID" db:"user_id"`
	Email      string  `json:"email" db:"email"`
	Sum        float32 `json:"sum" db:"sum"`
	Value      string  `json:"value" db:"value"`
	CreateDate string  `json:"createDate" db:"create_date"`
	LastChange string  `json:"lastChange" db:"last_change"`
	Status     string  `json:"status" db:"status"`
}
