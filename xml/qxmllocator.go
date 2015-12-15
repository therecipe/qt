package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlLocator struct {
	ptr unsafe.Pointer
}

type QXmlLocator_ITF interface {
	QXmlLocator_PTR() *QXmlLocator
}

func (p *QXmlLocator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLocator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLocator(ptr QXmlLocator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLocator_PTR().Pointer()
	}
	return nil
}

func NewQXmlLocatorFromPointer(ptr unsafe.Pointer) *QXmlLocator {
	var n = new(QXmlLocator)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlLocator_") {
		n.SetObjectNameAbs("QXmlLocator_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlLocator) QXmlLocator_PTR() *QXmlLocator {
	return ptr
}

func (ptr *QXmlLocator) ColumnNumber() int {
	defer qt.Recovering("QXmlLocator::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlLocator) LineNumber() int {
	defer qt.Recovering("QXmlLocator::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlLocator) DestroyQXmlLocator() {
	defer qt.Recovering("QXmlLocator::~QXmlLocator")

	if ptr.Pointer() != nil {
		C.QXmlLocator_DestroyQXmlLocator(ptr.Pointer())
	}
}

func (ptr *QXmlLocator) ObjectNameAbs() string {
	defer qt.Recovering("QXmlLocator::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLocator_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLocator) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlLocator::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlLocator_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
