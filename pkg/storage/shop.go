package storage

import (
	_ "database/sql"
	"fmt"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/jmoiron/sqlx"
)

type ShopPostgres struct {
	db *sqlx.DB
}

func NewShopPostgres(db *sqlx.DB) *ShopPostgres {
	return &ShopPostgres{db: db}
}

// CreateShop добавляет новый пункт самовывоза в базу данных.
func (s *ShopPostgres) CreateShop(shop *models.Shop) error {
	// SQL-запрос для вставки нового пункта самовывоза.
	query := fmt.Sprintf(`
        INSERT INTO %s (name, address, open_time, close_time)
        VALUES ($1, $2, $3, $4)
        RETURNING id`, shopsTable)

	// Выполняем запрос и получаем идентификатор нового пункта самовывоза.
	return s.db.QueryRow(query, shop.Name, shop.Address, shop.OpenTime, shop.CloseTime).Scan(&shop.ID)
}

// GetShopById возвращает пункт самовывоза по его идентификатору.
func (s *ShopPostgres) GetShopByName(shopName string) (*models.Shop, error) {
	shop := &models.Shop{}

	// SQL-запрос для получения пункта самовывоза по его идентификатору.
	query := fmt.Sprintf(`
        SELECT id, name, address, open_time, close_time
        FROM %s
        WHERE name = $1`, shopsTable)

	// Выполняем запрос и сканируем результат в структуру пункта самовывоза.
	err := s.db.QueryRow(query, shopName).Scan(&shop.ID, &shop.Name, &shop.Address, &shop.OpenTime, &shop.CloseTime)
	return shop, err
}

// GetAllShops возвращает все пункты самовывоза.
func (s *ShopPostgres) GetAllShops() ([]models.Shop, error) {
	// SQL-запрос для получения всех пунктов самовывоза.
	query := fmt.Sprintf(`
        SELECT id, name, address, open_time, close_time
        FROM %s`, shopsTable)

	// Выполняем запрос.
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shops []models.Shop

	// Обрабатываем результаты запроса построчно.
	for rows.Next() {
		var shop models.Shop

		err := rows.Scan(&shop.ID, &shop.Name, &shop.Address, &shop.OpenTime, &shop.CloseTime)
		if err != nil {
			return nil, err
		}

		shops = append(shops, shop)
	}

	return shops, nil
}
