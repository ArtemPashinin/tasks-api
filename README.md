## Локальный запуск

- createdb todo
- psql -U postgres -d todo -f db/migrations/001_create_tasks_table.up.sql
- go run main.go

## Запуск через docker

- docker-compose up --build