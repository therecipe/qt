package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QObject_") {
		n.SetObjectName("QObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QObject) QObject_PTR() *QObject {
	return ptr
}

func (ptr *QObject) InstallEventFilter(filterObj QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::installEventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_InstallEventFilter(ptr.Pointer(), PointerFromQObject(filterObj))
	}
}

func (ptr *QObject) ObjectName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::objectName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QObject_ObjectName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QObject) SetObjectName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::setObjectName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_SetObjectName(ptr.Pointer(), C.CString(name))
	}
}

func NewQObject(parent QObject_ITF) *QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::QObject")
		}
	}()

	return NewQObjectFromPointer(C.QObject_NewQObject(PointerFromQObject(parent)))
}

func (ptr *QObject) BlockSignals(block bool) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::blockSignals")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_BlockSignals(ptr.Pointer(), C.int(qt.GoBoolToInt(block))) != 0
	}
	return false
}

func (ptr *QObject) DeleteLater() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::deleteLater")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QObject) ConnectDestroyed(f func(obj *QObject)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::destroyed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_ConnectDestroyed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "destroyed", f)
	}
}

func (ptr *QObject) DisconnectDestroyed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::destroyed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_DisconnectDestroyed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "destroyed")
	}
}

//export callbackQObjectDestroyed
func callbackQObjectDestroyed(ptrName *C.char, obj unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::destroyed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "destroyed").(func(*QObject))(NewQObjectFromPointer(obj))
}

func (ptr *QObject) DumpObjectInfo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::dumpObjectInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_DumpObjectInfo(ptr.Pointer())
	}
}

func (ptr *QObject) DumpObjectTree() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::dumpObjectTree")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_DumpObjectTree(ptr.Pointer())
	}
}

func (ptr *QObject) Event(e QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::event")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_Event(ptr.Pointer(), PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QObject) EventFilter(watched QObject_ITF, event QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::eventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_EventFilter(ptr.Pointer(), PointerFromQObject(watched), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QObject) FindChild(name string, options Qt__FindChildOption) unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::findChild")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QObject_FindChild(ptr.Pointer(), C.CString(name), C.int(options)))
	}
	return nil
}

func (ptr *QObject) Inherits(className string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::inherits")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_Inherits(ptr.Pointer(), C.CString(className)) != 0
	}
	return false
}

func (ptr *QObject) IsWidgetType() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::isWidgetType")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_IsWidgetType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObject) IsWindowType() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::isWindowType")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_IsWindowType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObject) KillTimer(id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::killTimer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_KillTimer(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QObject) MetaObject() *QMetaObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::metaObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMetaObjectFromPointer(C.QObject_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QObject) MoveToThread(targetThread QThread_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::moveToThread")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_MoveToThread(ptr.Pointer(), PointerFromQThread(targetThread))
	}
}

func (ptr *QObject) ConnectObjectNameChanged(f func(objectName string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::objectNameChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_ConnectObjectNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "objectNameChanged", f)
	}
}

func (ptr *QObject) DisconnectObjectNameChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::objectNameChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_DisconnectObjectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "objectNameChanged")
	}
}

//export callbackQObjectObjectNameChanged
func callbackQObjectObjectNameChanged(ptrName *C.char, objectName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::objectNameChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "objectNameChanged").(func(string))(C.GoString(objectName))
}

func (ptr *QObject) Parent() *QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QObject_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QObject) Property(name string) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::property")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QObject_Property(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QObject) RemoveEventFilter(obj QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::removeEventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_RemoveEventFilter(ptr.Pointer(), PointerFromQObject(obj))
	}
}

func (ptr *QObject) SetParent(parent QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::setParent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_SetParent(ptr.Pointer(), PointerFromQObject(parent))
	}
}

func (ptr *QObject) SetProperty(name string, value QVariant_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::setProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_SetProperty(ptr.Pointer(), C.CString(name), PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QObject) StartTimer(interval int, timerType Qt__TimerType) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::startTimer")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QObject_StartTimer(ptr.Pointer(), C.int(interval), C.int(timerType)))
	}
	return 0
}

func (ptr *QObject) SignalsBlocked() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::signalsBlocked")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObject_SignalsBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObject) Thread() *QThread {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::thread")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQThreadFromPointer(C.QObject_Thread(ptr.Pointer()))
	}
	return nil
}

func QObject_Tr(sourceText string, disambiguation string, n int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::tr")
		}
	}()

	return C.GoString(C.QObject_QObject_Tr(C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}

func (ptr *QObject) DestroyQObject() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObject::~QObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObject_DestroyQObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
