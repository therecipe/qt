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

type QVariantAnimationITF interface {
	QAbstractAnimationITF
	QVariantAnimationPTR() *QVariantAnimation
}

func PointerFromQVariantAnimation(ptr QVariantAnimationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVariantAnimationPTR().Pointer()
	}
	return nil
}

func QVariantAnimationFromPointer(ptr unsafe.Pointer) *QVariantAnimation {
	var n = new(QVariantAnimation)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QVariantAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVariantAnimation) QVariantAnimationPTR() *QVariantAnimation {
	return ptr
}

func (ptr *QVariantAnimation) CurrentValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVariantAnimation_CurrentValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QVariantAnimation) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QVariantAnimation_Duration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVariantAnimation) EndValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVariantAnimation_EndValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QVariantAnimation) SetDuration(msecs int) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetDuration(C.QtObjectPtr(ptr.Pointer()), C.int(msecs))
	}
}

func (ptr *QVariantAnimation) SetEasingCurve(easing QEasingCurveITF) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetEasingCurve(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQEasingCurve(easing)))
	}
}

func (ptr *QVariantAnimation) SetEndValue(value string) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetEndValue(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QVariantAnimation) SetStartValue(value string) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetStartValue(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QVariantAnimation) StartValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVariantAnimation_StartValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQVariantAnimation(parent QObjectITF) *QVariantAnimation {
	return QVariantAnimationFromPointer(unsafe.Pointer(C.QVariantAnimation_NewQVariantAnimation(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QVariantAnimation) ConnectValueChanged(f func(value string)) {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_ConnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QVariantAnimation) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_DisconnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQVariantAnimationValueChanged
func callbackQVariantAnimationValueChanged(ptrName *C.char, value *C.char) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(string))(C.GoString(value))
}

func (ptr *QVariantAnimation) DestroyQVariantAnimation() {
	if ptr.Pointer() != nil {
		C.QVariantAnimation_DestroyQVariantAnimation(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
