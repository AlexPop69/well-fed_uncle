package storage

import (
	"fmt"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

// GetAdminByUsername получает администратора из базы данных по его username.
func (strg *AdminPostgres) GetAdminByUsername(username string) (*models.Admin, error) {
	admin := &models.Admin{}

	query := fmt.Sprintf(`
		SELECT id, username 
		FROM %s 
		WHERE username=$1`, adminsTable)

	// Выполняем запрос на выборку данных об администраторе по его username.
	err := strg.db.QueryRow(query, username).Scan(&admin.ID, &admin.Username)
	if err != nil {
		return nil, fmt.Errorf("GetAdminByUsername: Admin %s does not exist", username)
	}

	return admin, nil
}

// CreateAdmin добавляет нового администратора в базу данных.
func (strg *AdminPostgres) CreateAdmin(admin *models.Admin) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (username) 
		VALUES ($1) 
		RETURNING id`, adminsTable)

	// Выполняем запрос на вставку данных о новом администраторе.
	return strg.db.QueryRow(query, admin.Username).Scan(&admin.ID)
}

// DeleteAdmin удаляет администратора из базы данных по его username.
func (strg *AdminPostgres) DeleteAdmin(username string) error {
	query := fmt.Sprintf(`
		DELETE FROM %s 
		WHERE username=$1`, adminsTable)

	// Выполняем запрос на удаление администратора.
	_, err := strg.db.Exec(query, username)
	return err
}
