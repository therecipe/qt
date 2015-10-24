package widgets

//#include "qspinbox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSpinBox struct {
	QAbstractSpinBox
}

type QSpinBoxITF interface {
	QAbstractSpinBoxITF
	QSpinBoxPTR() *QSpinBox
}

func PointerFromQSpinBox(ptr QSpinBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpinBoxPTR().Pointer()
	}
	return nil
}

func QSpinBoxFromPointer(ptr unsafe.Pointer) *QSpinBox {
	var n = new(QSpinBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSpinBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSpinBox) QSpinBoxPTR() *QSpinBox {
	return ptr
}

func (ptr *QSpinBox) CleanText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_CleanText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSpinBox) DisplayIntegerBase() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_DisplayIntegerBase(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSpinBox) Maximum() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Maximum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSpinBox) Minimum() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Minimum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSpinBox) Prefix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Prefix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSpinBox) SetDisplayIntegerBase(base int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetDisplayIntegerBase(C.QtObjectPtr(ptr.Pointer()), C.int(base))
	}
}

func (ptr *QSpinBox) SetMaximum(max int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetMaximum(C.QtObjectPtr(ptr.Pointer()), C.int(max))
	}
}

func (ptr *QSpinBox) SetMinimum(min int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetMinimum(C.QtObjectPtr(ptr.Pointer()), C.int(min))
	}
}

func (ptr *QSpinBox) SetPrefix(prefix string) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetPrefix(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix))
	}
}

func (ptr *QSpinBox) SetSingleStep(val int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetSingleStep(C.QtObjectPtr(ptr.Pointer()), C.int(val))
	}
}

func (ptr *QSpinBox) SetSuffix(suffix string) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetSuffix(C.QtObjectPtr(ptr.Pointer()), C.CString(suffix))
	}
}

func (ptr *QSpinBox) SetValue(val int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetValue(C.QtObjectPtr(ptr.Pointer()), C.int(val))
	}
}

func (ptr *QSpinBox) SingleStep() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_SingleStep(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSpinBox) Suffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Suffix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSpinBox) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQSpinBox(parent QWidgetITF) *QSpinBox {
	return QSpinBoxFromPointer(unsafe.Pointer(C.QSpinBox_NewQSpinBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QSpinBox) SetRange(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetRange(C.QtObjectPtr(ptr.Pointer()), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QSpinBox) ConnectValueChanged(f func(i int)) {
	if ptr.Pointer() != nil {
		C.QSpinBox_ConnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QSpinBox) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QSpinBox_DisconnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQSpinBoxValueChanged
func callbackQSpinBoxValueChanged(ptrName *C.char, i C.int) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(i))
}

func (ptr *QSpinBox) DestroyQSpinBox() {
	if ptr.Pointer() != nil {
		C.QSpinBox_DestroyQSpinBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
