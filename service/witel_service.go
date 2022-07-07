package service

import (
	"context"
	"github.com/gtxiqbal/sac24/model/web"
)

type WitelService interface {
	FindAll(ctx context.Context) web.Response[[]web.WitelResponse]
	FindById(ctx context.Context, id string) web.Response[web.WitelResponse]
	FindByRegionalId(ctx context.Context, regionalId string) web.Response[[]web.WitelResponse]
}
