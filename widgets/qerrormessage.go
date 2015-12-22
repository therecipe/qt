package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QErrorMessage struct {
	QDialog
}

type QErrorMessage_ITF interface {
	QDialog_ITF
	QErrorMessage_PTR() *QErrorMessage
}

func PointerFromQErrorMessage(ptr QErrorMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QErrorMessage_PTR().Pointer()
	}
	return nil
}

func NewQErrorMessageFromPointer(ptr unsafe.Pointer) *QErrorMessage {
	var n = new(QErrorMessage)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QErrorMessage_") {
		n.SetObjectName("QErrorMessage_" + qt.Identifier())
	}
	return n
}

func (ptr *QErrorMessage) QErrorMessage_PTR() *QErrorMessage {
	return ptr
}

func NewQErrorMessage(parent QWidget_ITF) *QErrorMessage {
	defer qt.Recovering("QErrorMessage::QErrorMessage")

	return NewQErrorMessageFromPointer(C.QErrorMessage_NewQErrorMessage(PointerFromQWidget(parent)))
}

func (ptr *QErrorMessage) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQErrorMessageChangeEvent
func callbackQErrorMessageChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectDone(f func(a int)) {
	defer qt.Recovering("connect QErrorMessage::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QErrorMessage) DisconnectDone() {
	defer qt.Recovering("disconnect QErrorMessage::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQErrorMessageDone
func callbackQErrorMessageDone(ptrName *C.char, a C.int) bool {
	defer qt.Recovering("callback QErrorMessage::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(a))
		return true
	}
	return false

}

func QErrorMessage_QtHandler() *QErrorMessage {
	defer qt.Recovering("QErrorMessage::qtHandler")

	return NewQErrorMessageFromPointer(C.QErrorMessage_QErrorMessage_QtHandler())
}

func (ptr *QErrorMessage) ShowMessage(message string) {
	defer qt.Recovering("QErrorMessage::showMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QErrorMessage) ShowMessage2(message string, ty string) {
	defer qt.Recovering("QErrorMessage::showMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage2(ptr.Pointer(), C.CString(message), C.CString(ty))
	}
}

func (ptr *QErrorMessage) DestroyQErrorMessage() {
	defer qt.Recovering("QErrorMessage::~QErrorMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DestroyQErrorMessage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
