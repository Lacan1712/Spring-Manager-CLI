package createcommands;

import(
	"smc/src/providers/archive"
)

func CreateService(name string){
	archive.CarregarEntity(name)
}