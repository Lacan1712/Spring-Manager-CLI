package database

import(
	"fmt"
    "log"
    jsonservice "smc/src/services/json"
	_ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"               
    _ "github.com/denisenkom/go-mssqldb"
    models "smc/src/services/database/models"
	"database/sql"
)

func ConnectToDatabase(conn models.Connections) (*sql.DB, error) {
    var dsn string 

    switch conn.DriveDatabase {
    case "MySQL":
        dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", conn.Username, conn.Password, conn.Host, conn.DatabaseName)
    case "postgres":
        dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conn.Host, conn.Port, conn.Username, conn.Password, conn.DatabaseName)
    case "SQLServer":
        dsn = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", conn.Username, conn.Password, conn.Host, conn.DatabaseName)
    default:
        return nil, fmt.Errorf("unsupported database type: %s", conn.DriveDatabase)
    }

    db, err := sql.Open(conn.DriveDatabase, dsn)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}


func ListTablesDB(connectionName string) {
    var database models.Database

    err := jsonservice.MappingStructToJson("/home/rodrigo/Documentos Local/Projetos/Go/SMC/src/cmd/databasecommands/database.json", &database)
    if err != nil {
        fmt.Printf("Error reading JSON: %v", err)
		return
    }


    // Encontrar a conexão correta com base no connectionName
    var conn *models.Connections
    for _, connection := range database.Connections {
        if connection.ConnectionName == connectionName {
            conn = &connection
            break
        }
    }

    if conn == nil {
        fmt.Printf("Conexão com o nome '%s' não foi encontrada.\n", connectionName)
		return
	}

	db, err := ConnectToDatabase(*conn) 
        if err != nil {
            fmt.Printf("Failed to connect to %s: %v\n", conn.DatabaseName, err)
			return
        }

        fmt.Printf("Connected to database: %s\n", conn.DatabaseName)

        schema := "public"
        if conn.Schema != "" {
            schema = conn.Schema
        }

        query := fmt.Sprintf(`SELECT table_name FROM information_schema.tables WHERE table_schema = '%s'`, schema)

        rows, err := db.Query(query)
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

func ListColumnsDB(connectionName, tableName string) ([]models.Column, error) {
	var database models.Database

	err := jsonservice.MappingStructToJson("/home/rodrigo/Documentos Local/Projetos/Go/SMC/src/cmd/databasecommands/database.json", &database)
	if err != nil {
		return nil, fmt.Errorf("Error reading JSON: %v", err)
	}

	// Encontrar a conexão correta com base no connectionName
	var conn *models.Connections
	for _, connection := range database.Connections {
		if connection.ConnectionName == connectionName {
			conn = &connection
			break
		}
	}

	if conn == nil {
		return nil, fmt.Errorf("Conexão com o nome '%s' não foi encontrada.", connectionName)
	}

	db, err := ConnectToDatabase(*conn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to %s: %v", conn.DatabaseName, err)
	}
	defer db.Close()

	schema := "public"
	if conn.Schema != "" {
		schema = conn.Schema
	}

	// Query para listar as colunas da tabela especificada
	query := fmt.Sprintf(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = '%s' AND table_name = '%s'`, schema, tableName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []models.Column
	for rows.Next() {
		var columnName, dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			return nil, err
		}
		columns = append(columns, models.Column{
			Name: columnName,
			Type: dataType,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}