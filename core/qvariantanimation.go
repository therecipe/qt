package core

//#include "qvariantanimation.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QVariantAnimation struct {
	QAbstractAnimation
}

type QVariantAnimation_ITF interface {
	QAbstractAnimation_ITF
	QVariantAnimation_PTR() *QVariantAnimation
}

func PointerFromQVariantAnimation(ptr QVariantAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVariantAnimation_PTR().Pointer()
	}
	return nil
}

func NewQVariantAnimationFromPointer(ptr unsafe.Pointer) *QVariantAnimation {
	var n = new(QVariantAnimation)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QVariantAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVariantAnimation) QVariantAnimation_PTR() *QVariantAnimation {
	return ptr
}

func (ptr *QVariantAnimation) CurrentValue() *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_CurrentValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariantAnimation) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QVariantAnimation_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVariantAnimation) EasingCurve() *QEasingCurve {
	if ptr.Pointer() != nil {
		return NewQEasingCurveFromPointer(C.QVariantAnimation_EasingCurve(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariantAnimation) EndValue() *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_EndValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariantAnimation) SetDuration(msecs int) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetDuration(ptr.Pointer(), C.int(msecs))
	}
}

func (ptr *QVariantAnimation) SetEasingCurve(easing QEasingCurve_ITF) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetEasingCurve(ptr.Pointer(), PointerFromQEasingCurve(easing))
	}
}

func (ptr *QVariantAnimation) SetEndValue(value QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetEndValue(ptr.Pointer(), PointerFromQVariant(value))
	}
}

func (ptr *QVariantAnimation) SetStartValue(value QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetStartValue(ptr.Pointer(), PointerFromQVariant(value))
	}
}

func (ptr *QVariantAnimation) StartValue() *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_StartValue(ptr.Pointer()))
	}
	return nil
}

func NewQVariantAnimation(parent QObject_ITF) *QVariantAnimation {
	return NewQVariantAnimationFromPointer(C.QVariantAnimation_NewQVariantAnimation(PointerFromQObject(parent)))
}

func (ptr *QVariantAnimation) KeyValueAt(step float64) *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_KeyValueAt(ptr.Pointer(), C.double(step)))
	}
	return nil
}

func (ptr *QVariantAnimation) SetKeyValueAt(step float64, value QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetKeyValueAt(ptr.Pointer(), C.double(step), PointerFromQVariant(value))
	}
}

func (ptr *QVariantAnimation) ConnectValueChanged(f func(value *QVariant)) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QVariantAnimation) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQVariantAnimationValueChanged
func callbackQVariantAnimationValueChanged(ptrName *C.char, value unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(*QVariant))(NewQVariantFromPointer(value))
}

func (ptr *QVariantAnimation) DestroyQVariantAnimation() {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_DestroyQVariantAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
