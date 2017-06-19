// +build qml

package detail

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

func init() {
	detailController_QmlRegisterType2("Detail", 1, 0, "DetailController")
}

type detailController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//<-controller
	_ func()                                `signal:"showImageLabel"`
	_ func(profileLabelText string)         `signal:"showArtistProfile"`
	_ func(title string, elements []string) `signal:"showTitleAndAlbumDetails"`
}

func (d *detailController) init() {

	//<-controller
	controller.Instance.ConnectShowImageLabel(d.ShowImageLabel)
	controller.Instance.ConnectShowArtistProfile(d.ShowArtistProfile)
	controller.Instance.ConnectShowTitleAndAlbumDetails(d.ShowTitleAndAlbumDetails)
}
