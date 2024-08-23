package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/AlexPop69/well-fed_uncle/pkg/service"

	"github.com/sirupsen/logrus"
)

// Константа, содержащая базовый URL для Telegram API
const telegramAPI = "https://api.telegram.org/bot"

// Структура TelegramBot хранит токен бота и сервисы, с которыми будет работать бот.
type TelegramBot struct {
	token   string
	service service.Service
}

// NewBot создает новый экземпляр TelegramBot и возвращает его.
func NewBot(service *service.Service) *TelegramBot {
	token := os.Getenv("TELEGRAM_TOKEN")

	return &TelegramBot{
		token:   token,
		service: *service,
	}
}

// Start запускает основной цикл обработки обновлений от Telegram.
func (t *TelegramBot) Start() {
	offset := 0

	for {
		updates, err := t.getUpdates(offset)
		if err != nil {
			logrus.Errorf("Failed to get updates: %s", err)
			continue
		}

		for _, update := range updates {
			if update.Message != nil {
				t.handleUpdate(update.Message)

				offset = update.UpdateID + 1
			}
		}
	}
}

// getUpdates запрашивает новые обновления от Telegram API.
func (t *TelegramBot) getUpdates(offset int) ([]Update, error) {

	resp, err := http.Get(fmt.Sprintf("%s%s/getUpdates?offset=%d", telegramAPI, t.token, offset))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response UpdateResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

// Обработка входящих сообщений
func (t *TelegramBot) handleUpdate(message *Message) {
	logrus.Infof(`Received message "%s" from user: %s`, message.Text, message.From.UserName)

	if strings.HasPrefix(message.Text, adminPrefix) {
		if !t.isAdmin(message.From.UserName) {
			t.sendMessage(message.Chat.ID, "У вас нет прав для выполнения этой команды.")

			return
		}

		t.handleAdminCommands(message)

	} else {
		t.handleUserCommands(message)
	}
}

// Обработка команд пользователей
func (t *TelegramBot) handleUserCommands(message *Message) {
	switch strings.ToLower(message.Text) {
	case startCommand:
		t.sendMessage(message.Chat.ID, fmt.Sprintf("Добро пожаловать, %s! Выберите пункт самовывоза.", message.From.UserName))

		// Отправка изображения меню и кнопок с пунктами выдачи
		err := sendMenuImage(t, message.Chat.ID)
		if err != nil {
			logrus.Error(err)
			return
		}

		err = sendShopsButtons(t, message.Chat.ID)
		if err != nil {
			logrus.Error(err)
			return
		}

	default:
		t.sendMessage(message.Chat.ID, unknownCommandMessage)
	}
}

// Обработка команд администратора
func (t *TelegramBot) handleAdminCommands(message *Message) {
	args := parseCommandArgs(message.Text)

	if len(args) < 2 {
		t.sendMessage(message.Chat.ID, "Недостаточно аргументов для выполнения команды.")

		return
	}

	switch args[1] {
	case helpAdminCommand:
		helpMessages := adminHelp()

		for _, mes := range helpMessages {
			t.sendMessage(message.Chat.ID, mes)
		}

	case createAdminCommand:
		addAdmin(args, t, message)

	case deleteAdminCommand:
		deleteAdmin(args, t, message)

	case addShopCommand:
		addShop(args, t, message)

	// case "remove_shop":
	// 	if len(args) < 3 {
	// 		t.sendMessage(message.Chat.ID, "Использование: /admin remove_shop <name>")
	// 		return
	// 	}
	// 	name := args[2]
	// 	err := t.pickupPointService.RemoveShop(name)
	// 	if err != nil {
	// 		t.sendMessage(message.Chat.ID, "Не удалось удалить пункт выдачи.")
	// 		logrus.Error(err)
	// 		return
	// 	}
	// 	t.sendMessage(message.Chat.ID, fmt.Sprintf("Пункт выдачи %s успешно удален.", name))

	// case "update_shop_time":
	// 	if len(args) < 5 {
	// 		t.sendMessage(message.Chat.ID, "Использование: /admin update_shop_time <name> <open_time> <close_time>")
	// 		return
	// 	}
	// 	name := args[2]
	// 	openTime := args[3]
	// 	closeTime := args[4]
	// 	err := t.pickupPointService.UpdateShopTime(name, openTime, closeTime)
	// 	if err != nil {
	// 		t.sendMessage(message.Chat.ID, "Не удалось обновить время работы пункта выдачи.")
	// 		logrus.Error(err)
	// 		return
	// 	}
	// 	t.sendMessage(message.Chat.ID, fmt.Sprintf("Время работы пункта выдачи %s успешно обновлено.", name))

	// case "upload_menu_image":
	// 	// Пример: /admin upload_menu_image <image_path>
	// 	if len(args) < 3 {
	// 		t.sendMessage(message.Chat.ID, "Использование: /admin upload_menu_image <image_path>")
	// 		return
	// 	}
	// 	imagePath := args[2]
	// 	err := t.menuService.UploadMenuImage(imagePath)
	// 	if err != nil {
	// 		t.sendMessage(message.Chat.ID, "Не удалось загрузить изображение меню.")
	// 		logrus.Error(err)
	// 		return
	// 	}
	// 	t.sendMessage(message.Chat.ID, "Изображение меню успешно загружено.")

	// case "update_price":
	// 	if len(args) < 4 {
	// 		t.sendMessage(message.Chat.ID, "Использование: /admin update_price <item_name> <new_price>")
	// 		return
	// 	}
	// 	itemName := args[2]
	// 	newPrice := args[3]
	// 	err := t.menuService.UpdatePrice(itemName, newPrice)
	// 	if err != nil {
	// 		t.sendMessage(message.Chat.ID, "Не удалось обновить цену.")
	// 		logrus.Error(err)
	// 		return
	// 	}
	// 	t.sendMessage(message.Chat.ID, fmt.Sprintf("Цена товара %s успешно обновлена.", itemName))

	// case "stop_selling":
	// 	if len(args) < 3 {
	// 		t.sendMessage(message.Chat.ID, "Использование: /admin stop_selling <item_name>")
	// 		return
	// 	}
	// 	itemName := args[2]
	// 	err := t.menuService.StopSelling(itemName)
	// 	if err != nil {
	// 		t.sendMessage(message.Chat.ID, "Не удалось остановить продажу товара.")
	// 		logrus.Error(err)
	// 		return
	// 	}
	// 	t.sendMessage(message.Chat.ID, fmt.Sprintf("Продажа товара %s остановлена.", itemName))

	default:
		t.sendMessage(message.Chat.ID, "Неизвестная команда администратора.")
	}
}

// Разбор аргументов команды
func parseCommandArgs(input string) []string {
	re := regexp.MustCompile(`"([^"]*)"|\S+`)

	matches := re.FindAllString(input, -1)

	var args []string

	for _, match := range matches {
		// Убираем кавычки, если они есть
		args = append(args, strings.Trim(match, `"`))
	}

	return args
}

// sendMessage отправляет сообщение в Telegram чат.
func (t *TelegramBot) sendMessage(chatID int64, text string) error {
	data := url.Values{}
	data.Set("chat_id", fmt.Sprintf("%d", chatID))
	data.Set("text", text)

	resp, err := http.PostForm(fmt.Sprintf("%s%s/sendMessage", telegramAPI, t.token), data)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer resp.Body.Close()

	return nil
}
