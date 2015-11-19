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

type QCommandLinkButton_ITF interface {
	QPushButton_ITF
	QCommandLinkButton_PTR() *QCommandLinkButton
}

func PointerFromQCommandLinkButton(ptr QCommandLinkButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLinkButton_PTR().Pointer()
	}
	return nil
}

func NewQCommandLinkButtonFromPointer(ptr unsafe.Pointer) *QCommandLinkButton {
	var n = new(QCommandLinkButton)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCommandLinkButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCommandLinkButton) QCommandLinkButton_PTR() *QCommandLinkButton {
	return ptr
}

func (ptr *QCommandLinkButton) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLinkButton_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLinkButton) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QCommandLinkButton_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func NewQCommandLinkButton(parent QWidget_ITF) *QCommandLinkButton {
	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton(PointerFromQWidget(parent)))
}

func NewQCommandLinkButton2(text string, parent QWidget_ITF) *QCommandLinkButton {
	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton2(C.CString(text), PointerFromQWidget(parent)))
}

func NewQCommandLinkButton3(text string, description string, parent QWidget_ITF) *QCommandLinkButton {
	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton3(C.CString(text), C.CString(description), PointerFromQWidget(parent)))
}

func (ptr *QCommandLinkButton) DestroyQCommandLinkButton() {
	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DestroyQCommandLinkButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
