package transactionSummary

type TransactionSummaryResponse struct {
	MerchantId   int     `json:"merchantId"`
	MerchantName string  `json:"merchantName"`
	OutletId     int     `json:"outletId"`
	OutletName   string  `json:"outletName"`
	Omzet        float64 `json:"omzet"`
	Date         string  `json:"date"`
}
