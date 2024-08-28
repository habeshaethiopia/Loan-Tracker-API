package domain

import "time"

type Loan struct {
	ID           string    `bson:"_id,omitempty" json:"id"`
	UserID       string    `bson:"user_id" json:"user_id"`
	Amount       float64   `bson:"amount" json:"amount"`
	InterestRate float64   `bson:"interest_rate" json:"interest_rate"`
	Term         int       `bson:"term" json:"term"`
	Status       string    `bson:"status" json:"status"` // e.g., pending, approved, rejected
	AppliedAt    time.Time `bson:"applied_at" json:"applied_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
}
