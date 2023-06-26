package main

import (
	"course_beginner/config"
	"course_beginner/internal/handler"
	"course_beginner/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	uc := usecase.New()
	h := handler.New(uc)
	//_ = h
	r := gin.New()
	h.Register(r)

	r.GET("/", hello)

	log.Println("server started listening on port " + cfg.Port)
	err = r.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
