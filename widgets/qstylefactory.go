package widgets

//#include "qstylefactory.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QStyleFactory struct {
	ptr unsafe.Pointer
}

type QStyleFactory_ITF interface {
	QStyleFactory_PTR() *QStyleFactory
}

func (p *QStyleFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStyleFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStyleFactory(ptr QStyleFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleFactory_PTR().Pointer()
	}
	return nil
}

func NewQStyleFactoryFromPointer(ptr unsafe.Pointer) *QStyleFactory {
	var n = new(QStyleFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleFactory) QStyleFactory_PTR() *QStyleFactory {
	return ptr
}

func QStyleFactory_Create(key string) *QStyle {
	return NewQStyleFromPointer(C.QStyleFactory_QStyleFactory_Create(C.CString(key)))
}

func QStyleFactory_Keys() []string {
	return strings.Split(C.GoString(C.QStyleFactory_QStyleFactory_Keys()), "|")
}
