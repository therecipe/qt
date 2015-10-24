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

type QErrorMessageITF interface {
	QDialogITF
	QErrorMessagePTR() *QErrorMessage
}

func PointerFromQErrorMessage(ptr QErrorMessageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QErrorMessagePTR().Pointer()
	}
	return nil
}

func QErrorMessageFromPointer(ptr unsafe.Pointer) *QErrorMessage {
	var n = new(QErrorMessage)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QErrorMessage_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QErrorMessage) QErrorMessagePTR() *QErrorMessage {
	return ptr
}

func NewQErrorMessage(parent QWidgetITF) *QErrorMessage {
	return QErrorMessageFromPointer(unsafe.Pointer(C.QErrorMessage_NewQErrorMessage(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func QErrorMessage_QtHandler() *QErrorMessage {
	return QErrorMessageFromPointer(unsafe.Pointer(C.QErrorMessage_QErrorMessage_QtHandler()))
}

func (ptr *QErrorMessage) ShowMessage(message string) {
	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage(C.QtObjectPtr(ptr.Pointer()), C.CString(message))
	}
}

func (ptr *QErrorMessage) ShowMessage2(message string, ty string) {
	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage2(C.QtObjectPtr(ptr.Pointer()), C.CString(message), C.CString(ty))
	}
}

func (ptr *QErrorMessage) DestroyQErrorMessage() {
	if ptr.Pointer() != nil {
		C.QErrorMessage_DestroyQErrorMessage(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
