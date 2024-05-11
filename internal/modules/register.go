package modules

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

func RegisterHandlers(d *ext.Dispatcher) {

	d.AddHandler(handlers.NewCommand("start", start))
	d.AddHandler(handlers.NewCommand("help", help))
	d.AddHandler(handlers.NewCommand("settings", setting))
	d.AddHandler(handlers.NewCommand("about", about))
	d.AddHandler(handlers.NewCallback(callbackquery.Equal("about"), handleQuery))
}
