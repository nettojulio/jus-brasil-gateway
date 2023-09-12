package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nettojulio/jusbrasil-crawler/client"
)

type RequestBody struct {
	Url string `json:"url"`
}

func Start() {
	http.HandleFunc("/", handlePostRequest)

	port := 8080
	fmt.Printf("Server running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Server start error:", err)
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error to read request body", http.StatusInternalServerError)
		return
	}

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error to read JSON", http.StatusBadRequest)
		return
	}

	// Acesse a propriedade desejada na estrutura
	propriedade := requestBody.Url

	// Verifica se a propriedade foi recebida
	if propriedade == "" {
		http.Error(w, "Property 'url' does not exists", http.StatusBadRequest)
		return
	}

	html, err := client.GetRawHTML(propriedade)
	if err != nil {
		http.Error(w, "Error to get HTML", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(html)
}
