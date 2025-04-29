# University System

## Описание

**University System** — это полнофункциональное веб-приложение для управления университетом. Система поддерживает роли пользователей (студент, преподаватель, менеджер, администратор), регистрацию, аутентификацию (JWT), CRUD для пользователей, студентов, преподавателей, менеджеров, курсов, назначение преподавателей на курсы, зачисление студентов, выставление оценок и многое другое.

## Технологии
- Язык: Go (Golang)
- Веб-фреймворк: Gin
- ORM: sqlx
- БД: PostgreSQL
- Документация API: Swagger (swaggo)
- Docker (мультистейдж-сборка для оптимизации размера образа)

## Быстрый старт

### 1. Клонирование репозитория
```bash
git clone https://github.com/NurtileuRakhat/university_system
```

### 2. Настройка переменных окружения
Создайте файл `.env` или экспортируйте переменные:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=university
JWT_SECRET=your_jwt_secret
```

### 3. Сборка и запуск через Docker
```bash
docker build -t university-system .
docker run -p 8080:8080 --env-file .env university-system
```
И надо будет обновить docker-compose.yml, чтобы он использовал новые переменные окружения.

### 4. Локальный запуск (без Docker)
- Установите Go >= 1.18
- Установите PostgreSQL и создайте БД
- Установите зависимости:
  ```bash
  go mod tidy
  ```
- Запустите миграции (создание таблиц):
  ```bash
  go run ./cmd/migrate.go
  ```
- Запустите сервер:
  ```bash
  go run ./cmd/main.go
  ```
### 5. Test
 ```go test ./internal/university/services/...
```

### Swagger init
```swag init -g cmd/university/main.go --parseDependency --parseInternal```


## Основные возможности
- Регистрация и аутентификация пользователей (JWT)
- CRUD для студентов, преподавателей, менеджеров, пользователей
- CRUD для курсов
- Назначение преподавателей на курсы
- Зачисление студентов на курсы
- Выставление и просмотр оценок
- Swagger-документация по адресу `/swagger/index.html`

## Примеры curl для основных сценариев

### Регистрация студента (от имени администратора)
```bash
curl -X POST http://localhost:8080/students \
  -H "Authorization: Bearer <ADMIN_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "stud_test",
    "password": "pass1234",
    "firstname": "nurtileu",
    "lastname": "Студент",
    "email": "stud_test@mail.ru",
    "birthdate": "2004-09-01",
    "student_year": 2,
    "faculty": "ФИТ"
  }'
```

### Логин пользователя
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "stud_test",
    "password": "pass1234"
  }'
```

### Примеры для других ролей и операций см. в swagger или документации к API.

## Swagger
- Документация Swagger доступна по адресу: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Для обновления документации после изменения аннотаций:
  ```bash
  swag init
  ```

## Структура проекта
```
internal/
  domain/           # Модели и интерфейсы
  infrastructure/   # Репозитории (доступ к БД)
  university/
    controllers/    # Контроллеры (обработчики HTTP)
    services/       # Бизнес-логика
  routes/           # Роутинг (маршруты)
pkg/
  databases/        # Подключение и миграции БД
cmd/
  main.go           # Точка входа
  migrate.go        # Миграции
Dockerfile          # Docker-сборка
.env.example        # Пример переменных окружения
```

## Советы
- Для тестирования API удобно использовать Postman или Swagger UI.
- Все защищённые маршруты требуют Bearer JWT токен.
- После регистрации пользователя выполните логин для получения токена.

## Контакты
- Автор: Nurtileu Rakhat
- Telegram: @nurtileurk
- Email: raaaaahaaat@gmail.com

---

