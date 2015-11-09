package widgets

//#include "qkeysequenceedit.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QKeySequenceEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QKeySequenceEdit) QKeySequenceEdit_PTR() *QKeySequenceEdit {
	return ptr
}

func (ptr *QKeySequenceEdit) SetKeySequence(keySequence gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_SetKeySequence(ptr.Pointer(), gui.PointerFromQKeySequence(keySequence))
	}
}

func NewQKeySequenceEdit(parent QWidget_ITF) *QKeySequenceEdit {
	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit(PointerFromQWidget(parent)))
}

func NewQKeySequenceEdit2(keySequence gui.QKeySequence_ITF, parent QWidget_ITF) *QKeySequenceEdit {
	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit2(gui.PointerFromQKeySequence(keySequence), PointerFromQWidget(parent)))
}

func (ptr *QKeySequenceEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QKeySequenceEdit) ConnectEditingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectEditingFinished() {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQKeySequenceEditEditingFinished
func callbackQKeySequenceEditEditingFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QKeySequenceEdit) DestroyQKeySequenceEdit() {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DestroyQKeySequenceEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
