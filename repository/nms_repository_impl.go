package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
)

type NmsRepositoryImpl struct {
}

func NewNmsRepositoryImpl() NmsRepository {
	return &NmsRepositoryImpl{}
}

func (repository *NmsRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Nms {
	rows, err := db.QueryContext(ctx, "SELECT * FROM nms ORDER BY vendor, nama")
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var nmsList []domain.Nms
	for rows.Next() {
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		nms := domain.Nms{}

		err = rows.Scan(
			&nms.Id,
			&nms.Nama,
			&nms.Vendor,
			&nms.IpServer,
			&nms.PortTl1,
			&nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&nms.CreateAt,
			&nmsUpdateAt,
		)
		helper.PanicIfError(err)

		nms.Username = nmsUsername.String
		nms.Password = nmsPassword.String
		nms.UpdateAt = nmsUpdateAt.Time
		nmsList = append(nmsList, nms)
	}

	return nmsList
}

func (repository *NmsRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) (domain.Nms, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM nms WHERE id = $1 ORDER BY vendor, nama", id)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var nmsUsername sql.NullString
	var nmsPassword sql.NullString
	var nmsUpdateAt sql.NullTime
	nms := domain.Nms{}
	if rows.Next() {
		err = rows.Scan(
			&nms.Id,
			&nms.Nama,
			&nms.Vendor,
			&nms.IpServer,
			&nms.PortTl1,
			&nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&nms.CreateAt,
			&nmsUpdateAt,
		)
		helper.PanicIfError(err)

		nms.Username = nmsUsername.String
		nms.Password = nmsPassword.String
		nms.UpdateAt = nmsUpdateAt.Time

		return nms, nil
	}

	return nms, errors.New("nms not found by id")
}

func (repository *NmsRepositoryImpl) FindByNama(ctx context.Context, db *sql.DB, nama string) (domain.Nms, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM nms WHERE nama = $1 ORDER BY vendor, nama", nama)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var nmsUsername sql.NullString
	var nmsPassword sql.NullString
	var nmsUpdateAt sql.NullTime
	nms := domain.Nms{}
	if rows.Next() {
		err = rows.Scan(
			&nms.Id,
			&nms.Nama,
			&nms.Vendor,
			&nms.IpServer,
			&nms.PortTl1,
			&nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&nms.CreateAt,
			&nmsUpdateAt,
		)
		helper.PanicIfError(err)

		nms.Username = nmsUsername.String
		nms.Password = nmsPassword.String
		nms.UpdateAt = nmsUpdateAt.Time

		return nms, nil
	}

	return nms, errors.New("nms not found by nama")
}

func (repository *NmsRepositoryImpl) FindByVendor(ctx context.Context, db *sql.DB, vendor string) ([]domain.Nms, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM nms WHERE vendor = $1 ORDER BY vendor, nama", vendor)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var nmsList []domain.Nms
	for rows.Next() {
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		nms := domain.Nms{}

		err = rows.Scan(
			&nms.Id,
			&nms.Nama,
			&nms.Vendor,
			&nms.IpServer,
			&nms.PortTl1,
			&nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&nms.CreateAt,
			&nmsUpdateAt,
		)
		helper.PanicIfError(err)

		nms.Username = nmsUsername.String
		nms.Password = nmsPassword.String
		nms.UpdateAt = nmsUpdateAt.Time
		nmsList = append(nmsList, nms)
	}

	if len(nmsList) > 0 {
		return nmsList, nil
	}
	return nmsList, errors.New("")
}

func (repository *NmsRepositoryImpl) FindByIpServer(ctx context.Context, db *sql.DB, ipServer string) (domain.Nms, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM nms WHERE ip_server = $1 ORDER BY vendor, nama", ipServer)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var nmsUsername sql.NullString
	var nmsPassword sql.NullString
	var nmsUpdateAt sql.NullTime
	nms := domain.Nms{}
	if rows.Next() {
		err = rows.Scan(
			&nms.Id,
			&nms.Nama,
			&nms.Vendor,
			&nms.IpServer,
			&nms.PortTl1,
			&nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&nms.CreateAt,
			&nmsUpdateAt,
		)
		helper.PanicIfError(err)

		nms.Username = nmsUsername.String
		nms.Password = nmsPassword.String
		nms.UpdateAt = nmsUpdateAt.Time

		return nms, nil
	}

	return nms, errors.New("nms not found by ip server")
}
