package cmd

import (
	"github.com/spf13/cobra"
	"smc/src/cmd/databasecommands"
)

var (
	listTables bool
	connectionName string
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Recursos para manipulação de banco de dados",
	Run: func(cmd *cobra.Command, args []string) {
			switch {
				case listTables:
					databasecommands.ListTables(connectionName)
			}
	},
}

func init() {
	databaseCmd.Flags().BoolVarP(&listTables, "listables", "",false,"lista as tabelas do banco")
	databaseCmd.Flags().StringVarP(&connectionName, "name", "n","","nome da conexão criada")

	rootCmd.AddCommand(databaseCmd)
}
