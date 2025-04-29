package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	domainModels "university_system/internal/domain/models"
	domainRepo "university_system/internal/domain/repository"
)

type ManagerRepositoryImpl struct {
	DB *sqlx.DB
}

func NewManagerRepository(db *sqlx.DB) domainRepo.ManagerRepository {
	return &ManagerRepositoryImpl{DB: db}
}

func (r *ManagerRepositoryImpl) GetManagers(ctx context.Context) ([]domainModels.Manager, error) {
	var managers []domainModels.Manager
	query := `SELECT 
		m.id, 
		u.username, u.firstname, u.lastname, u.email, u.role, u.birthdate, 
		m.created_at, m.updated_at, m.deleted_at
	FROM managers m
	JOIN users u ON m.id = u.id`
	err := r.DB.SelectContext(ctx, &managers, query)
	if err != nil {
		return nil, err
	}
	return managers, nil
}

func (r *ManagerRepositoryImpl) GetManagerById(ctx context.Context, id string) (*domainModels.Manager, error) {
	var manager domainModels.Manager
	query := `SELECT 
		m.id, 
		u.username, u.firstname, u.lastname, u.email, u.role, u.birthdate, 
		m.created_at, m.updated_at, m.deleted_at
	FROM managers m
	JOIN users u ON m.id = u.id
	WHERE m.id = $1`
	err := r.DB.GetContext(ctx, &manager, query, id)
	if err != nil {
		return nil, err
	}
	return &manager, nil
}

func (r *ManagerRepositoryImpl) CreateManager(ctx context.Context, manager *domainModels.Manager) (*domainModels.Manager, error) {
	query := `INSERT INTO managers (id, department) VALUES (:id, :department) RETURNING id`
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var id string
	err = stmt.QueryRowxContext(ctx, manager).Scan(&id)
	if err != nil {
		return nil, err
	}
	manager.ID = id
	return manager, nil
}

func (r *ManagerRepositoryImpl) UpdateManager(ctx context.Context, manager domainModels.Manager) (*domainModels.Manager, error) {
	_, err := r.DB.NamedExecContext(ctx, `UPDATE managers SET id=:id WHERE id=:id`, &manager)
	if err != nil {
		return nil, err
	}
	return &manager, nil
}

func (r *ManagerRepositoryImpl) DeleteManager(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}

// AssignTeacherToCourse назначает преподавателя на курс
func (r *ManagerRepositoryImpl) AssignTeacherToCourse(ctx context.Context, teacherID string, courseID string) error {
	// Проверяем, существует ли такая запись
	var count int
	err := r.DB.GetContext(ctx, &count, 
		"SELECT COUNT(*) FROM teacher_courses WHERE teacher_id = $1 AND course_id = $2", 
		teacherID, courseID)
	
	if err != nil {
		return err
	}
	
	if count > 0 {
		// Запись уже существует, ничего не делаем
		return nil
	}
	
	// Создаем новую запись
	_, err = r.DB.ExecContext(ctx, 
		"INSERT INTO teacher_courses (teacher_id, course_id) VALUES ($1, $2)",
		teacherID, courseID)
	
	return err
}

func (r *ManagerRepositoryImpl) CreateUserWithRole(ctx context.Context, user domainModels.User, role string) (string, error) {
	var id string
	query := `INSERT INTO users (username, password, firstname, lastname, email, role, birthdate) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Password, user.Firstname, user.Lastname, user.Email, role, user.Birthdate).Scan(&id)
	return id, err
}
