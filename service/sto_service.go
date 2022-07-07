package service

import (
	"context"
	"github.com/gtxiqbal/sac24/model/web"
)

type StoService interface {
	FindAll(ctx context.Context) web.Response[[]web.StoResponse]
	FindById(ctx context.Context, id string) web.Response[web.StoResponse]
	FindByWitelId(ctx context.Context, witelId string) web.Response[[]web.StoResponse]
	FindByRegionalId(ctx context.Context, regionalId string) web.Response[[]web.StoResponse]
}
