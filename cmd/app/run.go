package app

import (
	"log"

	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user/handler"
	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user/repository"
	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user/usecase"
	"github.com/KotaroYamazaki/go-clean-arch-example/infra"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	userRepo := repository.New(infra.DB)
	userUC := usecase.New(userRepo)
	userH := handler.New(userUC)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	},
	)
	r.GET("/users/:id", userH.Get)
	r.POST("/singup", userH.Signup)

	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
