package repository

import (
	"context"
	"errors"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	row := r.db.QueryRow(ctx, `
        SELECT
            users.id,
            users.encrypted_password,
            users.created_at
        FROM
            users
        WHERE
            users.id = $1
    `, id,
	)

	var u entity.User
	err := row.Scan(&u.ID, &u.EncryptedPassword, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
