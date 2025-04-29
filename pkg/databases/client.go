package databases

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var Instance *sqlx.DB
var err error

func Migrate(ctx context.Context) error {
	// Создание таблицы пользователей
	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			firstname VARCHAR(255) NOT NULL,
			lastname VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			role VARCHAR(50) NOT NULL,
			birthdate TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		)
	`); err != nil {
		return err
	}

	// Создание таблицы студентов
	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS students (
			id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
			student_year INTEGER NOT NULL,
			faculty VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		)
	`); err != nil {
		return err
	}

	// Создание таблицы преподавателей
	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS teachers (
			id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
			department VARCHAR(255) NOT NULL,
			position VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		)
	`); err != nil {
		return err
	}

	// Создание таблицы курсов
	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS courses (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			code VARCHAR(50) UNIQUE NOT NULL,
			description TEXT,
			teacher_id INTEGER REFERENCES teachers(id),
			credits INTEGER NOT NULL DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		)
	`); err != nil {
		return err
	}

	// Создание связующей таблицы для студентов и курсов
	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS student_courses (
			student_id INTEGER REFERENCES students(id) ON DELETE CASCADE,
			course_id INTEGER REFERENCES courses(id) ON DELETE CASCADE,
			enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (student_id, course_id)
		)
	`); err != nil {
		return err
	}

	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS course_marks (
			id SERIAL PRIMARY KEY,
			student_id INTEGER REFERENCES students(id) ON DELETE CASCADE,
			course_id INTEGER REFERENCES courses(id) ON DELETE CASCADE,
			first_attestation FLOAT NOT NULL DEFAULT 0,
			second_attestation FLOAT NOT NULL DEFAULT 0,
			final_mark FLOAT NOT NULL DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE (student_id, course_id)
		)
	`); err != nil {
		return err
	}

	// Создание таблицы администраторов
	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS admins (
			id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		)
	`); err != nil {
		return err
	}

	// Создание таблицы менеджеров
	if _, err := Instance.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS managers (
			id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
			department VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		)
	`); err != nil {
		return err
	}

	if _, err := Instance.ExecContext(ctx, `
    CREATE TABLE IF NOT EXISTS teacher_courses (
        teacher_id INTEGER REFERENCES teachers(id) ON DELETE CASCADE,
        course_id INTEGER REFERENCES courses(id) ON DELETE CASCADE,
        assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (teacher_id, course_id)
    )
`); err != nil {
		return err
	}

	// Создание первого админа (ручная запись)
	if _, err := Instance.ExecContext(ctx, `
		INSERT INTO users (id, username, password, firstname, lastname, email, role, created_at, updated_at)
		VALUES (1, 'admin', 'admin', 'Admin', 'Admin', 'admin@admin.com', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		ON CONFLICT (id) DO NOTHING;
	`); err != nil {
		return err
	}
	if _, err := Instance.ExecContext(ctx, `
		INSERT INTO admins (id, created_at, updated_at)
		VALUES (1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		ON CONFLICT (id) DO NOTHING;
	`); err != nil {
		return err
	}

	if _, err := Instance.ExecContext(ctx, `
    SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
`); err != nil {
		return err
	}

	logrus.Println("Database Migration Completed...")
	return nil
}
