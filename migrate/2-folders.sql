CREATE TABLE folders (
    Id uuid NOT NULL,
    UserID VARCHAR(255),
    Title VARCHAR(255),
    Color VARCHAR(255),
    FolderType VARCHAR(255),
    CreatedDate TIME,
    PRIMARY KEY(Id)
    -- FOREIGN KEY(UserID)
    --     REFERENCES users(Id)
    --         ON DELETE CASCADE
);

INSERT INTO folders(Id) VALUES ('00000000-0000-0000-0000-000000000000');