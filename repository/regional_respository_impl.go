package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
)

type RegionalRepositoryImpl struct {
}

func NewRegionalRepositoryImpl() RegionalRepository {
	return &RegionalRepositoryImpl{}
}

func (repository *RegionalRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Regional {
	rows, err := db.QueryContext(ctx, "SELECT * FROM public.regional ORDER BY id")
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var regionals []domain.Regional
	for rows.Next() {
		regional := domain.Regional{}
		var updateAt sql.NullTime
		err := rows.Scan(
			&regional.Id,
			&regional.CreateAt,
			&regional.Nama,
			&updateAt,
		)
		helper.PanicIfError(err)
		regional.UpdateAt = updateAt.Time
		regionals = append(regionals, regional)
	}
	return regionals
}

func (repository *RegionalRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) (domain.Regional, error) {
	rows, err := db.QueryContext(
		ctx,
		"SELECT * FROM public.regional WHERE id = $1 ORDER BY ID",
		id,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	regional := domain.Regional{}
	if rows.Next() {
		var updateAt sql.NullTime
		err := rows.Scan(
			&regional.Id,
			&regional.CreateAt,
			&regional.Nama,
			&updateAt,
		)
		helper.PanicIfError(err)
		regional.UpdateAt = updateAt.Time
		return regional, nil
	}
	return regional, errors.New("regional not found by id")
}

func (repository *RegionalRepositoryImpl) Save(ctx context.Context, db *sql.DB, regional domain.Regional) (domain.Regional, error) {
	result, err := db.ExecContext(
		ctx,
		"INSERT INTO public.regional(id, create_at, nama) VALUES($1, now(), $2)",
		regional.Id, regional.Nama,
	)
	helper.PanicIfError(err)
	return save(result, regional, false)
}

func (repository *RegionalRepositoryImpl) SaveTx(ctx context.Context, tx *sql.Tx, regional domain.Regional) (domain.Regional, error) {
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO public.regional(id, create_at, nama) VALUES($1, now(), $2)",
		regional.Id, regional.Nama,
	)
	helper.PanicIfError(err)
	return save(result, regional, false)
}

func (repository *RegionalRepositoryImpl) SaveAll(ctx context.Context, db *sql.DB, regionals []domain.Regional) []domain.Regional {
	stmt, err := db.PrepareContext(
		ctx,
		"INSERT INTO public.regional(id, create_at, nama) VALUES($1, now(), $2)",
	)
	defer helper.StmtClose(stmt)
	return saveAll(ctx, stmt, regionals, err)
}

func (repository *RegionalRepositoryImpl) SaveAllTx(ctx context.Context, tx *sql.Tx, regionals []domain.Regional) []domain.Regional {
	stmt, err := tx.PrepareContext(
		ctx,
		"INSERT INTO public.regional(id, create_at, nama) VALUES($1, now(), $2)",
	)
	defer helper.StmtClose(stmt)
	return saveAll(ctx, stmt, regionals, err)
}

func (repository *RegionalRepositoryImpl) Update(ctx context.Context, db *sql.DB, regional domain.Regional) (domain.Regional, error) {
	result, err := db.ExecContext(
		ctx,
		"UPDATE public.regional SET nama=$1, update_at=now() WHERE id = $2",
		regional.Nama, regional.Id,
	)
	helper.PanicIfError(err)
	return save(result, regional, true)
}

func (repository *RegionalRepositoryImpl) UpdateTx(ctx context.Context, tx *sql.Tx, regional domain.Regional) (domain.Regional, error) {
	result, err := tx.ExecContext(
		ctx,
		"UPDATE public.regional SET nama=$1, update_at=now() WHERE id = $2",
		regional.Nama, regional.Id,
	)
	helper.PanicIfError(err)
	return save(result, regional, true)
}

func (repository *RegionalRepositoryImpl) UpdateAll(ctx context.Context, db *sql.DB, regionals []domain.Regional) []domain.Regional {
	stmt, err := db.PrepareContext(
		ctx,
		"UPDATE public.regional SET nama=$1, update_at=now() WHERE id = $1",
	)
	return updateAll(ctx, stmt, regionals, err)
}

func (repository *RegionalRepositoryImpl) UpdateAllTx(ctx context.Context, tx *sql.Tx, regionals []domain.Regional) []domain.Regional {
	stmt, err := tx.PrepareContext(
		ctx,
		"UPDATE public.regional SET nama=$1, update_at=now() WHERE id = $1",
	)
	return updateAll(ctx, stmt, regionals, err)
}

func (repository *RegionalRepositoryImpl) Delete(ctx context.Context, db *sql.DB, regional domain.Regional) {
	repository.DeleteById(ctx, db, regional.Id)
}

func (repository *RegionalRepositoryImpl) DeleteTx(ctx context.Context, tx *sql.Tx, regional domain.Regional) {
	repository.DeleteByIdTx(ctx, tx, regional.Id)
}

func (repository *RegionalRepositoryImpl) DeleteById(ctx context.Context, db *sql.DB, id string) {
	result, err := db.ExecContext(
		ctx,
		"DELETE FROM public.regional WHERE id = $1 ",
		id,
	)
	helper.PanicIfError(err)

	_, err = result.RowsAffected()
	helper.PanicIfError(err)
}

func (repository *RegionalRepositoryImpl) DeleteByIdTx(ctx context.Context, tx *sql.Tx, id string) {
	result, err := tx.ExecContext(
		ctx,
		"DELETE FROM public.regional WHERE id = $1 ",
		id,
	)
	helper.PanicIfError(err)

	_, err = result.RowsAffected()
	helper.PanicIfError(err)
}

func (repository *RegionalRepositoryImpl) DeleteAll(ctx context.Context, db *sql.DB, regionals []domain.Regional) {
	stmt, err := db.PrepareContext(ctx, "DELETE FROM public.regional WHERE id = $1 ")
	helper.PanicIfError(err)
	defer helper.StmtClose(stmt)
	for _, regional := range regionals {
		_, err := stmt.ExecContext(ctx, regional.Id)
		helper.PanicIfError(err)
	}
}

func (repository *RegionalRepositoryImpl) DeleteAllTx(ctx context.Context, tx *sql.Tx, regionals []domain.Regional) {
	stmt, err := tx.PrepareContext(ctx, "DELETE FROM public.regional WHERE id = $1 ")
	helper.PanicIfError(err)
	defer helper.StmtClose(stmt)
	for _, regional := range regionals {
		_, err := stmt.ExecContext(ctx, regional.Id)
		helper.PanicIfError(err)
	}
}

func (repository *RegionalRepositoryImpl) DeleteAllById(ctx context.Context, db *sql.DB, ids []string) {
	stmt, err := db.PrepareContext(ctx, "DELETE FROM public.regional WHERE id = $1 ")
	helper.PanicIfError(err)
	defer helper.StmtClose(stmt)
	for _, id := range ids {
		_, err := stmt.ExecContext(ctx, id)
		helper.PanicIfError(err)
	}
}

func (repository *RegionalRepositoryImpl) DeleteAllByIdTx(ctx context.Context, tx *sql.Tx, ids []string) {
	stmt, err := tx.PrepareContext(ctx, "DELETE FROM public.regional WHERE id = $1 ")
	helper.PanicIfError(err)
	defer helper.StmtClose(stmt)
	for _, id := range ids {
		_, err := stmt.ExecContext(ctx, id)
		helper.PanicIfError(err)
	}
}

func save(result sql.Result, regional domain.Regional, isUpdate bool) (domain.Regional, error) {
	rowsAffected, err := result.RowsAffected()
	helper.PanicIfError(err)

	if rowsAffected != 0 {
		return regional, nil
	}

	if isUpdate {
		return domain.Regional{}, errors.New("update data regional failed")
	}
	return domain.Regional{}, errors.New("insert data regional failed")
}

func saveAll(ctx context.Context, stmt *sql.Stmt, regionals []domain.Regional, err error) []domain.Regional {
	helper.PanicIfError(err)

	var regionalsSave []domain.Regional
	for _, regional := range regionals {
		result, err := stmt.ExecContext(ctx, regional.Id, regional.Nama)
		helper.PanicIfError(err)

		rowsAffected, err := result.RowsAffected()
		helper.PanicIfError(err)

		if rowsAffected != 0 {
			regionalsSave = append(regionalsSave, regional)
		}
	}

	return regionalsSave
}

func updateAll(ctx context.Context, stmt *sql.Stmt, regionals []domain.Regional, err error) []domain.Regional {
	helper.PanicIfError(err)
	defer helper.StmtClose(stmt)

	var regionalsSave []domain.Regional
	for _, regional := range regionals {
		result, err := stmt.ExecContext(ctx, regional.Nama, regional.Id)
		helper.PanicIfError(err)

		rowsAffected, err := result.RowsAffected()
		helper.PanicIfError(err)

		if rowsAffected != 0 {
			regionalsSave = append(regionalsSave, regional)
		}
	}

	return regionalsSave
}
