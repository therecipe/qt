package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QMetaProperty struct {
	ptr unsafe.Pointer
}

type QMetaProperty_ITF interface {
	QMetaProperty_PTR() *QMetaProperty
}

func (p *QMetaProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaProperty(ptr QMetaProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaProperty_PTR().Pointer()
	}
	return nil
}

func NewQMetaPropertyFromPointer(ptr unsafe.Pointer) *QMetaProperty {
	var n = new(QMetaProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaProperty) QMetaProperty_PTR() *QMetaProperty {
	return ptr
}

func (ptr *QMetaProperty) HasNotifySignal() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::hasNotifySignal")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_HasNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsConstant() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isConstant")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsConstant(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsDesignable(object QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isDesignable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsDesignable(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsEnumType() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isEnumType")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsEnumType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsFinal() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isFinal")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsFinal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsFlagType() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isFlagType")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsFlagType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsReadable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isReadable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsResettable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isResettable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsResettable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsScriptable(object QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isScriptable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsScriptable(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsStored(object QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isStored")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsStored(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsUser(object QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isUser")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsUser(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsWritable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::isWritable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaProperty) NotifySignalIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::notifySignalIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_NotifySignalIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaProperty) PropertyIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::propertyIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_PropertyIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaProperty) Read(object QObject_ITF) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::read")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QMetaProperty_Read(ptr.Pointer(), PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QMetaProperty) Reset(object QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::reset")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_Reset(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMetaProperty) Revision() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::revision")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_Revision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaProperty) UserType() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::userType")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_UserType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaProperty) Write(object QObject_ITF, value QVariant_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMetaProperty::write")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMetaProperty_Write(ptr.Pointer(), PointerFromQObject(object), PointerFromQVariant(value)) != 0
	}
	return false
}
