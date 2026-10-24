# WStore

Pet-проект интернет-магазина: fullstack-приложение с REST API на Go и SPA на React. Сейчас реализованы аутентификация пользователей, личный кабинет и витрина каталога (UI с mock-данными); модуль продаж и корзина — в разработке.

## Возможности

- **Аутентификация** — регистрация и вход по email/паролю (Argon2), JWT access token + refresh token в HttpOnly cookie
- **Пользователи** — список пользователей, профиль (`/user/me`), аватар
- **Frontend** — главная страница с каталогом, вход/регистрация, личный кабинет, настройки
- **i18n** — русский и английский интерфейс
- **API-документация** — Swagger UI на `/swagger/`
- **Google OAuth** — код подготовлен, временно отключён (`//go:build ignore`)

## Стек технологий

| Слой | Технологии |
|------|------------|
| Backend | Go 1.25, `net/http`, PostgreSQL, JWT, Argon2, Logrus, Swaggo |
| Frontend | React 19, Vite 8, React Router 7, ESLint |
| База данных | PostgreSQL (схема `accounts`, хранимые процедуры) |
| Инфраструктура | HTTPS (TLS / mkcert), env-конфигурация |

## Структура проекта

```
wstore/
├── backend/          # Go HTTP API
│   ├── cmd/wstore/   # точка входа, Swagger docs
│   └── internal/     # handlers, storage, config, app
├── frontend/         # React SPA (Vite)
├── migrations/       # SQL-миграции PostgreSQL
└── migrations.sql    # скрипт применения миграций
```

## Требования

- Go 1.25+
- Node.js 20+
- PostgreSQL 14+

## Быстрый старт

### 1. База данных

```bash
psql -U postgres -f migrations.sql
```

> Файл `migrations/roles.sql` нужен для настройки ролей БД и не включён в репозиторий — создайте его локально перед запуском миграций.

### 2. Backend

Создайте `.env` в корне или в `backend/` (загружается через `godotenv/autoload`):

```env
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_DATABASE=wstore

SERVER_JWT_SECRET=your-secret-key
SERVER_TLS=true
SERVER_CERT_FILE=certificates/server.crt
SERVER_KEY_FILE=certificates/server.key

FRONTEND_URL=https://localhost:8000
LOG_LEVEL=info
```

API: `https://localhost:8080/api/`  
Swagger: `https://localhost:8080/swagger/`

### 3. Frontend

```bash
cd frontend
npm install
npm run dev
```

Приложение: `https://localhost:8000`  
Запросы к `/api` проксируются на backend.

## API (основные эндпоинты)

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/service/health` | Health check |
| POST | `/api/auth/email/signup` | Регистрация |
| POST | `/api/auth/email/login` | Вход (возвращает access token) |
| GET | `/api/auth/refresh` | Обновление access token (cookie) |
| POST | `/api/auth/logout` | Выход |
| GET | `/api/users` | Список пользователей 🔒 |
| GET | `/api/user/{id}` | Профиль пользователя (`me` — текущий) 🔒 |
| GET | `/api/user/profile-image/{id}` | Аватар 🔒 |

🔒 — требуется заголовок `Authorization: Bearer <access_token>`

## Архитектура backend

- HTTP-сервер на стандартной библиотеке Go без фреймворка
- Маршруты группируются по доменам: `auth`, `users`, `service`
- Доступ к PostgreSQL через `database/sql` и хранимые функции
- Единый формат ответа: `{ success, data, error }`
- JWT middleware для защищённых маршрутов

## Архитектура frontend

- React SPA с React Router
- `AuthContext` — хранение access token, автоматический refresh
- `LanguageContext` — переключение RU/EN
- API-клиент с прокси через Vite dev server

## Статус проекта

Проект в активной разработке. Каталог товаров на главной странице использует статические mock-данные; интеграция с API продаж планируется.

## Лицензия

Apache 2.0
