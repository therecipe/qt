package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFocusFrame struct {
	QWidget
}

type QFocusFrame_ITF interface {
	QWidget_ITF
	QFocusFrame_PTR() *QFocusFrame
}

func PointerFromQFocusFrame(ptr QFocusFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFocusFrame_PTR().Pointer()
	}
	return nil
}

func NewQFocusFrameFromPointer(ptr unsafe.Pointer) *QFocusFrame {
	var n = new(QFocusFrame)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFocusFrame_") {
		n.SetObjectName("QFocusFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QFocusFrame) QFocusFrame_PTR() *QFocusFrame {
	return ptr
}

func NewQFocusFrame(parent QWidget_ITF) *QFocusFrame {
	defer qt.Recovering("QFocusFrame::QFocusFrame")

	return NewQFocusFrameFromPointer(C.QFocusFrame_NewQFocusFrame(PointerFromQWidget(parent)))
}

func (ptr *QFocusFrame) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFocusFramePaintEvent
func callbackQFocusFramePaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QFocusFrame) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QFocusFrame::setWidget")

	if ptr.Pointer() != nil {
		C.QFocusFrame_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QFocusFrame) Widget() *QWidget {
	defer qt.Recovering("QFocusFrame::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QFocusFrame_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFocusFrame) DestroyQFocusFrame() {
	defer qt.Recovering("QFocusFrame::~QFocusFrame")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DestroyQFocusFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
