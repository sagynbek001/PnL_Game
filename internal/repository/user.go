package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

type userRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) models.UserRepo {
	return &userRepo{conn: conn}
}

func (u *userRepo) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (login, password_hash) VALUES ($1, $2)`
	_, err := u.conn.Exec(ctx, query, user.Login, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("UserRepo error: %w", err)
	}

	return nil
}

func (u *userRepo) GetAll(ctx context.Context) ([]models.User, error) {
	sql := `SELECT * FROM users`
	rows, err := u.conn.Query(ctx, sql)

	if err != nil {
		return []models.User{}, fmt.Errorf("UserRepo error: %w", err)
	}

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Login, &user.PasswordHash)
		if err != nil {
			fmt.Printf("UserRepo error: %v", err)

			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepo) GetByID(ctx context.Context, id uint64) (models.User, error) {
	var user models.User

	query := `SELECT * FROM users WHERE id=$1`

	row := u.conn.QueryRow(ctx, query, id)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		return models.User{}, fmt.Errorf("UserRepo error: %w", err)
	}

	return user, nil
}

func (u *userRepo) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET  login=$2, password_hash=$3 WHERE id=$1`

	_, err := u.conn.Exec(ctx, query, user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		return fmt.Errorf("UserRepo error: %w", err)
	}

	return nil
}

func (u *userRepo) Delete(ctx context.Context, id uint64) error {
	sql := `DELETE FROM users WHERE id=$1`
	_, err := u.conn.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("UserRepo error: %w", err)
	}
	return nil
}

func (u *userRepo) FindByLogin(ctx context.Context, email string) (models.User, error) {
	sql := `SELECT * FROM users WHERE login=$1`
	rows := u.conn.QueryRow(ctx, sql, email)
	user := models.User{}
	err := rows.Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		return models.User{}, fmt.Errorf("UserRepo error: %w", err)
	}

	return user, nil
}
