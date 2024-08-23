package telegram

// comands
const (
	startCommand = "/start"
)

// Команды для управления администраторами, пунктами выдачи, меню
const (
	adminPrefix = "/admin"

	helpAdminCommand = "help"

	// операции с администраторами
	createAdminCommand = "add_admin"
	deleteAdminCommand = "del_admin"

	// операции с пунктами выдачи
	addShopCommand    = "add_shop"
	deleteShopCommand = "delete_shop"
)

// messages
const (
	unknownCommandMessage = "Команда не распознана. Попробуйте использовать /start."
)
