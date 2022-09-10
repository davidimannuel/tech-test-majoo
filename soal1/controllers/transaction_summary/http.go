package transactionSummary

import (
	"log"
	"net/http"
	"strconv"

	"majoo/soal1/controllers"
	transactionSummary "majoo/soal1/usecases/transaction_summary"
)

type transactionSummaryHTTPController struct {
	transactionSummaryUC transactionSummary.TransactionSummaryUsecase
}

func NewTransactionSummaryHTTPController(transactionSummaryUC transactionSummary.TransactionSummaryUsecase) *transactionSummaryHTTPController {
	return &transactionSummaryHTTPController{
		transactionSummaryUC: transactionSummaryUC,
	}
}

func (ctrl *transactionSummaryHTTPController) GetByMonth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	month := r.URL.Query().Get("month")
	year := r.FormValue("year")
	outletId, _ := strconv.Atoi(r.FormValue("outletId"))
	page, _ := strconv.Atoi(r.FormValue("page"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	log.Println("query param", month, year, outletId)

	result, p, err := ctrl.transactionSummaryUC.FindSummaryTransactionOutletByMonth(ctx, month, year, outletId, page, limit)
	if err != nil {
		controllers.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	res := []TransactionSummaryResponse{}
	for _, v := range result {
		res = append(res, TransactionSummaryResponse{
			MerchantId:   v.MerchantId,
			MerchantName: v.MerchantName,
			OutletId:     v.OutletId,
			OutletName:   v.OutletName,
			Omzet:        v.Omzet,
			Date:         v.Date,
		})
	}

	meta := controllers.Pagination{
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		TotalData:   p.TotalData,
		LastPage:    p.GetLastPage(),
	}

	controllers.WriteResponse(w, http.StatusOK, "", res, meta)
}
