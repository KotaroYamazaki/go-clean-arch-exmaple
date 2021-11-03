package handler

import (
	"net/http"
	"strconv"

	"github.com/KotaroYamazaki/go-cleanarchtecture/domains/user"
	"github.com/KotaroYamazaki/go-cleanarchtecture/domains/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

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

func (h *Handler) Signup(c *gin.Context) {
	params := &user.SignupParams{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := params.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.uc.Signup(c, params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
