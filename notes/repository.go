package notes

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

func (r *Repository) GetNotes(userID string) ([]models.Note, error) {
	var notes []models.Note
	query := `SELECT * FROM notes WHERE userId = $1 ORDER BY createdDate DESC`
	notes, err := r.fetch(query, userID, "")
	return notes, err
}

func (r *Repository) GetByFolder(folderID string, userID string) ([]models.Note, error) {
	var notes []models.Note
	query := `SELECT * FROM notes WHERE userId = $1 AND folderid = $2 ORDER BY createdDate DESC`
	notes, err := r.fetch(query, userID, folderID)
	return notes, err
}

func (r *Repository) GetNoteById(id string, userID string) (models.Note, error) {
	var answer models.Note

	query := `SELECT * FROM notes WHERE id = $1 AND userID = $2`
	answer, err := r.getOne(query, id, userID)
	return answer, err
}

func (r *Repository) Post(note *models.Note) (*models.Note, error) {
	statement := `INSERT INTO notes (id, userid, folderid, title, note, createdDate) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)`
	_, err := r.dbClient.Exec(statement, note.ID, note.UserID, note.FolderID, note.Title, note.Note)
	return note, err
}

func (r *Repository) updateNote(note *models.Note) (models.Note, error) {
	query := `UPDATE notes SET note = $1, title = $2, createdDate = CURRENT_TIMESTAMP WHERE userid = $3 AND id = $4`
	_, err := r.dbClient.Exec(query, note.Note, note.Title, note.UserID, note.ID)

	return *note, err
}

func (r *Repository) deleteNote(note models.Note) error {
	query := `DELETE FROM notes WHERE id = $1 AND userid = $2`
	_, err := r.dbClient.Exec(query, note.ID, note.UserID)
	return err
}

func (r *Repository) fetch(query string, userID string, folderID string) ([]models.Note, error) {
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
	result := make([]models.Note, 0)
	for rows.Next() {
		noteDB := models.NoteDB{}
		err := rows.Scan(&noteDB.ID, &noteDB.UserID, &noteDB.FolderID, &noteDB.Title, &noteDB.Note, &noteDB.CreatedDate)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			log.Infof("Fehler beim Lesen der Daten: %v", err)
			return result, err
		}
		result = append(result, noteDB.GetNote())
	}
	return result, nil
}

func (r *Repository) getOne(query string, id string, userID string) (models.Note, error) {
	noteDB := models.NoteDB{}
	var err error
	err = r.dbClient.QueryRow(query, id, userID).Scan(&noteDB.ID, &noteDB.UserID, &noteDB.FolderID, &noteDB.Title, &noteDB.Note, &noteDB.CreatedDate)
	if err != nil && err != sql.ErrNoRows {
		log.Infof("Fehler beim Lesen der Daten: %v", err)
	}
	return noteDB.GetNote(), err
}
