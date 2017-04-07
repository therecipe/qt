package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/printsupport"
	"github.com/therecipe/qt/quick"
)

type SlideView struct {
	quick.QQuickView

	_ func() `constructor:"init"`

	_ func(status quick.QQuickView__Status) `slot:"updateStatus"`
	_ func()                                `slot:"printCurrentSlide"`
	_ func()                                `slot:"goToNextSlide"`

	m_slidesLeft    int
	m_printedSlides int
	m_tid           int

	m_printer *printsupport.QPrinter
	m_painter *gui.QPainter
}

func (v *SlideView) init() {
	v.m_printer = printsupport.NewQPrinter(0)
	v.m_painter = gui.NewQPainter()

	v.ConnectStatusChanged(v.updateStatus)
	v.ConnectTimerEvent(v.timerEvent)
	v.ConnectPrintCurrentSlide(v.printCurrentSlide)
	v.ConnectGoToNextSlide(v.goToNextSlide)
}

func (v *SlideView) updateStatus(status quick.QQuickView__Status) {
	if v.m_slidesLeft > 0 {
		return
	}

	if status != quick.QQuickView__Ready {
		return
	}

	ri := v.RootObject()
	superClass := ri.MetaObject().SuperClass().ClassName()
	if !strings.HasPrefix(superClass, "Presentation") {
		fmt.Printf("Warning: Superclass appears to not be a Presentation: %s. ", superClass)
	} else {
		println("Found qml Presentation as rootObject")
	}

	ri.SetProperty("allowDelay", core.NewQVariant11(false)) //Disable partial reveals on slide pages
	slides := ri.Property("slides").ToList()
	v.m_slidesLeft = len(slides)

	println("SlideCount:", v.m_slidesLeft)
	rsize := v.m_printer.PageRect(1).Size()
	println("Printer's Page rect size (and suggested resolution of your presentation):", int(rsize.Width()), "x", int(rsize.Height()))

	v.m_printer.SetPageOrientation(gui.QPageLayout__Landscape)
	v.m_printer.SetFullPage(true)
	v.m_printer.SetOutputFileName("slides.pdf")
	v.m_painter.Begin(v.m_printer)

	// it would be better if we used the printer resolution here and forced
	// the presentation to be in the same resolution but when I try that,
	// the timer doesn't work properly for some reason?

	v.SetHeight(int(ri.Height()))
	v.SetWidth(int(ri.Width()))

	// Try uncommenting the below 4 lines and see what happens.
	//v.SetHeight(int(v.m_printer.PageRect(1).Height()))
	//v.SetWidth(int(v.m_printer.PageRect(1).Width()))
	//ri.SetHeight(float64(v.Height()))
	//ri.SetWidth(float64(v.Width()))

	v.m_tid = v.StartTimer(2000, 0)
}

func (v *SlideView) timerEvent(e *core.QTimerEvent) {
	v.printCurrentSlide()
	v.m_printedSlides++
	v.m_slidesLeft--

	if v.m_slidesLeft > 0 {
		v.m_printer.NewPage()
		v.goToNextSlide()
	} else {
		v.KillTimer(v.m_tid)
		v.m_painter.End()
		println("Printed to file: ", v.m_printer.OutputFileName())
		os.Exit(0)
	}
}

func (v *SlideView) printCurrentSlide() {
	pix := v.GrabWindow()
	println("Printing slide#", v.m_printedSlides+1, "Resolution:", pix.Size().Width(), "x", pix.Size().Height())

	pageRect := v.m_printer.PageRect(1)
	targetSize := pix.Size()

	targetSize.Scale2(pageRect.Size().ToSize(), core.Qt__KeepAspectRatio)
	v.m_painter.DrawImage(core.NewQRectF2(pageRect.TopLeft(), core.NewQSizeF2(targetSize)), pix, core.NewQRectF(), 0)
}

func (v *SlideView) goToNextSlide() {
	//workaround to invoke goToNextSlide in QML
	NewSlideViewFromPointer(v.RootObject().Pointer()).GoToNextSlide()
}
