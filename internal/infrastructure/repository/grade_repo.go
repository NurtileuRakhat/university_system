package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	domainModels "university_system/internal/domain/models"
	domainRepo "university_system/internal/domain/repository"
)

type GradeRepositoryImpl struct {
	DB *sqlx.DB
}

func NewGradeRepository(db *sqlx.DB) domainRepo.GradeRepository {
	return &GradeRepositoryImpl{DB: db}
}

func (r *GradeRepositoryImpl) GetStudentMarks(ctx context.Context, studentID string) ([]domainModels.Mark, error) {
	var marks []domainModels.Mark
	err := r.DB.SelectContext(ctx, &marks, "SELECT * FROM course_marks WHERE student_id = $1", studentID)
	if err != nil {
		return nil, err
	}
	return marks, nil
}

func (r *GradeRepositoryImpl) GetCourseMarks(ctx context.Context, courseID string) ([]domainModels.Mark, error) {
	var marks []domainModels.Mark
	err := r.DB.SelectContext(ctx, &marks, "SELECT * FROM course_marks WHERE course_id = $1", courseID)
	if err != nil {
		return nil, err
	}
	return marks, nil
}

// AddMark добавляет оценку указанного типа
func (r *GradeRepositoryImpl) AddMark(ctx context.Context, mark *domainModels.Mark, markType string) error {
	var query string
	var value float64

	switch markType {
	case "first_attestation":
		query = "UPDATE course_marks SET first_attestation = $1 WHERE student_id = $2 AND course_id = $3"
		value = mark.FirstAttestation
	case "second_attestation":
		query = "UPDATE course_marks SET second_attestation = $1 WHERE student_id = $2 AND course_id = $3"
		value = mark.SecondAttestation
	case "final":
		query = "UPDATE course_marks SET final_mark = $1 WHERE student_id = $2 AND course_id = $3"
		value = mark.FinalMark
	default:
		return domainModels.ErrInvalidMarkType
	}

	result, err := r.DB.ExecContext(ctx, query, value, mark.StudentID, mark.CourseID)
	if err != nil {
		return err
	}

	// Если не обновили ни одной записи, создаем новую запись
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		// Создаем новую запись с оценкой
		query := "INSERT INTO course_marks (student_id, course_id, first_attestation, second_attestation, final_mark) VALUES ($1, $2, $3, $4, $5)"
		
		firstAtt := 0.0
		secondAtt := 0.0
		finalMark := 0.0
		
		switch markType {
		case "first_attestation":
			firstAtt = value
		case "second_attestation":
			secondAtt = value
		case "final":
			finalMark = value
		}
		
		_, err = r.DB.ExecContext(ctx, query, mark.StudentID, mark.CourseID, firstAtt, secondAtt, finalMark)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsTeacherOfCourse проверяет, является ли преподаватель ведущим данного курса
func (r *GradeRepositoryImpl) IsTeacherOfCourse(ctx context.Context, teacherID string, courseID string) (bool, error) {
	var count int
	err := r.DB.GetContext(ctx, &count, 
		"SELECT COUNT(*) FROM teacher_courses WHERE teacher_id = $1 AND course_id = $2", 
		teacherID, courseID)
	
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}
