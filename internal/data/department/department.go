package department

import (
	"context"
	"database/sql"
)

type Department struct {
	DepartmentID int
	Name         int
}

type DepartmentModel struct {
	DB *sql.DB
}

func (d *DepartmentModel) AddDepartment(ctx context.Context, department *Department) error {
	stmt := `INSERT into departments (name) VALUES ($1) RETURNING department_id`

	err := d.DB.QueryRowContext(ctx, stmt, department.Name).Scan(&department.DepartmentID)
	if err != nil {
		return nil
	}
	return nil
}
