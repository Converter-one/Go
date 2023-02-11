package main

import (
"log"
"github.com/go-telegram-bot-api/telegram-bot-api"
"openai"
"types"

"github.com/go-telegram-bot-api/telegram-bot-api"
"github.com/openai/types"
"github.com/Gopher AI/gopher-ai"
"github.com/Gopher AI/gopher-ai/types"
"github.com/OpenAI/openai-go"
)


func main() {
bot, err := tgbotapi.NewBotAPI("Api")
if err != nil {
log.Panic(err)
}
bot.Debug = true

log.Printf("Authorized on account %s", bot.Self.UserName)

u := tgbotapi.NewUpdate(0)
u.Timeout = 60

updates, err := bot.GetUpdatesChan(u)
if err != nil {
	log.Panic(err)
}

for update := range updates {
	if update.Message == nil {
		continue
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
	bot.Send(msg)

	reply := types.InputMessageContent{
		Text: "Good Bye!",
	}
	openai.Query(&reply, update.Message.Chat.ID, update.Message.Text)
}}

openai.api_key = "Api"

if err != nil {
	panic(err)
}

bot.Debug = true

updateCfg := tgbotapi.NewUpdate(0)
updateCfg.Timeout = 60

updates, err := bot.GetUpdatesChan(updateCfg)

for update := range updates {
	tgMsg := update.Message
	if tgMsg == nil {
		continue
	}

	msg := gopher_ai.Message{
		Text: tgMsg.Text,
		User: types.User{
			ID:           tgMsg.From.ID,
			FirstName:    tgMsg.From.FirstName,
			Username:     tgMsg.From.UserName,
			LanguageCode: tgMsg.From.LanguageCode,
		},
		Chat: types.Chat{
			ID:        tgMsg.Chat.ID,
			Type:      tgMsg.Chat.Type,
			Title:     tgMsg.Chat.Title,
			FirstName: tgMsg.Chat.FirstName,
		},
	}

	reply, err := openai.Reply(msg)
	if err != nil {
		continue
	}

	msgCfg := tgbotapi.NewMessage(tgMsg.Chat.ID, reply.Text)

	bot.Send(msgCfg)
}
response, err := openai.Completion.Create(
	model: "text-davinci-003",
	prompt: "Soter",
	temperature: 0.5,
	max_tokens: 3000,
	top_p: 1.0,
	frequency_penalty: 0.5,
	presence_penalty: 0.0,
	)

if err != nil {
fmt.Println("Error: ", err)
} else {
fmt.Println("Response: ", response)
}

type MenuKeyboard struct {
	// profileBtn = "🗿 Профиль"
	exchangeBtn = "💵 Заработать"
	convertBtn  = "♾ Convert"
	walletBtn   = "💼 Профессии"
	settingsBtn = "⚙️ Настройки"
	toAdminBtn  = "⤵️ Перейти в Админ. панель"
  
	buttons        = [exchangeBtn, convertBtn, walletBtn, settingsBtn]
	adminButtons   = append(buttons, toAdminBtn)
	markupKeyboard = defaultKeyboard(buttons)
	adminMarkupKeyboard = defaultKeyboard(adminButtons)
  }
  
  type BackKeyboard struct {
	backBtn       = "⬅️ Назад"
	buttons       = [backBtn]
	markupKeyboard = defaultKeyboard(buttons)
  }
bot.StartPolling()