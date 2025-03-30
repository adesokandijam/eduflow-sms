-- migrations/000004_add_grade_id_to_students.up.sql

ALTER TABLE students
ADD COLUMN grade_id INT REFERENCES grades(grade_id);