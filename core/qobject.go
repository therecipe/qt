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

type QObject_ITF interface {
	QObject_PTR() *QObject
}

func (p *QObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQObject(ptr QObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func NewQObjectFromPointer(ptr unsafe.Pointer) *QObject {
	var n = new(QObject)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QObject) QObject_PTR() *QObject {
	return ptr
}

func (ptr *QObject) InstallEventFilter(filterObj QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_InstallEventFilter(ptr.Pointer(), PointerFromQObject(filterObj))
	}
}

func (ptr *QObject) ObjectName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QObject_ObjectName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QObject) SetObjectName(name string) {
	if ptr.Pointer() != nil {
		C.QObject_SetObjectName(ptr.Pointer(), C.CString(name))
	}
}

func NewQObject(parent QObject_ITF) *QObject {
	return NewQObjectFromPointer(C.QObject_NewQObject(PointerFromQObject(parent)))
}

func (ptr *QObject) BlockSignals(block bool) bool {
	if ptr.Pointer() != nil {
		return C.QObject_BlockSignals(ptr.Pointer(), C.int(qt.GoBoolToInt(block))) != 0
	}
	return false
}

func (ptr *QObject) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QObject_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QObject) ConnectDestroyed(f func(obj *QObject)) {
	if ptr.Pointer() != nil {
		C.QObject_ConnectDestroyed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "destroyed", f)
	}
}

func (ptr *QObject) DisconnectDestroyed() {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectDestroyed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "destroyed")
	}
}

//export callbackQObjectDestroyed
func callbackQObjectDestroyed(ptrName *C.char, obj unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "destroyed").(func(*QObject))(NewQObjectFromPointer(obj))
}

func (ptr *QObject) DumpObjectInfo() {
	if ptr.Pointer() != nil {
		C.QObject_DumpObjectInfo(ptr.Pointer())
	}
}

func (ptr *QObject) DumpObjectTree() {
	if ptr.Pointer() != nil {
		C.QObject_DumpObjectTree(ptr.Pointer())
	}
}

func (ptr *QObject) Event(e QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QObject_Event(ptr.Pointer(), PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QObject) EventFilter(watched QObject_ITF, event QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QObject_EventFilter(ptr.Pointer(), PointerFromQObject(watched), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QObject) FindChild(name string, options Qt__FindChildOption) unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QObject_FindChild(ptr.Pointer(), C.CString(name), C.int(options)))
	}
	return nil
}

func (ptr *QObject) Inherits(className string) bool {
	if ptr.Pointer() != nil {
		return C.QObject_Inherits(ptr.Pointer(), C.CString(className)) != 0
	}
	return false
}

func (ptr *QObject) IsWidgetType() bool {
	if ptr.Pointer() != nil {
		return C.QObject_IsWidgetType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObject) IsWindowType() bool {
	if ptr.Pointer() != nil {
		return C.QObject_IsWindowType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObject) KillTimer(id int) {
	if ptr.Pointer() != nil {
		C.QObject_KillTimer(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QObject) MetaObject() *QMetaObject {
	if ptr.Pointer() != nil {
		return NewQMetaObjectFromPointer(C.QObject_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QObject) MoveToThread(targetThread QThread_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_MoveToThread(ptr.Pointer(), PointerFromQThread(targetThread))
	}
}

func (ptr *QObject) ConnectObjectNameChanged(f func(objectName string)) {
	if ptr.Pointer() != nil {
		C.QObject_ConnectObjectNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "objectNameChanged", f)
	}
}

func (ptr *QObject) DisconnectObjectNameChanged() {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectObjectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "objectNameChanged")
	}
}

//export callbackQObjectObjectNameChanged
func callbackQObjectObjectNameChanged(ptrName *C.char, objectName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "objectNameChanged").(func(string))(C.GoString(objectName))
}

func (ptr *QObject) Parent() *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QObject_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QObject) Property(name string) *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QObject_Property(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QObject) RemoveEventFilter(obj QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_RemoveEventFilter(ptr.Pointer(), PointerFromQObject(obj))
	}
}

func (ptr *QObject) SetParent(parent QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_SetParent(ptr.Pointer(), PointerFromQObject(parent))
	}
}

func (ptr *QObject) SetProperty(name string, value QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QObject_SetProperty(ptr.Pointer(), C.CString(name), PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QObject) StartTimer(interval int, timerType Qt__TimerType) int {
	if ptr.Pointer() != nil {
		return int(C.QObject_StartTimer(ptr.Pointer(), C.int(interval), C.int(timerType)))
	}
	return 0
}

func (ptr *QObject) SignalsBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QObject_SignalsBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObject) Thread() *QThread {
	if ptr.Pointer() != nil {
		return NewQThreadFromPointer(C.QObject_Thread(ptr.Pointer()))
	}
	return nil
}

func QObject_Tr(sourceText string, disambiguation string, n int) string {
	return C.GoString(C.QObject_QObject_Tr(C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}

func (ptr *QObject) DestroyQObject() {
	if ptr.Pointer() != nil {
		C.QObject_DestroyQObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
