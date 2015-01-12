package ws

import (
	"encoding/json"
	"github.com/ansible-in/prototype-anarcher"
	"github.com/ansible-in/prototype-anarcher/router"
	"github.com/gorilla/mux"
	"github.com/igm/pubsub"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"

	r "github.com/dancannon/gorethink"
)

var chat pubsub.Publisher

type Content struct {
	Id       string `gorethink:"id,omitempty" json:"id,omitempty"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func Handler() *mux.Router {
	m := router.WebSocket()

	//WS
	//Not Good.
	m.Get(router.WS).Handler(sockjs.NewHandler("/ws", sockjs.DefaultOptions, wsHandler))

	return m
}

func wsHandler(session sockjs.Session) {
	log.Println("new ws session established")
	var closedSession = make(chan struct{})

	chat.Publish("[info]  new participant joined chat")
	defer chat.Publish("[info] participant left chat")

	go func() {
		reader, _ := chat.SubChannel(nil)
		for {
			select {
			case <-closedSession:
				return
			case content := <-reader:
				c, err := json.Marshal(content)
				if err != nil {
					log.Println(err)
					continue
				}

				if err := session.Send(string(c)); err != nil {
					return
				}
			}
		}
	}()

	for {
		if msg, err := session.Recv(); err == nil {
			var c Content
			err := json.Unmarshal([]byte(msg), &c)
			if err != nil {
				log.Println(err)
				continue
			}

			dbSession := ansible.DBSessions[0]
			_, err = r.Db(ansible.DB_NAME).Table(ansible.T_MESSAGES).Insert(c).RunWrite(dbSession)
			if err != nil {
				//TODO
				log.Println(err)
			}

			chat.Publish(c)
			continue
		}
		break
	}
	close(closedSession)
	log.Println("session closed")

}
