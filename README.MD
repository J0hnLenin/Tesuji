# Проект для просмотра партий по игре го

## Описание проекта
Веб-приложение для просмотра партий по игре го. На текущий момент партии импортируются с сервера KGS, в дальнейшем планируется добавить сервера FoxWeiqi, OGS.

## Технологии

### Backend
- **Golang**
- **PostgreSQL**
- **Redis**
- **Kafka**

### Frontend
- Vue
- Vite
- CSS

### Инфраструктура
- Docker & Docker Compose
- GitHub Actions (CI/CD)

## Покрытие тестами
Основные модули: 
- `gamelogic`  
- `gameservice`  
- `protoconvert`  

![Coverage Status](https://coveralls.io/repos/github/J0hnLenin/Tesuji/badge.svg?branch=main)
![Tests Status](https://github.com/J0hnLenin/Tesuji/actions/workflows/coverage.yml/badge.svg)

## Эндпоинты

### Backend (порт 8080)
```
GET /api/v1/games/{id} # Информация о конкретной игре
GET /api/v1/games # Список игр с пагинацией
?page=1&limit=10 # Параметры: страница и лимит
```

### Frontend (порт 4173)
```
GET /games # Страница со списком всех игр
GET /games/{id} # Страница просмотра конкретной gfhnbb
```

## Установка и запуск

Необходимые компоненты: git, docker

1. Клонировать репозиторий

```
git clone https://github.com/J0hnLenin/ReviewRequest
cd ./ReviewRequest
```

2. При необходимости поменять секреты в файле .env, но этот шаг необязательный
3. Собрать проект

```
docker compose up -d
```