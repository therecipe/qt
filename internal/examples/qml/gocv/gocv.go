package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"

	"gocv.io/x/gocv"
	"sync"
)

func init() {
	GoCV_QmlRegisterType2("CustomModule", 1, 0, "GoCV")
}

type GoCV struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`

	image *gui.QImage
	data  []byte
	mutex *sync.Mutex
}

func (p *GoCV) init() {

	//setup gocv
	webcam, _ := gocv.OpenVideoCapture(0)
	img := gocv.NewMat()

	//setup painting
	p.image = gui.NewQImage()
	p.mutex = new(sync.Mutex)
	p.SetRenderTarget(quick.QQuickPaintedItem__FramebufferObject)
	p.ConnectPaint(p.paint)

	go func() {
		for {
			webcam.Read(&img)
			p.mutex.Lock()
			p.data, _ = gocv.IMEncode(".jpg", img)
			p.mutex.Unlock()
			p.UpdateDefault()
		}
	}()
}

func (p *GoCV) paint(painter *gui.QPainter) {
	p.mutex.Lock()
	p.image.LoadFromData(p.data, len(p.data), "JPG")
	p.mutex.Unlock()

	painter.DrawImage6(core.NewQRect4(0, 0, int(p.Width()), int(p.Height())), p.image)
}
