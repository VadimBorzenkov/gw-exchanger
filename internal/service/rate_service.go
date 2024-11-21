package service

import (
	"github.com/VadimBorzenkov/gw-exchanger/internal/config"
	"github.com/VadimBorzenkov/gw-exchanger/internal/models"
	"github.com/VadimBorzenkov/gw-exchanger/internal/repository"
	"github.com/sirupsen/logrus"
)

type ExchangeService interface {
	GetAllRates() ([]models.ExchangeRate, error)
	GetRate(currency string) (*models.ExchangeRate, error)
	ConvertCurrency(fromCurrency, toCurrency string, amount float64) (float64, error)
}

type exchangeService struct {
	storage repository.PostgresStorage
	logger  *logrus.Logger
	cfg     *config.Config
}

func NewExchangeService(storage repository.PostgresStorage, logger *logrus.Logger, cfg *config.Config) ExchangeService {
	return &exchangeService{
		storage: storage,
		logger:  logger,
		cfg:     cfg,
	}
}

// GetAllRates - получение всех курсов обмена
func (s *exchangeService) GetAllRates() ([]models.ExchangeRate, error) {
	return s.storage.GetExchangeRates()
}

// GetRate - получение курса обмена для конкретной валюты
func (s *exchangeService) GetRate(currency string) (*models.ExchangeRate, error) {
	return s.storage.GetExchangeRate(currency)
}

// ConvertCurrency - конвертация валюты
func (s *exchangeService) ConvertCurrency(fromCurrency, toCurrency string, amount float64) (float64, error) {
	fromRate, err := s.storage.GetExchangeRate(fromCurrency)
	if err != nil {
		s.logger.Errorf("Failed to get rate for %s: %v", fromCurrency, err)
		return 0, err
	}

	toRate, err := s.storage.GetExchangeRate(toCurrency)
	if err != nil {
		s.logger.Errorf("Failed to get rate for %s: %v", toCurrency, err)
		return 0, err
	}

	// Конвертация валюты
	convertedAmount := (amount / fromRate.Rate) * toRate.Rate
	return convertedAmount, nil
}
