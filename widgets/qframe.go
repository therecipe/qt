package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFrame struct {
	QWidget
}

type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func PointerFromQFrame(ptr QFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func NewQFrameFromPointer(ptr unsafe.Pointer) *QFrame {
	var n = new(QFrame)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFrame_") {
		n.SetObjectName("QFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QFrame) QFrame_PTR() *QFrame {
	return ptr
}

//QFrame::Shadow
type QFrame__Shadow int64

const (
	QFrame__Plain  = QFrame__Shadow(0x0010)
	QFrame__Raised = QFrame__Shadow(0x0020)
	QFrame__Sunken = QFrame__Shadow(0x0030)
)

//QFrame::Shape
type QFrame__Shape int64

const (
	QFrame__NoFrame     = QFrame__Shape(0)
	QFrame__Box         = QFrame__Shape(0x0001)
	QFrame__Panel       = QFrame__Shape(0x0002)
	QFrame__WinPanel    = QFrame__Shape(0x0003)
	QFrame__HLine       = QFrame__Shape(0x0004)
	QFrame__VLine       = QFrame__Shape(0x0005)
	QFrame__StyledPanel = QFrame__Shape(0x0006)
)

//QFrame::StyleMask
type QFrame__StyleMask int64

var (
	QFrame__Shadow_Mask = QFrame__StyleMask(0x00f0)
	QFrame__Shape_Mask  = QFrame__StyleMask(0x000f)
)

func (ptr *QFrame) FrameRect() *core.QRect {
	defer qt.Recovering("QFrame::frameRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QFrame_FrameRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFrame) FrameShadow() QFrame__Shadow {
	defer qt.Recovering("QFrame::frameShadow")

	if ptr.Pointer() != nil {
		return QFrame__Shadow(C.QFrame_FrameShadow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameShape() QFrame__Shape {
	defer qt.Recovering("QFrame::frameShape")

	if ptr.Pointer() != nil {
		return QFrame__Shape(C.QFrame_FrameShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameWidth() int {
	defer qt.Recovering("QFrame::frameWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) LineWidth() int {
	defer qt.Recovering("QFrame::lineWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) MidLineWidth() int {
	defer qt.Recovering("QFrame::midLineWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_MidLineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) SetFrameRect(v core.QRect_ITF) {
	defer qt.Recovering("QFrame::setFrameRect")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameRect(ptr.Pointer(), core.PointerFromQRect(v))
	}
}

func (ptr *QFrame) SetFrameShadow(v QFrame__Shadow) {
	defer qt.Recovering("QFrame::setFrameShadow")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShadow(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetFrameShape(v QFrame__Shape) {
	defer qt.Recovering("QFrame::setFrameShape")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShape(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetLineWidth(v int) {
	defer qt.Recovering("QFrame::setLineWidth")

	if ptr.Pointer() != nil {
		C.QFrame_SetLineWidth(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetMidLineWidth(v int) {
	defer qt.Recovering("QFrame::setMidLineWidth")

	if ptr.Pointer() != nil {
		C.QFrame_SetMidLineWidth(ptr.Pointer(), C.int(v))
	}
}

func NewQFrame(parent QWidget_ITF, f core.Qt__WindowType) *QFrame {
	defer qt.Recovering("QFrame::QFrame")

	return NewQFrameFromPointer(C.QFrame_NewQFrame(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QFrame) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFrame) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFrameChangeEvent
func callbackQFrameChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QFrame) FrameStyle() int {
	defer qt.Recovering("QFrame::frameStyle")

	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFrame) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFramePaintEvent
func callbackQFramePaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QFrame) SetFrameStyle(style int) {
	defer qt.Recovering("QFrame::setFrameStyle")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QFrame) SizeHint() *core.QSize {
	defer qt.Recovering("QFrame::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QFrame_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFrame) DestroyQFrame() {
	defer qt.Recovering("QFrame::~QFrame")

	if ptr.Pointer() != nil {
		C.QFrame_DestroyQFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
