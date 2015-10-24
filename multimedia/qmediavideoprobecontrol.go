package multimedia

//#include "qmediavideoprobecontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaVideoProbeControl struct {
	QMediaControl
}

type QMediaVideoProbeControlITF interface {
	QMediaControlITF
	QMediaVideoProbeControlPTR() *QMediaVideoProbeControl
}

func PointerFromQMediaVideoProbeControl(ptr QMediaVideoProbeControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaVideoProbeControlPTR().Pointer()
	}
	return nil
}

func QMediaVideoProbeControlFromPointer(ptr unsafe.Pointer) *QMediaVideoProbeControl {
	var n = new(QMediaVideoProbeControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaVideoProbeControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaVideoProbeControl) QMediaVideoProbeControlPTR() *QMediaVideoProbeControl {
	return ptr
}

func (ptr *QMediaVideoProbeControl) ConnectFlush(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ConnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectFlush() {
	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DisconnectFlush(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaVideoProbeControlFlush
func callbackQMediaVideoProbeControlFlush(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QMediaVideoProbeControl) DestroyQMediaVideoProbeControl() {
	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
