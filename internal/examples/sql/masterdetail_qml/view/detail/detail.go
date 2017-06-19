// +build !qml

package detail

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

type detailController struct {
	widgets.QGroupBox

	_ func() `constructor:"init"`

	//<-controller
	_ func()                                `signal:"showImageLabel"`
	_ func(profileLabelText string)         `signal:"showArtistProfile"`
	_ func(title string, elements []string) `signal:"showTitleAndAlbumDetails"`

	//

	profileLabel *widgets.QLabel
	titleLabel   *widgets.QLabel
	iconLabel    *widgets.QLabel
	imageLabel   *widgets.QLabel
	trackList    *widgets.QListWidget
}

func (d *detailController) init() {

	d.profileLabel = widgets.NewQLabel(nil, 0)
	d.profileLabel.SetWordWrap(true)
	d.profileLabel.SetAlignment(core.Qt__AlignBottom)

	d.titleLabel = widgets.NewQLabel(nil, 0)
	d.titleLabel.SetWordWrap(true)
	d.titleLabel.SetAlignment(core.Qt__AlignBottom)

	d.iconLabel = widgets.NewQLabel(nil, 0)
	d.iconLabel.SetAlignment(core.Qt__AlignBottom | core.Qt__AlignRight)
	d.iconLabel.SetPixmap(gui.NewQPixmap5(":/images/icon.png", "", 0))

	d.imageLabel = widgets.NewQLabel(nil, 0)
	d.imageLabel.SetWordWrap(true)
	d.imageLabel.SetAlignment(core.Qt__AlignCenter)
	d.imageLabel.SetPixmap(gui.NewQPixmap5(":/images/image.png", "", 0))

	d.trackList = widgets.NewQListWidget(nil)

	layout := widgets.NewQGridLayout2()
	layout.AddWidget3(d.imageLabel, 0, 0, 3, 2, 0)
	layout.AddWidget(d.profileLabel, 0, 0, 0)
	layout.AddWidget(d.iconLabel, 0, 1, 0)
	layout.AddWidget3(d.titleLabel, 1, 0, 1, 2, 0)
	layout.AddWidget3(d.trackList, 2, 0, 1, 2, 0)
	layout.SetRowStretch(2, 1)
	d.SetLayout(layout)

	d.showImageLabel()

	//

	//<-controller
	controller.Instance.ConnectShowImageLabel(d.showImageLabel)
	controller.Instance.ConnectShowArtistProfile(d.showArtistProfile)
	controller.Instance.ConnectShowTitleAndAlbumDetails(d.showTitleAndAlbumDetails)
}

func (d *detailController) showImageLabel() {
	d.profileLabel.Hide()
	d.titleLabel.Hide()
	d.iconLabel.Hide()
	d.trackList.Hide()

	d.imageLabel.Show()
}

func (d *detailController) showArtistProfile(profileLabelText string) {
	d.profileLabel.SetText(profileLabelText)

	d.profileLabel.Show()
	d.iconLabel.Show()

	d.titleLabel.Hide()
	d.trackList.Hide()
	d.imageLabel.Hide()
}

func (d *detailController) showTitleAndAlbumDetails(title string, elements []string) {
	d.titleLabel.SetText(title)
	d.titleLabel.Show()
	d.trackList.Show()

	d.trackList.Clear()

	for _, v := range elements {
		item := widgets.NewQListWidgetItem(d.trackList, 0)
		item.SetText(v)
	}
}
