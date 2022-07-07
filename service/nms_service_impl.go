package service

import (
	"context"
	"database/sql"
	"github.com/dranikpg/dto-mapper"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/repository"
)

type NmsServiceImpl struct {
	*sql.DB
	repository.NmsRepository
}

func NewNmsServiceImpl(DB *sql.DB, nmsRepository repository.NmsRepository) NmsService {
	return &NmsServiceImpl{DB: DB, NmsRepository: nmsRepository}
}

func (service *NmsServiceImpl) FindAll(ctx context.Context) web.Response[[]web.NmsResponse] {
	db := service.DB
	nmsRepository := service.NmsRepository
	nmsList := nmsRepository.FindAll(ctx, db)

	var nmsListResponse []web.NmsResponse
	err := dto.Map(&nmsListResponse, nmsList)
	helper.PanicIfError(err)

	response := web.Response[[]web.NmsResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data NMS",
		Data:    nmsListResponse,
	}
	return response
}

func (service *NmsServiceImpl) FindById(ctx context.Context, id string) web.Response[web.NmsResponse] {
	db := service.DB
	nmsRepository := service.NmsRepository
	nms, err := nmsRepository.FindById(ctx, db, id)
	helper.PanicIfError(err)

	var nmsResponse web.NmsResponse
	err = dto.Map(&nmsResponse, nms)
	helper.PanicIfError(err)

	response := web.Response[web.NmsResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data NMS By Id",
		Data:    nmsResponse,
	}
	return response
}

func (service *NmsServiceImpl) FindByNama(ctx context.Context, nama string) web.Response[web.NmsResponse] {
	db := service.DB
	nmsRepository := service.NmsRepository
	nms, err := nmsRepository.FindByNama(ctx, db, nama)
	helper.PanicIfError(err)

	var nmsResponse web.NmsResponse
	err = dto.Map(&nmsResponse, nms)
	helper.PanicIfError(err)

	response := web.Response[web.NmsResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data NMS By Nama",
		Data:    nmsResponse,
	}
	return response
}

func (service *NmsServiceImpl) FindByVendor(ctx context.Context, vendor string) web.Response[[]web.NmsResponse] {
	db := service.DB
	nmsRepository := service.NmsRepository
	nmsList, err := nmsRepository.FindByVendor(ctx, db, vendor)
	helper.PanicIfError(err)

	var nmsListResponse []web.NmsResponse
	err = dto.Map(&nmsListResponse, nmsList)
	helper.PanicIfError(err)

	response := web.Response[[]web.NmsResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan List Data NMS By Vendor",
		Data:    nmsListResponse,
	}
	return response
}

func (service *NmsServiceImpl) FindByIpServer(ctx context.Context, ipServer string) web.Response[web.NmsResponse] {
	db := service.DB
	nmsRepository := service.NmsRepository
	nms, err := nmsRepository.FindByIpServer(ctx, db, ipServer)
	helper.PanicIfError(err)

	var nmsResponse web.NmsResponse
	err = dto.Map(&nmsResponse, nms)
	helper.PanicIfError(err)

	response := web.Response[web.NmsResponse]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Berhasil Menampilkan Data NMS By Ip Server",
		Data:    nmsResponse,
	}
	return response
}
