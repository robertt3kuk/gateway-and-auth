package app

import (
	"fmt"
	"log"
	"net"

	"smotri-auth/internal/delivery"
	"smotri-auth/internal/service"
	"smotri-auth/internal/service/repo"
	"smotri-auth/pkg/postgres"

	"google.golang.org/grpc"

	"smotri-auth/config"
	"smotri-auth/internal/pb"
	"smotri-auth/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	//Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	repos := repo.NewRepo(pg)
	usecase := service.NewUseCase(repos)
	//collection of UseCases for NewRouter
	server := delivery.Server{usecase}
	listener, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal(fmt.Errorf("app listener broken", err))
	}

	/*s := delivery.Server{usecase, pb.UnimplementedAuthServiceServer{}}
	pb.AuthServiceServer().Validate()*/

	var opt []grpc.ServerOption
	grpcServ := grpc.NewServer(opt...)

	pb.RegisterAuthServiceServer(grpcServ, &server)
	if err := grpcServ.Serve(listener); err != nil {
		l.Fatal("Failed to server:", err)
	}
	/*
		// Waiting signal
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

		select {
		case s := <-interrupt:
			l.Info("app - Run - signal: " + s.String())
		case err = <-httpServer.Notify():
			l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		}

		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}*/

}
