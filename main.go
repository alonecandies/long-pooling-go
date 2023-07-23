package main

import (
	"net/http"
	"time"

	"github.com/alonecandies/long-pooling-go/inject"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	injector *inject.Injector
}

func NewServer() *Server {
	deployTime := time.Now()
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Locale"},
		AllowAllOrigins: true,
	}))
	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"deployTime": deployTime,
			"timestamp":  time.Now(),
		})
	})

	injector := inject.NewInjector()

	server := &Server{
		engine:   engine,
		injector: injector,
	}
	server.registerUserRoutes()
	return server
}

func (s *Server) registerUserRoutes() {
	longPoolingAPI := s.injector.ProvideLongPoolingApi()

	v1 := s.engine.Group("v1")
	{
		v1.GET("/longPooling", longPoolingAPI.GetLongPooling)
	}

}

func (s *Server) Serve() error {
	return s.engine.Run(":8080")
}

func main() {
	server := NewServer()
	server.Serve()
}
