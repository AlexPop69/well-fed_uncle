package models

// Структура Client представляет клиента
type Client struct {
	ID            int    // Уникальный идентификатор клиента
	Name          string // Имя клиента
	Phone         string // Номер телефона клиента
	CountOfOrders int    // Количество заказов, сделанных клиентом
}
