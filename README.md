# url-shortener-go

Простой REST API для сокращения URL, написанный на Go с использованием стандартной библиотеки и базы данных для хранения коротких ссылок.

## Возможности

- Сокращение длинных URL.
- Получение оригинального URL по короткому коду.
- Простота использования и настройка.
- Поддержка работы с базой данных для хранения ссылок.
- Поддержка Docker для быстрого развёртывания.

## Технологии

- **Go** 
- **SQLite**
- **Docker** 
- **HTTP** 

## Структура проекта
```
todo-api/ 
├── main.go 
├── go.mod 
├── models/ 
    │ └── task.go 
├── handlers/ 
    │ └── task.go
├── database/ 
    │ └── setup.go 
├── Dockerfile
```

## Установка и запуск

### Локально
    git clone https://github.com/mariapylkova/url-shortener-go.git
    cd url-shortener-go
    go mod tidy
    go run main.go

### С помощью Docker
    docker build -t url-shortener-go .
    docker run -p 8080:8080 url-shortener-go


Теперь сервис будет доступен по адресу `http://localhost:8080`.

## Эндпоинты

| Метод   | URL            | Описание                                    | Тело запроса (если нужно)        |
|---------|----------------|---------------------------------------------|----------------------------------|
| **POST**    | `/shorten`      | Сократить длинный URL                      | `{ "original_url": "https://example.com" }` |
| **GET**     | `/:short_code`  | Получить оригинальный URL по короткому коду| –                                |
