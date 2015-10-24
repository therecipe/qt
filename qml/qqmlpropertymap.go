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

type QQmlPropertyMapITF interface {
	core.QObjectITF
	QQmlPropertyMapPTR() *QQmlPropertyMap
}

func PointerFromQQmlPropertyMap(ptr QQmlPropertyMapITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyMapPTR().Pointer()
	}
	return nil
}

func QQmlPropertyMapFromPointer(ptr unsafe.Pointer) *QQmlPropertyMap {
	var n = new(QQmlPropertyMap)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlPropertyMap_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlPropertyMap) QQmlPropertyMapPTR() *QQmlPropertyMap {
	return ptr
}

func NewQQmlPropertyMap(parent core.QObjectITF) *QQmlPropertyMap {
	return QQmlPropertyMapFromPointer(unsafe.Pointer(C.QQmlPropertyMap_NewQQmlPropertyMap(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQmlPropertyMap) Clear(key string) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_Clear(C.QtObjectPtr(ptr.Pointer()), C.CString(key))
	}
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_Contains(C.QtObjectPtr(ptr.Pointer()), C.CString(key)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlPropertyMap_Keys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlPropertyMap) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlPropertyMap) Value(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlPropertyMap_Value(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value string)) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQQmlPropertyMapValueChanged
func callbackQQmlPropertyMapValueChanged(ptrName *C.char, key *C.char, value *C.char) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(string, string))(C.GoString(key), C.GoString(value))
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMap(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
