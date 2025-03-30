package student

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"eduflow.dijam.io/internal/validator"
)

type Student struct {
	StudentID    int64
	FirstName    string
	LastName     string
	Gender       string
	Email        string
	PasswordHash []byte
	ImageURL     string
	ClassID      int64
	Status       string
}

type StudentModel struct {
	DB *sql.DB
}

func ValidateStudent(v *validator.Validator, student *Student) {
	v.Check((len(student.FirstName) > 0 && len(student.FirstName) < 50), "student first name", "must be greater than zero and less than 50")
	v.Check((len(student.LastName) > 0 && len(student.LastName) < 50), "student last name", "must be greater than zero and less than 50")
	v.Check(validator.In(student.Status, "active", "left", "punished", "withdrawn", "graduated"), "student status", "must be in the string")
	v.Check(validator.In(student.Gender, "male", "female"), "student gender", "must be male or female")
	v.Check(validator.Matches(student.Email, validator.EmailRX), "email", "must be valid")
}

func (u *StudentModel) AddStudent(ctx context.Context, student *Student) error {
	stmt := `INSERT INTO students (first_name, last_name, gender, email, image_url, status, class_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING student_id`
	err := u.DB.QueryRowContext(ctx, stmt, student.FirstName, student.LastName, student.Gender, student.Email, student.ImageURL, student.Status, student.ClassID).Scan(&student.StudentID)
	if err != nil {
		return err
	}
	return nil
}

func (u *StudentModel) AddStudentPassword(ctx context.Context, student *Student) error {
	stmt := `UPDATE students SET password_hash = $1 WHERE student_id = $2`
	_, err := u.DB.ExecContext(ctx, stmt, student.PasswordHash, student.StudentID)
	return err
}

func (u *StudentModel) GetStudentByID(ctx context.Context, id int64) (*Student, error) {
	stmt := `SELECT student_id, first_name, last_name, gender, email, password_hash, image_url, status, class_id FROM students WHERE student_id = $1`
	row := u.DB.QueryRowContext(ctx, stmt, id)

	student := &Student{}
	err := row.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.Gender, &student.Email, &student.PasswordHash, &student.ImageURL, &student.Status, &student.ClassID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	return student, nil
}

func (u *StudentModel) GetAllStudents(ctx context.Context) ([]*Student, error) {
	stmt := `SELECT student_id, first_name, last_name, gender, email, password_hash, image_url, status, class_id FROM students`
	rows, err := u.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []*Student{}
	for rows.Next() {
		student := &Student{}
		err := rows.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.Gender, &student.Email, &student.PasswordHash, &student.ImageURL, &student.Status, &student.ClassID)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (u *StudentModel) UpdateStudent(ctx context.Context, student *Student) error {
	stmt := `UPDATE students SET first_name = $1, last_name = $2, gender = $3, email = $4, image_url = $5, status = $6 WHERE student_id = $7`
	_, err := u.DB.ExecContext(ctx, stmt, student.FirstName, student.LastName, student.Gender, student.Email, student.ImageURL, student.Status, student.StudentID)
	return err
}

func (u *StudentModel) DeleteStudent(ctx context.Context, id int64) error {
	stmt := `DELETE FROM students WHERE student_id = $1`
	result, err := u.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("student with id %d not found", id)
	}

	return nil
}
