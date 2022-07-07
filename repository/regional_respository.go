package repository

import (
	"context"
	"database/sql"
	"github.com/gtxiqbal/sac24/model/domain"
)

type RegionalRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.Regional
	FindById(ctx context.Context, db *sql.DB, id string) (domain.Regional, error)
	Save(ctx context.Context, db *sql.DB, regional domain.Regional) (domain.Regional, error)
	SaveTx(ctx context.Context, tx *sql.Tx, regional domain.Regional) (domain.Regional, error)
	SaveAll(ctx context.Context, db *sql.DB, regionals []domain.Regional) []domain.Regional
	SaveAllTx(ctx context.Context, tx *sql.Tx, regionals []domain.Regional) []domain.Regional
	Update(ctx context.Context, db *sql.DB, regional domain.Regional) (domain.Regional, error)
	UpdateTx(ctx context.Context, tx *sql.Tx, regional domain.Regional) (domain.Regional, error)
	UpdateAll(ctx context.Context, db *sql.DB, regionals []domain.Regional) []domain.Regional
	UpdateAllTx(ctx context.Context, tx *sql.Tx, regionals []domain.Regional) []domain.Regional
	Delete(ctx context.Context, db *sql.DB, regional domain.Regional)
	DeleteTx(ctx context.Context, tx *sql.Tx, regional domain.Regional)
	DeleteById(ctx context.Context, db *sql.DB, id string)
	DeleteByIdTx(ctx context.Context, tx *sql.Tx, id string)
	DeleteAll(ctx context.Context, db *sql.DB, regionals []domain.Regional)
	DeleteAllTx(ctx context.Context, tx *sql.Tx, regionals []domain.Regional)
	DeleteAllById(ctx context.Context, db *sql.DB, ids []string)
	DeleteAllByIdTx(ctx context.Context, tx *sql.Tx, ids []string)
}
