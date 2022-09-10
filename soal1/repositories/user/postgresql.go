package user

import (
	"context"
	"database/sql"
)

type userPostgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(db *sql.DB) UserRepository {
	return &userPostgreRepository{
		db: db,
	}
}

func (repo *userPostgreRepository) FindOneByUserName(ctx context.Context, userName string) (row UserModel, err error) {

	sql := `SELECT id, name, user_name, password, created_at, created_by, updated_at, updated_by 
	FROM users WHERE user_name = $1`

	err = repo.db.QueryRowContext(ctx, sql, userName).Scan(&row.Id, &row.Name, &row.UserName, &row.Password,
		&row.CreatedAt, &row.CreatedBy, &row.UpdatedAt, &row.UpdatedBy)

	return
}
