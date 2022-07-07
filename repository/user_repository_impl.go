package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) FindByTelegramId(ctx context.Context, db *sql.DB, telegramId int) ([]domain.User, error) {
	rows, err := db.QueryContext(
		ctx,
		`SELECT username, create_at, default_user, email, enabled, expired_date, fail_pass, first_name, last_name, 
       				  limit_fail, limit_login, "password", telegram_id, update_at, expired_password_date 
			   FROM public.users 
			   WHERE telegram_id = $1 
			   ORDER BY username`,
		telegramId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var users []domain.User
	for rows.Next() {
		var userLastName sql.NullString
		var userUpdateAt sql.NullTime
		user := domain.User{}

		err := rows.Scan(
			&user.Username,
			&user.CreateAt,
			&user.DefaultUser,
			&user.Email,
			&user.Enabled,
			&user.ExpiredDate,
			&user.FailPass,
			&user.FirstName,
			&userLastName,
			&user.LimitFail,
			&user.LimitLogin,
			&user.Password,
			&user.TelegramId,
			&userUpdateAt,
			&user.ExpiredPasswordDate,
		)
		helper.PanicIfError(err)

		user.LastName = userLastName.String
		user.UpdateAt = userUpdateAt.Time
		users = append(users, user)
	}

	if len(users) > 0 {
		return users, nil
	}

	return users, errors.New("user not found by telegram id")
}
