package response

import "github.com/jackc/pgtype"

type ConsumedProduct struct {
	Product      Product     `json:"product"`
	Quantity     int         `json:"quantity"`
	ConsumedDate pgtype.Date `json:"consumedDate"`
}
