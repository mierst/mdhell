package bot

import ( 
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"

	db "../db"
	lib "../lib" 
)
 
type Bot struct {
	dg		*discordgo.Session
	status		string
}

var bot *Bot


func New(Token string) *Bot{

	bot = &Bot{
		status: "Being Built",
	}

	dg,err := discordgo.New("Bot " + Token) 
	bot.dg = dg

	if err != nil {
		fmt.Println(err)
	}
	
	return bot
}

func (bot *Bot) Run() {
	
	//db.DBSetup() 

	err := bot.dg.Open()


	if err == nil {
		fmt.Println("Bot Successfully Online Ctrl-C to close") 

		//bot.dg.AddHandler(core.MessageHandler)
		//go chron.WatchAudit
		
		db.SetupDB() 
		lib.SetupConfig() 

		bot.dg.UpdateStatus(0, bot.status) 

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
		bot.dg.Close()
	} else {
		fmt.Println(err)
	} 

}
