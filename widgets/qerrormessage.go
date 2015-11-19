package widgets

//#include "qerrormessage.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QErrorMessage_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QErrorMessage) QErrorMessage_PTR() *QErrorMessage {
	return ptr
}

func NewQErrorMessage(parent QWidget_ITF) *QErrorMessage {
	return NewQErrorMessageFromPointer(C.QErrorMessage_NewQErrorMessage(PointerFromQWidget(parent)))
}

func QErrorMessage_QtHandler() *QErrorMessage {
	return NewQErrorMessageFromPointer(C.QErrorMessage_QErrorMessage_QtHandler())
}

func (ptr *QErrorMessage) ShowMessage(message string) {
	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QErrorMessage) ShowMessage2(message string, ty string) {
	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage2(ptr.Pointer(), C.CString(message), C.CString(ty))
	}
}

func (ptr *QErrorMessage) DestroyQErrorMessage() {
	if ptr.Pointer() != nil {
		C.QErrorMessage_DestroyQErrorMessage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
