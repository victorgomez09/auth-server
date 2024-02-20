package services

import (
	"fmt"
	"net/http"

	"github.com/ESMO-ENTERPRISE/auth-server/database"
	"github.com/ESMO-ENTERPRISE/auth-server/dtos"
	"github.com/ESMO-ENTERPRISE/auth-server/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Conn *database.Connector
}

func (a *Auth) RegisterWithEmailAndPassword(c echo.Context) error {
	u := new(dtos.RegisterPayload)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Errorf(err.Error())
		return c.JSON(http.StatusBadRequest, "internal server error")
	}

	user := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: hashedPassword,
	}

	if userExist := a.Conn.DB.Where("email = ?", user.Email).First(&user); userExist != nil {
		return c.JSON(http.StatusBadRequest, "user already exists")
	}

	a.Conn.DB.Create(&user)

	return c.JSON(http.StatusCreated, map[string]string{
		"data": "user " + user.Name + " created",
	})
}
