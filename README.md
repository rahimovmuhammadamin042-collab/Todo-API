# Todo-API

Минимальный REST API сервер для управления задачами (Todo) с авторизацией через JWT.

Проект написан на Go и использует PostgreSQL для хранения данных.  
Поддерживает регистрацию пользователей, авторизацию и CRUD операции над задачами.

---

## Используемые технологии

- Go
- net/http
- Gorilla Mux
- PostgreSQL
- JSON Web Token (JWT)
- lib/pq (PostgreSQL driver)

---

# Архитектура проекта

- main.go — entry point приложения  
- config/
  - db.go — подключение к базе данных
- models/
  - user.go — модель пользователя
  - task.go — модель задачи
- handlers/
  - auth.go — регистрация и логин
  - tasks.go — CRUD операций с задачами
- middleware/
  - auth.go — JWT middleware
- utils/
  - jwt.go — генерация JWT токена
- router/
  - router.go — настройка маршрутов API
