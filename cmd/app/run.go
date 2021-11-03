package app

import (
	"log"

	"github.com/KotaroYamazaki/go-cleanarchtecture/domains/user/handler"
	"github.com/KotaroYamazaki/go-cleanarchtecture/domains/user/repository"
	"github.com/KotaroYamazaki/go-cleanarchtecture/domains/user/usecase"
	"github.com/KotaroYamazaki/go-cleanarchtecture/infra"
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
	r.POST("/users/singup", userH.Signup)

	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
