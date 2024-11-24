# gw-exchanger

ge-exchanger — это сервис на Go, предоставляющий функционал для работы с курсами валют. Сервис реализует хранение курсов валют в PostgreSQL и предоставляет API через gRPC для их получения. Проект поддерживает контейнеризацию с использованием Docker и Docker Compose.

---

## Основные функции

- Хранение курсов валют в базе данных PostgreSQL.
- Получение текущих курсов валют.
- Получение курса для конкретной валюты.
- Автоматическое применение миграций для настройки схемы базы данных.
- Поддержка gRPC для взаимодействия с клиентами.

---

## Стек технологий

- **Go**: язык программирования для разработки высокопроизводительных приложений.
- **gRPC**: протокол удаленного вызова процедур.
- **PostgreSQL**: база данных для хранения курсов валют.
- **Docker**: контейнеризация приложений.
- **Docker Compose**: оркестрация контейнеров.
- **Logrus**: логирование.

---

## Установка и запуск

### Клонирование репозитория
1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/VadimBorzenkov/gw-exchanger.git
   cd gw-exchanger

2. Скопируйте файл конфигурации:
    ````bash
    cp example.env .env

3. Настройте файл .env:
    -DB_HOST: адрес базы данных (например, localhost или db при использовании Docker).
    -DB_USER, DB_PASSWORD, DB_NAME: учетные данные PostgreSQL.
    -LOG_LEVEL: уровень логирования (debug, info, warn, error).
    -LOG_FORMAT: формат логов (text или json).


### Запуск через Docker
1. Убедитесь, что переменная DB_HOST установлена как db в .env.

2. Запустите сервисы:
    ````bash
    docker-compose up --build

3. Проверьте, что контейнеры запущены:
    ````bash
    docker ps

### Локальный запуск
1. Установите зависимости:
    ````bash
    go mod tidy

2. Запустите PostgreSQL локально или измените DB_HOST в .env на localhost.

3. Запустите приложение:
    ```bash
    Копировать код
    go run ./cmd/main.go

4. Приложение запустится на порту 50051 (или указанном в .env).