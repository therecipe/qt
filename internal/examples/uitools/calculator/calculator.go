package main

import (
	"os"

	"github.com/bluszcz/cutego/widgets"

	"github.com/bluszcz/cutego/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
