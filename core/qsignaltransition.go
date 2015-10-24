package core

//#include "qsignaltransition.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSignalTransition struct {
	QAbstractTransition
}

type QSignalTransitionITF interface {
	QAbstractTransitionITF
	QSignalTransitionPTR() *QSignalTransition
}

func PointerFromQSignalTransition(ptr QSignalTransitionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalTransitionPTR().Pointer()
	}
	return nil
}

func QSignalTransitionFromPointer(ptr unsafe.Pointer) *QSignalTransition {
	var n = new(QSignalTransition)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSignalTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSignalTransition) QSignalTransitionPTR() *QSignalTransition {
	return ptr
}

func NewQSignalTransition(sourceState QStateITF) *QSignalTransition {
	return QSignalTransitionFromPointer(unsafe.Pointer(C.QSignalTransition_NewQSignalTransition(C.QtObjectPtr(PointerFromQState(sourceState)))))
}

func NewQSignalTransition2(sender QObjectITF, signal string, sourceState QStateITF) *QSignalTransition {
	return QSignalTransitionFromPointer(unsafe.Pointer(C.QSignalTransition_NewQSignalTransition2(C.QtObjectPtr(PointerFromQObject(sender)), C.CString(signal), C.QtObjectPtr(PointerFromQState(sourceState)))))
}

func (ptr *QSignalTransition) SenderObject() *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QSignalTransition_SenderObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSignalTransition) ConnectSenderObjectChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_ConnectSenderObjectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "senderObjectChanged", f)
	}
}

func (ptr *QSignalTransition) DisconnectSenderObjectChanged() {
	if ptr.Pointer() != nil {
		C.QSignalTransition_DisconnectSenderObjectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "senderObjectChanged")
	}
}

//export callbackQSignalTransitionSenderObjectChanged
func callbackQSignalTransitionSenderObjectChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "senderObjectChanged").(func())()
}

func (ptr *QSignalTransition) SetSenderObject(sender QObjectITF) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_SetSenderObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(sender)))
	}
}

func (ptr *QSignalTransition) SetSignal(signal QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_SetSignal(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(signal)))
	}
}

func (ptr *QSignalTransition) ConnectSignalChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_ConnectSignalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "signalChanged", f)
	}
}

func (ptr *QSignalTransition) DisconnectSignalChanged() {
	if ptr.Pointer() != nil {
		C.QSignalTransition_DisconnectSignalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "signalChanged")
	}
}

//export callbackQSignalTransitionSignalChanged
func callbackQSignalTransitionSignalChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "signalChanged").(func())()
}

func (ptr *QSignalTransition) DestroyQSignalTransition() {
	if ptr.Pointer() != nil {
		C.QSignalTransition_DestroyQSignalTransition(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
