package db 

import(
	"fmt"
	lib "../lib"
)
func SetupDB() {
	for _,server := range lib.Config.Servers {
		cQ := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS server_%s_info(
			Name varchar(64) DEFAULT 'none', 
			ID	varchar(64) DEFAULT 'none', 
			MasterServer boolean DEFAULT false, 
			Warns int DEFAULT 0, 
			Bans int DEFAULT 0
			)`, server)
		lib.DBConn.Exec(cQ) 
	}

	
}