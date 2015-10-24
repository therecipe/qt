package qml

//#include "qqmlproperty.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlProperty struct {
	ptr unsafe.Pointer
}

type QQmlPropertyITF interface {
	QQmlPropertyPTR() *QQmlProperty
}

func (p *QQmlProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlProperty(ptr QQmlPropertyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyPTR().Pointer()
	}
	return nil
}

func QQmlPropertyFromPointer(ptr unsafe.Pointer) *QQmlProperty {
	var n = new(QQmlProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlProperty) QQmlPropertyPTR() *QQmlProperty {
	return ptr
}

//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int

var (
	QQmlProperty__InvalidCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          = QQmlProperty__PropertyTypeCategory(3)
)

//QQmlProperty::Type
type QQmlProperty__Type int

var (
	QQmlProperty__Invalid        = QQmlProperty__Type(0)
	QQmlProperty__Property       = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty = QQmlProperty__Type(2)
)

func NewQQmlProperty() *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty()))
}

func NewQQmlProperty2(obj core.QObjectITF) *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty2(C.QtObjectPtr(core.PointerFromQObject(obj)))))
}

func NewQQmlProperty3(obj core.QObjectITF, ctxt QQmlContextITF) *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty3(C.QtObjectPtr(core.PointerFromQObject(obj)), C.QtObjectPtr(PointerFromQQmlContext(ctxt)))))
}

func NewQQmlProperty4(obj core.QObjectITF, engine QQmlEngineITF) *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty4(C.QtObjectPtr(core.PointerFromQObject(obj)), C.QtObjectPtr(PointerFromQQmlEngine(engine)))))
}

func NewQQmlProperty5(obj core.QObjectITF, name string) *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty5(C.QtObjectPtr(core.PointerFromQObject(obj)), C.CString(name))))
}

func NewQQmlProperty6(obj core.QObjectITF, name string, ctxt QQmlContextITF) *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty6(C.QtObjectPtr(core.PointerFromQObject(obj)), C.CString(name), C.QtObjectPtr(PointerFromQQmlContext(ctxt)))))
}

func NewQQmlProperty7(obj core.QObjectITF, name string, engine QQmlEngineITF) *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty7(C.QtObjectPtr(core.PointerFromQObject(obj)), C.CString(name), C.QtObjectPtr(PointerFromQQmlEngine(engine)))))
}

func NewQQmlProperty8(other QQmlPropertyITF) *QQmlProperty {
	return QQmlPropertyFromPointer(unsafe.Pointer(C.QQmlProperty_NewQQmlProperty8(C.QtObjectPtr(PointerFromQQmlProperty(other)))))
}

func (ptr *QQmlProperty) ConnectNotifySignal(dest core.QObjectITF, slot string) bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(dest)), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QQmlProperty) ConnectNotifySignal2(dest core.QObjectITF, method int) bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(dest)), C.int(method)) != 0
	}
	return false
}

func (ptr *QQmlProperty) HasNotifySignal() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_HasNotifySignal(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) Index() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlProperty_Index(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlProperty) IsDesignable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsDesignable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsProperty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsProperty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsResettable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsResettable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsSignalProperty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsSignalProperty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsWritable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlProperty_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlProperty) NeedsNotifySignal() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_NeedsNotifySignal(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) Object() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlProperty_Object(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlProperty) PropertyType() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlProperty_PropertyType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyTypeCategory() QQmlProperty__PropertyTypeCategory {
	if ptr.Pointer() != nil {
		return QQmlProperty__PropertyTypeCategory(C.QQmlProperty_PropertyTypeCategory(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QQmlProperty_Read2(object core.QObjectITF, name string) string {
	return C.GoString(C.QQmlProperty_QQmlProperty_Read2(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(name)))
}

func QQmlProperty_Read3(object core.QObjectITF, name string, ctxt QQmlContextITF) string {
	return C.GoString(C.QQmlProperty_QQmlProperty_Read3(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(name), C.QtObjectPtr(PointerFromQQmlContext(ctxt))))
}

func QQmlProperty_Read4(object core.QObjectITF, name string, engine QQmlEngineITF) string {
	return C.GoString(C.QQmlProperty_QQmlProperty_Read4(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(name), C.QtObjectPtr(PointerFromQQmlEngine(engine))))
}

func (ptr *QQmlProperty) Read() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlProperty_Read(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlProperty) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_Reset(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) Type() QQmlProperty__Type {
	if ptr.Pointer() != nil {
		return QQmlProperty__Type(C.QQmlProperty_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QQmlProperty_Write2(object core.QObjectITF, name string, value string) bool {
	return C.QQmlProperty_QQmlProperty_Write2(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(name), C.CString(value)) != 0
}

func QQmlProperty_Write3(object core.QObjectITF, name string, value string, ctxt QQmlContextITF) bool {
	return C.QQmlProperty_QQmlProperty_Write3(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(name), C.CString(value), C.QtObjectPtr(PointerFromQQmlContext(ctxt))) != 0
}

func QQmlProperty_Write4(object core.QObjectITF, name string, value string, engine QQmlEngineITF) bool {
	return C.QQmlProperty_QQmlProperty_Write4(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(name), C.CString(value), C.QtObjectPtr(PointerFromQQmlEngine(engine))) != 0
}

func (ptr *QQmlProperty) Write(value string) bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_Write(C.QtObjectPtr(ptr.Pointer()), C.CString(value)) != 0
	}
	return false
}
