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

type QSignalTransition_ITF interface {
	QAbstractTransition_ITF
	QSignalTransition_PTR() *QSignalTransition
}

func PointerFromQSignalTransition(ptr QSignalTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalTransition_PTR().Pointer()
	}
	return nil
}

func NewQSignalTransitionFromPointer(ptr unsafe.Pointer) *QSignalTransition {
	var n = new(QSignalTransition)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSignalTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSignalTransition) QSignalTransition_PTR() *QSignalTransition {
	return ptr
}

func NewQSignalTransition(sourceState QState_ITF) *QSignalTransition {
	return NewQSignalTransitionFromPointer(C.QSignalTransition_NewQSignalTransition(PointerFromQState(sourceState)))
}

func NewQSignalTransition2(sender QObject_ITF, signal string, sourceState QState_ITF) *QSignalTransition {
	return NewQSignalTransitionFromPointer(C.QSignalTransition_NewQSignalTransition2(PointerFromQObject(sender), C.CString(signal), PointerFromQState(sourceState)))
}

func (ptr *QSignalTransition) SenderObject() *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalTransition_SenderObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSignalTransition) ConnectSenderObjectChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_ConnectSenderObjectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "senderObjectChanged", f)
	}
}

func (ptr *QSignalTransition) DisconnectSenderObjectChanged() {
	if ptr.Pointer() != nil {
		C.QSignalTransition_DisconnectSenderObjectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "senderObjectChanged")
	}
}

//export callbackQSignalTransitionSenderObjectChanged
func callbackQSignalTransitionSenderObjectChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "senderObjectChanged").(func())()
}

func (ptr *QSignalTransition) SetSenderObject(sender QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_SetSenderObject(ptr.Pointer(), PointerFromQObject(sender))
	}
}

func (ptr *QSignalTransition) SetSignal(signal QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_SetSignal(ptr.Pointer(), PointerFromQByteArray(signal))
	}
}

func (ptr *QSignalTransition) Signal() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QSignalTransition_Signal(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSignalTransition) ConnectSignalChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSignalTransition_ConnectSignalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalChanged", f)
	}
}

func (ptr *QSignalTransition) DisconnectSignalChanged() {
	if ptr.Pointer() != nil {
		C.QSignalTransition_DisconnectSignalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalChanged")
	}
}

//export callbackQSignalTransitionSignalChanged
func callbackQSignalTransitionSignalChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "signalChanged").(func())()
}

func (ptr *QSignalTransition) DestroyQSignalTransition() {
	if ptr.Pointer() != nil {
		C.QSignalTransition_DestroyQSignalTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
