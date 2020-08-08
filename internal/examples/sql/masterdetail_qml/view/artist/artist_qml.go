// +build qml

package artist

import (
	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/quick"

	"github.com/StarAurryon/qt/internal/examples/sql/masterdetail_qml/controller"
)

func init() {
	artistController_QmlRegisterType2("Artist", 1, 0, "ArtistController")
}

type artistController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"listModel"`

	//qml->controller
	_ func(row int) `signal:"changeArtist"`
}

func (a *artistController) init() {
	a.SetListModel(controller.Instance.ArtistModel())

	//qml<-controller
	a.ListModel().ConnectRowsInserted(func(*core.QModelIndex, int, int) {
		model := a.ListModel()
		a.SetListModel(nil)
		a.SetListModel(model)
	})

	//qml->controller
	a.ConnectChangeArtist(controller.Instance.ChangeArtist)
}
