package repository

import (
	"context"
	"database/sql"
	"github.com/gtxiqbal/sac24/model/domain"
)

type GponRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.Gpon
	FindById(ctx context.Context, db *sql.DB, id string) (domain.Gpon, error)
	FindByHostname(ctx context.Context, db *sql.DB, hostname string) (domain.Gpon, error)
	FindByIpAddress(ctx context.Context, db *sql.DB, ipAddress string) (domain.Gpon, error)
	FindByIpAddressIn(ctx context.Context, db *sql.DB, ipAddresses []string) ([]domain.Gpon, error)
	FindByIpAddressInAndProtocol(ctx context.Context, db *sql.DB, ipAddresses []string, protocol string) ([]domain.Gpon, error)
	FindByNmsId(ctx context.Context, db *sql.DB, nmsId string) ([]domain.Gpon, error)
	FindByNmsIpServer(ctx context.Context, db *sql.DB, nmsIpServer string) ([]domain.Gpon, error)
	FindByStoId(ctx context.Context, db *sql.DB, stoId string) ([]domain.Gpon, error)
	FindByWitelId(ctx context.Context, db *sql.DB, witelId string) ([]domain.Gpon, error)
	FindByRegionalId(ctx context.Context, db *sql.DB, regionalId string) ([]domain.Gpon, error)
	FindVendorByIpAddress(ctx context.Context, db *sql.DB, ipAddress string) (string, error)
	FindVendorByIpAddressIn(ctx context.Context, db *sql.DB, ipAddresses []string) ([]string, error)
	FindProtocolByIpAddress(ctx context.Context, db *sql.DB, ipAddress string) (string, error)
	FindProtocolByIpAddressIn(ctx context.Context, db *sql.DB, ipAddresses []string) ([]string, error)
}
