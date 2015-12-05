package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QPropertyAnimation_") {
		n.SetObjectName("QPropertyAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPropertyAnimation) QPropertyAnimation_PTR() *QPropertyAnimation {
	return ptr
}

func (ptr *QPropertyAnimation) PropertyName() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPropertyAnimation::propertyName")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QPropertyAnimation_PropertyName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPropertyAnimation) SetPropertyName(propertyName QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPropertyAnimation::setPropertyName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetPropertyName(ptr.Pointer(), PointerFromQByteArray(propertyName))
	}
}

func (ptr *QPropertyAnimation) SetTargetObject(target QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPropertyAnimation::setTargetObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetTargetObject(ptr.Pointer(), PointerFromQObject(target))
	}
}

func (ptr *QPropertyAnimation) TargetObject() *QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPropertyAnimation::targetObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QPropertyAnimation_TargetObject(ptr.Pointer()))
	}
	return nil
}

func NewQPropertyAnimation(parent QObject_ITF) *QPropertyAnimation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPropertyAnimation::QPropertyAnimation")
		}
	}()

	return NewQPropertyAnimationFromPointer(C.QPropertyAnimation_NewQPropertyAnimation(PointerFromQObject(parent)))
}

func NewQPropertyAnimation2(target QObject_ITF, propertyName QByteArray_ITF, parent QObject_ITF) *QPropertyAnimation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPropertyAnimation::QPropertyAnimation")
		}
	}()

	return NewQPropertyAnimationFromPointer(C.QPropertyAnimation_NewQPropertyAnimation2(PointerFromQObject(target), PointerFromQByteArray(propertyName), PointerFromQObject(parent)))
}

func (ptr *QPropertyAnimation) DestroyQPropertyAnimation() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPropertyAnimation::~QPropertyAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_DestroyQPropertyAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
