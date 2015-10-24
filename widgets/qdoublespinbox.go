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

type QDoubleSpinBoxITF interface {
	QAbstractSpinBoxITF
	QDoubleSpinBoxPTR() *QDoubleSpinBox
}

func PointerFromQDoubleSpinBox(ptr QDoubleSpinBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDoubleSpinBoxPTR().Pointer()
	}
	return nil
}

func QDoubleSpinBoxFromPointer(ptr unsafe.Pointer) *QDoubleSpinBox {
	var n = new(QDoubleSpinBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDoubleSpinBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDoubleSpinBox) QDoubleSpinBoxPTR() *QDoubleSpinBox {
	return ptr
}

func (ptr *QDoubleSpinBox) CleanText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_CleanText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDoubleSpinBox) Decimals() int {
	if ptr.Pointer() != nil {
		return int(C.QDoubleSpinBox_Decimals(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDoubleSpinBox) Prefix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Prefix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDoubleSpinBox) SetDecimals(prec int) {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetDecimals(C.QtObjectPtr(ptr.Pointer()), C.int(prec))
	}
}

func (ptr *QDoubleSpinBox) SetPrefix(prefix string) {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetPrefix(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix))
	}
}

func (ptr *QDoubleSpinBox) SetSuffix(suffix string) {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetSuffix(C.QtObjectPtr(ptr.Pointer()), C.CString(suffix))
	}
}

func (ptr *QDoubleSpinBox) Suffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Suffix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQDoubleSpinBox(parent QWidgetITF) *QDoubleSpinBox {
	return QDoubleSpinBoxFromPointer(unsafe.Pointer(C.QDoubleSpinBox_NewQDoubleSpinBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QDoubleSpinBox) DestroyQDoubleSpinBox() {
	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DestroyQDoubleSpinBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
