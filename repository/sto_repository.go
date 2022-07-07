package repository

import (
	"context"
	"database/sql"
	"github.com/gtxiqbal/sac24/model/domain"
)

type StoRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.Sto
	FindById(ctx context.Context, db *sql.DB, id string) (domain.Sto, error)
	FindByWitelId(ctx context.Context, db *sql.DB, witelId string) ([]domain.Sto, error)
	FindByRegionalId(ctx context.Context, db *sql.DB, regionalId string) ([]domain.Sto, error)
}
