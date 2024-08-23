package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/AlexPop69/well-fed_uncle/models"
	"github.com/sirupsen/logrus"
)

// Отправка кнопок с пунктами выдачи
func sendShopsButtons(t *TelegramBot, chatID int64) error {
	shops, err := t.service.GetOpenShops()
	if err != nil || len(shops) == 0 {
		logrus.Errorf("Failed to get shops: %s", err)

		t.sendMessage(chatID, "Нет доступных пунктов выдачи.")

		return err
	}

	var buttons [][]Button
	var text string
	for _, shop := range shops {
		text = fmt.Sprintf(`
		%s
		%s
		%s - %s`,
			shop.Name, shop.Address, shop.OpenTime.Format(timeLayout), shop.CloseTime.Format(timeLayout))

		buttons = append(buttons, []Button{{Text: text, CallbackData: shop.Name}})
	}

	keyboard := InlineKeyboardMarkup{InlineKeyboard: buttons}
	err = t.sendMessageWithKeyboard(chatID, "Выберите пункт выдачи:", keyboard)
	if err != nil {
		logrus.Errorf("Failed to send pickup points buttons: %s", err)
		return err
	}

	return nil
}

// Отправка текстового сообщения с клавиатурой
func (t *TelegramBot) sendMessageWithKeyboard(chatID int64, text string, keyboard InlineKeyboardMarkup) error {
	urlStr := fmt.Sprintf("%s%s/sendMessage", telegramAPI, t.token)

	data := url.Values{}
	data.Set("chat_id", fmt.Sprintf("%d", chatID))
	data.Set("text", text)
	data.Set("reply_markup", keyboard.String())

	resp, err := http.PostForm(urlStr, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message: %s", resp.Status)
	}

	return nil
}

// String преобразует структуру клавиатуры в строку JSON, необходимую для отправки в Telegram.
func (k InlineKeyboardMarkup) String() string {
	jsonStr, _ := json.Marshal(k)
	return string(jsonStr)
}

const timeLayout = "15:04"

func addShop(args []string, t *TelegramBot, message *Message) {
	if len(args) < 5 {
		t.sendMessage(message.Chat.ID,
			`Использование: /admin add_shop "name" "address" "open_time(15:04)"" "close_time(15:04)"`)
		return
	}
	openTime, err := time.Parse(timeLayout, args[4])
	if err != nil {
		t.sendMessage(message.Chat.ID, "Не удалось добавить пункт выдачи.")
		logrus.Error(err)
		return
	}

	closeTime, err := time.Parse(timeLayout, args[5])
	if err != nil {
		t.sendMessage(message.Chat.ID, "Не удалось добавить пункт выдачи.")
		logrus.Error(err)
		return
	}

	shop := models.Shop{
		Name:      args[2],
		Address:   args[3],
		OpenTime:  openTime,
		CloseTime: closeTime,
	}

	err = t.service.CreateShop(&shop)
	if err != nil {
		t.sendMessage(message.Chat.ID, "Не удалось добавить пункт выдачи.")
		logrus.Error(err)
		return
	}

	t.sendMessage(message.Chat.ID, fmt.Sprintf("Пункт выдачи %s успешно добавлен.", shop.Name))
	logrus.Infof("Пункт выдачи %s успешно добавлен.", shop.Name)
}
