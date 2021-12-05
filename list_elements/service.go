package list_elements

import (
	"notesBackend/models"

	"github.com/labstack/gommon/log"
	uuid "github.com/nu7hatch/gouuid"
)

type Service struct {
	ListElementRepository Repository
}

func NewService(ListElementRepository Repository) Service {
	return Service{ListElementRepository: ListElementRepository}
}

func (s *Service) GetAllByList(userID string, listID string) ([]models.ListElement, error) {
	listElements, err := s.ListElementRepository.GetAllByList(userID, listID)
	if err != nil {
		log.Warnf("ListElementsService.GetAllByList() Could not Get list elements: %s", err)
		return listElements, err
	}
	return listElements, err
}

func (s *Service) Post(ListElement models.ListElement, listID string) error {
	id, err := uuid.NewV4()
	if err != nil {
		log.Warnf("ListElementsService.Post() Could not create uuid: %s", err)
		return err
	}
	ListElement.ListID = listID
	ListElement.ID = id.String()
	ListElement, err = s.ListElementRepository.Post(ListElement)
	if err != nil {
		log.Warnf("ListService.Post() Could not post listElement: %s", err)
		return err
	}
	return err
}

// func (s *Service) Update(id string, element *models.ListElement, userID string) (models.ListElement, error) {
// 	element, err := s.ListElementRepository.Update(element)
// 	if err != nil {
// 		log.Warnf("ListService.Post() Could not post listElement: %s", err)
// 		return ListElement, err
// 	}
// 	return ListElement, err
// }

func (s *Service) Delete(ID string, userID string) error {
	err := s.ListElementRepository.Delete(ID, userID)
	if err != nil {
		log.Warnf("ListService.Post() Could not post listElement: %s", err)
		return err
	}
	return err
}

func (s *Service) DeleteAllFromList(ListID string, userID string) error {
	err := s.ListElementRepository.DeleteAllFromList(ListID, userID)
	if err != nil {
		log.Warnf("ListService.Post() Could not delete listElements: %s", err)
		return err
	}
	return err
}
