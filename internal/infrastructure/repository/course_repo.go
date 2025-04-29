package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	domainModels "university_system/internal/domain/models"
	domainRepo "university_system/internal/domain/repository"
)

type CourseRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) domainRepo.CourseRepository {
	return &CourseRepositoryImpl{DB: db}
}

func (r *CourseRepositoryImpl) CreateCourse(ctx context.Context, course *domainModels.Course) (*domainModels.Course, error) {
	query := `INSERT INTO courses (name, code, description, teacher_id, credits) VALUES (:name, :code, :description, :teacher_id, :credits) RETURNING id`
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var id string
	err = stmt.QueryRowxContext(ctx, course).Scan(&id)
	if err != nil {
		return nil, err
	}
	course.ID = id
	return course, nil
}

func (r *CourseRepositoryImpl) GetCourseByID(ctx context.Context, ID string) (*domainModels.Course, error) {
	var course domainModels.Course
	err := r.DB.GetContext(ctx, &course, "SELECT * FROM courses WHERE id = $1", ID)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepositoryImpl) GetAllCourses(ctx context.Context) ([]domainModels.Course, error) {
	var courses []domainModels.Course
	err := r.DB.SelectContext(ctx, &courses, "SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) UpdateCourse(ctx context.Context, course domainModels.Course) (*domainModels.Course, error) {
	_, err := r.DB.NamedExecContext(ctx, `UPDATE courses SET name=:name, code=:code, description=:description, teacher_id=:teacher_id, credits=:credits WHERE id=:id`, &course)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepositoryImpl) DeleteCourse(ctx context.Context, ID string) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM courses WHERE id = $1", ID)
	return err
}

func (r *CourseRepositoryImpl) GetCourseStudents(ctx context.Context, courseID string) ([]domainModels.Student, error) {
	var students []domainModels.Student
	err := r.DB.SelectContext(ctx, &students, "SELECT s.* FROM students s JOIN student_courses sc ON s.id = sc.student_id WHERE sc.course_id = $1", courseID)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *CourseRepositoryImpl) GetCourseTeachers(ctx context.Context, courseID string) ([]domainModels.Teacher, error) {
	var teachers []domainModels.Teacher
	err := r.DB.SelectContext(ctx, &teachers, "SELECT t.* FROM teachers t JOIN teacher_courses tc ON t.id = tc.teacher_id WHERE tc.course_id = $1", courseID)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}
