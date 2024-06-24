// ToDo :
// 1. Mendeklarasikan nama package pada file server.go
// 2. Mendeklarasikan struct bernama Server
// 3. Mendeklarasikan function initRoute
// 4. Membuat function Run
// 5. Memuat constructor bernama NewServer

package main

import (
	"fmt"
	"simple-clean-architecture/config"
	"simple-clean-architecture/controller"
	"simple-clean-architecture/repository"
	"simple-clean-architecture/usecase"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	bookUC usecase.BookUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")

	controller.NewBookController(s.bookUC, rg).Route()
	controller.NewAuthController(rg).Route()
}

func (s *Server) Run() {
	s.initRoute()

	err := s.engine.Run(s.host)

	if err != nil {
		panic(fmt.Errorf("server is not running on host %s, because error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)

	db, err := sql.Open(cfg.Driver, dsn)

	if err != nil {
		panic("connection error")
	}

	bookRepo := repository.NewBookRepository(db)
	bookUseCase := usecase.NewBookUseCase(bookRepo)

	engine := gin.Default()

	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		bookUC: bookUseCase,
		engine: engine,
		host:   host,
	}
}
