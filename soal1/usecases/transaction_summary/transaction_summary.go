package transaction_summary

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"majoo/soal1/controllers/middlewares"
	"majoo/soal1/repositories/outlet"
	transactionSummary "majoo/soal1/repositories/transaction_summary"
	"majoo/soal1/usecases"
	"majoo/soal1/utils"
	"time"
)

type transactionSummaryUsecase struct {
	transactionSummaryRepo transactionSummary.TransactionSummaryRepository
	outletRepo             outlet.OutletRepository
}

func NewTransactionSummaryUsecase(transactionSummaryRepo transactionSummary.TransactionSummaryRepository, outletRepo outlet.OutletRepository) TransactionSummaryUsecase {
	return &transactionSummaryUsecase{
		transactionSummaryRepo: transactionSummaryRepo,
		outletRepo:             outletRepo,
	}
}

func (uc *transactionSummaryUsecase) FindSummaryTransactionOutletByMonth(ctx context.Context, month, year string, outletID, page, limit int) (res []TransactionSummaryEntity, p usecases.Pagination, err error) {
	userID := ctx.Value(middlewares.CtxKeyUserID)
	user, ok := userID.(int)
	if !ok {
		err = errors.New("invalid user")
		return
	}

	log.Println("Info outletRepo.FindOneByIDAndUserID", outletID, user)
	outlet, err := uc.outletRepo.FindOneByIDAndUserID(ctx, outletID, user)
	if err != nil {
		log.Println("Err FindOneByIDAndUserID", err)
		if err == sql.ErrNoRows {
			err = errors.New("invalid outlet")
		}
		return
	}

	montInt := utils.StrToInt(month)
	yearInt := utils.StrToInt(year)

	firstDayOfThisMonth := time.Date(yearInt, time.Month(montInt), 1, 0, 0, 0, 0, time.Local)
	endOfThisMonth := time.Date(yearInt, time.Month(montInt+1), 0, 0, 0, 0, 0, time.Local)
	log.Println("firstDayOfThisMonth,endOfThisMonth", firstDayOfThisMonth, endOfThisMonth)
	p = usecases.Pagination{
		CurrentPage: page,
		PerPage:     limit,
		TotalData:   endOfThisMonth.Day(),
	}
	p.Validate() // validate pagination

	startDay := firstDayOfThisMonth.AddDate(0, 0, p.GetOffset())
	endDay := startDay.AddDate(0, 0, p.PerPage-1)
	sDay := startDay.Format("2006-01-02")
	eDay := endDay.Format("2006-01-02")
	log.Println("Info transactionSummaryRepo.FindAll", sDay, eDay, int(user), outletID)
	rows, err := uc.transactionSummaryRepo.FindAll(ctx, sDay, eDay, int(user), outletID)
	if err != nil {
		log.Println("Err FindByMonth", err)
		return
	}
	summaryMapByDate := make(map[string]transactionSummary.TransactionSummaryModel)

	for _, v := range rows {
		summaryMapByDate[v.Date] = v
	}

	i := 0
	for i < p.PerPage {
		date := startDay
		dateStr := date.AddDate(0, 0, i).Format("2006-01-02")

		temp := TransactionSummaryEntity{
			MerchantId:   outlet.MerchantId,
			MerchantName: outlet.MerchantName,
			OutletId:     outlet.Id,
			OutletName:   outlet.OutletName,
			Omzet:        0,
			Date:         dateStr,
		}
		tx, exist := summaryMapByDate[dateStr]
		if exist {
			temp.Omzet = tx.Omzet
		}
		res = append(res, temp)

		i++
	}

	return
}
