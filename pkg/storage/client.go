package storage

import (
	"database/sql"
	"errors"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/jmoiron/sqlx"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

// ErrClientNotFound определяет ошибку, которая возвращается, если клиент не найден.
var ErrClientNotFound = errors.New("client not found")

// CreateClient добавляет нового клиента в базу данных.
func (strg *ClientPostgres) CreateClient(client *models.Client) error {
	query := `INSERT INTO clients (name, phone, count_of_orders) VALUES ($1, $2, $3) RETURNING id`
	// Выполняем запрос на вставку данных в таблицу и возвращаем идентификатор нового клиента.
	return strg.db.QueryRow(query, client.Name, client.Phone, client.CountOfOrders).Scan(&client.ID)
}

// GetClientByChatID получает клиента из базы данных по идентификатору чата.
// func (strg *ClientPostgres) GetClientByChatID(chatID int64) (*models.Client, error) {
func (strg *ClientPostgres) GetClientById(id int64) (*models.Client, error) {
	client := &models.Client{}
	query := `SELECT id, name, phone, count_of_orders FROM clients WHERE id=$1`
	// Выполняем запрос на выборку данных о клиенте по идентификатору чата.
	err := strg.db.QueryRow(query, id).Scan(&client.ID, &client.Name, &client.Phone, &client.CountOfOrders)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrClientNotFound // Возвращаем ошибку, если клиент не найден.
		}
		return nil, err
	}
	return client, nil
}
