package modules

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func setting(b *gotgbot.Bot, ctx *ext.Context) error {
	text := `⚙️ <b>Settings</b>

📜 You can adjust the following settings:

💎 1. <b>Notification Preferences</b>: Choose how you want to receive notifications.
💎 2. <b>Language</b>: Set your preferred language for the bot's responses.

If you need assistance or have any questions about the settings, feel free to ask!
	`
	_, err := ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send setting message: %w", err)
	}
	return nil
}
