-- migrations/000005_create_classes_table.up.sql

CREATE TABLE classes (
    class_id SERIAL PRIMARY KEY,
    class_name VARCHAR(255) NOT NULL UNIQUE,
    class_description TEXT -- Optional: Add a description column
);