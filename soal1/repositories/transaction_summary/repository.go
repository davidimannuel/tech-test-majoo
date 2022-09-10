package transaction_summary

import (
	"context"
)

type TransactionSummaryRepository interface {
	FindAll(ctx context.Context, startDay, endDay string, userID, outletId int) (rows []TransactionSummaryModel, err error)
}

type TransactionSummaryModel struct {
	MerchantId   int
	MerchantName string
	OutletId     int
	OutletName   string
	Omzet        float64
	Date         string
}
