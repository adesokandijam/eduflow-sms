package teacher

import (
	"context"
	"database/sql"
)

type Teacher struct {
	TeacherID    string
	FirstName    string
	LastName     string
	Email        string
	DepartmentID int
	PasswordHash []byte
	Gender       string
	Status       string
	ImageURL     string
}

type TeacherModel struct {
	DB *sql.DB
}

func (t *TeacherModel) AddTeacher(ctx context.Context, teacher *Teacher) error {
	stmt := `INSERT INTO teachers(first_name, last_name, gender, email, image_url, status, department_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING teacher_id`
	err := t.DB.QueryRowContext(ctx, stmt, teacher.FirstName, teacher.LastName, teacher.Gender, teacher.Email, teacher.ImageURL, teacher.Status, teacher.Status).Scan(&teacher.TeacherID)
	if err != nil {
		return err
	}
	return nil
}
