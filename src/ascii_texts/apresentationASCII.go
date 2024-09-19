package ascii_apresentation

import (
	"github.com/common-nighthawk/go-figure"
)

func Apresentation() {
	styles := []string{"standard", "big", "small", "banner", "doom", "digital", "block"}

	fig := figure.NewFigure("SPRING MANAGER CLI", styles[1], true)
	fig.Print()

}
