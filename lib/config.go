package lib 

type BotConfig struct {
	MasterLogChannel	string
	Servers				[]string
	AuditChannel		string
}

var Config *BotConfig 

func SetupConfig() {
	Config = &BotConfig{
		MasterLogChannel: "", 
		Servers: []string{""},
		AuditChannel: "", 
	}
}