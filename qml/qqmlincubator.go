package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlIncubator struct {
	ptr unsafe.Pointer
}

type QQmlIncubator_ITF interface {
	QQmlIncubator_PTR() *QQmlIncubator
}

func (p *QQmlIncubator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlIncubator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlIncubator(ptr QQmlIncubator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubator_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubatorFromPointer(ptr unsafe.Pointer) *QQmlIncubator {
	var n = new(QQmlIncubator)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQmlIncubator_") {
		n.SetObjectNameAbs("QQmlIncubator_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlIncubator) QQmlIncubator_PTR() *QQmlIncubator {
	return ptr
}

//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int64

const (
	QQmlIncubator__Asynchronous         = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          = QQmlIncubator__IncubationMode(2)
)

//QQmlIncubator::Status
type QQmlIncubator__Status int64

const (
	QQmlIncubator__Null    = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   = QQmlIncubator__Status(1)
	QQmlIncubator__Loading = QQmlIncubator__Status(2)
	QQmlIncubator__Error   = QQmlIncubator__Status(3)
)

func NewQQmlIncubator(mode QQmlIncubator__IncubationMode) *QQmlIncubator {
	defer qt.Recovering("QQmlIncubator::QQmlIncubator")

	return NewQQmlIncubatorFromPointer(C.QQmlIncubator_NewQQmlIncubator(C.int(mode)))
}

func (ptr *QQmlIncubator) Clear() {
	defer qt.Recovering("QQmlIncubator::clear")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_Clear(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) ForceCompletion() {
	defer qt.Recovering("QQmlIncubator::forceCompletion")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_ForceCompletion(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {
	defer qt.Recovering("QQmlIncubator::incubationMode")

	if ptr.Pointer() != nil {
		return QQmlIncubator__IncubationMode(C.QQmlIncubator_IncubationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) IsError() bool {
	defer qt.Recovering("QQmlIncubator::isError")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsLoading() bool {
	defer qt.Recovering("QQmlIncubator::isLoading")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsNull() bool {
	defer qt.Recovering("QQmlIncubator::isNull")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsReady() bool {
	defer qt.Recovering("QQmlIncubator::isReady")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) Object() *core.QObject {
	defer qt.Recovering("QQmlIncubator::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlIncubator_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlIncubator) ConnectSetInitialState(f func(object *core.QObject)) {
	defer qt.Recovering("connect QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setInitialState", f)
	}
}

func (ptr *QQmlIncubator) DisconnectSetInitialState() {
	defer qt.Recovering("disconnect QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setInitialState")
	}
}

//export callbackQQmlIncubatorSetInitialState
func callbackQQmlIncubatorSetInitialState(ptrName *C.char, object unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlIncubator::setInitialState")

	var signal = qt.GetSignal(C.GoString(ptrName), "setInitialState")
	if signal != nil {
		defer signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
		return true
	}
	return false

}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {
	defer qt.Recovering("QQmlIncubator::status")

	if ptr.Pointer() != nil {
		return QQmlIncubator__Status(C.QQmlIncubator_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) ConnectStatusChanged(f func(status QQmlIncubator__Status)) {
	defer qt.Recovering("connect QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "statusChanged", f)
	}
}

func (ptr *QQmlIncubator) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "statusChanged")
	}
}

//export callbackQQmlIncubatorStatusChanged
func callbackQQmlIncubatorStatusChanged(ptrName *C.char, status C.int) bool {
	defer qt.Recovering("callback QQmlIncubator::statusChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "statusChanged")
	if signal != nil {
		defer signal.(func(QQmlIncubator__Status))(QQmlIncubator__Status(status))
		return true
	}
	return false

}

func (ptr *QQmlIncubator) ObjectNameAbs() string {
	defer qt.Recovering("QQmlIncubator::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlIncubator_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlIncubator) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlIncubator::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
