ALTER TABLE students
ADD COLUMN class_id INT REFERENCES classes(class_id);

-- If you want class_id to be NOT NULL and provide a default, do it here.
-- Example:
-- INSERT INTO classes (class_name) VALUES ('Default Class'); -- Create a default class
-- UPDATE students SET class_id = (SELECT class_id FROM classes WHERE class_name = 'Default Class');
-- ALTER TABLE students ALTER COLUMN class_id SET NOT NULL;

-- Remove the grade_id column
ALTER TABLE students
DROP COLUMN IF EXISTS grade_id;