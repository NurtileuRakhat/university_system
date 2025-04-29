package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	domainModels "university_system/internal/domain/models"
	domainRepo "university_system/internal/domain/repository"
)

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domainRepo.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) GetUsers(ctx context.Context) ([]domainModels.User, error) {
	var users []domainModels.User
	err := r.DB.SelectContext(ctx, &users, "SELECT id, username, firstname, lastname, email, password, birthdate, role FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id string) (*domainModels.User, error) {
	var user domainModels.User
	err := r.DB.GetContext(ctx, &user, "SELECT id, username, firstname, lastname, email, password, birthdate, role FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*domainModels.User, error) {
	var user domainModels.User
	err := r.DB.GetContext(ctx, &user, "SELECT id, username, firstname, lastname, email, password, birthdate, role FROM users WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domainModels.User, error) {
	var user domainModels.User
	err := r.DB.GetContext(ctx, &user, "SELECT id, username, firstname, lastname, email, password, birthdate, role FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *domainModels.User) (*domainModels.User, error) {
	query := `INSERT INTO users (username, firstname, lastname, email, password, birthdate, role) VALUES (:username, :firstname, :lastname, :email, :password, :birthdate, :role) RETURNING id`
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var id string
	err = stmt.QueryRowxContext(ctx, user).Scan(&id)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user domainModels.User) (*domainModels.User, error) {
	_, err := r.DB.NamedExecContext(ctx, `UPDATE users SET username=:username, firstname=:firstname, lastname=:lastname, email=:email, password=:password, birthdate=:birthdate, role=:role WHERE id=:id`, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}
