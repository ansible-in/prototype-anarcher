package ansible

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"log"
	"os"
)

var DBSessions []*r.Session

const (
	DB_NAME    = "ansible"
	T_MESSAGES = "messages"
	T_CHANNELS = "channels"
	T_USERS    = "users"
)

func InitDB() {
	addr := os.Getenv("DB_1_PORT_28015_TCP_ADDR")
	if addr == "" {
		addr = "localhost"
	}
	port := os.Getenv("DB_1_PORT_28015_TCP_PORT")
	if port == "" {
		port = "28015"
	}

	addr = fmt.Sprintf("%s:%s", addr, port)

	session, err := r.Connect(r.ConnectOpts{
		Address: addr,
	})

	//DB
	err = r.DbCreate(DB_NAME).Exec(session)
	if err != nil {
		log.Println(err)
	}

	//Tables
	tableNames := []string{T_MESSAGES, T_USERS, T_CHANNELS}
	for _, tableName := range tableNames {
		err = r.Db(DB_NAME).TableCreate(tableName).Exec(session)
		if err != nil {
			log.Println(err)
		}
	}

	DBSessions = append(DBSessions, session)

}
