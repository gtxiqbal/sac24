package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
)

type StoRepositoryImpl struct {
}

func NewStoRepositoryImpl() StoRepository {
	return &StoRepositoryImpl{}
}

func (repository *StoRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Sto {
	rows, err := db.QueryContext(ctx,
		`SELECT s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM sto s, witel w, regional r 
				WHERE s.witel_id = w.id
				  AND w.regional_id = r.id 
				ORDER BY r.id, w.alias, s.alias `,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var stos []domain.Sto
	for rows.Next() {
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		sto := domain.Sto{}

		err := rows.Scan(
			&sto.Id,
			&sto.Alias,
			&sto.CreateAt,
			&sto.Nama,
			&stoUpdateAt,
			&sto.Witel.Id,
			&sto.Witel.Alias,
			&sto.Witel.CreateAt,
			&sto.Witel.Nama,
			&witelUpdateAt,
			&sto.Witel.Regional.Id,
			&sto.Witel.Regional.CreateAt,
			&sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		sto.UpdateAt = stoUpdateAt.Time
		sto.Witel.UpdateAt = witelUpdateAt.Time
		sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		stos = append(stos, sto)
	}

	return stos
}

func (repository *StoRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) (domain.Sto, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM sto s, witel w, regional r 
				WHERE s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND s.id = $1
				ORDER BY r.id, w.alias, s.alias `,
		id,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var sto domain.Sto
	if rows.Next() {
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime

		err = rows.Scan(
			&sto.Id,
			&sto.Alias,
			&sto.CreateAt,
			&sto.Nama,
			&stoUpdateAt,
			&sto.Witel.Id,
			&sto.Witel.Alias,
			&sto.Witel.CreateAt,
			&sto.Witel.Nama,
			&witelUpdateAt,
			&sto.Witel.Regional.Id,
			&sto.Witel.Regional.CreateAt,
			&sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		sto.UpdateAt = stoUpdateAt.Time
		sto.Witel.UpdateAt = witelUpdateAt.Time
		sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		return sto, nil
	}

	return sto, errors.New("sto not found by id")
}

func (repository *StoRepositoryImpl) FindByWitelId(ctx context.Context, db *sql.DB, witelId string) ([]domain.Sto, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM sto s, witel w, regional r 
				WHERE s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND w.id = $1
				ORDER BY r.id, w.alias, s.alias `,
		witelId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var stos []domain.Sto
	for rows.Next() {
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		sto := domain.Sto{}

		err := rows.Scan(
			&sto.Id,
			&sto.Alias,
			&sto.CreateAt,
			&sto.Nama,
			&stoUpdateAt,
			&sto.Witel.Id,
			&sto.Witel.Alias,
			&sto.Witel.CreateAt,
			&sto.Witel.Nama,
			&witelUpdateAt,
			&sto.Witel.Regional.Id,
			&sto.Witel.Regional.CreateAt,
			&sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		sto.UpdateAt = stoUpdateAt.Time
		sto.Witel.UpdateAt = witelUpdateAt.Time
		sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		stos = append(stos, sto)
	}

	if len(stos) > 0 {
		return stos, nil
	}
	return stos, errors.New("sto not found by witel id")
}

func (repository *StoRepositoryImpl) FindByRegionalId(ctx context.Context, db *sql.DB, regionalId string) ([]domain.Sto, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM sto s, witel w, regional r 
				WHERE s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND r.id = $1
				ORDER BY r.id, w.alias, s.alias `,
		regionalId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var stos []domain.Sto
	for rows.Next() {
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		sto := domain.Sto{}

		err = rows.Scan(
			&sto.Id,
			&sto.Alias,
			&sto.CreateAt,
			&sto.Nama,
			&stoUpdateAt,
			&sto.Witel.Id,
			&sto.Witel.Alias,
			&sto.Witel.CreateAt,
			&sto.Witel.Nama,
			&witelUpdateAt,
			&sto.Witel.Regional.Id,
			&sto.Witel.Regional.CreateAt,
			&sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		sto.UpdateAt = stoUpdateAt.Time
		sto.Witel.UpdateAt = witelUpdateAt.Time
		sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		stos = append(stos, sto)
	}

	if len(stos) > 0 {
		return stos, nil
	}
	return stos, errors.New("sto not found by regional id")
}
