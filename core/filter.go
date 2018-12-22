package core 

import(
	"regexp" 
	"time"
	"github.com/bwmarrin/discordgo"
	obj "../obj"
) 

func discordLinkFilter(messageContent string) bool /*found boolean*/ {
	discordLinkRegex := `(http://|https//|)(discord|discordapp)(\.gg|\.com)\/(invite\/|)[A-Za-z1-9]+`
	re := regexp.MustCompile(discordLinkRegex) 
	if re.MatchString(messageContent) {
		return true
	}

	return false 
}

func Filter(m *discordgo.Message, s *discordgo.Session) bool /*removed boolean*/  {
	if discordLinkFilter(m.Content) {

		l := &obj.Log {
			Moderator: s.State.User.ID,  
			ModName: s.State.User.Username, 
			User: m.Author.ID, 
			UserName: m.Author.Username, 
			Server: m.GuildID,
			Command: "none", 
			Reason: "Discord Link Removal", 
			Time: int(time.Now().Unix()),
		}

		l.Save() 

		s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + " You can't send discord links here") 
		s.ChannelMessageDelete(m.ChannelID, m.ID) 
		return true
	}

	return false 
}

 