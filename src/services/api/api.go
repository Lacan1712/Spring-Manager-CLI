package api

import (
	"io"
	"log"
	"net/http"
)

func Get(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao realizar chamada para %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Erro: Status da resposta: %v", resp.Status)
	}

	return resp.Body, nil
}
