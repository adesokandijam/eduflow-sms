-- migrations/000002_create_departments_and_teachers_tables.up.sql

CREATE TABLE departments (
    department_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE teachers (
    teacher_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash BYTEA,
    image_url VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    department_id INT REFERENCES departments(department_id)
);

