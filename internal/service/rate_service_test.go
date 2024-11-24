package service_test

import (
	"errors"
	"testing"

	"github.com/VadimBorzenkov/gw-exchanger/internal/config"
	"github.com/VadimBorzenkov/gw-exchanger/internal/repository/mocks"
	"github.com/VadimBorzenkov/gw-exchanger/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestExchangeService_GetAllRates(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockExchangeRateStorage(ctrl)
	mockLogger := logrus.New()
	cfg := &config.Config{}

	expectedRates := map[string]float64{"USD": 1.0, "EUR": 0.85}
	mockStorage.EXPECT().GetExchangeRates().Return(expectedRates, nil)

	service := service.NewExchangeService(mockStorage, mockLogger, cfg)

	rates, err := service.GetAllRates()

	assert.NoError(t, err)
	assert.Equal(t, expectedRates, rates)
}

func TestExchangeService_GetAllRates_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockExchangeRateStorage(ctrl)
	mockLogger := logrus.New()
	cfg := &config.Config{}

	mockStorage.EXPECT().GetExchangeRates().Return(nil, errors.New("storage error"))

	service := service.NewExchangeService(mockStorage, mockLogger, cfg)

	rates, err := service.GetAllRates()

	assert.Error(t, err)
	assert.Nil(t, rates)
}

func TestExchangeService_GetRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockExchangeRateStorage(ctrl)
	mockLogger := logrus.New()
	cfg := &config.Config{}

	expectedRate := 0.85
	mockStorage.EXPECT().GetExchangeRate("EUR").Return(expectedRate, nil)

	service := service.NewExchangeService(mockStorage, mockLogger, cfg)

	rate, err := service.GetRate("EUR")

	assert.NoError(t, err)
	assert.Equal(t, expectedRate, rate)
}

func TestExchangeService_GetRate_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockExchangeRateStorage(ctrl)
	mockLogger := logrus.New()
	cfg := &config.Config{}

	mockStorage.EXPECT().GetExchangeRate("EUR").Return(0.0, errors.New("rate not found"))

	service := service.NewExchangeService(mockStorage, mockLogger, cfg)

	rate, err := service.GetRate("EUR")

	assert.Error(t, err)
	assert.Equal(t, 0.0, rate)
}

func TestExchangeService_ConvertCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockExchangeRateStorage(ctrl)
	mockLogger := logrus.New()
	cfg := &config.Config{}

	mockStorage.EXPECT().GetExchangeRate("USD").Return(1.0, nil)
	mockStorage.EXPECT().GetExchangeRate("EUR").Return(0.85, nil)

	service := service.NewExchangeService(mockStorage, mockLogger, cfg)

	amount := 100.0
	expectedAmount := 85.0

	result, err := service.ConvertCurrency("USD", "EUR", amount)

	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, result)
}

func TestExchangeService_ConvertCurrency_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockExchangeRateStorage(ctrl)
	mockLogger := logrus.New()
	cfg := &config.Config{}

	mockStorage.EXPECT().GetExchangeRate("USD").Return(0.0, errors.New("rate not found"))

	service := service.NewExchangeService(mockStorage, mockLogger, cfg)

	result, err := service.ConvertCurrency("USD", "EUR", 100.0)

	assert.Error(t, err)
	assert.Equal(t, 0.0, result)
}
