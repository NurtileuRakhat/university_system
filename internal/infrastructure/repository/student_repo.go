package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	domainModels "university_system/internal/domain/models"
	domainRepo "university_system/internal/domain/repository"
)

type StudentRepositoryImpl struct {
	DB *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) domainRepo.StudentRepository {
	return &StudentRepositoryImpl{DB: db}
}

func (r *StudentRepositoryImpl) GetStudents(ctx context.Context) ([]domainModels.Student, error) {
	var students []domainModels.Student
	query := `SELECT 
		s.id, 
		u.username, u.firstname, u.lastname, u.email, u.role, u.birthdate, 
		s.student_year, s.faculty, s.created_at, s.updated_at, s.deleted_at
	FROM students s
	JOIN users u ON s.id = u.id`
	err := r.DB.SelectContext(ctx, &students, query)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentRepositoryImpl) GetStudentById(ctx context.Context, id string) (*domainModels.Student, error) {
	var student domainModels.Student
	query := `SELECT 
		s.id, 
		u.username, u.firstname, u.lastname, u.email, u.role, u.birthdate, 
		s.student_year, s.faculty, s.created_at, s.updated_at, s.deleted_at
	FROM students s
	JOIN users u ON s.id = u.id
	WHERE s.id = $1`
	err := r.DB.GetContext(ctx, &student, query, id)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepositoryImpl) CreateStudent(ctx context.Context, student *domainModels.Student) (*domainModels.Student, error) {
	query := `INSERT INTO students (id, student_year, faculty) VALUES (:id, :student_year, :faculty) RETURNING id`
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var id string
	err = stmt.QueryRowxContext(ctx, student).Scan(&id)
	if err != nil {
		return nil, err
	}
	student.ID = id
	return student, nil
}

func (r *StudentRepositoryImpl) UpdateStudent(ctx context.Context, student domainModels.Student) (*domainModels.Student, error) {
	_, err := r.DB.NamedExecContext(ctx, `UPDATE students SET student_year=:student_year, faculty=:faculty WHERE id=:id`, &student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepositoryImpl) DeleteStudent(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}

func (r *StudentRepositoryImpl) EnrollStudentToCourse(ctx context.Context, studentID, courseID string) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO student_courses (student_id, course_id) VALUES ($1, $2)", studentID, courseID)
	return err
}

func (r *StudentRepositoryImpl) GetStudentCourses(ctx context.Context, studentID string) ([]domainModels.Course, error) {
	var courses []domainModels.Course
	err := r.DB.SelectContext(ctx, &courses, "SELECT c.* FROM courses c JOIN student_courses sc ON c.id = sc.course_id WHERE sc.student_id = $1", studentID)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *StudentRepositoryImpl) CreateUserWithRole(ctx context.Context, user domainModels.User, role string) (string, error) {
	var id string
	query := `INSERT INTO users (username, password, firstname, lastname, email, role, birthdate) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Password, user.Firstname, user.Lastname, user.Email, role, user.Birthdate).Scan(&id)
	return id, err
}
