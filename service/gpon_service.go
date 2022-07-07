package service

import (
	"context"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/repository"
)

type GponService interface {
	GetRepository() repository.GponRepository
	FindAll(ctx context.Context) web.Response[[]web.GponResponse]
	FindById(ctx context.Context, id string) web.Response[web.GponResponse]
	FindByHostname(ctx context.Context, hostname string) web.Response[web.GponResponse]
	FindByIpAddress(ctx context.Context, ipAddress string) web.Response[web.GponResponse]
	FindByNmsId(ctx context.Context, nmsId string) web.Response[[]web.GponResponse]
	FindByNmsIpServer(ctx context.Context, nmsIpServer string) web.Response[[]web.GponResponse]
	FindByStoId(ctx context.Context, stoId string) web.Response[[]web.GponResponse]
	FindByWitelId(ctx context.Context, witelId string) web.Response[[]web.GponResponse]
	FindByRegionalId(ctx context.Context, regionalId string) web.Response[[]web.GponResponse]
}
