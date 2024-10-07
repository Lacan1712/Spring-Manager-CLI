package databasecommands

import (
    db "smc/src/services/database"
)

func ListTables(connectionName string) {
   db.ListTablesDB(connectionName)
}
