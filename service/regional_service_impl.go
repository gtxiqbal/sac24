package service

import (
	"context"
	"database/sql"
	"github.com/dranikpg/dto-mapper"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/repository"
)

type RegionalServiceImpl struct {
	*sql.DB
	repository.RegionalRepository
}

func NewRegionalServiceImpl(DB *sql.DB, regionalRepository repository.RegionalRepository) RegionalService {
	return &RegionalServiceImpl{DB: DB, RegionalRepository: regionalRepository}
}

func (service *RegionalServiceImpl) FindAll(ctx context.Context) web.Response[[]web.RegionalResponse] {
	db := service.DB
	regionalRepository := service.RegionalRepository
	regionals := regionalRepository.FindAll(ctx, db)

	var regionalResponse []web.RegionalResponse
	err := dto.Map(&regionalResponse, regionals)
	helper.PanicIfError(err)

	webResponse := web.Response[[]web.RegionalResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data Regional",
		Data:    regionalResponse,
	}

	return webResponse
}

func (service *RegionalServiceImpl) FindById(ctx context.Context, id string) web.Response[web.RegionalResponse] {
	db := service.DB
	regionalRepository := service.RegionalRepository
	regionals, err := regionalRepository.FindById(ctx, db, id)
	helper.PanicIfError(err)

	var regionalResponse web.RegionalResponse
	err = dto.Map(&regionalResponse, regionals)
	helper.PanicIfError(err)

	webResponse := web.Response[web.RegionalResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data Regional By ID " + id,
		Data:    regionalResponse,
	}

	return webResponse
}

func (service *RegionalServiceImpl) Create(ctx context.Context, request web.RegionalRequest) web.Response[web.RegionalResponse] {
	regionalRepository := service.RegionalRepository
	regional := domain.Regional{}
	err := dto.Map(&regional, request)
	helper.PanicIfError(err)

	db := service.DB
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	regionalSaved, err := regionalRepository.SaveTx(ctx, tx, regional)
	helper.PanicIfError(err)

	var regionalResponse web.RegionalResponse
	err = dto.Map(&regionalResponse, regionalSaved)
	helper.PanicIfError(err)

	webResponse := web.Response[web.RegionalResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menambahkan Data Regional",
		Data:    regionalResponse,
	}
	return webResponse
}
