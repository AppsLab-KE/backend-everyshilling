package server

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/repositories"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/services"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/storage"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/platform"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/routes/handlers"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
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
	postgresClient, err := platform.NewDBServiceClient(s.cfg.DB)
	if err != nil {
		log.Panic("error:", err)
	}

	// Create User storage
	dbStorage := storage.NewDBStorage(postgresClient)

	// create  repository
	accountsRepo := repositories.NewAccountRepository(dbStorage)
	exchangeRepo := repositories.NewExchangeRepository(dbStorage)
	tradeRepo := repositories.NewTradeRepository(dbStorage)
	transactionsRepo := repositories.NewTransactionRepository(dbStorage)

	// create services
	accountsService := services.NewAccountsService(accountsRepo)
	exchangeService := services.NewExchangeService(exchangeRepo)
	tradeService := services.NewTradeService(tradeRepo, accountsRepo, exchangeRepo, transactionsRepo)
	transactionService := services.NewTransactionService(transactionsRepo, accountsRepo)

	// create handler
	grpcHandler := handlers.NewHandler(exchangeService, accountsService, tradeService, transactionService)

	// run server
	lis, err := net.Listen("tcp", ":"+s.cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer()
	exchange.RegisterExchangeServiceServer(grpcServer, grpcHandler)

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
