// No menu DEMOS > Echo Test do site http://websocket.org é possível o funcionamento
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Uma das propriedades de o upgrader é checar a origem do trafego
var upgrader = websocket.Upgrader{

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {

	http.HandleFunc("/", handler)     // Declaramos o path e o handler
	http.ListenAndServe(":3000", nil) // Indicamos a porta em que nosso WebServer irá responder

}

// Tratamos nossa requisição: o writer irá permitir definir o que desejamos escrever de volta para
// o client que eniou a requisição
func handler(writer http.ResponseWriter, request *http.Request) {

	socket, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println(err)
	}

	for {
		// Le a mensagem recebida via WebSocket
		msgType, msg, err := socket.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}

		// Loga no console do WebWriter
		fmt.Println("Mensagem recebida: ", string(msg))

		// Devolve a mensagem recebida de volta para o cliente
		err = socket.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
