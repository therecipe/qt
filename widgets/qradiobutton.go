package widgets

//#include "qradiobutton.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRadioButton struct {
	QAbstractButton
}

type QRadioButtonITF interface {
	QAbstractButtonITF
	QRadioButtonPTR() *QRadioButton
}

func PointerFromQRadioButton(ptr QRadioButtonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioButtonPTR().Pointer()
	}
	return nil
}

func QRadioButtonFromPointer(ptr unsafe.Pointer) *QRadioButton {
	var n = new(QRadioButton)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRadioButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRadioButton) QRadioButtonPTR() *QRadioButton {
	return ptr
}

func NewQRadioButton(parent QWidgetITF) *QRadioButton {
	return QRadioButtonFromPointer(unsafe.Pointer(C.QRadioButton_NewQRadioButton(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQRadioButton2(text string, parent QWidgetITF) *QRadioButton {
	return QRadioButtonFromPointer(unsafe.Pointer(C.QRadioButton_NewQRadioButton2(C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QRadioButton) DestroyQRadioButton() {
	if ptr.Pointer() != nil {
		C.QRadioButton_DestroyQRadioButton(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
