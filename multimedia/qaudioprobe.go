package multimedia

//#include "qaudioprobe.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioProbe struct {
	core.QObject
}

type QAudioProbeITF interface {
	core.QObjectITF
	QAudioProbePTR() *QAudioProbe
}

func PointerFromQAudioProbe(ptr QAudioProbeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioProbePTR().Pointer()
	}
	return nil
}

func QAudioProbeFromPointer(ptr unsafe.Pointer) *QAudioProbe {
	var n = new(QAudioProbe)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioProbe_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioProbe) QAudioProbePTR() *QAudioProbe {
	return ptr
}

func NewQAudioProbe(parent core.QObjectITF) *QAudioProbe {
	return QAudioProbeFromPointer(unsafe.Pointer(C.QAudioProbe_NewQAudioProbe(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAudioProbe) ConnectFlush(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioProbe_ConnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QAudioProbe) DisconnectFlush() {
	if ptr.Pointer() != nil {
		C.QAudioProbe_DisconnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQAudioProbeFlush
func callbackQAudioProbeFlush(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QAudioProbe) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QAudioProbe_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource(source QMediaObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaObject(source))) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource2(mediaRecorder QMediaRecorderITF) bool {
	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaRecorder(mediaRecorder))) != 0
	}
	return false
}

func (ptr *QAudioProbe) DestroyQAudioProbe() {
	if ptr.Pointer() != nil {
		C.QAudioProbe_DestroyQAudioProbe(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
