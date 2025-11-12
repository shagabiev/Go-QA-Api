# Go QA API

**API для системы вопросов и ответов**  
на **Go + PostgreSQL + GORM + Docker**  

---

## Особенности

- Использован net/http
- Работа с БД через GORM
- Использован PostgreSQL
- Использованы миграции с помощью goose
- Обернуто в Docker, для запуска docker-compose
- **Чистая архитектура** — `handlers`, `repository`, `models`

---

## API Endpoints

| Метод | Путь | Описание | Тело запроса |
|------|------|---------|-------------|
| `GET` | `/questions` | Список вопросов | — |
| `POST` | `/questions` | Создать вопрос | `{ "text": "..." }` |
| `GET` | `/questions/{id}` | Вопрос + ответы | — |
| `DELETE` | `/questions/{id}` | Удалить вопрос | — |
| `POST` | `/questions/{id}/answers` | Добавить ответ | `{ "text": "..." }` |
| `GET` | `/answers/{id}` | Получить ответ | — |
| `DELETE` | `/answers/{id}` | Удалить ответ | — |

---

## Запуск проекта

1. Клонируйте репозиторий
2. Запуск с помощью Docker
   docker-compose up --build
   API будет доступ: http://localhost:8080
3. Для остановки контейнеров использовать
   docker-compose down
4. Для удаления БД
   docker-compose down -v

---

## Тестирование endpoints в Postman

1. GET http://localhost:8080/questions
   []
2. POST http://localhost:8080/questions
   JSON:
   { "text": "Который час?" }
   {
    "id": 1,
    "text": "Который час?",
    "created_at": "2025-11-12T12:43:50.480532291Z"
   }
3. POST http://localhost:8080/questions/1/answers
   JSON:
   { "text": "15:45" }
   {
    "id": 1,
    "question_id": 1,
    "user_id": "5a862315-e046-4fc5-bd90-ae39ccf8bdd4",
    "text": "15:45",
    "created_at": "2025-11-12T12:45:13.205612603Z"
   }
4. POST http://localhost:8080/questions
   JSON:
   { "text": "Какой сегодня день недели?" }
   {
    "id": 2,
    "text": "Какой сегодня день недели?",
    "created_at": "2025-11-12T12:43:50.480532291Z"
   }
5. DELETE http://localhost:8080/questions/1
6. GET http://localhost:8080/questions
   [
    {
        "id": 4,
        "text": "Какой сегодня день недели?",
        "created_at": "2025-11-12T12:46:25.362899Z"
    }
  ]
