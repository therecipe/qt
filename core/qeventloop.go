package core

//#include "qeventloop.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QEventLoop struct {
	QObject
}

type QEventLoop_ITF interface {
	QObject_ITF
	QEventLoop_PTR() *QEventLoop
}

func PointerFromQEventLoop(ptr QEventLoop_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventLoop_PTR().Pointer()
	}
	return nil
}

func NewQEventLoopFromPointer(ptr unsafe.Pointer) *QEventLoop {
	var n = new(QEventLoop)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QEventLoop_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QEventLoop) QEventLoop_PTR() *QEventLoop {
	return ptr
}

//QEventLoop::ProcessEventsFlag
type QEventLoop__ProcessEventsFlag int64

const (
	QEventLoop__AllEvents              = QEventLoop__ProcessEventsFlag(0x00)
	QEventLoop__ExcludeUserInputEvents = QEventLoop__ProcessEventsFlag(0x01)
	QEventLoop__ExcludeSocketNotifiers = QEventLoop__ProcessEventsFlag(0x02)
	QEventLoop__WaitForMoreEvents      = QEventLoop__ProcessEventsFlag(0x04)
	QEventLoop__X11ExcludeTimers       = QEventLoop__ProcessEventsFlag(0x08)
	QEventLoop__EventLoopExec          = QEventLoop__ProcessEventsFlag(0x20)
	QEventLoop__DialogExec             = QEventLoop__ProcessEventsFlag(0x40)
)

func NewQEventLoop(parent QObject_ITF) *QEventLoop {
	return NewQEventLoopFromPointer(C.QEventLoop_NewQEventLoop(PointerFromQObject(parent)))
}

func (ptr *QEventLoop) Event(event QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QEventLoop_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QEventLoop) Exec(flags QEventLoop__ProcessEventsFlag) int {
	if ptr.Pointer() != nil {
		return int(C.QEventLoop_Exec(ptr.Pointer(), C.int(flags)))
	}
	return 0
}

func (ptr *QEventLoop) Exit(returnCode int) {
	if ptr.Pointer() != nil {
		C.QEventLoop_Exit(ptr.Pointer(), C.int(returnCode))
	}
}

func (ptr *QEventLoop) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QEventLoop_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QEventLoop) ProcessEvents(flags QEventLoop__ProcessEventsFlag) bool {
	if ptr.Pointer() != nil {
		return C.QEventLoop_ProcessEvents(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QEventLoop) ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxTime int) {
	if ptr.Pointer() != nil {
		C.QEventLoop_ProcessEvents2(ptr.Pointer(), C.int(flags), C.int(maxTime))
	}
}

func (ptr *QEventLoop) Quit() {
	if ptr.Pointer() != nil {
		C.QEventLoop_Quit(ptr.Pointer())
	}
}

func (ptr *QEventLoop) WakeUp() {
	if ptr.Pointer() != nil {
		C.QEventLoop_WakeUp(ptr.Pointer())
	}
}

func (ptr *QEventLoop) DestroyQEventLoop() {
	if ptr.Pointer() != nil {
		C.QEventLoop_DestroyQEventLoop(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
