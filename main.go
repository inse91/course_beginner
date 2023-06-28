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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("hello world1"))
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

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
