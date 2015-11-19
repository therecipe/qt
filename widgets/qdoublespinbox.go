package widgets

//#include "qdoublespinbox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDoubleSpinBox struct {
	QAbstractSpinBox
}

type QDoubleSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QDoubleSpinBox_PTR() *QDoubleSpinBox
}

func PointerFromQDoubleSpinBox(ptr QDoubleSpinBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDoubleSpinBox_PTR().Pointer()
	}
	return nil
}

func NewQDoubleSpinBoxFromPointer(ptr unsafe.Pointer) *QDoubleSpinBox {
	var n = new(QDoubleSpinBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDoubleSpinBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDoubleSpinBox) QDoubleSpinBox_PTR() *QDoubleSpinBox {
	return ptr
}

func (ptr *QDoubleSpinBox) CleanText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_CleanText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDoubleSpinBox) Decimals() int {
	if ptr.Pointer() != nil {
		return int(C.QDoubleSpinBox_Decimals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDoubleSpinBox) Prefix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDoubleSpinBox) SetDecimals(prec int) {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetDecimals(ptr.Pointer(), C.int(prec))
	}
}

func (ptr *QDoubleSpinBox) SetPrefix(prefix string) {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QDoubleSpinBox) SetSuffix(suffix string) {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QDoubleSpinBox) Suffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Suffix(ptr.Pointer()))
	}
	return ""
}

func NewQDoubleSpinBox(parent QWidget_ITF) *QDoubleSpinBox {
	return NewQDoubleSpinBoxFromPointer(C.QDoubleSpinBox_NewQDoubleSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QDoubleSpinBox) DestroyQDoubleSpinBox() {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DestroyQDoubleSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
