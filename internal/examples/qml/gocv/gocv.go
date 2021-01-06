package main

import (
	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/gui"
	"github.com/dev-drprasad/qt/quick"

	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"sync"
)

func init() {
	GoCV_QmlRegisterType2("CustomModule", 1, 0, "GoCV")
}

type GoCV struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`
	_ bool   `property:"detect"`

	image *gui.QImage
	data  []byte
	mutex *sync.Mutex
}

func (p *GoCV) init() {
	p.SetDetect(false)

	xmlFile := "./facedetect.xml"

	//setup gocv
	webcam, _ := gocv.OpenVideoCapture(0)
	// prepare image matrix
	img := gocv.NewMat()

	//setup painting
	p.image = gui.NewQImage()
	p.mutex = new(sync.Mutex)
	p.SetRenderTarget(quick.QQuickPaintedItem__FramebufferObject)
	p.ConnectPaint(p.paint)

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	go func() {
		for {
			webcam.Read(&img)
			if img.Empty() {
				fmt.Println("empty, continuing")
				continue
			}
			p.mutex.Lock()

			if p.IsDetect() {
				// detect faces
				rects := classifier.DetectMultiScale(img)
				fmt.Printf("found %d faces\n", len(rects))

				// draw a rectangle around each face on the original image,
				// along with text identifing as "Human"
				for _, r := range rects {
					gocv.Rectangle(&img, r, blue, 3)

					size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
					pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
					gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
				}
			}

			p.data, _ = gocv.IMEncode(".jpg", img)

			p.mutex.Unlock()
			p.Update()
		}
	}()
}

func (p *GoCV) paint(painter *gui.QPainter) {
	p.mutex.Lock()
	p.image.LoadFromData(p.data, len(p.data), "JPG")
	p.mutex.Unlock()

	painter.DrawImage6(core.NewQRect4(0, 0, int(p.Width()), int(p.Height())), p.image)
}
