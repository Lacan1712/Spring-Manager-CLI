package createcommands;

import(
	"smc/src/providers/archive"
)

func CreateRepository(name string){
	archive.CarregarRepository(name)
}