package multimedia

//#include "qvideoprobe.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVideoProbe struct {
	core.QObject
}

type QVideoProbe_ITF interface {
	core.QObject_ITF
	QVideoProbe_PTR() *QVideoProbe
}

func PointerFromQVideoProbe(ptr QVideoProbe_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoProbe_PTR().Pointer()
	}
	return nil
}

func NewQVideoProbeFromPointer(ptr unsafe.Pointer) *QVideoProbe {
	var n = new(QVideoProbe)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QVideoProbe_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoProbe) QVideoProbe_PTR() *QVideoProbe {
	return ptr
}

func NewQVideoProbe(parent core.QObject_ITF) *QVideoProbe {
	return NewQVideoProbeFromPointer(C.QVideoProbe_NewQVideoProbe(core.PointerFromQObject(parent)))
}

func (ptr *QVideoProbe) ConnectFlush(f func()) {
	if ptr.Pointer() != nil {
		C.QVideoProbe_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QVideoProbe) DisconnectFlush() {
	if ptr.Pointer() != nil {
		C.QVideoProbe_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQVideoProbeFlush
func callbackQVideoProbeFlush(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QVideoProbe) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QVideoProbe_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource(source QMediaObject_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource(ptr.Pointer(), PointerFromQMediaObject(source)) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource2(mediaRecorder QMediaRecorder_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource2(ptr.Pointer(), PointerFromQMediaRecorder(mediaRecorder)) != 0
	}
	return false
}

func (ptr *QVideoProbe) DestroyQVideoProbe() {
	if ptr.Pointer() != nil {
		C.QVideoProbe_DestroyQVideoProbe(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
