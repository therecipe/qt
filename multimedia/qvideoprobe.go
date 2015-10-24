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

type QVideoProbeITF interface {
	core.QObjectITF
	QVideoProbePTR() *QVideoProbe
}

func PointerFromQVideoProbe(ptr QVideoProbeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoProbePTR().Pointer()
	}
	return nil
}

func QVideoProbeFromPointer(ptr unsafe.Pointer) *QVideoProbe {
	var n = new(QVideoProbe)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QVideoProbe_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoProbe) QVideoProbePTR() *QVideoProbe {
	return ptr
}

func NewQVideoProbe(parent core.QObjectITF) *QVideoProbe {
	return QVideoProbeFromPointer(unsafe.Pointer(C.QVideoProbe_NewQVideoProbe(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QVideoProbe) ConnectFlush(f func()) {
	if ptr.Pointer() != nil {
		C.QVideoProbe_ConnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QVideoProbe) DisconnectFlush() {
	if ptr.Pointer() != nil {
		C.QVideoProbe_DisconnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQVideoProbeFlush
func callbackQVideoProbeFlush(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QVideoProbe) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QVideoProbe_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource(source QMediaObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaObject(source))) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource2(mediaRecorder QMediaRecorderITF) bool {
	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaRecorder(mediaRecorder))) != 0
	}
	return false
}

func (ptr *QVideoProbe) DestroyQVideoProbe() {
	if ptr.Pointer() != nil {
		C.QVideoProbe_DestroyQVideoProbe(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
