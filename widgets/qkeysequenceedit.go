package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QKeySequenceEdit struct {
	QWidget
}

type QKeySequenceEdit_ITF interface {
	QWidget_ITF
	QKeySequenceEdit_PTR() *QKeySequenceEdit
}

func PointerFromQKeySequenceEdit(ptr QKeySequenceEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeySequenceEdit_PTR().Pointer()
	}
	return nil
}

func NewQKeySequenceEditFromPointer(ptr unsafe.Pointer) *QKeySequenceEdit {
	var n = new(QKeySequenceEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QKeySequenceEdit_") {
		n.SetObjectName("QKeySequenceEdit_" + qt.Identifier())
	}
	return n
}

func (ptr *QKeySequenceEdit) QKeySequenceEdit_PTR() *QKeySequenceEdit {
	return ptr
}

func (ptr *QKeySequenceEdit) SetKeySequence(keySequence gui.QKeySequence_ITF) {
	defer qt.Recovering("QKeySequenceEdit::setKeySequence")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_SetKeySequence(ptr.Pointer(), gui.PointerFromQKeySequence(keySequence))
	}
}

func NewQKeySequenceEdit(parent QWidget_ITF) *QKeySequenceEdit {
	defer qt.Recovering("QKeySequenceEdit::QKeySequenceEdit")

	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit(PointerFromQWidget(parent)))
}

func NewQKeySequenceEdit2(keySequence gui.QKeySequence_ITF, parent QWidget_ITF) *QKeySequenceEdit {
	defer qt.Recovering("QKeySequenceEdit::QKeySequenceEdit")

	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit2(gui.PointerFromQKeySequence(keySequence), PointerFromQWidget(parent)))
}

func (ptr *QKeySequenceEdit) Clear() {
	defer qt.Recovering("QKeySequenceEdit::clear")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QKeySequenceEdit) ConnectEditingFinished(f func()) {
	defer qt.Recovering("connect QKeySequenceEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectEditingFinished() {
	defer qt.Recovering("disconnect QKeySequenceEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQKeySequenceEditEditingFinished
func callbackQKeySequenceEditEditingFinished(ptrName *C.char) {
	defer qt.Recovering("callback QKeySequenceEdit::editingFinished")

	var signal = qt.GetSignal(C.GoString(ptrName), "editingFinished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QKeySequenceEdit) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQKeySequenceEditKeyPressEvent
func callbackQKeySequenceEditKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QKeySequenceEdit::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QKeySequenceEdit) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQKeySequenceEditKeyReleaseEvent
func callbackQKeySequenceEditKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QKeySequenceEdit::keyReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QKeySequenceEdit) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQKeySequenceEditTimerEvent
func callbackQKeySequenceEditTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QKeySequenceEdit::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QKeySequenceEdit) DestroyQKeySequenceEdit() {
	defer qt.Recovering("QKeySequenceEdit::~QKeySequenceEdit")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DestroyQKeySequenceEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
