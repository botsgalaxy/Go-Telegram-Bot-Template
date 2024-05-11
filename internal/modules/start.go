package modules

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	text := fmt.Sprintf(`👋 Hello %s!

Welcome to our Telegram bot 🤖

You can use this bot to [describe what your bot does or its main features].
To get started, simply type /help to see the list of available commands.
If you have any questions or need assistance, feel free to reach out to us.

Good Day! 😊`, user.FirstName)

	startButton := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				gotgbot.InlineKeyboardButton{Text: "📝 About", CallbackData: "about"},
				gotgbot.InlineKeyboardButton{Text: "🆘 Help", CallbackData: "help"},
			},
			{
				gotgbot.InlineKeyboardButton{Text: "⚙️ Settings", CallbackData: "settings"},
			},
		},
	}

	_, err := ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: startButton,
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
