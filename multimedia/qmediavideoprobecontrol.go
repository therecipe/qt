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
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaVideoProbeControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaVideoProbeControl) QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl {
	return ptr
}

func (ptr *QMediaVideoProbeControl) ConnectFlush(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectFlush() {
	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaVideoProbeControlFlush
func callbackQMediaVideoProbeControlFlush(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "flush").(func())()
}

func (ptr *QMediaVideoProbeControl) DestroyQMediaVideoProbeControl() {
	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
