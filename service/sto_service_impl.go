package service

import (
	"context"
	"database/sql"
	"github.com/dranikpg/dto-mapper"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/repository"
)

type StoServiceImpl struct {
	*sql.DB
	repository.StoRepository
}

func NewStoServiceImpl(DB *sql.DB, stoRepository repository.StoRepository) StoService {
	return &StoServiceImpl{DB: DB, StoRepository: stoRepository}
}

func (service *StoServiceImpl) FindAll(ctx context.Context) web.Response[[]web.StoResponse] {
	db := service.DB
	stoRepository := service.StoRepository
	stos := stoRepository.FindAll(ctx, db)

	var stosResponse []web.StoResponse
	err := dto.Map(&stosResponse, stos)
	helper.PanicIfError(err)

	response := web.Response[[]web.StoResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data STO",
		Data:    stosResponse,
	}
	return response
}

func (service *StoServiceImpl) FindById(ctx context.Context, id string) web.Response[web.StoResponse] {
	db := service.DB
	stoRepository := service.StoRepository
	sto, err := stoRepository.FindById(ctx, db, id)
	helper.PanicIfError(err)

	var stoResponse web.StoResponse
	err = dto.Map(&stoResponse, sto)
	helper.PanicIfError(err)

	response := web.Response[web.StoResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data STO By Id",
		Data:    stoResponse,
	}
	return response
}

func (service *StoServiceImpl) FindByWitelId(ctx context.Context, witelId string) web.Response[[]web.StoResponse] {
	db := service.DB
	stoRepository := service.StoRepository
	stos, err := stoRepository.FindByWitelId(ctx, db, witelId)
	helper.PanicIfError(err)

	var stosResponse []web.StoResponse
	err = dto.Map(&stosResponse, stos)
	helper.PanicIfError(err)

	response := web.Response[[]web.StoResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data STO By Witel Id",
		Data:    stosResponse,
	}
	return response
}

func (service *StoServiceImpl) FindByRegionalId(ctx context.Context, regionalId string) web.Response[[]web.StoResponse] {
	db := service.DB
	stoRepository := service.StoRepository
	stos, err := stoRepository.FindByRegionalId(ctx, db, regionalId)
	helper.PanicIfError(err)

	var stosResponse []web.StoResponse
	err = dto.Map(&stosResponse, stos)
	helper.PanicIfError(err)

	response := web.Response[[]web.StoResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data STO By Regional Id",
		Data:    stosResponse,
	}
	return response
}
