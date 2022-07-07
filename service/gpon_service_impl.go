package service

import (
	"context"
	"database/sql"
	"github.com/dranikpg/dto-mapper"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/repository"
)

type GponServiceImpl struct {
	*sql.DB
	repository.GponRepository
}

func NewGponServiceImpl(DB *sql.DB, gponRepository repository.GponRepository) GponService {
	return &GponServiceImpl{DB: DB, GponRepository: gponRepository}
}

func (service *GponServiceImpl) GetRepository() repository.GponRepository {
	return service.GponRepository
}

func (service *GponServiceImpl) FindAll(ctx context.Context) web.Response[[]web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpons := gponRepository.FindAll(ctx, db)

	var gponsResponse []web.GponResponse
	err := dto.Map(&gponsResponse, gpons)
	helper.PanicIfError(err)

	response := web.Response[[]web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data GPON",
		Data:    gponsResponse,
	}
	return response
}

func (service *GponServiceImpl) FindById(ctx context.Context, id string) web.Response[web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpon, err := gponRepository.FindById(ctx, db, id)
	helper.PanicIfError(err)

	var gponResponse web.GponResponse
	err = dto.Map(&gponResponse, gpon)
	helper.PanicIfError(err)

	response := web.Response[web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data GPON By Id",
		Data:    gponResponse,
	}
	return response
}

func (service *GponServiceImpl) FindByHostname(ctx context.Context, hostname string) web.Response[web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpon, err := gponRepository.FindByHostname(ctx, db, hostname)
	helper.PanicIfError(err)

	var gponResponse web.GponResponse
	err = dto.Map(&gponResponse, gpon)
	helper.PanicIfError(err)

	response := web.Response[web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data GPON By Hostname",
		Data:    gponResponse,
	}
	return response
}

func (service *GponServiceImpl) FindByIpAddress(ctx context.Context, ipAddress string) web.Response[web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpon, err := gponRepository.FindByIpAddress(ctx, db, ipAddress)
	helper.PanicIfError(err)

	var gponResponse web.GponResponse
	err = dto.Map(&gponResponse, gpon)
	helper.PanicIfError(err)

	response := web.Response[web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data GPON By Ip Address",
		Data:    gponResponse,
	}
	return response
}

func (service *GponServiceImpl) FindByNmsId(ctx context.Context, nmsId string) web.Response[[]web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpons, err := gponRepository.FindByNmsId(ctx, db, nmsId)
	helper.PanicIfError(err)

	var gponsResponse []web.GponResponse
	err = dto.Map(&gponsResponse, gpons)
	helper.PanicIfError(err)

	response := web.Response[[]web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data GPON By NMS Id",
		Data:    gponsResponse,
	}
	return response
}

func (service *GponServiceImpl) FindByNmsIpServer(ctx context.Context, nmsIpServer string) web.Response[[]web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpons, err := gponRepository.FindByNmsIpServer(ctx, db, nmsIpServer)
	helper.PanicIfError(err)

	var gponsResponse []web.GponResponse
	err = dto.Map(&gponsResponse, gpons)
	helper.PanicIfError(err)

	response := web.Response[[]web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data GPON By NMS Ip Server",
		Data:    gponsResponse,
	}
	return response
}

func (service *GponServiceImpl) FindByStoId(ctx context.Context, stoId string) web.Response[[]web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpons, err := gponRepository.FindByStoId(ctx, db, stoId)
	helper.PanicIfError(err)

	var gponsResponse []web.GponResponse
	err = dto.Map(&gponsResponse, gpons)
	helper.PanicIfError(err)

	response := web.Response[[]web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data GPON By STO Id",
		Data:    gponsResponse,
	}
	return response
}

func (service *GponServiceImpl) FindByWitelId(ctx context.Context, witelId string) web.Response[[]web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpons, err := gponRepository.FindByWitelId(ctx, db, witelId)
	helper.PanicIfError(err)

	var gponsResponse []web.GponResponse
	err = dto.Map(&gponsResponse, gpons)
	helper.PanicIfError(err)

	response := web.Response[[]web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data GPON By Witel Id",
		Data:    gponsResponse,
	}
	return response
}

func (service *GponServiceImpl) FindByRegionalId(ctx context.Context, regionalId string) web.Response[[]web.GponResponse] {
	db := service.DB
	gponRepository := service.GponRepository
	gpons, err := gponRepository.FindByRegionalId(ctx, db, regionalId)
	helper.PanicIfError(err)

	var gponsResponse []web.GponResponse
	err = dto.Map(&gponsResponse, gpons)
	helper.PanicIfError(err)

	response := web.Response[[]web.GponResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data GPON By Regional Id",
		Data:    gponsResponse,
	}
	return response
}
