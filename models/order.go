package models

import "time"

// Структура Order представляет заказ
type Order struct {
	ID            int        // Уникальный идентификатор заказа
	ClientID      int        // Внешний ключ для клиента
	PickupPointID int        // Внешний ключ для пункта самовывоза
	Items         []MenuItem // Список элементов меню в заказе
	CreatedAt     time.Time  // Дата и время создания заказа
}
