package storage

import (
	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/jmoiron/sqlx"
)

type MenuPostgres struct {
	db *sqlx.DB
}

func NewMenuPostgres(db *sqlx.DB) *MenuPostgres {
	return &MenuPostgres{db: db}
}

// CreateMenuItem добавляет новый элемент меню в базу данных.
func (strg *MenuPostgres) CreateMenuItem(item *models.MenuItem) error {
	query := `INSERT INTO menu (name, price) VALUES ($1, $2) RETURNING id`
	// Выполняем запрос на вставку данных о новом элементе меню.
	return strg.db.QueryRow(query, item.Name, item.Price).Scan(&item.ID)
}

// GetMenuItemById получает элемент меню из базы данных по его идентификатору.
func (strg *MenuPostgres) GetMenuItemById(itemID int) (*models.MenuItem, error) {
	item := &models.MenuItem{}
	query := `SELECT id, name, price FROM menu WHERE id=$1`
	// Выполняем запрос на выборку данных об элементе меню по его идентификатору.
	err := strg.db.QueryRow(query, itemID).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// GetAllMenuItems возвращает все элементы меню из базы данных.
func (s *MenuPostgres) GetAllMenuItems() ([]models.MenuItem, error) {
	query := `SELECT id, name, price FROM menu`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.MenuItem
	for rows.Next() {
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// UpdateMenuItem обновляет информацию об элементе меню в базе данных.
func (strg *MenuPostgres) UpdateMenuItem(itemID int, name string, price float64) error {
	query := `UPDATE menu SET name=$1, price=$2 WHERE id=$3`
	// Выполняем запрос на обновление данных элемента меню.
	_, err := strg.db.Exec(query, name, price, itemID)
	return err
}

// DeleteMenuItem удаляет элемент меню из базы данных по его идентификатору.
func (strg *MenuPostgres) DeleteMenuItem(itemID int) error {
	query := `DELETE FROM menu WHERE id=$1`
	// Выполняем запрос на удаление элемента меню.
	_, err := strg.db.Exec(query, itemID)
	return err
}
