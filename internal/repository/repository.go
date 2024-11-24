package repository

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type ExchangeRateStorage interface {
	GetExchangeRates() (map[string]float64, error)
	GetExchangeRate(currency string) (float64, error)
}

type PostgresStorage struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewPostgresStorage(db *sql.DB, logger *logrus.Logger) *PostgresStorage {
	return &PostgresStorage{
		db:     db,
		logger: logger,
	}
}

func (s *PostgresStorage) GetExchangeRates() (map[string]float64, error) {
	// Инициализируем мапу для хранения результатов
	rates := make(map[string]float64)

	// Выполняем SQL-запрос для получения данных
	rows, err := s.db.Query("SELECT currency_code, rate FROM exchange_rates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Проходим по каждой строке результата
	for rows.Next() {
		var currencyCode string
		var rate float64

		// Считываем данные из строки
		if err := rows.Scan(&currencyCode, &rate); err != nil {
			return nil, err
		}

		// Добавляем данные в мапу
		rates[currencyCode] = rate
	}

	// Проверяем на ошибки после обработки строк
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rates, nil
}

func (s *PostgresStorage) GetExchangeRate(currency_code string) (float64, error) {
	var rate float64
	query := "SELECT rate FROM exchange_rates WHERE currency_code = $1"

	err := s.db.QueryRow(query, currency_code).Scan(&rate)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // Курс не найден
		}
		return 0, err
	}

	return rate, nil
}
