package modules

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var aboutText = `<b> ℹ️ About </b>

This bot is built using Go Programming Language by 𝐏𝐫𝐢𝐦𝐞 𝐀𝐤𝐚𝐬𝐡 🇧🇩. You can find the source code on GitHub.

`
var aboutButton = gotgbot.InlineKeyboardMarkup{
	InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
		{
			gotgbot.InlineKeyboardButton{Text: "👨‍💻 Author", Url: "https://t.me/primeakash"},
			gotgbot.InlineKeyboardButton{Text: "</> Github", Url: "https://github.com/botsgalaxy"},
		},
		{
			gotgbot.InlineKeyboardButton{Text: "📢 Updates Channel", Url: "https://github.com/botsgalaxyOfficial"},
		},
	},
}

func about(b *gotgbot.Bot, ctx *ext.Context) error {

	_, err := ctx.EffectiveMessage.Reply(
		b,
		aboutText,
		&gotgbot.SendMessageOpts{
			ParseMode:   "html",
			ReplyMarkup: aboutButton,
		},
	)
	return err
}
