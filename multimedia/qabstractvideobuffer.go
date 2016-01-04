package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractVideoBuffer struct {
	ptr unsafe.Pointer
}

type QAbstractVideoBuffer_ITF interface {
	QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer
}

func (p *QAbstractVideoBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractVideoBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractVideoBuffer(ptr QAbstractVideoBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoBuffer_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractVideoBuffer {
	var n = new(QAbstractVideoBuffer)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractVideoBuffer_") {
		n.SetObjectNameAbs("QAbstractVideoBuffer_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractVideoBuffer) QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer {
	return ptr
}

//QAbstractVideoBuffer::HandleType
type QAbstractVideoBuffer__HandleType int64

const (
	QAbstractVideoBuffer__NoHandle         = QAbstractVideoBuffer__HandleType(0)
	QAbstractVideoBuffer__GLTextureHandle  = QAbstractVideoBuffer__HandleType(1)
	QAbstractVideoBuffer__XvShmImageHandle = QAbstractVideoBuffer__HandleType(2)
	QAbstractVideoBuffer__CoreImageHandle  = QAbstractVideoBuffer__HandleType(3)
	QAbstractVideoBuffer__QPixmapHandle    = QAbstractVideoBuffer__HandleType(4)
	QAbstractVideoBuffer__EGLImageHandle   = QAbstractVideoBuffer__HandleType(5)
	QAbstractVideoBuffer__UserHandle       = QAbstractVideoBuffer__HandleType(1000)
)

//QAbstractVideoBuffer::MapMode
type QAbstractVideoBuffer__MapMode int64

const (
	QAbstractVideoBuffer__NotMapped = QAbstractVideoBuffer__MapMode(0x00)
	QAbstractVideoBuffer__ReadOnly  = QAbstractVideoBuffer__MapMode(0x01)
	QAbstractVideoBuffer__WriteOnly = QAbstractVideoBuffer__MapMode(0x02)
	QAbstractVideoBuffer__ReadWrite = QAbstractVideoBuffer__MapMode(QAbstractVideoBuffer__ReadOnly | QAbstractVideoBuffer__WriteOnly)
)

func (ptr *QAbstractVideoBuffer) Handle() *core.QVariant {
	defer qt.Recovering("QAbstractVideoBuffer::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractVideoBuffer_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoBuffer) HandleType() QAbstractVideoBuffer__HandleType {
	defer qt.Recovering("QAbstractVideoBuffer::handleType")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QAbstractVideoBuffer_HandleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) MapMode() QAbstractVideoBuffer__MapMode {
	defer qt.Recovering("QAbstractVideoBuffer::mapMode")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__MapMode(C.QAbstractVideoBuffer_MapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) ConnectRelease(f func()) {
	defer qt.Recovering("connect QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "release", f)
	}
}

func (ptr *QAbstractVideoBuffer) DisconnectRelease() {
	defer qt.Recovering("disconnect QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "release")
	}
}

//export callbackQAbstractVideoBufferRelease
func callbackQAbstractVideoBufferRelease(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoBuffer::release")

	if signal := qt.GetSignal(C.GoString(ptrName), "release"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractVideoBufferFromPointer(ptr).ReleaseDefault()
	}
}

func (ptr *QAbstractVideoBuffer) Release() {
	defer qt.Recovering("QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Release(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) ReleaseDefault() {
	defer qt.Recovering("QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_ReleaseDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) Unmap() {
	defer qt.Recovering("QAbstractVideoBuffer::unmap")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Unmap(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) DestroyQAbstractVideoBuffer() {
	defer qt.Recovering("QAbstractVideoBuffer::~QAbstractVideoBuffer")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractVideoBuffer::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractVideoBuffer_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractVideoBuffer) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractVideoBuffer::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
