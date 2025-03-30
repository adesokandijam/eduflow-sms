package data

import (
	"database/sql"

	"eduflow.dijam.io/internal/data/student"
	"eduflow.dijam.io/internal/data/teacher"
)

type Models struct {
	StudentModel student.StudentModel
	TeacherModel teacher.TeacherModel
}

func NewModels(DB *sql.DB) Models {
	return Models{
		StudentModel: student.StudentModel{
			DB: DB,
		},
		TeacherModel: teacher.TeacherModel{
			DB: DB,
		},
	}
}
