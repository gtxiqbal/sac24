package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
)

type WitelRepositoryImpl struct {
}

func NewWitelRepositoryImpl() WitelRepository {
	return &WitelRepositoryImpl{}
}

func (repository *WitelRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Witel {
	rows, err := db.QueryContext(ctx,
		`SELECT w.id, w.alias, w.create_at, w.nama, w.update_at,
						r.id, r.create_at, r.nama, r.update_at 
				FROM witel w, regional r 
				WHERE w.regional_id = r.id 
				ORDER BY alias `,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var witels []domain.Witel
	for rows.Next() {
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		witel := domain.Witel{}

		err := rows.Scan(
			&witel.Id,
			&witel.Alias,
			&witel.CreateAt,
			&witel.Nama,
			&witelUpdateAt,
			&witel.Regional.Id,
			&witel.Regional.CreateAt,
			&witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		witel.UpdateAt = witelUpdateAt.Time
		witel.Regional.UpdateAt = regionalUpdateAt.Time
		witels = append(witels, witel)
	}

	return witels
}

func (repository *WitelRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) (domain.Witel, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT w.id, w.alias, w.create_at, w.nama, w.update_at,
						r.id, r.create_at, r.nama, r.update_at 
				FROM witel w, regional r 
				WHERE w.regional_id = r.id 
					AND w.id = $1
				ORDER BY r.id, w.alias `,
		id,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var witelUpdateAt sql.NullTime
	var regionalUpdateAt sql.NullTime
	witel := domain.Witel{}

	if rows.Next() {
		err := rows.Scan(
			&witel.Id,
			&witel.Alias,
			&witel.CreateAt,
			&witel.Nama,
			&witelUpdateAt,
			&witel.Regional.Id,
			&witel.Regional.CreateAt,
			&witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		witel.UpdateAt = witelUpdateAt.Time
		witel.Regional.UpdateAt = regionalUpdateAt.Time

		return witel, nil
	}

	return witel, errors.New("witel not found by id")
}

func (repository *WitelRepositoryImpl) FindByRegionalId(ctx context.Context, db *sql.DB, regionalId string) ([]domain.Witel, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT w.id, w.alias, w.create_at, w.nama, w.update_at,
       				r.id, r.create_at, r.nama, r.update_at
				FROM witel w, regional r 
				WHERE w.regional_id = r.id 
				  AND r.id = $1
				ORDER BY r.id, w.alias `,
		regionalId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var witels []domain.Witel
	for rows.Next() {
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		witel := domain.Witel{}

		err := rows.Scan(
			&witel.Id,
			&witel.Alias,
			&witel.CreateAt,
			&witel.Nama,
			&witelUpdateAt,
			&witel.Regional.Id,
			&witel.Regional.CreateAt,
			&witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		witel.UpdateAt = witelUpdateAt.Time
		witel.Regional.UpdateAt = regionalUpdateAt.Time
		witels = append(witels, witel)
	}

	if len(witels) > 0 {
		return witels, nil
	}
	return witels, errors.New("witel not found by regional id")
}
