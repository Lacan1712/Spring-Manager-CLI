package archive

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"smc/src/services/database"
	"smc/src/services/database/models"

)

type EntityFB struct {
	PackageName string
	EntityName  string
	Columns     []models.Column
}

// Mapear tipos SQL para tipos Java
func mapSQLToJavaType(sqlType string) string {
    switch sqlType {
    case "integer", "int", "bigint", "smallint":
        return "Integer"
    case "boolean":
        return "boolean"
    case "character varying", "varchar", "text", "char":
        return "String"
    case "timestamp", "date", "datetime":
        return "LocalDate"
    case "float", "real", "double":
        return "double"
    default:
        return "String"
    }
}

func CarregarEntityFromDB(entityPath, connectionName string) {
	templatePath := setupEntityPathsFB()

	entityPath = normalizeEntityPathFB(entityPath)

	dir, entityName := extractDirectoryAndEntityNameFB(entityPath)

	packageName := convertDirToPackageNameEntityFB(dir)

	createDirectoryIfNotExistsEntityFB(dir)

	// Carregar as colunas do banco de dados
	columns, err := database.ListColumnsDB(connectionName, entityName)
	if err != nil {
		log.Fatalf("Erro ao listar as colunas: %v", err)
	}

	tmpl := loadTemplateEntityFB(templatePath)

	// Escrever o arquivo da entidade baseado no template
	writeEntityFileFB(tmpl, dir, entityName, packageName, columns)
}

func setupEntityPathsFB() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Erro ao obter o caminho do executável: %v", err)
	}

	exeDir := filepath.Dir(exePath)
	templatePath := filepath.Join(exeDir, "src", "templates", "database", "Entity.tpl")
	return templatePath
}

func normalizeEntityPathFB(entityPath string) string {
	if !filepath.IsAbs(entityPath) {
		entityPath = filepath.Join(".", entityPath) // Usar diretório local
	}
	return strings.ReplaceAll(entityPath, ".", "/")
}

func extractDirectoryAndEntityNameFB(entityPath string) (string, string) {
	parts := strings.Split(entityPath, "/")
	file := parts[len(parts)-1] // O último item é o nome do arquivo
	dir := filepath.Join(parts[:len(parts)-1]...) // O restante forma o diretório

	entityName := strings.TrimSuffix(file, filepath.Ext(file))
	if entityName == "" {
		entityName = "Entity" // Nome padrão caso não seja especificado
	}
	return dir, entityName
}

func convertDirToPackageNameEntityFB(dir string) string {
	packageName := strings.ReplaceAll(filepath.ToSlash(dir), "/", ".")
	return strings.Trim(strings.ToLower(packageName), ".")
}

func createDirectoryIfNotExistsEntityFB(dir string) {
	if dir == "" {
		dir = "entity"
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalf("Erro ao criar o diretório %s: %v", dir, err)
	}
}

func loadTemplateEntityFB(templatePath string) *template.Template {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Erro ao carregar o template: %v", err)
	}
	return tmpl
}

func writeEntityFileFB(tmpl *template.Template, dir, entityName, packageName string, columns []models.Column) {
	outputFilePath := filepath.Join(dir, entityName+".java")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo de saída: %v", err)
	}
	defer outputFile.Close()

	for i := range columns {
        columns[i].Type = mapSQLToJavaType(columns[i].Type)
    }

	data := EntityFB{
		PackageName: packageName,
		EntityName:  entityName,
		Columns:     columns, // Passar as colunas para o template
	}

	if err := tmpl.Execute(outputFile, data); err != nil {
		log.Fatalf("Erro ao executar o template: %v", err)
	}

	fmt.Printf("Entity %s criado com sucesso em %s\n", entityName, outputFilePath)
}
