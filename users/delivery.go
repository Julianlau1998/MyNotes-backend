// package users

// import (
// 	"fmt"
// 	"net/http"
// 	"notesBackend/models"

// 	"github.com/labstack/echo/v4"
// )

// type Delivery struct {
// 	userService Service
// }

// func NewDelivery(userService Service) Delivery {
// 	return Delivery{userService: userService}
// }

// func (d *Delivery) GetAll(c echo.Context) error {
// 	username := c.Request().Header.Get("username")
// 	password := c.Request().Header.Get("password")
// 	lists, err := d.userService.GetUser(username, password)
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, lists)
// }

// func (d *Delivery) Post(c echo.Context) error {
// 	requestBody := new(models.User)
// 	if err := c.Bind(requestBody); err != nil {
// 		fmt.Print(err)
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}

// 	user, err := d.userService.Post(requestBody)
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.String(http.StatusOK, user.ID)
// }