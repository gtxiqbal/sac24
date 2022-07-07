package repository

import (
	"context"
	"database/sql"
	"github.com/gtxiqbal/sac24/model/domain"
)

type UserRepository interface {
	FindByTelegramId(ctx context.Context, db *sql.DB, telegramId int) ([]domain.User, error)
}
