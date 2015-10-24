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

type QMediaAudioProbeControlITF interface {
	QMediaControlITF
	QMediaAudioProbeControlPTR() *QMediaAudioProbeControl
}

func PointerFromQMediaAudioProbeControl(ptr QMediaAudioProbeControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAudioProbeControlPTR().Pointer()
	}
	return nil
}

func QMediaAudioProbeControlFromPointer(ptr unsafe.Pointer) *QMediaAudioProbeControl {
	var n = new(QMediaAudioProbeControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaAudioProbeControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaAudioProbeControl) QMediaAudioProbeControlPTR() *QMediaAudioProbeControl {
	return ptr
}

func (ptr *QMediaAudioProbeControl) ConnectFlush(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ConnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectFlush() {
	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DisconnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaAudioProbeControlFlush
func callbackQMediaAudioProbeControlFlush(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QMediaAudioProbeControl) DestroyQMediaAudioProbeControl() {
	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
