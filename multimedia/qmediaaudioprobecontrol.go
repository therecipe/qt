package multimedia

//#include "qmediaaudioprobecontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMediaAudioProbeControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaAudioProbeControl) QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl {
	return ptr
}

func (ptr *QMediaAudioProbeControl) ConnectFlush(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectFlush() {
	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaAudioProbeControlFlush
func callbackQMediaAudioProbeControlFlush(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QMediaAudioProbeControl) DestroyQMediaAudioProbeControl() {
	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
