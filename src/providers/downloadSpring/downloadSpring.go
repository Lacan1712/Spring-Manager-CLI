package providers

import (
	"SMC/src/providers/zip"
	"SMC/src/services/api"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
)

type ProjectConfig struct {
	GroupId     string
	ArtifactId  string
	Name        string
	Description string
	PackageName string
	Version     string
	BootVersion string
}

func DownloadSpringBootProject(config ProjectConfig) {

	baseURL := "https://start.spring.io/starter.zip"
	params := url.Values{}
	params.Add("type", "maven-project")
	params.Add("groupId", config.GroupId)
	params.Add("artifactId", config.ArtifactId)
	params.Add("name", config.Name)
	params.Add("description", config.Description)
	params.Add("packageName", config.PackageName)
	params.Add("version", config.Version)

	// Adiciona a versão do Spring Boot, se especificada
	if config.BootVersion != "" {
		params.Add("bootVersion", config.BootVersion)
	}

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	body, err := api.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição: %v", err)
	}

	defer body.Close()

	outFile, err := os.Create("demo.zip")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, body)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
		return
	}

	if err := zip.Unzip("demo.zip", "./App"); err != nil {
		fmt.Println("Erro ao descompactar o arquivo:", err)
		return
	}

	if err := os.Remove("demo.zip"); err != nil {
		fmt.Println("Erro ao excluir o arquivo:", err)
		return
	}

	fmt.Println("Projeto Spring Boot baixado com sucesso como demo.zip!")

}
