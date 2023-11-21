package mapping

import (
	"github.com/jackc/pgtype"
)

type ConsumedProduct struct {
	ID           int         `db:"id"`
	UserID       int         `db:"user_id"`
	ProductID    int         `db:"product_id"`
	Quantity     int         `db:"quantity"`
	ConsumedDate pgtype.Date `db:"consumed_date"`
}
