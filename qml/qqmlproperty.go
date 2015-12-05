package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QQmlProperty struct {
	ptr unsafe.Pointer
}

type QQmlProperty_ITF interface {
	QQmlProperty_PTR() *QQmlProperty
}

func (p *QQmlProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlProperty(ptr QQmlProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyFromPointer(ptr unsafe.Pointer) *QQmlProperty {
	var n = new(QQmlProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlProperty) QQmlProperty_PTR() *QQmlProperty {
	return ptr
}

//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int64

const (
	QQmlProperty__InvalidCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          = QQmlProperty__PropertyTypeCategory(3)
)

//QQmlProperty::Type
type QQmlProperty__Type int64

const (
	QQmlProperty__Invalid        = QQmlProperty__Type(0)
	QQmlProperty__Property       = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty = QQmlProperty__Type(2)
)

func NewQQmlProperty() *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty())
}

func NewQQmlProperty2(obj core.QObject_ITF) *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty2(core.PointerFromQObject(obj)))
}

func NewQQmlProperty3(obj core.QObject_ITF, ctxt QQmlContext_ITF) *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty3(core.PointerFromQObject(obj), PointerFromQQmlContext(ctxt)))
}

func NewQQmlProperty4(obj core.QObject_ITF, engine QQmlEngine_ITF) *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty4(core.PointerFromQObject(obj), PointerFromQQmlEngine(engine)))
}

func NewQQmlProperty5(obj core.QObject_ITF, name string) *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty5(core.PointerFromQObject(obj), C.CString(name)))
}

func NewQQmlProperty6(obj core.QObject_ITF, name string, ctxt QQmlContext_ITF) *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty6(core.PointerFromQObject(obj), C.CString(name), PointerFromQQmlContext(ctxt)))
}

func NewQQmlProperty7(obj core.QObject_ITF, name string, engine QQmlEngine_ITF) *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty7(core.PointerFromQObject(obj), C.CString(name), PointerFromQQmlEngine(engine)))
}

func NewQQmlProperty8(other QQmlProperty_ITF) *QQmlProperty {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::QQmlProperty")
		}
	}()

	return NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty8(PointerFromQQmlProperty(other)))
}

func (ptr *QQmlProperty) ConnectNotifySignal(dest core.QObject_ITF, slot string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::connectNotifySignal")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal(ptr.Pointer(), core.PointerFromQObject(dest), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QQmlProperty) ConnectNotifySignal2(dest core.QObject_ITF, method int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::connectNotifySignal")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal2(ptr.Pointer(), core.PointerFromQObject(dest), C.int(method)) != 0
	}
	return false
}

func (ptr *QQmlProperty) HasNotifySignal() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::hasNotifySignal")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_HasNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Index() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::index")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QQmlProperty_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) IsDesignable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::isDesignable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsDesignable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsProperty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::isProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsResettable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::isResettable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsResettable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsSignalProperty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::isSignalProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsSignalProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsWritable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::isWritable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlProperty_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlProperty) NeedsNotifySignal() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::needsNotifySignal")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_NeedsNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Object() *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::object")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlProperty_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlProperty) PropertyType() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::propertyType")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QQmlProperty_PropertyType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyTypeCategory() QQmlProperty__PropertyTypeCategory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::propertyTypeCategory")
		}
	}()

	if ptr.Pointer() != nil {
		return QQmlProperty__PropertyTypeCategory(C.QQmlProperty_PropertyTypeCategory(ptr.Pointer()))
	}
	return 0
}

func QQmlProperty_Read2(object core.QObject_ITF, name string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::read")
		}
	}()

	return core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), C.CString(name)))
}

func QQmlProperty_Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::read")
		}
	}()

	return core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), C.CString(name), PointerFromQQmlContext(ctxt)))
}

func QQmlProperty_Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::read")
		}
	}()

	return core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), C.CString(name), PointerFromQQmlEngine(engine)))
}

func (ptr *QQmlProperty) Read() *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::read")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlProperty_Read(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlProperty) Reset() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::reset")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Type() QQmlProperty__Type {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::type")
		}
	}()

	if ptr.Pointer() != nil {
		return QQmlProperty__Type(C.QQmlProperty_Type(ptr.Pointer()))
	}
	return 0
}

func QQmlProperty_Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::write")
		}
	}()

	return C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), C.CString(name), core.PointerFromQVariant(value)) != 0
}

func QQmlProperty_Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::write")
		}
	}()

	return C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), C.CString(name), core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt)) != 0
}

func QQmlProperty_Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::write")
		}
	}()

	return C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), C.CString(name), core.PointerFromQVariant(value), PointerFromQQmlEngine(engine)) != 0
}

func (ptr *QQmlProperty) Write(value core.QVariant_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlProperty::write")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlProperty_Write(ptr.Pointer(), core.PointerFromQVariant(value)) != 0
	}
	return false
}
