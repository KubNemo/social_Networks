package db

import (
	"context"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB() (*pgxpool.Pool, error) {
	err := godotenv.Load("../env")
	if err != nil {
		log.Fatal("Не удалось прочесть .env файл")
	}
	DB, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("не удалось подключиться к БД:", err)
		return nil, err
	}

	log.Println("Вы успешно подключились к ДБ")
	return DB, nil
}
func ApplyMigrations() {
	m, err := migrate.New("file://internal/migrations", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)

	}
	log.Println("миграции успешно выполнены")
}
