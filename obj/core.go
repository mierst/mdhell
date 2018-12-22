package obj 

import(
	lib "../lib"
) 

type Log struct {
	LogID		string
	Moderator	string
	ModName		string
	User		string
	UserName	string
	Server		string
	Command		string
	Reason		string
	Time		int 
}

type Warn struct {
	WarnID		string 
	Moderator	string
	ModName		string
	User		string 
	UserName	string
	Reason		string
	Time		int
}

type Mute struct {

}

type Ban struct {
	UserID	string 
	UserName	string
	Moderator	string
	ModName		string 
	Server		string
	ServerName	string
	Reason		string
	Time		int
}

func (l Log) Save() {
	iQ := `
	INSERT logs SET
	Moderator=?,
	ModName=?,
	User=?,
	UserName=?,
	Server=?,
	Command=?,
	Reason=?,
	Time=?`
	
	stmt,_ := lib.DBConn.Prepare(iQ) 
	stmt.Exec(
		l.Moderator, 
		l.ModName, 
		l.User,
		l.UserName,
		l.Server,
		l.Command,
		l.Reason,
		l.Time,
	)
}