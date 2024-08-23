package storage

import (
	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/jmoiron/sqlx"
)

type Client interface {
	CreateClient(client *models.Client) error
	GetClientById(id int64) (*models.Client, error)
	//GetClientByChatID(chatID int64) (*models.Client, error)

	// UpdateClient(client *models.Client) error
	// DeleteClient(clientID int) error
}

type Order interface {
	CreateOrder(order *models.Order) error
	GetOrderById(orderID int) (*models.Order, error)
	GetOrdersByClientId(clientID int) ([]models.Order, error)
	//UpdateOrder(order *models.Order) error
}

type Shop interface {
	CreateShop(shop *models.Shop) error
	GetShopByName(shopName string) (*models.Shop, error)
	GetAllShops() ([]models.Shop, error)

	// UpdatePickupPoint(pickupPoint *models.PickupPoint) error
	// DeletePickupPoint(pickupPointID int) error
}

type Admin interface {
	CreateAdmin(admin *models.Admin) error
	GetAdminByUsername(username string) (*models.Admin, error)
	DeleteAdmin(username string) error
}

type Menu interface {
	CreateMenuItem(item *models.MenuItem) error
	GetMenuItemById(itemID int) (*models.MenuItem, error)
	GetAllMenuItems() ([]models.MenuItem, error)
	UpdateMenuItem(itemID int, name string, price float64) error
	DeleteMenuItem(itemID int) error
}

type Storage struct {
	Client
	Order
	Shop
	Admin
	Menu
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Menu:   NewMenuPostgres(db),
		Admin:  NewAdminPostgres(db),
		Client: NewClientPostgres(db),
		Order:  NewOrderPostgres(db),
		Shop:   NewShopPostgres(db),
	}
}
