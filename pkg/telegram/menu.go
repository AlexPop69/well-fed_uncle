package telegram

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const (
	menuImagePath = "assets/menu.png"
)

// Отправка изображения меню
func sendMenuImage(t *TelegramBot, chatID int64) error {
	err := sendPhoto(t.token, chatID, menuImagePath)
	if err != nil {
		logrus.Errorf("Failed to send menu image: %s", err)
		t.sendMessage(chatID, "Не удалось загрузить меню.")

		return err
	}

	return nil
}

// Отправка изображения
func sendPhoto(token string, chatID int64, photoPath string) error {
	file, err := os.Open(photoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("photo", filepath.Base(photoPath))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s/sendPhoto", telegramAPI, token)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	query := req.URL.Query()
	query.Add("chat_id", fmt.Sprintf("%d", chatID))
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send photo: %s", resp.Status)
	}

	return nil
}

// // Отправка изображения меню
// func (t *TelegramBot) sendShops(chatID int64) {

// 	// Создаем кнопки для пунктов выдачи
// 	pickupPoints, err := t.service.GetOpenShops()
// 	if err != nil || len(pickupPoints) == 0 {
// 		logrus.Errorf("Failed to get pickup points: %s", err)
// 		t.sendMessage(chatID, "Нет доступных пунктов выдачи.")
// 		return
// 	}

// 	var buttons [][]Button
// 	for _, point := range pickupPoints {
// 		buttons = append(buttons, []Button{{Text: point.Name, CallbackData: point.Name}})
// 	}

// 	keyboard := InlineKeyboardMarkup{InlineKeyboard: buttons}
// 	t.sendMessageWithKeyboard(chatID, "Выберите пункт выдачи:", keyboard)
// }
