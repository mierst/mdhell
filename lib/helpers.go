package lib 

import(
	"github.com/bwmarrin/discordgo" 
)


func HasPermission(m *discordgo.Message, s *discordgo.Session, permission int) (bool, error) {
	g, err := s.State.Guild(m.GuildID) 
	gM, err := s.GuildMember(m.GuildID, m.Author.ID) 

	if err != nil {
		return false, err
	}
 
	if g.OwnerID == gM.User.ID {
		return true, nil 
	}
	
	for _, roleID := range gM.Roles { 
		role, err := s.State.Role(m.GuildID, roleID) 
		if err != nil { 
			return false, err 
		}
		if role.Permissions&permission != 0{
			return true, nil 
		}
	}

	return false, nil 
}