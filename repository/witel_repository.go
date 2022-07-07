package repository

import (
	"context"
	"database/sql"
	"github.com/gtxiqbal/sac24/model/domain"
)

type WitelRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.Witel
	FindById(ctx context.Context, db *sql.DB, id string) (domain.Witel, error)
	FindByRegionalId(ctx context.Context, db *sql.DB, regionalId string) ([]domain.Witel, error)
}
