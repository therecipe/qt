package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QMediaVideoProbeControl struct {
	QMediaControl
}

type QMediaVideoProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl
}

func PointerFromQMediaVideoProbeControl(ptr QMediaVideoProbeControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaVideoProbeControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaVideoProbeControlFromPointer(ptr unsafe.Pointer) *QMediaVideoProbeControl {
	var n = new(QMediaVideoProbeControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaVideoProbeControl_") {
		n.SetObjectName("QMediaVideoProbeControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaVideoProbeControl) QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl {
	return ptr
}

func (ptr *QMediaVideoProbeControl) ConnectFlush(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaVideoProbeControl::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectFlush() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaVideoProbeControl::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaVideoProbeControlFlush
func callbackQMediaVideoProbeControlFlush(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaVideoProbeControl::flush")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QMediaVideoProbeControl) DestroyQMediaVideoProbeControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaVideoProbeControl::~QMediaVideoProbeControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
