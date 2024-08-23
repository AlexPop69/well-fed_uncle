package main

import (
	"os"

	"github.com/AlexPop69/well-fed_uncle/pkg/service"
	"github.com/AlexPop69/well-fed_uncle/pkg/storage"
	"github.com/AlexPop69/well-fed_uncle/pkg/telegram"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Подключение к базе данных PostgreSQL
	db, err := storage.NewPostgresDB(&storage.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		DbName:   viper.GetString("database.dbname"),
		SSLmode:  viper.GetString("database.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize database: %s", err.Error())
	}

	// Инициализация сервисов
	stor := storage.NewStorage(db)
	services := service.NewService(stor)

	// Инициализация Telegram-бота с передачей всех сервисов
	bot := telegram.NewBot(services)

	// Запуск бота
	logrus.Info("Starting Telegram bot...")
	bot.Start()

}

func initConfig() error {
	// Настройка логгирования
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
		return err
	}

	// Загрузка конфигурации из файла config.yml
	viper.SetConfigFile("./configs/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file: %s", err)
		return err
	}

	return nil
}
