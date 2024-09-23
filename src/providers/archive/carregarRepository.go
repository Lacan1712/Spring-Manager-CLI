package archive

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "text/template"
    "log"
)

type Repository struct {
    PackageName string
    RepositoryName   string
}

func CarregarRepository(repositoryPath string) {
    //Caminho do executável go
    exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Erro ao obter o caminho do executável: %v", err)
	}

    exeDir := filepath.Dir(exePath)
    templatePath := filepath.Join(exeDir,"src", "templates", "repository", "Repository.tpl")

    // Se o usuário forneceu apenas o nome do repository (sem caminho)
    if !strings.Contains(repositoryPath, "/") && !strings.Contains(repositoryPath, "\\") {
        repositoryPath = "./" + repositoryPath // Usar diretório local
    }

    // Divide o caminho em diretório e nome do repositoryPath
    dir, file := filepath.Split(repositoryPath)

    // Separa o nome do arquivo e remove a extensão, se houver
    repositoryName := strings.TrimSuffix(file, filepath.Ext(file))

    // Converte o caminho do diretório para um formato de pacote
    packageName := strings.ReplaceAll(filepath.ToSlash(dir), "/", ".")

    // Ajusta o nome do pacote para não conter um ponto no início ou no final
    if strings.HasPrefix(packageName, ".") {
        packageName = strings.ToLower(packageName[1:]) 
    }
    if strings.HasSuffix(packageName, ".") {
        packageName = strings.ToLower(packageName[:len(packageName)-1])
    }

    // Cria o diretório, se não existir
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        log.Fatalf("Erro ao criar o diretório %s: %v", dir, err)
    }

    // Carrega o template
    tmpl, err := template.ParseFiles(templatePath)
    if err != nil {
        log.Fatalf("Erro ao carregar o template: %v", err)
    }

    // Prepara os dados para o template
    data := Repository{
        PackageName: packageName,
        RepositoryName:   repositoryName,
    }

    // Cria o arquivo de saída no diretório especificado
    outputFilePath := filepath.Join(dir, repositoryName+".java")
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        log.Fatalf("Erro ao criar o arquivo de saída: %v", err)
    }
    defer outputFile.Close()

    // Executa o template e escreve no arquivo
    err = tmpl.Execute(outputFile, data)
    if err != nil {
        log.Fatalf("Erro ao executar o template: %v", err)
    }

    fmt.Printf("repository %s criado com sucesso em %s\n", repositoryName, outputFilePath)
}
