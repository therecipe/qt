package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::QQmlIncubator")
		}
	}()

	return NewQQmlIncubatorFromPointer(C.QQmlIncubator_NewQQmlIncubator(C.int(mode)))
}

func (ptr *QQmlIncubator) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlIncubator_Clear(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) ForceCompletion() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::forceCompletion")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlIncubator_ForceCompletion(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::incubationMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QQmlIncubator__IncubationMode(C.QQmlIncubator_IncubationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) IsError() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::isError")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsLoading() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::isLoading")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsReady() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::isReady")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) Object() *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::object")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlIncubator_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubator::status")
		}
	}()

	if ptr.Pointer() != nil {
		return QQmlIncubator__Status(C.QQmlIncubator_Status(ptr.Pointer()))
	}
	return 0
}
