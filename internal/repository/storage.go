package repository

import (
	"database/sql"

	"github.com/VadimBorzenkov/gw-exchanger/internal/models"
	"github.com/sirupsen/logrus"
)

type ExchangeRateStorage interface {
	GetExchangeRates() (map[string]float64, error)
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
}

type PostgresStorage struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewPostgresStorage(db *sql.DB, logger *logrus.Logger) PostgresStorage {
	return PostgresStorage{
		db:     db,
		logger: logger,
	}
}

func (s *PostgresStorage) GetExchangeRates() ([]models.ExchangeRate, error) {
	rates := []models.ExchangeRate{}

	rows, err := s.db.Query("SELECT currency, rate, updated_at FROM exchange_rates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rate models.ExchangeRate
		if err := rows.Scan(&rate.CurrencyCode, &rate.Rate, &rate.UpdatedAt); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}

	return rates, nil
}

func (s *PostgresStorage) GetExchangeRate(currency string) (*models.ExchangeRate, error) {
	var rate models.ExchangeRate
	query := "SELECT currency, rate, updated_at FROM exchange_rates WHERE currency = $1"

	err := s.db.QueryRow(query, currency).Scan(&rate.CurrencyCode, &rate.Rate, &rate.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Курс не найден
		}
		return nil, err
	}

	return &rate, nil
}
