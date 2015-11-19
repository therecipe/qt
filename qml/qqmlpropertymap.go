package qml

//#include "qqmlpropertymap.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QQmlPropertyMap struct {
	core.QObject
}

type QQmlPropertyMap_ITF interface {
	core.QObject_ITF
	QQmlPropertyMap_PTR() *QQmlPropertyMap
}

func PointerFromQQmlPropertyMap(ptr QQmlPropertyMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyMap_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyMapFromPointer(ptr unsafe.Pointer) *QQmlPropertyMap {
	var n = new(QQmlPropertyMap)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlPropertyMap_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap {
	return ptr
}

func NewQQmlPropertyMap(parent core.QObject_ITF) *QQmlPropertyMap {
	return NewQQmlPropertyMapFromPointer(C.QQmlPropertyMap_NewQQmlPropertyMap(core.PointerFromQObject(parent)))
}

func (ptr *QQmlPropertyMap) Clear(key string) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_Clear(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_Contains(ptr.Pointer(), C.CString(key)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlPropertyMap_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlPropertyMap) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlPropertyMap) Value(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlPropertyMap_Value(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value *core.QVariant)) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQQmlPropertyMapValueChanged
func callbackQQmlPropertyMapValueChanged(ptrName *C.char, key *C.char, value unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
