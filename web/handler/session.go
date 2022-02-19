package handler

import (
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService: userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *sessionHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)

	if err != nil {

		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	user, err := h.userService.Login(input)

	if err != nil || user.Role != "admin" {
		fmt.Println(err.Error())
		c.Redirect(http.StatusFound, "/login")
		return
	}

	session := sessions.Default(c)

	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/users")

}

func (h *sessionHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
