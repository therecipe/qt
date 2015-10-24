package core

//#include "qobject.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QObject struct {
	ptr unsafe.Pointer
}

type QObjectITF interface {
	QObjectPTR() *QObject
}

func (p *QObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQObject(ptr QObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QObjectPTR().Pointer()
	}
	return nil
}

func QObjectFromPointer(ptr unsafe.Pointer) *QObject {
	var n = new(QObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QObject) QObjectPTR() *QObject {
	return ptr
}

func (ptr *QObject) InstallEventFilter(filterObj QObjectITF) {
	if ptr.Pointer() != nil {
		C.QObject_InstallEventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(filterObj)))
	}
}

func (ptr *QObject) ObjectName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QObject_ObjectName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QObject) SetObjectName(name string) {
	if ptr.Pointer() != nil {
		C.QObject_SetObjectName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func NewQObject(parent QObjectITF) *QObject {
	return QObjectFromPointer(unsafe.Pointer(C.QObject_NewQObject(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QObject) BlockSignals(block bool) bool {
	if ptr.Pointer() != nil {
		return C.QObject_BlockSignals(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(block))) != 0
	}
	return false
}

func (ptr *QObject) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QObject_DeleteLater(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QObject) ConnectDestroyed(f func(obj QObjectITF)) {
	if ptr.Pointer() != nil {
		C.QObject_ConnectDestroyed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "destroyed", f)
	}
}

func (ptr *QObject) DisconnectDestroyed() {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectDestroyed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "destroyed")
	}
}

//export callbackQObjectDestroyed
func callbackQObjectDestroyed(ptrName *C.char, obj unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "destroyed").(func(*QObject))(QObjectFromPointer(obj))
}

func (ptr *QObject) DumpObjectInfo() {
	if ptr.Pointer() != nil {
		C.QObject_DumpObjectInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QObject) DumpObjectTree() {
	if ptr.Pointer() != nil {
		C.QObject_DumpObjectTree(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QObject) Event(e QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QObject_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QObject) EventFilter(watched QObjectITF, event QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QObject_EventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(watched)), C.QtObjectPtr(PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QObject) Inherits(className string) bool {
	if ptr.Pointer() != nil {
		return C.QObject_Inherits(C.QtObjectPtr(ptr.Pointer()), C.CString(className)) != 0
	}
	return false
}

func (ptr *QObject) IsWidgetType() bool {
	if ptr.Pointer() != nil {
		return C.QObject_IsWidgetType(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QObject) IsWindowType() bool {
	if ptr.Pointer() != nil {
		return C.QObject_IsWindowType(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QObject) KillTimer(id int) {
	if ptr.Pointer() != nil {
		C.QObject_KillTimer(C.QtObjectPtr(ptr.Pointer()), C.int(id))
	}
}

func (ptr *QObject) MetaObject() *QMetaObject {
	if ptr.Pointer() != nil {
		return QMetaObjectFromPointer(unsafe.Pointer(C.QObject_MetaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QObject) MoveToThread(targetThread QThreadITF) {
	if ptr.Pointer() != nil {
		C.QObject_MoveToThread(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQThread(targetThread)))
	}
}

func (ptr *QObject) ConnectObjectNameChanged(f func(objectName string)) {
	if ptr.Pointer() != nil {
		C.QObject_ConnectObjectNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "objectNameChanged", f)
	}
}

func (ptr *QObject) DisconnectObjectNameChanged() {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectObjectNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "objectNameChanged")
	}
}

//export callbackQObjectObjectNameChanged
func callbackQObjectObjectNameChanged(ptrName *C.char, objectName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "objectNameChanged").(func(string))(C.GoString(objectName))
}

func (ptr *QObject) Parent() *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QObject_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QObject) Property(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QObject_Property(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QObject) RemoveEventFilter(obj QObjectITF) {
	if ptr.Pointer() != nil {
		C.QObject_RemoveEventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(obj)))
	}
}

func (ptr *QObject) SetParent(parent QObjectITF) {
	if ptr.Pointer() != nil {
		C.QObject_SetParent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(parent)))
	}
}

func (ptr *QObject) SetProperty(name string, value string) bool {
	if ptr.Pointer() != nil {
		return C.QObject_SetProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QObject) StartTimer(interval int, timerType Qt__TimerType) int {
	if ptr.Pointer() != nil {
		return int(C.QObject_StartTimer(C.QtObjectPtr(ptr.Pointer()), C.int(interval), C.int(timerType)))
	}
	return 0
}

func (ptr *QObject) SignalsBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QObject_SignalsBlocked(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QObject) Thread() *QThread {
	if ptr.Pointer() != nil {
		return QThreadFromPointer(unsafe.Pointer(C.QObject_Thread(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func QObject_Tr(sourceText string, disambiguation string, n int) string {
	return C.GoString(C.QObject_QObject_Tr(C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}

func (ptr *QObject) DestroyQObject() {
	if ptr.Pointer() != nil {
		C.QObject_DestroyQObject(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
