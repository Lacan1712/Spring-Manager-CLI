package createcommands;

import(
	"smc/src/providers/archive"
)

func CreateEntity(name string, databaseName string){
	if databaseName == ""{
		archive.CarregarEntity(name)
	} else {
		archive.CarregarEntityFromDB(name, databaseName)
	}
}