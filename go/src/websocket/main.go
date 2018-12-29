package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)     // Declaramos o path e o handler
	http.ListenAndServe(":3000", nil) // Indicamos a porta em que nosso WebServer irá responder

}

// Tratamos nossa requisição: o writer irá permitir definir o que desejamos escrever de volta para
// o client que eniou a requisição
func handler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprint(writer, "Hello")

}
