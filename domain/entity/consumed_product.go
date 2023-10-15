package entity

import "github.com/jackc/pgtype"

type ConsumedProduct struct {
	Product       Product     `json:"product"`
	Quantity      int         `json:"quantity"`
	Consumed_Date pgtype.Date `json:"consumed_date"`
}
