package cmd

import(
	"github.com/spf13/cobra"
	"fmt"
)

var initCmd1 = &cobra.Command{
	Use:   "init-i",
	Short: "Baixa um projeto Spring Boot da API pública",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executando o comando init-i...")
		// Adicione aqui a lógica para o comando init-i, se necessário
	},
}

func init(){
	rootCmd.AddCommand(initCmd1)
}