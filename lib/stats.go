package lib 

import (
	"github.com/bwmarrin/discordgo" 
	"runtime" 
	"time" 
	"fmt"
)

func DisplayStats(m *discordgo.Message, s *discordgo.Session) { 
	check,_ := HasPermission(m,s,8) 
	if check { 
		var dbBytes float64
		sQ := `SELECT SUM(data_length + index_length) FROM information_schema.tables WHERE table_schema='jayneBot'`
		DBConn.QueryRow(sQ).Scan(&dbBytes) 
		dbSize := ((dbBytes*8)/1024)/1024
		cTime := time.Now() 
		s.Channel(m.ChannelID) 
		now := time.Now() 
		timeDiff := now.Sub(cTime) 
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats) 
		alloc := (stats.Alloc)/1024/1024
		total := (stats.TotalAlloc)/1024/1024
		systm := (stats.Sys)/1024/1024
		gc := (stats.NumGC) 
		msg := fmt.Sprintf("```\nPing: %v \nDB Size: %vmb \nAllocated Memory: %vmb \nTotal Memory: %vmb \nSystem Memory: %vmb \nGarbage Cycles: %v```", timeDiff, dbSize, alloc, total, systm, gc) 
		s.ChannelMessageSend(m.ChannelID, msg) 
		return
	} else {
		s.ChannelMessageSend(m.ChannelID, "You don't have that permission.") 
		return 
	}
}
