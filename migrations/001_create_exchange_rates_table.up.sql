CREATE TABLE exchange_rates (
    id SERIAL PRIMARY KEY,
    currency_code VARCHAR(3) NOT NULL UNIQUE, -- Код валюты (например, USD, EUR)
    rate DECIMAL(20, 8) NOT NULL, -- Курс валюты
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Время последнего обновления
);