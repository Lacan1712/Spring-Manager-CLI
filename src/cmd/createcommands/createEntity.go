package createcommands;

import(
	"smc/src/providers/archive"
)

func CreateEntity(name string){
	archive.CarregarEntity(name)
}