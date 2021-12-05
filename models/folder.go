package models

import (
	"database/sql"
	"notesBackend/utility"
)

type Folder struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Color       string `json:"color"`
	Type        string `json:"type"`
	CreatedDate string `json:"created_date"`
}

type FolderDB struct {
	ID          string
	UserID      sql.NullString
	Title       sql.NullString
	Color       sql.NullString
	Type        sql.NullString
	CreatedDate sql.NullString
}

func (dbV *FolderDB) GetFolder() (f Folder) {
	f.ID = dbV.ID
	f.UserID = utility.GetStringValue(dbV.UserID)
	f.Title = utility.GetStringValue(dbV.Title)
	f.Color = utility.GetStringValue(dbV.Color)
	f.Type = utility.GetStringValue(dbV.Type)
	f.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	return f
}
