package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QMediaAudioProbeControl struct {
	QMediaControl
}

type QMediaAudioProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl
}

func PointerFromQMediaAudioProbeControl(ptr QMediaAudioProbeControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAudioProbeControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaAudioProbeControlFromPointer(ptr unsafe.Pointer) *QMediaAudioProbeControl {
	var n = new(QMediaAudioProbeControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaAudioProbeControl_") {
		n.SetObjectName("QMediaAudioProbeControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaAudioProbeControl) QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl {
	return ptr
}

func (ptr *QMediaAudioProbeControl) ConnectFlush(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaAudioProbeControl::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectFlush() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaAudioProbeControl::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaAudioProbeControlFlush
func callbackQMediaAudioProbeControlFlush(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaAudioProbeControl::flush")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QMediaAudioProbeControl) DestroyQMediaAudioProbeControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaAudioProbeControl::~QMediaAudioProbeControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
