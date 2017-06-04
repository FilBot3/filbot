package filbot

import(
    "github.com/go-chat-bot/bot"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestEcho(t *testing.T) {
    Convey("Given a string", t, func(){
        cmd := &bot.PassiveCmd{}
        Convey("It echos the random string", func(){
            cmd.Raw = "This is a string to be echoed."
            result, err := echo(cmd)
            So(err, ShouldBeNil)
            So(result, ShouldEqual, "This is a string to be echoed.")
        })
    })
}