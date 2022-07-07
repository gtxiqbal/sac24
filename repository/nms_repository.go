package repository

import (
	"context"
	"database/sql"
	"github.com/gtxiqbal/sac24/model/domain"
)

type NmsRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.Nms
	FindById(ctx context.Context, db *sql.DB, id string) (domain.Nms, error)
	FindByNama(ctx context.Context, db *sql.DB, nama string) (domain.Nms, error)
	FindByVendor(ctx context.Context, db *sql.DB, vendor string) ([]domain.Nms, error)
	FindByIpServer(ctx context.Context, db *sql.DB, ipServer string) (domain.Nms, error)
}
