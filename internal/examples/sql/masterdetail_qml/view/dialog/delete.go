// +build !qml

package dialog

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

type deleteDialogController struct {
	core.QObject

	_ func() `constructor:"init"`

	//<-controller
	_ func(title, artist string) `signal:"deleteAlbumShowRequest"`

	//->controller
	_ func() `signal:"deleteAlbumCommand"`
}

func (d *deleteDialogController) init() {

	//<-controller
	controller.Instance.ConnectDeleteAlbumShowRequest(d.deleteAlbumShowRequest)

	//->controller
	d.ConnectDeleteAlbumCommand(controller.Instance.DeleteAlbumCommand)
}

func (d *deleteDialogController) deleteAlbumShowRequest(title, artist string) {
	var button widgets.QMessageBox__StandardButton
	button = widgets.QMessageBox_Question(nil, "Delete Album",
		fmt.Sprintf("Are you sure you want to delete '%v' by '%v'?", title, artist),
		widgets.QMessageBox__Yes|widgets.QMessageBox__No, 0)

	if button == widgets.QMessageBox__Yes {
		d.DeleteAlbumCommand()
	} else {
		widgets.QMessageBox_Information(nil, "Delete Album",
			"Select the album you want to delete.", 0, 0)
	}
}
