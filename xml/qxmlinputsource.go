package xml

//#include "qxmlinputsource.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QXmlInputSource struct {
	ptr unsafe.Pointer
}

type QXmlInputSource_ITF interface {
	QXmlInputSource_PTR() *QXmlInputSource
}

func (p *QXmlInputSource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlInputSource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlInputSource(ptr QXmlInputSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlInputSource_PTR().Pointer()
	}
	return nil
}

func NewQXmlInputSourceFromPointer(ptr unsafe.Pointer) *QXmlInputSource {
	var n = new(QXmlInputSource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlInputSource) QXmlInputSource_PTR() *QXmlInputSource {
	return ptr
}

func NewQXmlInputSource() *QXmlInputSource {
	return NewQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource())
}

func NewQXmlInputSource2(dev core.QIODevice_ITF) *QXmlInputSource {
	return NewQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource2(core.PointerFromQIODevice(dev)))
}

func (ptr *QXmlInputSource) Data() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) FetchData() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchData(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) Reset() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) SetData2(dat core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData2(ptr.Pointer(), core.PointerFromQByteArray(dat))
	}
}

func (ptr *QXmlInputSource) SetData(dat string) {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) DestroyQXmlInputSource() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_DestroyQXmlInputSource(ptr.Pointer())
	}
}
