package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextFrame struct {
	QTextObject
}

type QTextFrame_ITF interface {
	QTextObject_ITF
	QTextFrame_PTR() *QTextFrame
}

func PointerFromQTextFrame(ptr QTextFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrame_PTR().Pointer()
	}
	return nil
}

func NewQTextFrameFromPointer(ptr unsafe.Pointer) *QTextFrame {
	var n = new(QTextFrame)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTextFrame_") {
		n.SetObjectName("QTextFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextFrame) QTextFrame_PTR() *QTextFrame {
	return ptr
}

func NewQTextFrame(document QTextDocument_ITF) *QTextFrame {
	defer qt.Recovering("QTextFrame::QTextFrame")

	return NewQTextFrameFromPointer(C.QTextFrame_NewQTextFrame(PointerFromQTextDocument(document)))
}

func (ptr *QTextFrame) FirstPosition() int {
	defer qt.Recovering("QTextFrame::firstPosition")

	if ptr.Pointer() != nil {
		return int(C.QTextFrame_FirstPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrame) LastPosition() int {
	defer qt.Recovering("QTextFrame::lastPosition")

	if ptr.Pointer() != nil {
		return int(C.QTextFrame_LastPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrame) ParentFrame() *QTextFrame {
	defer qt.Recovering("QTextFrame::parentFrame")

	if ptr.Pointer() != nil {
		return NewQTextFrameFromPointer(C.QTextFrame_ParentFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextFrame) SetFrameFormat(format QTextFrameFormat_ITF) {
	defer qt.Recovering("QTextFrame::setFrameFormat")

	if ptr.Pointer() != nil {
		C.QTextFrame_SetFrameFormat(ptr.Pointer(), PointerFromQTextFrameFormat(format))
	}
}

func (ptr *QTextFrame) DestroyQTextFrame() {
	defer qt.Recovering("QTextFrame::~QTextFrame")

	if ptr.Pointer() != nil {
		C.QTextFrame_DestroyQTextFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextFrame) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTextFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTextFrame) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTextFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTextFrameTimerEvent
func callbackQTextFrameTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextFrame::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextFrame) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTextFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTextFrame) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTextFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTextFrameChildEvent
func callbackQTextFrameChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextFrame::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextFrame) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTextFrame) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTextFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTextFrameCustomEvent
func callbackQTextFrameCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextFrame::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
