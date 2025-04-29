package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	domainModels "university_system/internal/domain/models"
	domainRepo "university_system/internal/domain/repository"
)

type TeacherRepositoryImpl struct {
	DB *sqlx.DB
}

func NewTeacherRepository(db *sqlx.DB) domainRepo.TeacherRepository {
	return &TeacherRepositoryImpl{DB: db}
}

func (r *TeacherRepositoryImpl) GetTeachers(ctx context.Context) ([]domainModels.Teacher, error) {
	var teachers []domainModels.Teacher
	query := `SELECT 
		t.id, 
		u.username, u.firstname, u.lastname, u.email, u.role, u.birthdate, 
		t.department, t.position, t.created_at, t.updated_at, t.deleted_at
	FROM teachers t
	JOIN users u ON t.id = u.id`
	err := r.DB.SelectContext(ctx, &teachers, query)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *TeacherRepositoryImpl) GetTeacherById(ctx context.Context, id string) (*domainModels.Teacher, error) {
	var teacher domainModels.Teacher
	query := `SELECT 
		t.id, 
		u.username, u.firstname, u.lastname, u.email, u.role, u.birthdate, 
		t.department, t.position, t.created_at, t.updated_at, t.deleted_at
	FROM teachers t
	JOIN users u ON t.id = u.id
	WHERE t.id = $1`
	err := r.DB.GetContext(ctx, &teacher, query, id)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepositoryImpl) CreateTeacher(ctx context.Context, teacher *domainModels.Teacher) (*domainModels.Teacher, error) {
	query := `INSERT INTO teachers (id, department, position) VALUES (:id, :department, :position) RETURNING id`
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var id string
	err = stmt.QueryRowxContext(ctx, teacher).Scan(&id)
	if err != nil {
		return nil, err
	}
	teacher.ID = id
	return teacher, nil
}

func (r *TeacherRepositoryImpl) UpdateTeacher(ctx context.Context, teacher domainModels.Teacher) (*domainModels.Teacher, error) {
	_, err := r.DB.NamedExecContext(ctx, `UPDATE teachers SET department=:department, position=:position WHERE id=:id`, &teacher)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepositoryImpl) DeleteTeacher(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}

func (r *TeacherRepositoryImpl) AssignTeacherToCourse(ctx context.Context, teacherID, courseID string) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO teacher_courses (teacher_id, course_id) VALUES ($1, $2)", teacherID, courseID)
	return err
}

func (r *TeacherRepositoryImpl) GetTeacherCourses(ctx context.Context, teacherID string) ([]domainModels.Course, error) {
	var courses []domainModels.Course
	err := r.DB.SelectContext(ctx, &courses, "SELECT c.* FROM courses c JOIN teacher_courses tc ON c.id = tc.course_id WHERE tc.teacher_id = $1", teacherID)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *TeacherRepositoryImpl) CreateUserWithRole(ctx context.Context, user domainModels.User, role string) (string, error) {
	var id string
	query := `INSERT INTO users (username, password, firstname, lastname, email, role, birthdate) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Password, user.Firstname, user.Lastname, user.Email, role, user.Birthdate).Scan(&id)
	return id, err
}
