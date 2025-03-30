-- migrations/000003_create_grades_table.up.sql

CREATE TABLE grades (
    grade_id SERIAL PRIMARY KEY,
    grade_name VARCHAR(255) NOT NULL UNIQUE
);