package lists

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

func (r *Repository) GetLists(userID string) ([]models.List, error) {
	var lists []models.List
	query := `SELECT * FROM lists WHERE userId = $1 ORDER BY createdDate DESC`
	lists, err := r.fetch(query, userID, "")
	return lists, err
}

func (r *Repository) GetListById(id string, userId string) (models.List, error) {
	query := `SELECT * FROM lists WHERE id = $1 AND userID = $2`
	answer, err := r.getOne(query, id, userId)
	return answer, err
}

func (r *Repository) GetByFolder(folderID string, userID string) ([]models.List, error) {
	var lists []models.List
	query := `SELECT * FROM lists WHERE userId = $1 AND folderid = $2 ORDER BY createdDate DESC`
	lists, err := r.fetch(query, userID, folderID)
	return lists, err
}

func (r *Repository) PostList(list *models.List) (*models.List, error) {
	statement := `INSERT INTO lists (id, userid, folderid, title, createdDate) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)`
	_, err := r.dbClient.Exec(statement, list.ID, list.UserID, list.FolderID, list.Title)
	return list, err
}

func (r *Repository) updateList(list *models.List, ID string, userID string) (models.List, error) {
	query := `UPDATE lists SET title = $1, createdDate = CURRENT_TIMESTAMP WHERE userid = $2 AND id = $3`
	_, err := r.dbClient.Exec(query, list.Title, userID, ID)

	return *list, err
}

func (r *Repository) DeleteList(list models.List, id string, userID string) error {
	query := `DELETE FROM lists WHERE id = $1 AND userid = $2`
	_, err := r.dbClient.Exec(query, id, userID)
	return err
}

func (r *Repository) fetch(query string, userID string, folderID string) ([]models.List, error) {
	var rows *sql.Rows
	var err error
	if len(folderID) > 0 {
		rows, err = r.dbClient.Query(query, userID, folderID)
	} else {
		rows, err = r.dbClient.Query(query, userID)
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Errorf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
		}
	}()
	result := make([]models.List, 0)
	for rows.Next() {
		listDB := models.ListDB{}
		err := rows.Scan(&listDB.ID, &listDB.UserID, &listDB.FolderID, &listDB.Title, &listDB.CreatedDate)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			log.Infof("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, listDB.GetList())
	}
	return result, nil
}

func (r *Repository) getOne(query string, id string, userID string) (models.List, error) {
	listDB := models.ListDB{}
	var err error
	err = r.dbClient.QueryRow(query, id, userID).Scan(&listDB.ID, &listDB.UserID, &listDB.FolderID, &listDB.Title, &listDB.CreatedDate)
	if err != nil && err != sql.ErrNoRows {
		log.Infof("Fehler beim Lesen der Daten: %v", err)
	}
	return listDB.GetList(), err
}
