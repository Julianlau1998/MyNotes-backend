package folders

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

func (r *Repository) GetFolders(userID string) ([]models.Folder, error) {
	var folders []models.Folder
	query := `SELECT * FROM folders WHERE userId = $1 ORDER BY createdDate DESC`
	folders, err := r.fetch(query, userID)
	return folders, err
}

func (r *Repository) GetById(id string, userID string) (models.Folder, error) {
	var folder models.Folder

	query := `SELECT * FROM folders WHERE id = $1 AND userID = $2`
	folder, err := r.getOne(query, id, userID)
	return folder, err
}

func (r *Repository) Post(folder *models.Folder) (*models.Folder, error) {
	statement := `INSERT INTO folders (id, userid, title, color, folderType, createdDate) VALUES ($1, $2, $3, $4, $5, CURRENT_TIME)`
	_, err := r.dbClient.Exec(statement, folder.ID, folder.UserID, folder.Title, folder.Color, folder.Type)
	return folder, err
}

func (r *Repository) Update(folder *models.Folder, ID string, userID string) (models.Folder, error) {
	query := `UPDATE folders SET title = $1, color = $2, createdDate = Now() WHERE userid = $3 AND id = $4`
	_, err := r.dbClient.Exec(query, folder.Title, folder.Color, folder.ID)

	return *folder, err
}

func (r *Repository) Delete(folder models.Folder, id string) error {
	query := `DELETE FROM folders WHERE id = $1 AND userid = $2`
	_, err := r.dbClient.Exec(query, folder.ID, folder.UserID)
	return err
}

func (r *Repository) fetch(query string, userID string) ([]models.Folder, error) {
	var rows *sql.Rows
	var err error
	rows, err = r.dbClient.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Errorf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
		}
	}()
	result := make([]models.Folder, 0)
	for rows.Next() {
		folderDB := models.FolderDB{}
		err := rows.Scan(&folderDB.ID, &folderDB.UserID, &folderDB.Title, &folderDB.Color, &folderDB.Type, &folderDB.CreatedDate)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			log.Infof("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, folderDB.GetFolder())
	}
	return result, nil
}

func (r *Repository) getOne(query string, id string, userID string) (models.Folder, error) {
	folderDB := models.FolderDB{}
	var err error
	err = r.dbClient.QueryRow(query, id, userID).Scan(&folderDB.ID, &folderDB.UserID, &folderDB.Title, &folderDB.Color, &folderDB.Type, &folderDB.CreatedDate)
	if err != nil && err != sql.ErrNoRows {
		log.Infof("Fehler beim Lesen der Daten: %v", err)
	}
	return folderDB.GetFolder(), err
}
