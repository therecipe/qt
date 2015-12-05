package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceContent struct {
	ptr unsafe.Pointer
}

type QPlaceContent_ITF interface {
	QPlaceContent_PTR() *QPlaceContent
}

func (p *QPlaceContent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceContent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceContent(ptr QPlaceContent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContent_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContentFromPointer(ptr unsafe.Pointer) *QPlaceContent {
	var n = new(QPlaceContent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceContent) QPlaceContent_PTR() *QPlaceContent {
	return ptr
}

//QPlaceContent::Type
type QPlaceContent__Type int64

const (
	QPlaceContent__NoType        = QPlaceContent__Type(0)
	QPlaceContent__ImageType     = QPlaceContent__Type(1)
	QPlaceContent__ReviewType    = QPlaceContent__Type(2)
	QPlaceContent__EditorialType = QPlaceContent__Type(3)
)
