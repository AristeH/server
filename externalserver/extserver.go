package externalserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

// MessageCH структура содержащая
type MessageCH struct {
	Message *Message
	Client  *Client
}

// СН канал входящих сообщений
var СН = make(chan *MessageCH)

// VСН канал исходящих сообщение
var VСН = make(chan *MessageCH)

// Message структура содержащая
type Message struct {
	Sender     string   `json:"sender,omitempty"`
	Recipient  string   `json:"recipient,mitempty"`
	Content    string   `json:"content,omitempty"`
	Action     string   `json:"action,omitempty"`
	Parameters []string `json:"parameters,omitempty"`
}

// Client структура содержащая
// id уникальный номер клиента
// socket вебсокет
// send канал для обмена сообщениями с клиентом
type Client struct {
	ID     string
	socket *websocket.Conn
	send   chan []byte
}

//ClientManager - список клиентов
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Manager структура содержащая список подключенных клиентов
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

// Start функция  для работы с клиентами

func (manager *ClientManager) Start(s string) {
	for {
		select {
		case conn := <-manager.Register:
			fmt.Println("A new socket has connected....")
			manager.Clients[conn] = true
			//jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			
			s1 := "login;" + s

			manager.send([]byte(s1), conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				//close(conn.send)
				delete(manager.Clients, conn)
				//	jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				fmt.Println("A socket has disconnected....")
				//	manager.send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range Manager.Clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

// Init инициализация
func Init(sOpt string) {
	fmt.Println("Приложение запущено...")
	go Manager.Start("")

	http.HandleFunc("/telephon", wsPage)
	http.ListenAndServe(sOpt, nil)

}

func (c *Client) read() {

	defer func() {
		Manager.Unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.socket.Close()
			break
		}
		res1 := string(message)
		println("1! " + res1)
		res := Message{}

		json.Unmarshal(message, &res)
		//println(res.Parameters[0])
		println("3! " + res.Action)

		ms := MessageCH{&res, c}
		СН <- &ms //
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

var upgrader = websocket.Upgrader{
	EnableCompression: true,
} // use default options

func wsPage(res http.ResponseWriter, req *http.Request) {

	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	uu := "54467"
	client := &Client{ID: uu, socket: conn, send: make(chan []byte)}

	Manager.Register <- client

	go client.read()
	go client.write()
}
