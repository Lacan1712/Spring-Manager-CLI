package cmd

import (
	createProject "SMC/src/providers/downloadSpring"
	ascii_apresentation "SMC/src/ascii_texts"
	"fmt"

	"github.com/spf13/cobra"
)

// Comando init
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Baixa um projeto Spring Boot da API pública",
	Run: func(cmd *cobra.Command, args []string) {
		custom, _ := cmd.Flags().GetBool("custom")
		ascii_apresentation.Apresentation()
		if custom {
			fmt.Println("Instalação customizada!")

			// Solicitar informações do usuário
			var groupId, artifactId, name, description, packageName, version, bootVersion string

			fmt.Print("Digite o Group ID: ")
			fmt.Scanln(&groupId)

			fmt.Print("Digite o Artifact ID: ")
			fmt.Scanln(&artifactId)

			fmt.Print("Digite o Nome do Projeto: ")
			fmt.Scanln(&name)

			fmt.Print("Digite a Descrição do Projeto: ")
			fmt.Scanln(&description)

			fmt.Print("Digite o Package Name: ")
			fmt.Scanln(&packageName)

			fmt.Print("Digite a Versão do Projeto: ")
			fmt.Scanln(&version)

			fmt.Print("Digite a Versão do Spring (deixe vazio se não quiser especificar): ")
			fmt.Scanln(&bootVersion)

			config := createProject.ProjectConfig{
				GroupId:     groupId,
				ArtifactId:  artifactId,
				Name:        name,
				Description: description,
				PackageName: packageName,
				Version:     version,
				BootVersion: bootVersion,
			}
			createProject.DownloadSpringBootProject(config)
		} else {
			fmt.Println("Baixando o projeto Spring Boot com configuração padrão...")
			fmt.Print("Baixando projeto spring...")
			config := createProject.ProjectConfig{
				GroupId:     "smc.example",
				ArtifactId:  "A project create by SMC",
				Name:        "smc-demo",
				Description: "A project create by SMC",
				PackageName: "smc.example.demo",
				Version:     "0.0.1-SNAPSHOT",
				BootVersion: "",
			}
			createProject.DownloadSpringBootProject(config)
		}
	},
}

func init() {
	initCmd.Flags().BoolP("custom", "c", false, "Ativa o comportamento personalizado")
	RootCmd.AddCommand(initCmd)
}