package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDomNamedNodeMap struct {
	ptr unsafe.Pointer
}

type QDomNamedNodeMap_ITF interface {
	QDomNamedNodeMap_PTR() *QDomNamedNodeMap
}

func (p *QDomNamedNodeMap) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomNamedNodeMap) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomNamedNodeMap(ptr QDomNamedNodeMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNamedNodeMap_PTR().Pointer()
	}
	return nil
}

func NewQDomNamedNodeMapFromPointer(ptr unsafe.Pointer) *QDomNamedNodeMap {
	var n = new(QDomNamedNodeMap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomNamedNodeMap) QDomNamedNodeMap_PTR() *QDomNamedNodeMap {
	return ptr
}

func NewQDomNamedNodeMap() *QDomNamedNodeMap {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::QDomNamedNodeMap")
		}
	}()

	return NewQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap())
}

func NewQDomNamedNodeMap2(n QDomNamedNodeMap_ITF) *QDomNamedNodeMap {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::QDomNamedNodeMap")
		}
	}()

	return NewQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap2(PointerFromQDomNamedNodeMap(n)))
}

func (ptr *QDomNamedNodeMap) Contains(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_Contains(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Length() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::length")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) Size() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::size")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) DestroyQDomNamedNodeMap() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomNamedNodeMap::~QDomNamedNodeMap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomNamedNodeMap_DestroyQDomNamedNodeMap(ptr.Pointer())
	}
}
