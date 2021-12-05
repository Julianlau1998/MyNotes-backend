package list_elements

import (
	"database/sql"
	"notesBackend/models"

	"github.com/labstack/gommon/log"
)

type Repository struct {
	dbClient *sql.DB
}

func NewRepository(dbClient *sql.DB) Repository {
	return Repository{dbClient: dbClient}
}

func (r *Repository) GetAllByList(userID string, listID string) ([]models.ListElement, error) {
	var elements []models.ListElement
	query := `SELECT * FROM list_elements WHERE userId = $1 AND listid = $2 ORDER BY position ASC`
	elements, err := r.fetch(query, userID, listID)
	return elements, err
}

func (r *Repository) Post(element models.ListElement) (models.ListElement, error) {
	statement := `INSERT INTO list_elements (id, userid, listId, element, deleted, position, createdDate) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIME)`
	_, err := r.dbClient.Exec(statement, element.ID, element.UserID, element.ListID, element.Element, element.Deleted, element.Position)
	return element, err
}

// func (r *Repository) updateList(list *models.List, ID string, userID string) (models.List, error) {
// 	query := `UPDATE list_elements SET title = $1, createdDate = Now() WHERE userid = $2 AND id = $3`
// 	_, err := r.dbClient.Exec(query, list.Title, ID, userID)

// 	return *list, err
// }

func (r *Repository) Delete(id string, userID string) error {
	query := `DELETE FROM list_elements WHERE id = $1 AND userid = $2`
	_, err := r.dbClient.Exec(query, id, userID)
	return err
}

func (r *Repository) DeleteAllFromList(listID string, userID string) error {
	query := `DELETE FROM list_elements WHERE listid = $1 AND userid = $2`
	_, err := r.dbClient.Exec(query, listID, userID)
	return err
}

func (r *Repository) fetch(query string, userID string, listID string) ([]models.ListElement, error) {
	var rows *sql.Rows
	var err error
	rows, err = r.dbClient.Query(query, userID, listID)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Errorf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
		}
	}()
	result := make([]models.ListElement, 0)
	for rows.Next() {
		ListElementDB := models.ListElementDB{}
		err := rows.Scan(&ListElementDB.ID, &ListElementDB.UserID, &ListElementDB.ListID, &ListElementDB.Element, &ListElementDB.Deleted, &ListElementDB.Position, &ListElementDB.CreatedDate)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			log.Infof("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, ListElementDB.GetListElement())
	}
	return result, nil
}

func (r *Repository) getOne(query string, id string, userID string) (models.ListElement, error) {
	ListElementDB := models.ListElementDB{}
	var err error
	err = r.dbClient.QueryRow(query, id, userID).Scan(&ListElementDB.ID, &ListElementDB.UserID, &ListElementDB.ListID, &ListElementDB.Element, &ListElementDB.Position, &ListElementDB.CreatedDate)
	if err != nil && err != sql.ErrNoRows {
		log.Infof("Fehler beim Lesen der Daten: %v", err)
	}
	return ListElementDB.GetListElement(), err
}
