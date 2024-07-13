// (c) Jisin0

package plugins

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var Dispatcher *ext.Dispatcher = ext.NewDispatcher(&ext.DispatcherOpts{
	// If an error is returned by a handler, log it and continue going.
	Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
		fmt.Println("an error occurred while handling update:", err.Error())
		return ext.DispatcherActionNoop
	},
	MaxRoutines: ext.DefaultMaxRoutines,
})

const (
	commandHandlerGroup  = 2
	callbackHandlerGroup = 1
)

func init() {
	// Dispatcher.AddHandlerToGroup(handlers.NewCommand("start", Start), commandHandlerGroup)

	// // Static Commands.
	// Dispatcher.AddHandlerToGroup(handlers.NewMessage(allCommand, CommandHandler), commandHandlerGroup)
}

func allCommand(msg *gotgbot.Message) bool {
	ents := msg.GetEntities()
	if len(ents) != 0 && ents[0].Offset == 0 && ents[0].Type != "bot_command" {
		return false
	}

	text := msg.GetText()

	if r, _ := utf8.DecodeRuneInString(text); r != '/' {
		return false
	}

	split := strings.Split(strings.ToLower(strings.Fields(text)[0]), "@")
	cmd := split[0][1:]

	return cmd != ""
}
