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

type QXmlInputSourceITF interface {
	QXmlInputSourcePTR() *QXmlInputSource
}

func (p *QXmlInputSource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlInputSource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlInputSource(ptr QXmlInputSourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlInputSourcePTR().Pointer()
	}
	return nil
}

func QXmlInputSourceFromPointer(ptr unsafe.Pointer) *QXmlInputSource {
	var n = new(QXmlInputSource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlInputSource) QXmlInputSourcePTR() *QXmlInputSource {
	return ptr
}

func NewQXmlInputSource() *QXmlInputSource {
	return QXmlInputSourceFromPointer(unsafe.Pointer(C.QXmlInputSource_NewQXmlInputSource()))
}

func NewQXmlInputSource2(dev core.QIODeviceITF) *QXmlInputSource {
	return QXmlInputSourceFromPointer(unsafe.Pointer(C.QXmlInputSource_NewQXmlInputSource2(C.QtObjectPtr(core.PointerFromQIODevice(dev)))))
}

func (ptr *QXmlInputSource) Data() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_Data(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlInputSource) FetchData() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchData(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlInputSource) Reset() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlInputSource) SetData2(dat core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(dat)))
	}
}

func (ptr *QXmlInputSource) SetData(dat string) {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData(C.QtObjectPtr(ptr.Pointer()), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) DestroyQXmlInputSource() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_DestroyQXmlInputSource(C.QtObjectPtr(ptr.Pointer()))
	}
}
