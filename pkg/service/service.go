package service

import (
	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/AlexPop69/well-fed_uncle/pkg/storage"
)

type Client interface {
	StartInteraction(chatID int64, username string) (*models.Client, error)
	GetClient(chatID int64) (*models.Client, error)

	// CreateClient(client *models.Client) error
	// GetClientByChatID(chatID int64) (*models.Client, error)
	// UpdateClient(client *models.Client) error
	// DeleteClient(clientID int) error
}

type Order interface {
	CreateOrder(clientID, pickupPointID int, items []models.MenuItem) (*models.Order, error)
	GetOrder(orderID int) (*models.Order, error)
	ListOrders(clientID int) ([]models.Order, error)

	// CreateOrder(order *models.Order) error
	// GetOrderById(orderID int) (*models.Order, error)
	// GetOrdersByClientId(clientID int) ([]models.Order, error)
	//UpdateOrder(order *models.Order) error
}

type Shop interface {
	//ListShops() ([]models.Shop, error)
	//IsOpen(shopID int) (bool, error)

	CreateShop(shop *models.Shop) error

	GetShopByName(shopName string) (*models.Shop, error)
	GetAllShops() ([]models.Shop, error)
	GetOpenShops() ([]models.Shop, error)

	// GetPickupPointById(pickupPointID int) (*models.PickupPoint, error)
	// GetAllPickupPoints() ([]models.PickupPoint, error)
	// UpdatePickupPoint(pickupPoint *models.PickupPoint) error
	// DeletePickupPoint(pickupPointID int) error
}

type Admin interface {
	Authentication(username string) (*models.Admin, error)
	AddAdmin(username string) (*models.Admin, error)
	DeleteAdmin(username string) error
}

type Menu interface {
	GetMenuItems() ([]models.MenuItem, error)
	GetMenuItem(id int) (*models.MenuItem, error)
	AddMenuItem(name string, price float64) (*models.MenuItem, error)
	UpdateMenuItem(id int, name string, price float64) error
	DeleteMenuItem(id int) error

	// CreateMenuItem(item *models.MenuItem) error
	// GetMenuItemById(itemID int) (*models.MenuItem, error)
	// GetAllMenuItems() ([]models.MenuItem, error)
}

type Service struct {
	Client
	Order
	Shop
	Admin
	Menu
}

func NewService(strg *storage.Storage) *Service {
	return &Service{
		Menu:   NewMenuService(strg),
		Admin:  NewAdminService(strg),
		Client: NewClientService(strg),
		Order:  NewOrderService(strg),
		Shop:   NewShopService(strg),
	}
}
