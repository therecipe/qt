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

type QKeySequenceEditITF interface {
	QWidgetITF
	QKeySequenceEditPTR() *QKeySequenceEdit
}

func PointerFromQKeySequenceEdit(ptr QKeySequenceEditITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeySequenceEditPTR().Pointer()
	}
	return nil
}

func QKeySequenceEditFromPointer(ptr unsafe.Pointer) *QKeySequenceEdit {
	var n = new(QKeySequenceEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QKeySequenceEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QKeySequenceEdit) QKeySequenceEditPTR() *QKeySequenceEdit {
	return ptr
}

func (ptr *QKeySequenceEdit) SetKeySequence(keySequence gui.QKeySequenceITF) {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_SetKeySequence(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQKeySequence(keySequence)))
	}
}

func NewQKeySequenceEdit(parent QWidgetITF) *QKeySequenceEdit {
	return QKeySequenceEditFromPointer(unsafe.Pointer(C.QKeySequenceEdit_NewQKeySequenceEdit(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQKeySequenceEdit2(keySequence gui.QKeySequenceITF, parent QWidgetITF) *QKeySequenceEdit {
	return QKeySequenceEditFromPointer(unsafe.Pointer(C.QKeySequenceEdit_NewQKeySequenceEdit2(C.QtObjectPtr(gui.PointerFromQKeySequence(keySequence)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QKeySequenceEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QKeySequenceEdit) ConnectEditingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ConnectEditingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectEditingFinished() {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DisconnectEditingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQKeySequenceEditEditingFinished
func callbackQKeySequenceEditEditingFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QKeySequenceEdit) DestroyQKeySequenceEdit() {
	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DestroyQKeySequenceEdit(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
