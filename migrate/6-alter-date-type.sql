ALTER TABLE notes
ALTER COLUMN CreatedDate TYPE TIMESTAMP;

ALTER TABLE lists
ALTER COLUMN CreatedDate TYPE TIMESTAMP;


ALTER TABLE folders
ALTER COLUMN CreatedDate TYPE TIMESTAMP;

ALTER TABLE list_elements
ALTER COLUMN CreatedDate TYPE TIMESTAMP;