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

type QSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QSpinBox_PTR() *QSpinBox
}

func PointerFromQSpinBox(ptr QSpinBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpinBox_PTR().Pointer()
	}
	return nil
}

func NewQSpinBoxFromPointer(ptr unsafe.Pointer) *QSpinBox {
	var n = new(QSpinBox)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSpinBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSpinBox) QSpinBox_PTR() *QSpinBox {
	return ptr
}

func (ptr *QSpinBox) CleanText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_CleanText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) DisplayIntegerBase() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_DisplayIntegerBase(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Maximum() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Minimum() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Prefix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) SetDisplayIntegerBase(base int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetDisplayIntegerBase(ptr.Pointer(), C.int(base))
	}
}

func (ptr *QSpinBox) SetMaximum(max int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetMaximum(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QSpinBox) SetMinimum(min int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetMinimum(ptr.Pointer(), C.int(min))
	}
}

func (ptr *QSpinBox) SetPrefix(prefix string) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QSpinBox) SetSingleStep(val int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetSingleStep(ptr.Pointer(), C.int(val))
	}
}

func (ptr *QSpinBox) SetSuffix(suffix string) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QSpinBox) SetValue(val int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetValue(ptr.Pointer(), C.int(val))
	}
}

func (ptr *QSpinBox) SingleStep() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_SingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Suffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Suffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Value(ptr.Pointer()))
	}
	return 0
}

func NewQSpinBox(parent QWidget_ITF) *QSpinBox {
	return NewQSpinBoxFromPointer(C.QSpinBox_NewQSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QSpinBox) SetRange(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QSpinBox_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QSpinBox) ConnectValueChanged(f func(i int)) {
	if ptr.Pointer() != nil {
		C.QSpinBox_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QSpinBox) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QSpinBox_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQSpinBoxValueChanged
func callbackQSpinBoxValueChanged(ptrName *C.char, i C.int) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(i))
}

func (ptr *QSpinBox) DestroyQSpinBox() {
	if ptr.Pointer() != nil {
		C.QSpinBox_DestroyQSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
