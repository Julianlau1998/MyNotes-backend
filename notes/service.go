package notes

import (
	"notesBackend/models"

	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	noteRepository Repository
}

func NewService(noteRepository Repository) Service {
	return Service{noteRepository: noteRepository}
}

func (s *Service) GetNotes(userID string) ([]models.Note, error) {
	notes, err := s.noteRepository.GetNotes(userID)
	if err != nil {
		log.Warnf("NotesService.GetNotes(): Could not get notes: %s", err)
		return notes, err
	}
	return notes, err
}

func (s *Service) GetByFolder(folderID string, userID string) ([]models.Note, error) {
	notes, err := s.noteRepository.GetByFolder(folderID, userID)
	if err != nil {
		log.Warnf("NotesService.GetByFolder(): Could not get notes by folder: %s", err)
		return notes, err
	}
	return notes, nil
}

func (s *Service) GetNoteById(id string, userID string) (models.Note, error) {
	note, err := s.noteRepository.GetNoteById(id, userID)
	if err != nil {
		log.Warnf("NotesService.GetNoteById(): Could not get note by Id: %s", err)
		return note, err
	}
	return note, err
}

func (s *Service) Post(note *models.Note) (*models.Note, error) {
	id, err := uuid.NewV4()
	note.ID = id.String()
	if note.FolderID == "" {
		note.FolderID = "00000000-0000-0000-0000-000000000000"
	}
	note, err = s.noteRepository.Post(note)
	if err != nil {
		log.Warnf("NotesService.Post(): Could not post note: %s", err)
		return note, err
	}
	return note, err
}

func (s *Service) updateNote(ID string, note *models.Note, userID string) (models.Note, error) {
	note.ID = ID
	note.UserID = userID
	newNote, err := s.noteRepository.updateNote(note)
	if err != nil {
		log.Warnf("NotesService.UpdateNote(): Could not update note: %s", err)
		return newNote, err
	}
	return newNote, err
}

func (s *Service) DeleteNote(ID string, userID string) error {
	var note models.Note
	note.ID = ID
	note.UserID = userID
	err := s.noteRepository.deleteNote(note)
	if err != nil {
		log.Warnf("NotesService.DeleteNote(): Could not delete note: %s", err)
		return err
	}
	return err
}
