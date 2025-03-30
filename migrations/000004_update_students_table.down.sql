-- migrations/000004_add_grade_id_to_students.down.sql

ALTER TABLE students
DROP COLUMN grade_id;