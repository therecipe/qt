package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QSpinBox_") {
		n.SetObjectName("QSpinBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QSpinBox) QSpinBox_PTR() *QSpinBox {
	return ptr
}

func (ptr *QSpinBox) CleanText() string {
	defer qt.Recovering("QSpinBox::cleanText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_CleanText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) DisplayIntegerBase() int {
	defer qt.Recovering("QSpinBox::displayIntegerBase")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_DisplayIntegerBase(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Maximum() int {
	defer qt.Recovering("QSpinBox::maximum")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Minimum() int {
	defer qt.Recovering("QSpinBox::minimum")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Prefix() string {
	defer qt.Recovering("QSpinBox::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) SetDisplayIntegerBase(base int) {
	defer qt.Recovering("QSpinBox::setDisplayIntegerBase")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetDisplayIntegerBase(ptr.Pointer(), C.int(base))
	}
}

func (ptr *QSpinBox) SetMaximum(max int) {
	defer qt.Recovering("QSpinBox::setMaximum")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetMaximum(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QSpinBox) SetMinimum(min int) {
	defer qt.Recovering("QSpinBox::setMinimum")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetMinimum(ptr.Pointer(), C.int(min))
	}
}

func (ptr *QSpinBox) SetPrefix(prefix string) {
	defer qt.Recovering("QSpinBox::setPrefix")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QSpinBox) SetSingleStep(val int) {
	defer qt.Recovering("QSpinBox::setSingleStep")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetSingleStep(ptr.Pointer(), C.int(val))
	}
}

func (ptr *QSpinBox) SetSuffix(suffix string) {
	defer qt.Recovering("QSpinBox::setSuffix")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QSpinBox) SetValue(val int) {
	defer qt.Recovering("QSpinBox::setValue")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetValue(ptr.Pointer(), C.int(val))
	}
}

func (ptr *QSpinBox) SingleStep() int {
	defer qt.Recovering("QSpinBox::singleStep")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_SingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Suffix() string {
	defer qt.Recovering("QSpinBox::suffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Suffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) Value() int {
	defer qt.Recovering("QSpinBox::value")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Value(ptr.Pointer()))
	}
	return 0
}

func NewQSpinBox(parent QWidget_ITF) *QSpinBox {
	defer qt.Recovering("QSpinBox::QSpinBox")

	return NewQSpinBoxFromPointer(C.QSpinBox_NewQSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QSpinBox) SetRange(minimum int, maximum int) {
	defer qt.Recovering("QSpinBox::setRange")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QSpinBox) ConnectValueChanged(f func(i int)) {
	defer qt.Recovering("connect QSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QSpinBox_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QSpinBox) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QSpinBox_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQSpinBoxValueChanged
func callbackQSpinBoxValueChanged(ptrName *C.char, i C.int) {
	defer qt.Recovering("callback QSpinBox::valueChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "valueChanged")
	if signal != nil {
		signal.(func(int))(int(i))
	}

}

func (ptr *QSpinBox) DestroyQSpinBox() {
	defer qt.Recovering("QSpinBox::~QSpinBox")

	if ptr.Pointer() != nil {
		C.QSpinBox_DestroyQSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
