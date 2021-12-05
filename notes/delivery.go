package notes

import (
	"fmt"
	"net/http"
	"notesBackend/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Delivery struct {
	noteService Service
}

func NewDelivery(noteService Service) Delivery {
	return Delivery{noteService: noteService}
}

func (d *Delivery) GetAll(c echo.Context) error {
	userID := c.Request().Context().Value("currentUser").(jwt.MapClaims)["sub"].(string)
	notes, err := d.noteService.GetNotes(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notes)
}

func (d *Delivery) GetByFolder(c echo.Context) error {
	userID := c.Request().Context().Value("currentUser").(jwt.MapClaims)["sub"].(string)
	folderID := c.Request().Header.Get("folderID")
	notes, err := d.noteService.GetByFolder(folderID, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notes)
}

func (d *Delivery) GetById(c echo.Context) error {
	userID := c.Request().Context().Value("currentUser").(jwt.MapClaims)["sub"].(string)
	id := c.Param("id")
	note, err := d.noteService.GetNoteById(id, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if note.ID == "" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, note)
}

func (d *Delivery) Post(c echo.Context) error {
	userID := c.Request().Context().Value("currentUser").(jwt.MapClaims)["sub"].(string)
	requestBody := new(models.Note)
	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	requestBody.UserID = userID

	note, err := d.noteService.Post(requestBody)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, note.ID)
}

func (d *Delivery) Update(c echo.Context) (err error) {
	userID := c.Request().Context().Value("currentUser").(jwt.MapClaims)["sub"].(string)
	ID := c.Param("id")
	requestBody := new(models.Note)
	if err = c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	requestBody.UserID = userID
	fmt.Print(requestBody)

	note, err := d.noteService.updateNote(ID, requestBody, userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if note.ID == "" {
		return c.String(http.StatusForbidden, "Not Authorized")
	}
	return c.JSON(http.StatusOK, note)
}

func (d *Delivery) Delete(c echo.Context) (err error) {
	userID := c.Request().Context().Value("currentUser").(jwt.MapClaims)["sub"].(string)
	id := c.Param("id")

	err = d.noteService.DeleteNote(id, userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, err)
}
