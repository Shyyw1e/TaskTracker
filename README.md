# TaskTracker (REST API + Telegram Bot)

TaskTracker — это backend-приложение на Go, в котором реализован сервис для управления задачами. Пользователь может работать с задачами как через REST API, так и через Telegram-бота. Все данные хранятся в базе (PostgreSQL или SQLite) и привязаны к конкретному пользователю.

---

## 🚀 Возможности

- Создание, просмотр, редактирование и удаление задач
- Telegram-бот с поддержкой основных команд (`/new`, `/list`, `/done`, `/delete`)
- Хранение задач в базе данных
- Привязка задач к Telegram-пользователю
- REST API с CRUD-операциями
- Конфигурация через `.env`
- Логирование всех операций

---

## ⚙️ Стек

- **Язык**: Go
- **API**: `net/http` или `chi`
- **База данных**: PostgreSQL или SQLite
- **ORM**: `sqlx`, `gorm` или `bun`
- **Telegram Bot API**: [`go-telegram-bot-api`](https://github.com/go-telegram-bot-api/telegram-bot-api)
- **Конфигурация**: `godotenv`
- **Логирование**: `logrus` или `zap`

---

## 📦 REST API

| Метод | Путь | Описание |
|-------|------|----------|
| `GET`    | `/tasks?user_id=`        | Получить список задач |
| `POST`   | `/tasks`                 | Создать новую задачу |
| `PUT`    | `/tasks/{id}`            | Обновить задачу |
| `DELETE` | `/tasks/{id}`            | Удалить задачу |
| `PATCH`  | `/tasks/{id}/done`       | Отметить задачу выполненной |

---

## 🤖 Telegram команды

- `/new <название>` — создать задачу  
- `/list` — список всех задач  
- `/done <id>` — отметить задачу как выполненную  
- `/delete <id>` — удалить задачу  

---

## 🛠 Запуск проекта

```bash
git clone https://github.com/yourusername/tasktracker.git
cd tasktracker

# Настрой .env
cp .env.example .env

# Установи зависимости
go mod tidy

# Собери и запусти
go run ./cmd/server
