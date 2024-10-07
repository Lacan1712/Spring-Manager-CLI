package models

type Connections struct {
	ConnectionName string `json:"connectionName"`
	DriveDatabase  string `json:"driveDatabase"` 
    DatabaseName   string `json:"databaseName"`
    Host   		   string `json:"host"`
	Port   		   string `json:"port"`
	Username 	   string `json:"username"`
	Password 	   string `json:"password"`
	Schema		   string `json:"schema"`
}

type Database struct {
    Connections []Connections `json:"connections"`
}

type Column struct {
	Name         string
	Type         string
	IsPrimaryKey bool
	IsNullable   bool
	IsUnique     bool
}