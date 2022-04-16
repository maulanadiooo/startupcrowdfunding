package handler

import (
	"net/http"
	"startup/helper"
	"startup/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// menangkap input dari user
	// map input tersebut ke struct RegisterUserInput
	// hasil struct di parsing sebagai parameter service

	var input user.RegisteruserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormateValidationError(err)

		errorMessage := gin.H{"errors:": errors}

		response := helper.APIResponse("Account has been failed to register", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Account has been failed to register", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "hahahahahahaha")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// masukkan input oleh user [email dan password]
	// input ditangkap oleh handler
	// mapping dari input user ke input struct
	// input struct di pasrsing ke service
	// di service mencari degan bantuan repositroy user dengan email yang di input
	// jika ketemu cocokkan passwordnya

	var input user.LoginInput

	err := c.ShouldBindJSON(input)

	if err != nil {

		errors := helper.FormateValidationError(err)

		errorMessage := gin.H{"errors:": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors:": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	formatter := user.FormatUser(loggedinUser, "jajajajajaja")

	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
