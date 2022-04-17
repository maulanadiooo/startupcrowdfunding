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

	err := c.ShouldBindJSON(&input)

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

func (h *userHandler) CheckEmailAvailable(c *gin.Context) {
	// ada input email dari user,
	// input email di mapping ke struct input
	// struct input di passing ke service
	// service akan memanggil repository untuk mengecek apakah email sudah ada atau belum
	// repository - db

	var input user.CheckEmailInnput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormateValidationError(err)

		errorMessage := gin.H{"errors:": errors}

		response := helper.APIResponse("Checking email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsemailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors:": "Server error"}

		response := helper.APIResponse("Checking email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	// formatter := user.FormatUser(loggedinUser, "jajajajajaja")

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email is registered"

	if isEmailAvailable == true {
		metaMessage = "Email is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
