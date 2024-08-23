package telegram

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Проверка является ли пользователь администратором
func (t *TelegramBot) isAdmin(username string) bool {
	admin, err := t.service.Authentication(username)
	if err != nil || admin == nil {
		logrus.Error(err)

		return false
	}

	return true
}

// Отправка текста с доступными командами администраторов
func adminHelp() []string {
	helpText := []string{
		`Доступные команды администратора:`,
		`/admin help - Показать это сообщение`,
		`/admin add_admin <username> - Добавить нового администратора`,
		`/admin del_admin <username> - Удалить администратора`,
		`/admin add_shop "name" "address" <open_time(15:04)> <close_time(15:04)> - Добавить новый пункт выдачи`,
		`/admin remove_shop "name" - Удалить пункт выдачи`,
		`/admin update_shop "name" <open_time(15:04)> <close_time(15:04)> - Изменить время работы пункта выдачи`,
		`/admin upload_menu_image <image_path> - Загрузить новое изображение меню`,
		`/admin update_price <item_name> <new_price> - Изменить цену на товар`,
		`/admin stop_selling <item_name> - Остановить продажу товара`}
	return helpText
}

// Добавление нового администратора
func addAdmin(args []string, t *TelegramBot, message *Message) {
	if len(args) < 3 {
		t.sendMessage(message.Chat.ID, "Использование: /admin add_admin <username>")

		return
	}

	username := args[2]

	admin, err := t.service.AddAdmin(username)
	if err != nil || admin == nil {
		t.sendMessage(message.Chat.ID, "Не удалось добавить администратора.")
		t.sendMessage(message.Chat.ID, "Использование: /admin add_admin <username>")
		logrus.Error(err)

		return
	}

	t.sendMessage(message.Chat.ID, fmt.Sprintf("Администратор %s успешно добавлен.", username))
	logrus.Infof("Администратор %s успешно добавлен.", username)
}

// Удаление администратора
func deleteAdmin(args []string, t *TelegramBot, message *Message) {
	if len(args) < 3 {
		t.sendMessage(message.Chat.ID, "Использование: /admin del_admin <username>")
		return
	}

	username := args[2]

	err := t.service.DeleteAdmin(username)
	if err != nil {
		t.sendMessage(message.Chat.ID, "Не удалось удалить администратора.")
		logrus.Error(err)

		return
	}

	t.sendMessage(message.Chat.ID, fmt.Sprintf("Администратор %s успешно удален.", username))
	logrus.Infof("Администратор %s успешно удален.", username)
}
