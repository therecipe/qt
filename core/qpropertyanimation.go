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

type QPropertyAnimation_ITF interface {
	QVariantAnimation_ITF
	QPropertyAnimation_PTR() *QPropertyAnimation
}

func PointerFromQPropertyAnimation(ptr QPropertyAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPropertyAnimation_PTR().Pointer()
	}
	return nil
}

func NewQPropertyAnimationFromPointer(ptr unsafe.Pointer) *QPropertyAnimation {
	var n = new(QPropertyAnimation)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPropertyAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPropertyAnimation) QPropertyAnimation_PTR() *QPropertyAnimation {
	return ptr
}

func (ptr *QPropertyAnimation) PropertyName() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QPropertyAnimation_PropertyName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPropertyAnimation) SetPropertyName(propertyName QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetPropertyName(ptr.Pointer(), PointerFromQByteArray(propertyName))
	}
}

func (ptr *QPropertyAnimation) SetTargetObject(target QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetTargetObject(ptr.Pointer(), PointerFromQObject(target))
	}
}

func (ptr *QPropertyAnimation) TargetObject() *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QPropertyAnimation_TargetObject(ptr.Pointer()))
	}
	return nil
}

func NewQPropertyAnimation(parent QObject_ITF) *QPropertyAnimation {
	return NewQPropertyAnimationFromPointer(C.QPropertyAnimation_NewQPropertyAnimation(PointerFromQObject(parent)))
}

func NewQPropertyAnimation2(target QObject_ITF, propertyName QByteArray_ITF, parent QObject_ITF) *QPropertyAnimation {
	return NewQPropertyAnimationFromPointer(C.QPropertyAnimation_NewQPropertyAnimation2(PointerFromQObject(target), PointerFromQByteArray(propertyName), PointerFromQObject(parent)))
}

func (ptr *QPropertyAnimation) DestroyQPropertyAnimation() {
	if ptr.Pointer() != nil {
		C.QPropertyAnimation_DestroyQPropertyAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
