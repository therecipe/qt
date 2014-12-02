package qt

//#include "qobject.h"
import "C"

type qobject struct {
	ptr C.QtObjectPtr
}

type QObject interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Destroy()
	BlockSignals_Bool(block bool) bool
	Disconnect_String_QObject_String(signal string, receiver QObject, method string) bool
	Disconnect_QObject_String(receiver QObject, method string) bool
	DumpObjectInfo()
	DumpObjectTree()
	Inherits_String(className string) bool
	InstallEventFilter_QObject(filterObj QObject)
	IsWidgetType() bool
	IsWindowType() bool
	KillTimer_Int(id int)
	MoveToThread_QThread(targetThread QThread)
	ObjectName() string
	Parent() QObject
	RemoveEventFilter_QObject(obj QObject)
	SetObjectName_String(name string)
	SetParent_QObject(parent QObject)
	SignalsBlocked() bool
	StartTimer_Int_TimerType(interval int, timerType TimerType) int
	Thread() QThread
	ConnectSlotDeleteLater()
	DisconnectSlotDeleteLater()
	SlotDeleteLater()
	ConnectSignalDestroyed(f func())
	DisconnectSignalDestroyed()
	SignalDestroyed() func()
	ConnectSignalObjectNameChanged(f func())
	DisconnectSignalObjectNameChanged()
	SignalObjectNameChanged() func()
}

func (p *qobject) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qobject) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQObject_QObject(parent QObject) QObject {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qobject = new(qobject)
	qobject.SetPointer(C.QObject_New_QObject(parentPtr))
	qobject.SetObjectName_String("QObject_" + randomIdentifier())
	return qobject
}

func (p *qobject) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QObject_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qobject) BlockSignals_Bool(block bool) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QObject_BlockSignals_Bool(p.Pointer(), goBoolToCInt(block)) != 0
	}
}

func (p *qobject) Disconnect_String_QObject_String(signal string, receiver QObject, method string) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var receiverPtr C.QtObjectPtr = nil
		if receiver != nil {
			receiverPtr = receiver.Pointer()
		}
		return C.QObject_Disconnect_String_QObject_String(p.Pointer(), C.CString(signal), receiverPtr, C.CString(method)) != 0
	}
}

func (p *qobject) Disconnect_QObject_String(receiver QObject, method string) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var receiverPtr C.QtObjectPtr = nil
		if receiver != nil {
			receiverPtr = receiver.Pointer()
		}
		return C.QObject_Disconnect_QObject_String(p.Pointer(), receiverPtr, C.CString(method)) != 0
	}
}

func (p *qobject) DumpObjectInfo() {
	if p.Pointer() != nil {
		C.QObject_DumpObjectInfo(p.Pointer())
	}
}

func (p *qobject) DumpObjectTree() {
	if p.Pointer() != nil {
		C.QObject_DumpObjectTree(p.Pointer())
	}
}

func (p *qobject) Inherits_String(className string) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QObject_Inherits_String(p.Pointer(), C.CString(className)) != 0
	}
}

func (p *qobject) InstallEventFilter_QObject(filterObj QObject) {
	if p.Pointer() == nil {
	} else {
		var filterObjPtr C.QtObjectPtr = nil
		if filterObj != nil {
			filterObjPtr = filterObj.Pointer()
		}
		C.QObject_InstallEventFilter_QObject(p.Pointer(), filterObjPtr)
	}
}

func (p *qobject) IsWidgetType() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QObject_IsWidgetType(p.Pointer()) != 0
	}
}

func (p *qobject) IsWindowType() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QObject_IsWindowType(p.Pointer()) != 0
	}
}

func (p *qobject) KillTimer_Int(id int) {
	if p.Pointer() != nil {
		C.QObject_KillTimer_Int(p.Pointer(), C.int(id))
	}
}

func (p *qobject) MoveToThread_QThread(targetThread QThread) {
	if p.Pointer() == nil {
	} else {
		var targetThreadPtr C.QtObjectPtr = nil
		if targetThread != nil {
			targetThreadPtr = targetThread.Pointer()
		}
		C.QObject_MoveToThread_QThread(p.Pointer(), targetThreadPtr)
	}
}

func (p *qobject) ObjectName() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QObject_ObjectName(p.Pointer()))
	}
}

func (p *qobject) Parent() QObject {
	if p.Pointer() == nil {
		return nil
	} else {
		var qobject = new(qobject)
		qobject.SetPointer(C.QObject_Parent(p.Pointer()))
		if qobject.ObjectName() == "" {
			qobject.SetObjectName_String("QObject_" + randomIdentifier())
		}
		return qobject
	}
}

func (p *qobject) RemoveEventFilter_QObject(obj QObject) {
	if p.Pointer() == nil {
	} else {
		var objPtr C.QtObjectPtr = nil
		if obj != nil {
			objPtr = obj.Pointer()
		}
		C.QObject_RemoveEventFilter_QObject(p.Pointer(), objPtr)
	}
}

func (p *qobject) SetObjectName_String(name string) {
	if p.Pointer() != nil {
		C.QObject_SetObjectName_String(p.Pointer(), C.CString(name))
	}
}

func (p *qobject) SetParent_QObject(parent QObject) {
	if p.Pointer() == nil {
	} else {
		var parentPtr C.QtObjectPtr = nil
		if parent != nil {
			parentPtr = parent.Pointer()
		}
		C.QObject_SetParent_QObject(p.Pointer(), parentPtr)
	}
}

func (p *qobject) SignalsBlocked() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QObject_SignalsBlocked(p.Pointer()) != 0
	}
}

func (p *qobject) StartTimer_Int_TimerType(interval int, timerType TimerType) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QObject_StartTimer_Int_TimerType(p.Pointer(), C.int(interval), C.int(timerType)))
	}
}

func (p *qobject) Thread() QThread {
	if p.Pointer() == nil {
		return nil
	} else {
		var qthread = new(qthread)
		qthread.SetPointer(C.QObject_Thread(p.Pointer()))
		if qthread.ObjectName() == "" {
			qthread.SetObjectName_String("QThread_" + randomIdentifier())
		}
		return qthread
	}
}

func (p *qobject) ConnectSlotDeleteLater() {
	C.QObject_ConnectSlotDeleteLater(p.Pointer())
}

func (p *qobject) DisconnectSlotDeleteLater() {
	C.QObject_DisconnectSlotDeleteLater(p.Pointer())
}

func (p *qobject) SlotDeleteLater() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QObject_DeleteLater(p.Pointer())
	}
}

func (p *qobject) ConnectSignalDestroyed(f func()) {
	C.QObject_ConnectSignalDestroyed(p.Pointer())
	connectSignal(p.ObjectName(), "destroyed", f)
}

func (p *qobject) DisconnectSignalDestroyed() {
	C.QObject_DisconnectSignalDestroyed(p.Pointer())
	disconnectSignal(p.ObjectName(), "destroyed")
}

func (p *qobject) SignalDestroyed() func() {
	return getSignal(p.ObjectName(), "destroyed")
}

func (p *qobject) ConnectSignalObjectNameChanged(f func()) {
	C.QObject_ConnectSignalObjectNameChanged(p.Pointer())
	connectSignal(p.ObjectName(), "objectNameChanged", f)
}

func (p *qobject) DisconnectSignalObjectNameChanged() {
	C.QObject_DisconnectSignalObjectNameChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "objectNameChanged")
}

func (p *qobject) SignalObjectNameChanged() func() {
	return getSignal(p.ObjectName(), "objectNameChanged")
}

func QObject_Disconnect_QObject_String_QObject_String(sender QObject, signal string, receiver QObject, method string) bool {
	var senderPtr C.QtObjectPtr = nil
	if sender != nil {
		senderPtr = sender.Pointer()
	}
	var receiverPtr C.QtObjectPtr = nil
	if receiver != nil {
		receiverPtr = receiver.Pointer()
	}
	return C.QObject_Disconnect_QObject_String_QObject_String(senderPtr, C.CString(signal), receiverPtr, C.CString(method)) != 0
}

func QObject_Tr_String_String_Int(sourceText string, disambiguation string, n int) string {
	return C.GoString(C.QObject_Tr_String_String_Int(C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}

//export callbackQObject
func callbackQObject(ptr C.QtObjectPtr, msg *C.char) {
	var qobject = new(qobject)
	qobject.SetPointer(ptr)
	getSignal(qobject.ObjectName(), C.GoString(msg))()
}
