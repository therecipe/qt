package core

//#include "qpropertyanimation.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPropertyAnimation struct {
	QVariantAnimation
}

type QPropertyAnimationITF interface {
	QVariantAnimationITF
	QPropertyAnimationPTR() *QPropertyAnimation
}

func PointerFromQPropertyAnimation(ptr QPropertyAnimationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPropertyAnimationPTR().Pointer()
	}
	return nil
}

func QPropertyAnimationFromPointer(ptr unsafe.Pointer) *QPropertyAnimation {
	var n = new(QPropertyAnimation)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPropertyAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPropertyAnimation) QPropertyAnimationPTR() *QPropertyAnimation {
	return ptr
}

func (ptr *QPropertyAnimation) SetPropertyName(propertyName QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetPropertyName(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(propertyName)))
	}
}

func (ptr *QPropertyAnimation) SetTargetObject(target QObjectITF) {
	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetTargetObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(target)))
	}
}

func (ptr *QPropertyAnimation) TargetObject() *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QPropertyAnimation_TargetObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQPropertyAnimation(parent QObjectITF) *QPropertyAnimation {
	return QPropertyAnimationFromPointer(unsafe.Pointer(C.QPropertyAnimation_NewQPropertyAnimation(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQPropertyAnimation2(target QObjectITF, propertyName QByteArrayITF, parent QObjectITF) *QPropertyAnimation {
	return QPropertyAnimationFromPointer(unsafe.Pointer(C.QPropertyAnimation_NewQPropertyAnimation2(C.QtObjectPtr(PointerFromQObject(target)), C.QtObjectPtr(PointerFromQByteArray(propertyName)), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QPropertyAnimation) DestroyQPropertyAnimation() {
	if ptr.Pointer() != nil {
		C.QPropertyAnimation_DestroyQPropertyAnimation(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
