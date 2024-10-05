package database

import(
	"fmt"
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

