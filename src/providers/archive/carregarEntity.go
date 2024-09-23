package archive

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "text/template"
    "log"
)

type Entity struct {
    PackageName string
    EntityName   string
}

func CarregarEntity(entityPath string) {
    //Caminho do executável go
    exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Erro ao obter o caminho do executável: %v", err)
	}

    exeDir := filepath.Dir(exePath)
    templatePath := filepath.Join(exeDir,"src", "templates", "entity", "Entity.tpl")

    // Se o usuário forneceu apenas o nome do repository (sem caminho)
    if !strings.Contains(entityPath, "/") && !strings.Contains(entityPath, "\\") {
        entityPath = "./" + entityPath // Usar diretório local
    }

    // Divide o caminho em diretório e nome do entityPath
    dir, file := filepath.Split(entityPath)

    // Separa o nome do arquivo e remove a extensão, se houver
    entityName := strings.TrimSuffix(file, filepath.Ext(file))

	// Nome padrão caso não seja especificado
	if (entityName == ""){
		entityName = "Entity"
	}

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
    data := Entity{
        PackageName: packageName,
        EntityName:   entityName,
    }

    // Cria o arquivo de saída no diretório especificado
    outputFilePath := filepath.Join(dir, entityName+".java")
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

    fmt.Printf("repository %s criado com sucesso em %s\n", entityName, outputFilePath)
}
