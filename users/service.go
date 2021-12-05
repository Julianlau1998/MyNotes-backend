// package users

// import (
// 	"notesBackend/models"

// 	uuid "github.com/nu7hatch/gouuid"
// )

// type Service struct {
// 	userRepository Repository
// }

// func NewService(userRepository Repository) Service {
// 	return Service{userRepository: userRepository}
// }

// func (s *Service) GetUser(username string, password string) (models.User, error) {
// 	users, err := s.userRepository.GetUsers()
// 	var emptyUser models.User
// 	for _, user := range users {
// 		if user.Username == username && user.Password == password {
// 			return user, err
// 		} else {
// 			emptyUser.ID = ""
// 			emptyUser.Username = ""
// 			emptyUser.Password = ""
// 		}
// 	}
// 	return emptyUser, err
// }

// func (s *Service) Post(user *models.User) (*models.User, error) {
// 	id, err := uuid.NewV4()
// 	if err != nil {
// 		return user, err
// 	}
// 	user.ID = id.String()
// 	return s.userRepository.PostUser(user)
// }