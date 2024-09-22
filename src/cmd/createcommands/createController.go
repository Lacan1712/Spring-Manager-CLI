package createcommands;

import(
	"SMC/src/providers/archive"
)

func CreateController(name string){
	archive.CarregarController(name)
}