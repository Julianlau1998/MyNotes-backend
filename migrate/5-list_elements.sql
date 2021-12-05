CREATE TABLE list_elements (
    Id uuid NOT NULL,
    UserID VARCHAR(255),
    ListId uuid,
    element VARCHAR(255),
    deleted BOOLEAN NOT NULL DEFAULT false,
    position SMALLINT,
    CreatedDate TIME,
    PRIMARY KEY(Id),
    -- FOREIGN KEY(UserID)
    --     REFERENCES users(Id)
    --         ON DELETE CASCADE,
    FOREIGN KEY(ListId)
        REFERENCES lists(Id)
            ON DELETE CASCADE
);