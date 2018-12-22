package core 

import(
	"github.com/bwmarrin/discordgo"
	str "strings"
	lib "../lib"
) 

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return 
	}

	if m.Author.Bot {
		return 
	}

	removed := Filter(m.Message, s) 
	if !removed {
		if len(m.Message.Content) > 0 {
			if string(m.Message.Content[0]) == "!" {
				check,_ := lib.HasPermission(m.Message,s,8)
				if check {
					switch str.Fields(m.Content)[0] {
						case "!mdlog":
							ModLog(m.Message, s) 
							break
						case "!ban":
							Ban(m.Message, s) 
							break 
						case "!audit": 
							Audit(m.Message,s)
							break
					}
					} else {
					s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + " You do not have that permission") 
					return 
				}
			}
		}
	} else {
		return 
	}
}

func PresenceHandler(s *discordgo.Session, pU *discordgo.PresenceUpdate) {

}

func VoiceHandler(s *discordgo.Session, vS *discordgo.VoiceStateUpdate) {

}

