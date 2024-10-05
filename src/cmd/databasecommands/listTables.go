package databasecommands

import (
    "fmt"
	"log"
	jsonservice "smc/src/services/json"
	db "smc/src/services/database"
	_ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"               
    _ "github.com/denisenkom/go-mssqldb"
	models "smc/src/services/database/models"
)

func ListTables() {
	var database models.Database
	
	err := jsonservice.MappingStructToJson("/home/rodrigo/Documentos Local/Projetos/Go/SMC/src/cmd/databasecommands/database.json", &database)
	if err != nil{
        fmt.Printf("Error reading JSON: %v", err)
	}

    for _, conn := range database.Connections {
        db, err := db.ConnectToDatabase(conn) 
        if err != nil {
            fmt.Printf("Failed to connect to %s: %v\n", conn.DatabaseName, err)
            continue
        }

        fmt.Printf("Connected to database: %s\n", conn.DatabaseName)

		rows, err := db.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		fmt.Println("Tabelas encontradas:")
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				log.Fatal(err)
			}
			fmt.Println(tableName)
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
		db.Close()
    }
	
}