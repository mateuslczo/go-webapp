package main

import (
	"env"
	"fmt"
	"html/template"
	"net/http"
)

// Renderiza pagina principal do site
func index(wResp http.ResponseWriter, wReq *http.Request) {
	webpage, _ := template.ParseFiles("html/index.html")
	data := map[string]string{"Title": "Go Store :)"}

	wResp.WriteHeader(http.StatusOK)
	webpage.Execute(wResp, data)

}

func main() {
	port := env.GetEnvOrDefault("PORT", "4200")
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port %s.\n", port)
	http.ListenAndServe(":"+port, nil)
}
