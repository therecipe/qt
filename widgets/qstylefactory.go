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

type QStyleFactoryITF interface {
	QStyleFactoryPTR() *QStyleFactory
}

func (p *QStyleFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStyleFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStyleFactory(ptr QStyleFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleFactoryPTR().Pointer()
	}
	return nil
}

func QStyleFactoryFromPointer(ptr unsafe.Pointer) *QStyleFactory {
	var n = new(QStyleFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleFactory) QStyleFactoryPTR() *QStyleFactory {
	return ptr
}

func QStyleFactory_Create(key string) *QStyle {
	return QStyleFromPointer(unsafe.Pointer(C.QStyleFactory_QStyleFactory_Create(C.CString(key))))
}

func QStyleFactory_Keys() []string {
	return strings.Split(C.GoString(C.QStyleFactory_QStyleFactory_Keys()), "|")
}
