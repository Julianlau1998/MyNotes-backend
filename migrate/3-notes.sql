CREATE TABLE notes (
    Id uuid NOT NULL,
    UserID VARCHAR(255),
    FolderID uuid,
    Title VARCHAR(255),
    Note TEXT,
    CreatedDate TIME,
    PRIMARY KEY(Id),
    -- FOREIGN KEY(UserID)
    --     REFERENCES users(Id)
    --         ON DELETE CASCADE,
    FOREIGN KEY(FolderId)
        REFERENCES folders(Id)
            ON DELETE CASCADE    
);