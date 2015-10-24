package xml

//#include "qdomnamednodemap.h"
import "C"
import (
	"unsafe"
)

type QDomNamedNodeMap struct {
	ptr unsafe.Pointer
}

type QDomNamedNodeMapITF interface {
	QDomNamedNodeMapPTR() *QDomNamedNodeMap
}

func (p *QDomNamedNodeMap) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomNamedNodeMap) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomNamedNodeMap(ptr QDomNamedNodeMapITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNamedNodeMapPTR().Pointer()
	}
	return nil
}

func QDomNamedNodeMapFromPointer(ptr unsafe.Pointer) *QDomNamedNodeMap {
	var n = new(QDomNamedNodeMap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomNamedNodeMap) QDomNamedNodeMapPTR() *QDomNamedNodeMap {
	return ptr
}

func NewQDomNamedNodeMap() *QDomNamedNodeMap {
	return QDomNamedNodeMapFromPointer(unsafe.Pointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap()))
}

func NewQDomNamedNodeMap2(n QDomNamedNodeMapITF) *QDomNamedNodeMap {
	return QDomNamedNodeMapFromPointer(unsafe.Pointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap2(C.QtObjectPtr(PointerFromQDomNamedNodeMap(n)))))
}

func (ptr *QDomNamedNodeMap) Contains(name string) bool {
	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_Contains(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) DestroyQDomNamedNodeMap() {
	if ptr.Pointer() != nil {
		C.QDomNamedNodeMap_DestroyQDomNamedNodeMap(C.QtObjectPtr(ptr.Pointer()))
	}
}
