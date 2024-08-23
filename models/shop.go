package models

import "time"

// Структура Shop представляет пункт самовывоза
type Shop struct {
	ID        int       // Уникальный идентификатор пункта самовывоза
	Name      string    // Название пункта самовывоза
	Address   string    // Адрес пункта самовывоза
	OpenTime  time.Time // Время открытия
	CloseTime time.Time // Время закрытия
}
