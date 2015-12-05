package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QSocketNotifier struct {
	QObject
}

type QSocketNotifier_ITF interface {
	QObject_ITF
	QSocketNotifier_PTR() *QSocketNotifier
}

func PointerFromQSocketNotifier(ptr QSocketNotifier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSocketNotifier_PTR().Pointer()
	}
	return nil
}

func NewQSocketNotifierFromPointer(ptr unsafe.Pointer) *QSocketNotifier {
	var n = new(QSocketNotifier)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSocketNotifier_") {
		n.SetObjectName("QSocketNotifier_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSocketNotifier) QSocketNotifier_PTR() *QSocketNotifier {
	return ptr
}

//QSocketNotifier::Type
type QSocketNotifier__Type int64

const (
	QSocketNotifier__Read      = QSocketNotifier__Type(0)
	QSocketNotifier__Write     = QSocketNotifier__Type(1)
	QSocketNotifier__Exception = QSocketNotifier__Type(2)
)

func (ptr *QSocketNotifier) ConnectActivated(f func(socket int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSocketNotifier::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSocketNotifier_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QSocketNotifier) DisconnectActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSocketNotifier::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSocketNotifier_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQSocketNotifierActivated
func callbackQSocketNotifierActivated(ptrName *C.char, socket C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSocketNotifier::activated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activated").(func(int))(int(socket))
}

func (ptr *QSocketNotifier) IsEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSocketNotifier::isEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSocketNotifier_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSocketNotifier) SetEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSocketNotifier::setEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSocketNotifier_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QSocketNotifier) Type() QSocketNotifier__Type {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSocketNotifier::type")
		}
	}()

	if ptr.Pointer() != nil {
		return QSocketNotifier__Type(C.QSocketNotifier_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSocketNotifier) DestroyQSocketNotifier() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSocketNotifier::~QSocketNotifier")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSocketNotifier_DestroyQSocketNotifier(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
