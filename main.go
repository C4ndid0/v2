package main

import (
	"ACBrLibCEP-API-Go/acbr"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CEPResponse struct {
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Municipio string `json:"municipio"`
	UF string `json:"uf"`
	CEP string `json:"cep"`
	CodigoIBGE string `json:"codigo_ibge"`
	CodigoIBGEMunicipio string `json:"codigo_ibge_municipio"`
	CodigoIBGEUF string `json:"codigo_ibge_uf"`
}

func buscarCEPHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP não fornecido", http.StatusBadRequest)
		return
	}

	// Buscar informações do CEP
	result, err := acbr.BuscarPorCEP(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Transformar a resposta em JSON
	var response CEPResponse
	if err := json.Unmarshal([]byte(result), &response); err != nil {
		http.Error(w, "Erro ao processar a resposta do CEP", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Inicializar a biblioteca ACBr
	err := acbr.Inicializar()
	if err != nil {
		log.Fatalf("Erro ao inicializar a biblioteca: %v", err)
	}
	defer acbr.Finalizar()

	// Configurar o servidor HTTP
	http.HandleFunc("/cep", buscarCEPHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
