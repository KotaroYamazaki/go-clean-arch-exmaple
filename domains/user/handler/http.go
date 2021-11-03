package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user"
	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// Get godoc
// @Summary Get a user
// @Description get user by ID
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} user.User
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} int
// @Failure 500 {object} int
// @Failure default {object} int
// @Router /users/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	res, err := h.uc.Get(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// SIgnup godoc
// @Summary Signup
// @Description signup
// @ID singup
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} user.User
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} int
// @Failure 500 {object} int
// @Failure default {object} int
// @Router /signup [post]
func (h *Handler) Signup(c *gin.Context) {
	params := &user.SignupParams{}
	if err := c.BindJSON(&params); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := params.Validate(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.uc.Signup(c, params); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
