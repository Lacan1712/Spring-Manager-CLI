package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"SMC/src/cmd/createcommands" // Ajuste para o caminho correto
)

var (
	repositoryName string
	controllerName string
	serviceName    string
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
		case serviceName != "":
			createcommands.CreateService(serviceName) // Chama a função de criar serviço
		default:
			fmt.Println("Por favor, forneça um nome para repository, controller ou service.")
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&repositoryName, "repository", "r", "", "Cria um repositório com o nome especificado")
	createCmd.Flags().StringVarP(&controllerName, "controller", "l", "", "Cria um controlador com o nome especificado")
	createCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Cria um serviço com o nome especificado")

	RootCmd.AddCommand(createCmd)
}
