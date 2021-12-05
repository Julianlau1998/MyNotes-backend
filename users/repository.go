// package users

// import (
// 	"database/sql"
// 	"notesBackend/models"

// 	"github.com/labstack/gommon/log"
// )

// type Repository struct {
// 	dbClient *sql.DB
// }

// func NewRepository(dbClient *sql.DB) Repository {
// 	return Repository{dbClient: dbClient}
// }

// func (r *Repository) GetUsers() ([]models.User, error) {
// 	var users []models.User
// 	query := `SELECT * FROM users`
// 	users, err := r.fetch(query)
// 	return users, err
// }

// func (r *Repository) PostUser(user *models.User) (*models.User, error) {
// 	statement := `INSERT INTO users (id, username, passw) VALUES ($1, $2, $3)`
// 	_, err := r.dbClient.Exec(statement, user.ID, user.Username, user.Password)
// 	return user, err
// }

// func (r *Repository) fetch(query string) ([]models.User, error) {
// 	rows, err := r.dbClient.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		err := rows.Close()
// 		if err != nil {
// 			log.Errorf("Datenbankverbindung konnte nicht korrekt geschlossen werden: %v", err)
// 		}
// 	}()
// 	result := make([]models.User, 0)
// 	for rows.Next() {
// 		userDB := models.UserDB{}
// 		err := rows.Scan(&userDB.ID, &userDB.Username, &userDB.Password)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				continue
// 			}
// 			log.Infof("Fehler beim Lesen der Daten: %v", err)
// 			return result, err
// 		}
// 		result = append(result, userDB.GetUser())
// 	}
// 	return result, nil
// }
