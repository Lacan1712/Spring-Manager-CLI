package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smc",
	Short: "Spring Manager CLI para gerenciar aplicações Java Spring Boot",
	Long:  "Uma ferramenta de linha de comando para gerenciar projetos Spring Boot.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
