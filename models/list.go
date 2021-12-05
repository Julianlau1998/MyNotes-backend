package models

import (
	"database/sql"
	"notesBackend/utility"
)

type List struct {
	ID          string   `json:"id"`
	UserID      string   `json:"user_id"`
	FolderID    string   `json:"folder_id"`
	Title       string   `json:"title"`
	List        []string `json:"list" `
	DoneItems   []string `json:"doneItems"`
	CreatedDate string   `json:"created_date"`
}

type ListDB struct {
	ID          string
	UserID      sql.NullString
	FolderID    sql.NullString
	Title       sql.NullString
	CreatedDate sql.NullString
}

func (dbV *ListDB) GetList() (l List) {
	l.ID = dbV.ID
	l.UserID = utility.GetStringValue(dbV.UserID)
	l.FolderID = utility.GetStringValue(dbV.FolderID)
	l.Title = utility.GetStringValue(dbV.Title)
	l.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	return l
}
