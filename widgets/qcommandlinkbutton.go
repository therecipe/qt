package widgets

//#include "qcommandlinkbutton.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCommandLinkButton struct {
	QPushButton
}

type QCommandLinkButtonITF interface {
	QPushButtonITF
	QCommandLinkButtonPTR() *QCommandLinkButton
}

func PointerFromQCommandLinkButton(ptr QCommandLinkButtonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLinkButtonPTR().Pointer()
	}
	return nil
}

func QCommandLinkButtonFromPointer(ptr unsafe.Pointer) *QCommandLinkButton {
	var n = new(QCommandLinkButton)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCommandLinkButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCommandLinkButton) QCommandLinkButtonPTR() *QCommandLinkButton {
	return ptr
}

func (ptr *QCommandLinkButton) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLinkButton_Description(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCommandLinkButton) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QCommandLinkButton_SetDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(description))
	}
}

func NewQCommandLinkButton(parent QWidgetITF) *QCommandLinkButton {
	return QCommandLinkButtonFromPointer(unsafe.Pointer(C.QCommandLinkButton_NewQCommandLinkButton(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQCommandLinkButton2(text string, parent QWidgetITF) *QCommandLinkButton {
	return QCommandLinkButtonFromPointer(unsafe.Pointer(C.QCommandLinkButton_NewQCommandLinkButton2(C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQCommandLinkButton3(text string, description string, parent QWidgetITF) *QCommandLinkButton {
	return QCommandLinkButtonFromPointer(unsafe.Pointer(C.QCommandLinkButton_NewQCommandLinkButton3(C.CString(text), C.CString(description), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QCommandLinkButton) DestroyQCommandLinkButton() {
	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DestroyQCommandLinkButton(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
