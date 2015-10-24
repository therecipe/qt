package multimedia

//#include "qvideofilterrunnable.h"
import "C"
import (
	"unsafe"
)

type QVideoFilterRunnable struct {
	ptr unsafe.Pointer
}

type QVideoFilterRunnableITF interface {
	QVideoFilterRunnablePTR() *QVideoFilterRunnable
}

func (p *QVideoFilterRunnable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoFilterRunnable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoFilterRunnable(ptr QVideoFilterRunnableITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoFilterRunnablePTR().Pointer()
	}
	return nil
}

func QVideoFilterRunnableFromPointer(ptr unsafe.Pointer) *QVideoFilterRunnable {
	var n = new(QVideoFilterRunnable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoFilterRunnable) QVideoFilterRunnablePTR() *QVideoFilterRunnable {
	return ptr
}

//QVideoFilterRunnable::RunFlag
type QVideoFilterRunnable__RunFlag int

var (
	QVideoFilterRunnable__LastInChain = QVideoFilterRunnable__RunFlag(0x01)
)
