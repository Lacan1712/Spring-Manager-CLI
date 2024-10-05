package models

type Connections struct {
	DriveDatabase string `json:"driveDatabase"` 
    DatabaseName  string `json:"databaseName"`
    Host   		  string `json:"host"`
	Port   		  string `json:"port"`
	Username 	  string `json:"username"`
	Password 	  string `json:"password"`
}

type Database struct {
    Connections []Connections `json:"connections"`
}