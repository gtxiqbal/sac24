package service

import (
	"context"
	"github.com/gtxiqbal/sac24/model/web"
)

type NmsService interface {
	FindAll(ctx context.Context) web.Response[[]web.NmsResponse]
	FindById(ctx context.Context, id string) web.Response[web.NmsResponse]
	FindByNama(ctx context.Context, nama string) web.Response[web.NmsResponse]
	FindByVendor(ctx context.Context, vendor string) web.Response[[]web.NmsResponse]
	FindByIpServer(ctx context.Context, ipServer string) web.Response[web.NmsResponse]
}
