package archive

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "text/template"
)

type Repository struct {
    PackageName    string
    RepositoryName string
}

func CarregarRepository(repositoryPath string) {
    templatePath := setupRepositoryPaths()

    repositoryPath = normalizeRepositoryPath(repositoryPath)

    dir, repositoryName := extractDirectoryAndRepositoryName(repositoryPath)

    packageName := convertDirToPackageNameRepository(dir)

    createDirectoryIfNotExistsRepository(dir)

    tmpl := loadTemplateRepository(templatePath)

    writeRepositoryFile(tmpl, dir, repositoryName, packageName)
}

func setupRepositoryPaths() string {
    exePath, err := os.Executable()
    if err != nil {
        log.Fatalf("Erro ao obter o caminho do executável: %v", err)
    }

    exeDir := filepath.Dir(exePath)
    templatePath := filepath.Join(exeDir, "src", "templates", "repository", "Repository.tpl")
    return templatePath
}

func normalizeRepositoryPath(repositoryPath string) string {
    if !filepath.IsAbs(repositoryPath) {
        repositoryPath = filepath.Join(".", repositoryPath) // Usar diretório local
    }
    return strings.ReplaceAll(repositoryPath, ".", "/")
}

func extractDirectoryAndRepositoryName(repositoryPath string) (string, string) {
    parts := strings.Split(repositoryPath, "/")
    file := parts[len(parts)-1] // O último item é o nome do arquivo
    dir := filepath.Join(parts[:len(parts)-1]...) // O restante forma o diretório

    repositoryName := strings.TrimSuffix(file, filepath.Ext(file))
    if repositoryName == "" {
        repositoryName = "Repository" // Nome padrão caso não seja especificado
    }
    return dir, repositoryName
}

func convertDirToPackageNameRepository(dir string) string {
    packageName := strings.ReplaceAll(filepath.ToSlash(dir), "/", ".")
    return strings.Trim(strings.ToLower(packageName), ".")
}

func createDirectoryIfNotExistsRepository(dir string) {
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        log.Fatalf("Erro ao criar o diretório %s: %v", dir, err)
    }
}

func loadTemplateRepository(templatePath string) *template.Template {
    tmpl, err := template.ParseFiles(templatePath)
    if err != nil {
        log.Fatalf("Erro ao carregar o template: %v", err)
    }
    return tmpl
}

func writeRepositoryFile(tmpl *template.Template, dir, repositoryName, packageName string) {
    outputFilePath := filepath.Join(dir, repositoryName+".java")
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        log.Fatalf("Erro ao criar o arquivo de saída: %v", err)
    }
    defer outputFile.Close()

    data := Repository{
        PackageName: packageName,
        RepositoryName:  repositoryName,
    }

    if err := tmpl.Execute(outputFile, data); err != nil {
        log.Fatalf("Erro ao executar o template: %v", err)
    }

    fmt.Printf("Repository %s criado com sucesso em %s\n", repositoryName, outputFilePath)
}
