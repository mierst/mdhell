package lib 

import ( 
	"database/sql" 
	_ "github.com/go-sql-driver/mysql"  
	"fmt"
)

var ( 
	DBConn *sql.DB 
)

func DBConnect() (error) {
	var err error 
	DBConn, err = sql.Open("mysql", "root:password@/modBot") 
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Exists(sql string, args ...string) bool {
	var exists bool 
	query := fmt.Sprintf("SELECT exists (%s)", sql)
	queryArgs := make([]interface{}, len(args)) 
	for i,v := range args {
		queryArgs[i] = v
	}
	err := DBConn.QueryRow(query, queryArgs...).Scan(&exists) 
	if err != nil {
		fmt.Println(err)
		return false
	}
	return exists
}