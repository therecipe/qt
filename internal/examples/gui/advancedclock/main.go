//author: https://github.com/omac777

package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type AnalogClockQwdgt struct {
	widgets.QWidget

	_ func() `constructor:"init"`

	_ func() `slot:"renderLater"`
	_ func() `slot:"renderNow"`

	render func(painter *gui.QPainter)

	contextMenu    *widgets.QMenu
	m_dragPosition *core.QPoint
	m_timerId      int
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	var window = NewAnalogClockQwdgt(nil, core.Qt__FramelessWindowHint|core.Qt__WindowSystemMenuHint)
	window.Show()
	widgets.QApplication_Exec()
}

func (w *AnalogClockQwdgt) init() {
	w.Resize2(200, 200)
	w.m_timerId = w.StartTimer(1000, 0)
	w.ConnectTimerEvent(w.timerEvent)

	w.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	w.ConnectCustomContextMenuRequested(w.customContextMenuRequested)
	w.SetToolTip("Drag the clock with the left mouse button.\nUse the right mouse button to open a context menu.")

	w.ConnectRenderLater(w.renderLater)
	w.ConnectPaintEvent(w.renderNow)
	w.ConnectResizeEvent(w.resizeEvent)

	w.ConnectMouseMoveEvent(w.QWdgtMouseMoveEvent)
	w.ConnectMousePressEvent(w.QWdgtMousePressEvent)
	w.ConnectMouseReleaseEvent(w.QWdgtMouseReleaseEvent)

	w.render = w.renderFunc
}

func (w *AnalogClockQwdgt) timerEvent(event *core.QTimerEvent) {
	fmt.Printf("timerEvent()...\n")
	if event.TimerId() == w.m_timerId {
		fmt.Printf("timerEvent()...event.TimerId() == w.m_timerId\n")
		w.renderLater()
	}
}

func (w *AnalogClockQwdgt) customContextMenuRequested(p *core.QPoint) {
	fmt.Printf("customContextMenuRequested()...\n")
	if w.contextMenu == nil {
		fmt.Printf("customContextMenuRequested()...w.contextMenu == nil\n")
		w.contextMenu = widgets.NewQMenu(w)
		menuActionExit := w.contextMenu.AddAction("E&xit")
		menuActionExit.SetShortcut(gui.NewQKeySequence2("Ctrl+X", gui.QKeySequence__NativeText))
		menuActionExit.ConnectTriggered(func(checked bool) { w.Close() })
	}

	w.contextMenu.Exec2(w.MapToGlobal(p), nil)
}

func (w *AnalogClockQwdgt) renderFunc(p *gui.QPainter) {
	fmt.Printf("renderFunc()...\n")
	var hourHand = gui.NewQPolygon3([]*core.QPoint{
		core.NewQPoint2(7, 8),
		core.NewQPoint2(-7, 8),
		core.NewQPoint2(0, -40),
	})

	var minuteHand = gui.NewQPolygon4([]*core.QPoint{
		core.NewQPoint2(7, 8),
		core.NewQPoint2(-7, 8),
		core.NewQPoint2(0, -70),
	})

	var hourColor = gui.NewQColor3(127, 0, 127, 255)   //opaque 255  ...transparent 0
	var minuteColor = gui.NewQColor3(0, 127, 127, 191) //a bit of transparency
	var arcColor = gui.NewQColor3(255, 0, 0, 191)      //transparent red

	p.SetRenderHint(gui.QPainter__Antialiasing, true)
	p.SetPen3(core.Qt__NoPen)
	p.SetBrush(gui.NewQBrush3(hourColor, 1))

	firstTranslateX := float64(w.Width()) / 2
	firstTranslateY := float64(w.Height()) / 2
	fmt.Printf("firstTranslateX:<<%v>>\n", firstTranslateX)
	fmt.Printf("firstTranslateY:<<%v>>\n", firstTranslateY)

	p.Translate3(firstTranslateX, firstTranslateY)

	nSide := MyQMin(w.Width(), w.Height())
	fmt.Printf("nSide:<<%v>>\n", nSide)

	fFirstScaleX := float64(nSide) / 200.0
	fFirstScaleY := float64(nSide) / 200.0
	fmt.Printf("fFirstScaleX:<<%v>>\n", fFirstScaleX)
	fmt.Printf("fFirstScaleY:<<%v>>\n", fFirstScaleY)

	p.Scale(fFirstScaleX, fFirstScaleY)

	var tTimeNow = core.QTime_CurrentTime()

	p.Save() //push 1
	myHourNow := tTimeNow.Hour()
	myMinuteNow := tTimeNow.Minute()
	fmt.Printf("myHourNow:<<%v>>\n", myHourNow)
	fmt.Printf("myMinuteNow:<<%v>>\n", myMinuteNow)
	myFirstRotate := float64(30.0 * (myHourNow + myMinuteNow/60.0))
	fmt.Printf("myFirstRotate for the hour hand:<<%v>>\n", myFirstRotate)
	p.Rotate(myFirstRotate)
	p.DrawConvexPolygon4(hourHand)
	p.Restore() //pop 1

	p.Save() //push 1
	p.SetPen2(hourColor)
	fmt.Printf("setting hour ticks...\n")
	for i := 0; i < 12; i++ {
		p.DrawLine3(88, 0, 96, 0)
		mySecondRotate := 30.0
		fmt.Printf("mySecondRotate for hour ticks:<<%v>>\n", mySecondRotate)
		p.Rotate(mySecondRotate)
	}

	p.SetPen3(core.Qt__NoPen)
	p.SetBrush(gui.NewQBrush3(minuteColor, 1))
	p.Restore() //pop 1

	p.Save() //push 1
	myThirdRotate := float64(6.0 * (tTimeNow.Minute() + tTimeNow.Second()/60.0))
	fmt.Printf("myThirdRotate for the minute hand:<<%v>>\n", myThirdRotate)
	p.Rotate(myThirdRotate)
	p.DrawConvexPolygon4(minuteHand)
	p.Restore() //pop 1

	p.Save() //push 1
	p.SetPen2(minuteColor)
	fmt.Printf("setting minute ticks...\n")
	for j := 0; j < 60; j++ {
		if (j % 5) != 0 {
			p.DrawLine3(92, 0, 96, 0)
		}

		myFourthRotate := 6.0
		fmt.Printf("myFourthRotate for the minute ticks:<<%v>>\n", myFourthRotate)
		p.Rotate(myFourthRotate)

	}
	p.Restore() //pop 1

	p.SetPen2(arcColor)
	p.SetBrush(gui.NewQBrush3(arcColor, 1))

	//x,y is the centerpoint
	//r is the radius
	//rect is the bounding rectangle
	//myRegionXPos, myRegionYPos, w.Width(), w.Height()
	fmt.Printf("w.Width():<<%v>>\n", w.Width())
	fmt.Printf("w.Height():<<%v>>\n", w.Height())
	//nSide := MyQMin(w.Width(), w.Height())
	myRegionXPos := float64(w.Width())/2.0 - float64(nSide)/2.0
	myRegionYPos := float64(w.Height())/2.0 - float64(nSide)/2.0
	theRadius := float64(w.Width()) / 2.0
	//theCenterX := float64(w.Width())/2.0
	//theCenterY := float64(w.Height())/2.0
	xPosForArc := myRegionXPos - theRadius
	yPosForArc := myRegionYPos - theRadius
	rectBoundingForArc := core.NewQRectF4(xPosForArc, yPosForArc, 2*theRadius, 2*theRadius)
	// QPainterPath piePath;
	// NewQPainterPath() *QPainterPath
	myPiePainterPath := gui.NewQPainterPath()
	//We are already translated with p so we don't need to move/translate any further.
	//myPiePainterPath.MoveTo2(firstTranslateX, firstTranslateY)
	//myPiePainterPath.MoveTo2(theCenterX, theCenterY) //should move to the center
	//p.Translate3(0.0, 0.0)

	// piePath.arcTo(20.0, 30.0, 60.0, 40.0, 60.0, 240.0);
	// (ptr *QPainterPath) ArcTo(qrectF, startAngle float64, sweepLengthAngle float64)
	myPiePainterPath.ArcTo(rectBoundingForArc.QRectF_PTR(), 40.0, 90.0)
	myPiePainterPath.CloseSubpath()
	p.DrawPath(myPiePainterPath)
}

func (w *AnalogClockQwdgt) QWdgtMouseMoveEvent(ev *gui.QMouseEvent) {
	fmt.Printf("QWdgtMouseMoveEvent()...\n")
	nMouseState := ev.MouseState()
	bLeftButtonPressed := bool(nMouseState == core.Qt__LeftButton)
	if bLeftButtonPressed {
		qptMoveHere := Point1MinusPoint2(ev.GlobalPos(), w.m_dragPosition)
		//w.SetPosition(qptMoveHere)
		w.Move(qptMoveHere)
	}
}

func (w *AnalogClockQwdgt) QWdgtMousePressEvent(ev *gui.QMouseEvent) {
	fmt.Printf("QWdgtMousePressEvent()...\n")
	nMouseState := ev.MouseState()
	bLeftButtonPressed := bool(nMouseState == core.Qt__LeftButton)
	if bLeftButtonPressed {
		qptGlobalPos := ev.GlobalPos()
		qptTopLeft := w.FrameGeometry().TopLeft()
		w.m_dragPosition = Point1MinusPoint2(qptGlobalPos, qptTopLeft)
	}
}

func (w *AnalogClockQwdgt) QWdgtMouseReleaseEvent(ev *gui.QMouseEvent) {
	fmt.Printf("QWdgtMouseReleaseEvent()...\n")
}

func Point1MinusPoint2(pt1 *core.QPoint, pt2 *core.QPoint) *core.QPoint {
	return core.NewQPoint2(pt1.X()-pt2.X(), pt1.Y()-pt2.Y())
}

func (w *AnalogClockQwdgt) renderNow(ev *gui.QPaintEvent) {
	fmt.Printf("renderNow()...\n")

	nSide := MyQMin(w.Width(), w.Height())
	myRegionXPos := w.Width()/2 - nSide/2
	myRegionYPos := w.Height()/2 - nSide/2
	var rect = core.NewQRect4(myRegionXPos, myRegionYPos, w.Width(), w.Height())
	qr3 := gui.NewQRegion3(rect, gui.QRegion__Ellipse) //set the ellipse within the qregion

	w.BackingStore().BeginPaint(qr3)
	var device = w.BackingStore().PaintDevice()
	var painter = gui.NewQPainter2(device)
	defer painter.DestroyQPainter()
	w.render(painter)
	w.BackingStore().EndPaint()
	w.BackingStore().Flush(gui.NewQRegion3(rect, gui.QRegion__Ellipse), nil, core.NewQPoint())
}

func (w *AnalogClockQwdgt) resizeEvent(ev *gui.QResizeEvent) {
	fmt.Printf("resizeEvent()...\n")
	nSide := MyQMin(w.Width(), w.Height())
	myRegionXPos := w.Width()/2 - nSide/2
	myRegionYPos := w.Height()/2 - nSide/2
	var rect = core.NewQRect4(myRegionXPos, myRegionYPos, w.Width(), w.Height())
	qr3 := gui.NewQRegion3(rect, gui.QRegion__Ellipse) //set the ellipse within the qregion
	w.SetMask2(qr3)

	w.BackingStore().Resize(ev.Size())

	w.RenderLater()
}

func (w *AnalogClockQwdgt) renderLater() {
	fmt.Printf("renderLater()...\n")
	w.Update()
}

func MyQMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
