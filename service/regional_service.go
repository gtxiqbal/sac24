package service

import (
	"context"
	"github.com/gtxiqbal/sac24/model/web"
)

type RegionalService interface {
	FindAll(ctx context.Context) web.Response[[]web.RegionalResponse]
	FindById(ctx context.Context, id string) web.Response[web.RegionalResponse]
	Create(ctx context.Context, regional web.RegionalRequest) web.Response[web.RegionalResponse]
}
