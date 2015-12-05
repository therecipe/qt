package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QRadioButton struct {
	QAbstractButton
}

type QRadioButton_ITF interface {
	QAbstractButton_ITF
	QRadioButton_PTR() *QRadioButton
}

func PointerFromQRadioButton(ptr QRadioButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioButton_PTR().Pointer()
	}
	return nil
}

func NewQRadioButtonFromPointer(ptr unsafe.Pointer) *QRadioButton {
	var n = new(QRadioButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRadioButton_") {
		n.SetObjectName("QRadioButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRadioButton) QRadioButton_PTR() *QRadioButton {
	return ptr
}

func NewQRadioButton(parent QWidget_ITF) *QRadioButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioButton::QRadioButton")
		}
	}()

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton(PointerFromQWidget(parent)))
}

func NewQRadioButton2(text string, parent QWidget_ITF) *QRadioButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioButton::QRadioButton")
		}
	}()

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QRadioButton) DestroyQRadioButton() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioButton::~QRadioButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioButton_DestroyQRadioButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
