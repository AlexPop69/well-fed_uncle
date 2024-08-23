package telegram

// Структуры для обработки данных от Telegram API

// UpdateResponse представляет ответ от Telegram API на запрос обновлений.
type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update представляет одно обновление от Telegram, содержащее сообщение или другую информацию.
type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

// Message представляет сообщение от пользователя.
type Message struct {
	MessageID int    `json:"message_id"`
	From      *User  `json:"from"`
	Chat      *Chat  `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

// User представляет пользователя, отправившего сообщение.
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

// Chat представляет чат, в который было отправлено сообщение.
type Chat struct {
	ID int64 `json:"id"`
}

// Структура Button представляет собой кнопку на клавиатуре в Telegram.
type Button struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

// InlineKeyboardMarkup представляет разметку клавиатуры с кнопками.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]Button `json:"inline_keyboard"`
}
