package models

import (
	"database/sql"
	"notesBackend/utility"
)

type ListElement struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	ListID      string `json:"list_id"`
	Element     string `json:"element"`
	Deleted     bool   `json:"deleted"`
	Position    int    `json:"position"`
	CreatedDate string `json:"created_date"`
}

type ListElementDB struct {
	ID          string
	UserID      sql.NullString
	ListID      sql.NullString
	Element     sql.NullString
	Deleted     bool
	Position    int
	CreatedDate sql.NullString
}

func (dbV *ListElementDB) GetListElement() (l ListElement) {
	l.ID = dbV.ID
	l.UserID = utility.GetStringValue(dbV.UserID)
	l.ListID = utility.GetStringValue(dbV.ListID)
	l.Element = utility.GetStringValue(dbV.Element)
	l.Deleted = dbV.Deleted
	l.Position = dbV.Position
	l.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	return l
}
