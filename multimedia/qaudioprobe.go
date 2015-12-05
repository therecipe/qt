package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAudioProbe struct {
	core.QObject
}

type QAudioProbe_ITF interface {
	core.QObject_ITF
	QAudioProbe_PTR() *QAudioProbe
}

func PointerFromQAudioProbe(ptr QAudioProbe_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioProbe_PTR().Pointer()
	}
	return nil
}

func NewQAudioProbeFromPointer(ptr unsafe.Pointer) *QAudioProbe {
	var n = new(QAudioProbe)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioProbe_") {
		n.SetObjectName("QAudioProbe_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioProbe) QAudioProbe_PTR() *QAudioProbe {
	return ptr
}

func NewQAudioProbe(parent core.QObject_ITF) *QAudioProbe {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::QAudioProbe")
		}
	}()

	return NewQAudioProbeFromPointer(C.QAudioProbe_NewQAudioProbe(core.PointerFromQObject(parent)))
}

func (ptr *QAudioProbe) ConnectFlush(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioProbe_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QAudioProbe) DisconnectFlush() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioProbe_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQAudioProbeFlush
func callbackQAudioProbeFlush(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::flush")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QAudioProbe) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAudioProbe_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource(source QMediaObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::setSource")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource(ptr.Pointer(), PointerFromQMediaObject(source)) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource2(mediaRecorder QMediaRecorder_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::setSource")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource2(ptr.Pointer(), PointerFromQMediaRecorder(mediaRecorder)) != 0
	}
	return false
}

func (ptr *QAudioProbe) DestroyQAudioProbe() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioProbe::~QAudioProbe")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioProbe_DestroyQAudioProbe(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
