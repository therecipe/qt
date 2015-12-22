package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QRubberBand struct {
	QWidget
}

type QRubberBand_ITF interface {
	QWidget_ITF
	QRubberBand_PTR() *QRubberBand
}

func PointerFromQRubberBand(ptr QRubberBand_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRubberBand_PTR().Pointer()
	}
	return nil
}

func NewQRubberBandFromPointer(ptr unsafe.Pointer) *QRubberBand {
	var n = new(QRubberBand)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRubberBand_") {
		n.SetObjectName("QRubberBand_" + qt.Identifier())
	}
	return n
}

func (ptr *QRubberBand) QRubberBand_PTR() *QRubberBand {
	return ptr
}

//QRubberBand::Shape
type QRubberBand__Shape int64

const (
	QRubberBand__Line      = QRubberBand__Shape(0)
	QRubberBand__Rectangle = QRubberBand__Shape(1)
)

func (ptr *QRubberBand) SetGeometry(rect core.QRect_ITF) {
	defer qt.Recovering("QRubberBand::setGeometry")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func NewQRubberBand(s QRubberBand__Shape, p QWidget_ITF) *QRubberBand {
	defer qt.Recovering("QRubberBand::QRubberBand")

	return NewQRubberBandFromPointer(C.QRubberBand_NewQRubberBand(C.int(s), PointerFromQWidget(p)))
}

func (ptr *QRubberBand) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QRubberBand::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQRubberBandChangeEvent
func callbackQRubberBandChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRubberBand) Move2(p core.QPoint_ITF) {
	defer qt.Recovering("QRubberBand::move")

	if ptr.Pointer() != nil {
		C.QRubberBand_Move2(ptr.Pointer(), core.PointerFromQPoint(p))
	}
}

func (ptr *QRubberBand) Move(x int, y int) {
	defer qt.Recovering("QRubberBand::move")

	if ptr.Pointer() != nil {
		C.QRubberBand_Move(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QRubberBand) ConnectMoveEvent(f func(v *gui.QMoveEvent)) {
	defer qt.Recovering("connect QRubberBand::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QRubberBand::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQRubberBandMoveEvent
func callbackQRubberBandMoveEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QRubberBand::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QRubberBand::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQRubberBandPaintEvent
func callbackQRubberBandPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRubberBand) Resize2(size core.QSize_ITF) {
	defer qt.Recovering("QRubberBand::resize")

	if ptr.Pointer() != nil {
		C.QRubberBand_Resize2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QRubberBand) Resize(width int, height int) {
	defer qt.Recovering("QRubberBand::resize")

	if ptr.Pointer() != nil {
		C.QRubberBand_Resize(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQRubberBandResizeEvent
func callbackQRubberBandResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRubberBand) SetGeometry2(x int, y int, width int, height int) {
	defer qt.Recovering("QRubberBand::setGeometry")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) Shape() QRubberBand__Shape {
	defer qt.Recovering("QRubberBand::shape")

	if ptr.Pointer() != nil {
		return QRubberBand__Shape(C.QRubberBand_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRubberBand) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QRubberBand::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QRubberBand::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQRubberBandShowEvent
func callbackQRubberBandShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRubberBand) DestroyQRubberBand() {
	defer qt.Recovering("QRubberBand::~QRubberBand")

	if ptr.Pointer() != nil {
		C.QRubberBand_DestroyQRubberBand(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
