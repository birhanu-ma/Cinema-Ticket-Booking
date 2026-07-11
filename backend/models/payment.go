package models

import (
	"time"

	"github.com/google/uuid"
)
type Payment struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	TicketID  uuid.UUID `json:"ticket_id" db:"ticket_id"`
	TxRef     string    `json:"tx_ref" db:"tx_ref"`
	Amount    float64   `json:"amount" db:"amount"`
	Status    string    `json:"status" db:"status"` // pending/success/failed
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}