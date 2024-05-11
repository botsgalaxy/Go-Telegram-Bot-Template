package modules

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func help(b *gotgbot.Bot, ctx *ext.Context) error {
	text := `ℹ️ <b>Help Menu</b>

🔁 Here are the available commands you can use:

🔰 /help - Display this help message.
🔰 /start - Start interacting with the bot.
🔰 /settings - Update your bot setting.
🔰 /about - about this bot .
🔰 [add other commands and their descriptions here].

If you need further assistance or have any questions, don't hesitate to ask!

Enjoy using our bot! 😊
	`
	_, err := ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send help message: %w", err)
	}
	return nil
}
