package createcommands;

import(
	"smc/src/providers/archive"
)

func CreateController(name string){
	archive.CarregarController(name)
}