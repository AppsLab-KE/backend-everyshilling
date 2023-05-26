package server

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/config"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg config.Config
}

func (s *Server) Run() {
	log.Infof("DATABASE GRPC server initialising")

	// create database client
	postgresClient, err := postgres.NewClient(s.cfg.Postgres)
	if err != nil {
		log.Panic("error:", err)
	}

	// Create User storage
	userStorage := storage.NewUserStorage(postgresClient)
	userCache := storage.NewUserCacheStorage()

	// rates storage
	ratesStorage := storage.NewRatesPostgresStorage(postgresClient)

	// create  repository
	userRepo := repository.NewUserRepo(userStorage, userCache)
	ratesRepo := repository.NewRatesRepo(ratesStorage)

	// create handler
	grpcHandler := handlers.NewHandler(userRepo, ratesRepo)

	// run server
	lis, err := net.Listen("tcp", ":"+s.cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer()
	db.RegisterDbServiceServer(grpcServer, grpcHandler)

	// Run the server

	// wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("cannot start apps server:", err)
			return
		}
	}()

	<-quit

}

func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}
