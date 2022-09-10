package transaction_summary

import (
	"context"
	"majoo/soal1/usecases"
)

type TransactionSummaryUsecase interface {
	FindSummaryTransactionOutletByMonth(ctx context.Context, month, year string, outletId, page, limit int) (res []TransactionSummaryEntity, p usecases.Pagination, err error)
}

type TransactionSummaryEntity struct {
	MerchantId   int
	MerchantName string
	OutletId     int
	OutletName   string
	Omzet        float64
	Date         string
}
