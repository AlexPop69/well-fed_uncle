package service

import (
	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/AlexPop69/well-fed_uncle/pkg/storage"
	"github.com/sirupsen/logrus"
)

// AdminService предоставляет методы для управления администраторами.
type AdminService struct {
	storage *storage.Storage // Слой доступа к данным
}

// NewAdminService создает новый экземпляр AdminService.
func NewAdminService(storage *storage.Storage) *AdminService {
	return &AdminService{storage: storage}
}

// Authentication выполняет аутентификацию администратора.
func (s *AdminService) Authentication(username string) (*models.Admin, error) {
	return s.storage.GetAdminByUsername(username)
}

// AddAdmin добавляет нового администратора.
func (s *AdminService) AddAdmin(username string) (*models.Admin, error) {
	admin := &models.Admin{
		Username: username,
	}

	err := s.storage.CreateAdmin(admin)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return admin, nil
}

// DeleteAdmin удялет администратора из базы данных.
func (s *AdminService) DeleteAdmin(username string) error {
	err := s.storage.DeleteAdmin(username)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
