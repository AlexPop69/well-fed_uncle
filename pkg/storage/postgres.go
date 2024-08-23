package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL для работы с базой данных
	"github.com/sirupsen/logrus"
)

const (
	adminsTable  = "admins"
	shopsTable   = "shops"
	clientsTable = "clients"
	menuTable    = "menu"
	ordersTable  = "orders"
)

// Конфигурация базы данных
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLmode  string
}

// NewPostgresDB устанавливает соединение с базой данных PostgreSQL и возвращает экземпляр sqlx.DB.
func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSLmode))
	if err != nil {
		logrus.Fatalf("Failed to open database connection: %s", err)
		return nil, err
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		logrus.Fatalf("Failed to ping database: %s", err)
		return nil, err
	}

	return db, nil
}
