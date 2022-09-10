package outlet

import (
	"context"
	"database/sql"
	"log"
)

type outletPostgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(db *sql.DB) OutletRepository {
	return &outletPostgreRepository{
		db: db,
	}
}

func (repo *outletPostgreRepository) FindOneByIDAndUserID(ctx context.Context, id, userId int) (row OutletModel, err error) {

	sql := `SELECT outlet.id, outlet.merchant_id, merchant.merchant_name,
	outlet.outlet_name, outlet.created_at, outlet.created_by, outlet.updated_at, outlet.updated_by
	FROM Outlets outlet
	LEFT JOIN Merchants merchant ON merchant.id = outlet.merchant_id 
	LEFT JOIN Users "user" ON "user".id = merchant.user_id 
	WHERE outlet.id = $1 AND "user".id = $2
	`
	log.Println("FindOneByIDAndUserID", id, userId)
	err = repo.db.QueryRowContext(ctx, sql, id, userId).Scan(&row.Id, &row.MerchantId, &row.MerchantName, &row.OutletName,
		&row.CreatedAt, &row.CreatedBy, &row.UpdatedAt, &row.UpdatedBy)

	return
}
