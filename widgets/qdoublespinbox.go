package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QDoubleSpinBox_") {
		n.SetObjectName("QDoubleSpinBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QDoubleSpinBox) QDoubleSpinBox_PTR() *QDoubleSpinBox {
	return ptr
}

func (ptr *QDoubleSpinBox) CleanText() string {
	defer qt.Recovering("QDoubleSpinBox::cleanText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_CleanText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDoubleSpinBox) Decimals() int {
	defer qt.Recovering("QDoubleSpinBox::decimals")

	if ptr.Pointer() != nil {
		return int(C.QDoubleSpinBox_Decimals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDoubleSpinBox) Prefix() string {
	defer qt.Recovering("QDoubleSpinBox::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDoubleSpinBox) SetDecimals(prec int) {
	defer qt.Recovering("QDoubleSpinBox::setDecimals")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetDecimals(ptr.Pointer(), C.int(prec))
	}
}

func (ptr *QDoubleSpinBox) SetPrefix(prefix string) {
	defer qt.Recovering("QDoubleSpinBox::setPrefix")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QDoubleSpinBox) SetSuffix(suffix string) {
	defer qt.Recovering("QDoubleSpinBox::setSuffix")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QDoubleSpinBox) Suffix() string {
	defer qt.Recovering("QDoubleSpinBox::suffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Suffix(ptr.Pointer()))
	}
	return ""
}

func NewQDoubleSpinBox(parent QWidget_ITF) *QDoubleSpinBox {
	defer qt.Recovering("QDoubleSpinBox::QDoubleSpinBox")

	return NewQDoubleSpinBoxFromPointer(C.QDoubleSpinBox_NewQDoubleSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QDoubleSpinBox) ConnectValueChanged2(f func(text string)) {
	defer qt.Recovering("connect QDoubleSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ConnectValueChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged2", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectValueChanged2() {
	defer qt.Recovering("disconnect QDoubleSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DisconnectValueChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged2")
	}
}

//export callbackQDoubleSpinBoxValueChanged2
func callbackQDoubleSpinBoxValueChanged2(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QDoubleSpinBox::valueChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "valueChanged2")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QDoubleSpinBox) DestroyQDoubleSpinBox() {
	defer qt.Recovering("QDoubleSpinBox::~QDoubleSpinBox")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DestroyQDoubleSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
