package outlet

import (
	"context"
	"time"
)

type OutletRepository interface {
	FindOneByIDAndUserID(ctx context.Context, id, userId int) (row OutletModel, err error)
}

type OutletModel struct {
	Id           int
	MerchantId   int
	MerchantName string
	OutletName   string
	CreatedAt    time.Time
	CreatedBy    int
	UpdatedAt    time.Time
	UpdatedBy    int
}
