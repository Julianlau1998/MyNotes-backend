package models

import (
	"database/sql"
	"notesBackend/utility"
)

type Note struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	FolderID    string `json:"folder_id"`
	Title       string `json:"title"`
	Note        string `json:"body" `
	CreatedDate string `json:"created_date"`
}

type NoteDB struct {
	ID          string
	UserID      sql.NullString
	FolderID    sql.NullString
	Title       sql.NullString
	Note        sql.NullString
	CreatedDate sql.NullString
}

func (dbV *NoteDB) GetNote() (n Note) {
	n.ID = dbV.ID
	n.UserID = utility.GetStringValue(dbV.UserID)
	n.FolderID = utility.GetStringValue(dbV.FolderID)
	n.Title = utility.GetStringValue(dbV.Title)
	n.Note = utility.GetStringValue(dbV.Note)
	n.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	return n
}
