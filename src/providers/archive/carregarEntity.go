package archive

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "text/template"
)

type Entity struct {
    PackageName    string
    EntityName string
}

func CarregarEntity(entityPath string) {
    templatePath := setupEntityPaths()

    entityPath = normalizeEntityPath(entityPath)

    dir, entityName := extractDirectoryAndEntityName(entityPath)

    packageName := convertDirToPackageNameEntity(dir)

    createDirectoryIfNotExistsEntity(dir)

    tmpl := loadTemplateEntity(templatePath)

    writeEntityFile(tmpl, dir, entityName, packageName)
}

func setupEntityPaths() string {
    exePath, err := os.Executable()
    if err != nil {
        log.Fatalf("Erro ao obter o caminho do executável: %v", err)
    }

    exeDir := filepath.Dir(exePath)
    templatePath := filepath.Join(exeDir, "src", "templates", "entity", "Entity.tpl")
    return templatePath
}

func normalizeEntityPath(entityPath string) string {
    if !filepath.IsAbs(entityPath) {
        entityPath = filepath.Join(".", entityPath) // Usar diretório local
    }
    return strings.ReplaceAll(entityPath, ".", "/")
}

func extractDirectoryAndEntityName(entityPath string) (string, string) {
    parts := strings.Split(entityPath, "/")
    file := parts[len(parts)-1] // O último item é o nome do arquivo
    dir := filepath.Join("src", "main", "java",filepath.Join(parts[:len(parts)-1]...))

    entityName := strings.TrimSuffix(file, filepath.Ext(file))
    if entityName == "" {
        entityName = "Entity" // Nome padrão caso não seja especificado
    }
    return dir, entityName
}

func convertDirToPackageNameEntity(dir string) string {
    packageName := strings.ReplaceAll(filepath.ToSlash(dir), "/", ".")
    return strings.Trim(strings.ToLower(packageName), ".")
}

func createDirectoryIfNotExistsEntity(dir string) {
    if(dir == ""){
        dir = "entity"
    }
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        log.Fatalf("Erro ao criar o diretório %s: %v", dir, err)
    }
}

func loadTemplateEntity(templatePath string) *template.Template {
    tmpl, err := template.ParseFiles(templatePath)
    if err != nil {
        log.Fatalf("Erro ao carregar o template: %v", err)
    }
    return tmpl
}

func writeEntityFile(tmpl *template.Template, dir, entityName, packageName string) {
    outputFilePath := filepath.Join(dir, entityName+".java")
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        log.Fatalf("Erro ao criar o arquivo de saída: %v", err)
    }
    defer outputFile.Close()

    data := Entity{
        PackageName: packageName,
        EntityName:  entityName,
    }

    if err := tmpl.Execute(outputFile, data); err != nil {
        log.Fatalf("Erro ao executar o template: %v", err)
    }

    fmt.Printf("Entity %s criado com sucesso em %s\n", entityName, outputFilePath)
}
