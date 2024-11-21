package app

import (
	"net"

	"github.com/VadimBorzenkov/gw-exchanger/internal/config"
	"github.com/VadimBorzenkov/gw-exchanger/internal/db"
	"github.com/VadimBorzenkov/gw-exchanger/internal/repository"
	"github.com/VadimBorzenkov/gw-exchanger/internal/service"
	"github.com/VadimBorzenkov/gw-exchanger/pkg/logger"
	"github.com/VadimBorzenkov/gw-exchanger/pkg/migrator"
	pb "github.com/VadimBorzenkov/proto-exchange/exchange"
	"google.golang.org/grpc"
)

// server - структура для gRPC сервера
type server struct {
	pb.UnimplementedExchangeServiceServer
	service service.ExchangeService
}

// Run - запуск gRPC сервера
func Run() {
	log := logger.InitLogger()

	// Загрузка конфигурации
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Инициализация подключения к базе данных
	dbase, err := db.Init(config)
	if err != nil {
		log.Fatalf("Could not initialize DB connection: %v", err)
	}
	defer func() {
		if err := db.Close(dbase); err != nil {
			log.Errorf("Failed to close database: %v", err)
		}
	}()

	if err := migrator.RunDatabaseMigrations(dbase); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Создание объекта для работы с хранилищем
	repository := repository.NewPostgresStorage(dbase, log)

	// Создание сервиса для обмена
	exchangeService := service.NewExchangeService(repository, log, config)

	// Создание нового gRPC сервера
	grpcServer := grpc.NewServer()

	// Регистрация сервиса
	pb.RegisterExchangeServiceServer(grpcServer, &server{service: exchangeService})

	// Ожидание на подключение по порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("gRPC server is running on port :50051")

	// Запуск gRPC сервера
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
