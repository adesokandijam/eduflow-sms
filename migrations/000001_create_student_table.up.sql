-- migrations/000001_create_students_table.up.sql

CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash BYTEA,
    image_url VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active'
);