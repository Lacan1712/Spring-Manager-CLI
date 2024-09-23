package archive

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "text/template"
)

type Controller struct {
    PackageName    string
    ControllerName string
}

func CarregarController(controllerPath string) {
    templatePath := setupControllerPaths()

    controllerPath = normalizeControllerPath(controllerPath)

    dir, controllerName := extractDirectoryAndControllerName(controllerPath)

    packageName := convertDirToPackageNameController(dir)

    createDirectoryIfNotExistsController(dir)

    tmpl := loadTemplateController(templatePath)

    writeControllerFile(tmpl, dir, controllerName, packageName)
}

func setupControllerPaths() string {
    exePath, err := os.Executable()
    if err != nil {
        log.Fatalf("Erro ao obter o caminho do executável: %v", err)
    }

    exeDir := filepath.Dir(exePath)
    templatePath := filepath.Join(exeDir, "src", "templates", "controllers", "Controller.tpl")
    return templatePath
}

func normalizeControllerPath(controllerPath string) string {
    if !filepath.IsAbs(controllerPath) {
        controllerPath = filepath.Join(".", controllerPath) // Usar diretório local
    }
    return strings.ReplaceAll(controllerPath, ".", "/")
}

func extractDirectoryAndControllerName(controllerPath string) (string, string) {
    parts := strings.Split(controllerPath, "/")
    file := parts[len(parts)-1] // O último item é o nome do arquivo
    dir := filepath.Join(parts[:len(parts)-1]...) // O restante forma o diretório

    controllerName := strings.TrimSuffix(file, filepath.Ext(file))
    if controllerName == "" {
        controllerName = "Controller" // Nome padrão caso não seja especificado
    }
    return dir, controllerName
}

func convertDirToPackageNameController(dir string) string {
    packageName := strings.ReplaceAll(filepath.ToSlash(dir), "/", ".")
    return strings.Trim(strings.ToLower(packageName), ".")
}

func createDirectoryIfNotExistsController(dir string) {
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        log.Fatalf("Erro ao criar o diretório %s: %v", dir, err)
    }
}

func loadTemplateController(templatePath string) *template.Template {
    tmpl, err := template.ParseFiles(templatePath)
    if err != nil {
        log.Fatalf("Erro ao carregar o template: %v", err)
    }
    return tmpl
}

func writeControllerFile(tmpl *template.Template, dir, controllerName, packageName string) {
    outputFilePath := filepath.Join(dir, controllerName+".java")
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        log.Fatalf("Erro ao criar o arquivo de saída: %v", err)
    }
    defer outputFile.Close()

    data := Controller{
        PackageName:    packageName,
        ControllerName: controllerName,
    }

    if err := tmpl.Execute(outputFile, data); err != nil {
        log.Fatalf("Erro ao executar o template: %v", err)
    }

    fmt.Printf("Controller %s criado com sucesso em %s\n", controllerName, outputFilePath)
}
