package filbot

import(
    "fmt"
    "github.com/go-chat-bot/bot"
)

// This echo function just repeats everything that is said in the channel.
// This can be used for just echoing commands, or channel logging. 
func echo(command *bot.PassiveCmd) (msg string, err error){
    msg = fmt.Sprintf(command.Raw)
    return
}

func init() {
    bot.RegisterPassiveCommand(
        "", 
        echo,
    )
}