package transaction_summary

import (
	"context"
	"database/sql"
	"log"
)

type transactionSummaryPostgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(db *sql.DB) TransactionSummaryRepository {
	return &transactionSummaryPostgreRepository{
		db: db,
	}
}

func (repo *transactionSummaryPostgreRepository) FindAll(ctx context.Context, startDay, endDay string, userID, outletId int) (rows []TransactionSummaryModel, err error) {
	sql := `SELECT 
	tr.merchant_id, merchant.merchant_name, 
	tr.outlet_id, outlet.outlet_name, 
	sum(tr.bill_total) omzet, to_char(tr.created_at ,'YYYY-MM-DD') "date" 
	FROM transactions tr
	LEFT JOIN outlets outlet ON outlet.id =  tr.outlet_id
	LEFT JOIN merchants merchant ON merchant.id = outlet.merchant_id 
	LEFT JOIN users "user" ON "user".id = merchant.user_id 
	WHERE DATE_TRUNC('day',tr.created_at AT TIME ZONE 'Asia/Jakarta') >= $1 AND DATE_TRUNC('day',tr.created_at AT TIME ZONE 'Asia/Jakarta') <= $2 
	AND "user".id = $3 AND outlet.id = $4
	GROUP BY tr.merchant_id, merchant.id, tr.outlet_id, outlet.id, "date" 
	ORDER BY "date" ASC 
	`
	log.Println("info repo FindAll ", sql, startDay, endDay, userID, outletId)
	sqlRows, err := repo.db.QueryContext(ctx, sql, startDay, endDay, userID, outletId)
	if err != nil {
		return
	}
	defer sqlRows.Close()

	for sqlRows.Next() {
		temp := TransactionSummaryModel{}
		err = sqlRows.Scan(
			&temp.MerchantId, &temp.MerchantName, &temp.OutletId, &temp.OutletName, &temp.Omzet, &temp.Date,
		)
		if err != nil {
			return
		}
		rows = append(rows, temp)
	}

	return
}
