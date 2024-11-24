package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetExchangeRates(t *testing.T) {
	// Создаем мок-базу данных и объект sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Логгер
	logger := logrus.New()

	// Создаем экземпляр репозитория
	repo := NewPostgresStorage(db, logger)

	// Настраиваем ожидаемые строки результата
	rows := sqlmock.NewRows([]string{"currency_code", "rate"}).
		AddRow("USD", 100.0).
		AddRow("EUR", 105.0)

	mock.ExpectQuery("SELECT currency_code, rate FROM exchange_rates").
		WillReturnRows(rows)

	// Вызываем метод репозитория
	rates, err := repo.GetExchangeRates()

	// Проверяем результаты
	assert.NoError(t, err)
	assert.Equal(t, map[string]float64{"USD": 100.0, "EUR": 105.0}, rates)

	// Убедитесь, что все ожидания выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetExchangeRate(t *testing.T) {
	// Создаем мок-базу данных и объект sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Логгер
	logger := logrus.New()

	// Создаем экземпляр репозитория
	repo := NewPostgresStorage(db, logger)

	// Настраиваем ожидаемое поведение для запроса
	mock.ExpectQuery("SELECT rate FROM exchange_rates WHERE currency_code = \\$1").
		WithArgs("USD").
		WillReturnRows(sqlmock.NewRows([]string{"rate"}).AddRow(1.0))

	// Вызываем метод репозитория
	rate, err := repo.GetExchangeRate("USD")

	// Проверяем результаты
	assert.NoError(t, err)
	assert.Equal(t, 1.0, rate)

	// Убедитесь, что все ожидания выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetExchangeRate_NotFound(t *testing.T) {
	// Создаем мок-базу данных и объект sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Логгер
	logger := logrus.New()

	// Создаем экземпляр репозитория
	repo := NewPostgresStorage(db, logger)

	// Настраиваем ожидаемое поведение для случая, когда курс не найден
	mock.ExpectQuery("SELECT rate FROM exchange_rates WHERE currency_code = \\$1").
		WithArgs("JPY").
		WillReturnError(sql.ErrNoRows)

	// Вызываем метод репозитория
	rate, err := repo.GetExchangeRate("JPY")

	// Проверяем результаты
	assert.NoError(t, err)
	assert.Equal(t, 0.0, rate) // Ожидаем, что для несуществующей валюты возвращается 0

	// Убедитесь, что все ожидания выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}
