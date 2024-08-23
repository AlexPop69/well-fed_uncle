package storage

import (
	"encoding/json"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

// CreateOrder добавляет новый заказ в базу данных.
func (strg *OrderPostgres) CreateOrder(order *models.Order) error {
	// Преобразуем элементы заказа в JSON-строку для хранения.
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	// SQL-запрос для вставки нового заказа.
	query := `
        INSERT INTO orders (client_id, pickup_point_id, items, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	// Выполняем запрос и получаем идентификатор нового заказа.
	return strg.db.QueryRow(query, order.ClientID, order.PickupPointID, itemsJSON, order.CreatedAt).Scan(&order.ID)
}

// GetOrderById возвращает заказ по его идентификатору.
func (strg *OrderPostgres) GetOrderById(orderID int) (*models.Order, error) {
	order := &models.Order{}
	var itemsJSON []byte

	// SQL-запрос для получения заказа по идентификатору.
	query := `
        SELECT id, client_id, pickup_point_id, items, created_at
        FROM orders
        WHERE id = $1`

	// Выполняем запрос и сканируем результат в структуру заказа.
	err := strg.db.QueryRow(query, orderID).Scan(&order.ID, &order.ClientID, &order.PickupPointID, &itemsJSON, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	// Декодируем JSON-строку в структуру элементов заказа.
	err = json.Unmarshal(itemsJSON, &order.Items)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetOrdersByClientId возвращает все заказы клиента по его идентификатору.
func (strg *OrderPostgres) GetOrdersByClientId(clientID int) ([]models.Order, error) {
	// SQL-запрос для получения заказов клиента.
	query := `
        SELECT id, client_id, pickup_point_id, items, created_at
        FROM orders
        WHERE client_id = $1`

	// Выполняем запрос.
	rows, err := strg.db.Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	// Обрабатываем результаты запроса построчно.
	for rows.Next() {
		var order models.Order
		var itemsJSON []byte

		err := rows.Scan(&order.ID, &order.ClientID, &order.PickupPointID, &itemsJSON, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		// Декодируем JSON-строку в структуру элементов заказа.
		err = json.Unmarshal(itemsJSON, &order.Items)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
