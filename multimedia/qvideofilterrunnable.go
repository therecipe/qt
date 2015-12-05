package multimedia

//#include "multimedia.h"
import "C"
import (
	"unsafe"
)

type QVideoFilterRunnable struct {
	ptr unsafe.Pointer
}

type QVideoFilterRunnable_ITF interface {
	QVideoFilterRunnable_PTR() *QVideoFilterRunnable
}

func (p *QVideoFilterRunnable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoFilterRunnable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoFilterRunnable(ptr QVideoFilterRunnable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoFilterRunnable_PTR().Pointer()
	}
	return nil
}

func NewQVideoFilterRunnableFromPointer(ptr unsafe.Pointer) *QVideoFilterRunnable {
	var n = new(QVideoFilterRunnable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoFilterRunnable) QVideoFilterRunnable_PTR() *QVideoFilterRunnable {
	return ptr
}

//QVideoFilterRunnable::RunFlag
type QVideoFilterRunnable__RunFlag int64

const (
	QVideoFilterRunnable__LastInChain = QVideoFilterRunnable__RunFlag(0x01)
)
