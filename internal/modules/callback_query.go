package modules

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func handleQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.Update.CallbackQuery
	_, er := query.Answer(b, &gotgbot.AnswerCallbackQueryOpts{})
	if er != nil {
		return er
	}
	fmt.Println(query.Data)
	if query.Data == "about" {
		_, _, err := query.Message.EditText(
			b,
			aboutText,
			&gotgbot.EditMessageTextOpts{
				ParseMode:   "html",
				ReplyMarkup: aboutButton,
			},
		)
		if err != nil {
			return err
		}

	}
	return nil

}
