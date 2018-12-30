package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// - Rotas
func main() {

	// StrictSlash true permite trabalhar com nossas rotas da seguinte forma /rota e desta /rota/
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/", getAll).Methods("GET")

	rotas.HandleFunc("/persons", create).Methods("POST")

	// variável que define para qual porta faremos requisições
	var port = ":3000"

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))

}

// Módulo vulgo(Entidade) pessoa
type Person struct {
	Name string
}

// Preenchendo o módulo com tres valores
var persons = []Person{

	Person{Name: "Marcos"},
	Person{Name: "Lucenzo"},
	Person{Name: "Mariana"},
}

// Método listar todos
func getAll(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	json.NewEncoder(responseWriter).Encode(persons)

}

// Método criar pessoa
func create(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	// Configura o header com o Content-type
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Define uma variavel com valores do model
	var person Person

	// Recebe valores do request e passa para as variáveis body, err
	body, err := ioutil.ReadAll(httpRequest.Body)

	// valida se err é diferente de nill
	if err != nil {

		panic(err)

	}

	// valida se o conteudo não está vazio
	if err := httpRequest.Body.Close(); err != nil {

		panic(err)

	}

	// Retorna dados para o usuario
	if err := json.Unmarshal(body, &person); err != nil {

		responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
		responseWriter.WriteHeader(422)

		if err := json.NewEncoder(responseWriter).Encode(err); err != nil {

			panic(err)

		}

	}

	// Faz um parse do objeto json
	json.Unmarshal(body, &person)

	// adiciona valores no arrary
	persons = append(persons, person)

	// seta e retorna o status 201 de criado e o objeto enviado pelo usuário
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(responseWriter).Encode(person); err != nil {

		panic(err)
	}

}
