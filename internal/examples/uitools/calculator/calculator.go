package main

import (
	"os"

	"github.com/dev-drprasad/qt/widgets"

	"github.com/dev-drprasad/qt/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
