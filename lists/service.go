package lists

import (
	"notesBackend/list_elements"
	"notesBackend/models"

	"github.com/labstack/gommon/log"
	uuid "github.com/nu7hatch/gouuid"
)

type Service struct {
	listRepository     Repository
	ListElementService list_elements.Service
}

func NewService(listRepository Repository, listElementService list_elements.Service) Service {
	return Service{
		listRepository:     listRepository,
		ListElementService: listElementService,
	}
}

func (s *Service) GetLists(userID string) ([]models.List, error) {
	lists, err := s.listRepository.GetLists(userID)
	if err != nil {
		log.Warnf("ListService.GetListById() Could not Get lists by id: %s", err)
		return lists, err
	}
	return lists, err
}

func (s *Service) GetListById(id string, userID string) (models.List, error) {
	list, err := s.listRepository.GetListById(id, userID)
	listElements, err := s.ListElementService.GetAllByList(list.UserID, list.ID)
	if err != nil {
		log.Warnf("ListService.GetListById() Could get list by id: %s", err)
		return list, err
	}
	for _, element := range listElements {
		if element.Deleted == true {
			list.DoneItems = append(list.DoneItems, element.Element)
		} else {
			list.List = append(list.List, element.Element)
		}
	}
	if list.DoneItems == nil {
		list.DoneItems = []string{}
	}
	if err != nil {
		log.Warnf("ListService.GetListById() Could not Get list by id: %s", err)
		return list, err
	}
	return list, err
}

func (s *Service) GetByFolder(folderID string, userID string) ([]models.List, error) {
	lists, err := s.listRepository.GetByFolder(folderID, userID)
	if err != nil {
		log.Warnf("ListService.Delete() Could not load lists by folder: %s", err)
		return lists, err
	}
	return lists, nil
}

func (s *Service) Post(list *models.List, userID string) (*models.List, error) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Warnf("ListService.Post() Could not delete list: %s", err)
		return list, err
	}
	if list.FolderID == "" {
		list.FolderID = "00000000-0000-0000-0000-000000000000"
	}
	list.ID = id.String()
	list.UserID = userID
	list, err = s.listRepository.PostList(list)
	if err != nil {
		log.Warnf("ListService.Post() Could not post list: %s", err)
		return list, err
	}
	for index, element := range list.List {
		var listElement models.ListElement
		listElement.Element = element
		listElement.Deleted = false
		listElement.Position = index
		listElement.UserID = userID
		err = s.ListElementService.Post(listElement, list.ID)
	}
	for index, element := range list.DoneItems {
		var listElement models.ListElement
		listElement.Element = element
		listElement.Deleted = true
		listElement.Position = index
		listElement.UserID = userID
		err = s.ListElementService.Post(listElement, list.ID)
	}
	return list, err
}

func (s *Service) UpdateList(id string, list *models.List, userID string) (models.List, error) {
	newList, err := s.listRepository.updateList(list, id, userID)
	if err != nil {
		log.Warnf("ListService.Update() Could not update list: %s", err)
		return newList, err
	}
	err = s.ListElementService.DeleteAllFromList(list.ID, userID)
	if err != nil {
		log.Warnf("ListService.UpdateList() Could not delete list_elements: %s", err)
		return newList, err
	}
	for index, element := range list.List {
		var listElement models.ListElement
		listElement.Element = element
		listElement.Deleted = false
		listElement.Position = index
		listElement.UserID = userID
		err = s.ListElementService.Post(listElement, list.ID)
	}
	for index, element := range list.DoneItems {
		var listElement models.ListElement
		listElement.Element = element
		listElement.Deleted = true
		listElement.Position = index
		listElement.UserID = userID
		err = s.ListElementService.Post(listElement, list.ID)
	}
	return newList, err
}

func (s *Service) DeleteList(ID string, userID string) error {
	var list models.List
	list.ID = ID
	err := s.listRepository.DeleteList(list, ID, userID)
	if err != nil {
		log.Warnf("ListService.Delete() Could not delete list: %s", err)
		return err
	}
	return err
}
