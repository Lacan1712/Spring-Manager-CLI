package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"smc/src/cmd/createcommands"
)

var (
	repositoryName string
	controllerName string
	entityName    string
	databaseName string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Cria recursos para a aplicação Spring",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case repositoryName != "":
			createcommands.CreateRepository(repositoryName) // Chama a função de criar repositório
		case controllerName != "":
			createcommands.CreateController(controllerName) // Chama a função de criar controlador
		case entityName != "":
			createcommands.CreateEntity(entityName, databaseName) // Chama a função de criar serviço
		default:
			fmt.Println("Por favor, forneça um nome para repository, controller ou service.")
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&repositoryName, "repository", "r", "", "Cria um repositório com o nome especificado")
	createCmd.Flags().StringVarP(&controllerName, "controller", "c", "", "Cria um controlador com o nome especificado")
	createCmd.Flags().StringVarP(&entityName, "entity", "e", "", "Cria uma entity com o nome especificado")
	createCmd.Flags().StringVarP(&databaseName, "database", "d", "", "Especifica o nome da conexão com o banco de dados ao criar a entity")

	rootCmd.AddCommand(createCmd)
}
