package service

import (
	"context"
	"database/sql"
	"github.com/dranikpg/dto-mapper"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/repository"
)

type WitelServiceImpl struct {
	*sql.DB
	repository.WitelRepository
}

func NewWitelServiceImpl(DB *sql.DB, witelRepository repository.WitelRepository) WitelService {
	return &WitelServiceImpl{DB: DB, WitelRepository: witelRepository}
}

func (service *WitelServiceImpl) FindAll(ctx context.Context) web.Response[[]web.WitelResponse] {
	db := service.DB
	witelRepository := service.WitelRepository
	witels := witelRepository.FindAll(ctx, db)

	var witelsResponse []web.WitelResponse
	err := dto.Map(&witelsResponse, witels)
	helper.PanicIfError(err)

	response := web.Response[[]web.WitelResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data Witel",
		Data:    witelsResponse,
	}
	return response
}

func (service *WitelServiceImpl) FindById(ctx context.Context, id string) web.Response[web.WitelResponse] {
	db := service.DB
	witelRepository := service.WitelRepository
	witel, err := witelRepository.FindById(ctx, db, id)
	helper.PanicIfError(err)

	var witelResponse web.WitelResponse
	err = dto.Map(&witelResponse, witel)
	helper.PanicIfError(err)

	response := web.Response[web.WitelResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data Witel By Id " + witelResponse.Id,
		Data:    witelResponse,
	}
	return response
}

func (service *WitelServiceImpl) FindByRegionalId(ctx context.Context, regionalId string) web.Response[[]web.WitelResponse] {
	db := service.DB
	witelRepository := service.WitelRepository
	witels, err := witelRepository.FindByRegionalId(ctx, db, regionalId)
	helper.PanicIfError(err)

	var witelsResponse []web.WitelResponse
	err = dto.Map(&witelsResponse, witels)
	helper.PanicIfError(err)

	response := web.Response[[]web.WitelResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data Witel By Regional Id " + regionalId,
		Data:    witelsResponse,
	}
	return response
}
