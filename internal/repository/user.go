package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/news/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) models.UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(c context.Context, user models.UserRequest) (int, error) {
	var userID int
	userQuery := `INSERT INTO users (email, password, name)
VALUES ($1, $2, $3) returning id;`
	err := ur.db.QueryRow(c, userQuery, user.Email, user.Password, user.Name).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (ur *UserRepository) EditUser(c context.Context, user models.User) (int, error) {
	userQuery := `UPDATE users SET name = $1 WHERE id = $2;`
	_, err := ur.db.Exec(c, userQuery, user.Name, user.ID)
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil
}
func (ur *UserRepository) DeleteUser(c context.Context, userID int) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := ur.db.Exec(c, query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(c context.Context, email string) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, name,role_id FROM users where email = $1`
	row := ur.db.QueryRow(c, query, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.RoleID)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByID(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, name,role_id FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.RoleID)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetProfile(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, name, role_id FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.RoleID)

	if err != nil {
		return user, err
	}

	return user, nil
}