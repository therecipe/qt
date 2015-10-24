package location

//#include "qplacecontent.h"
import "C"
import (
	"unsafe"
)

type QPlaceContent struct {
	ptr unsafe.Pointer
}

type QPlaceContentITF interface {
	QPlaceContentPTR() *QPlaceContent
}

func (p *QPlaceContent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceContent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceContent(ptr QPlaceContentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentPTR().Pointer()
	}
	return nil
}

func QPlaceContentFromPointer(ptr unsafe.Pointer) *QPlaceContent {
	var n = new(QPlaceContent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceContent) QPlaceContentPTR() *QPlaceContent {
	return ptr
}

//QPlaceContent::Type
type QPlaceContent__Type int

var (
	QPlaceContent__NoType        = QPlaceContent__Type(0)
	QPlaceContent__ImageType     = QPlaceContent__Type(1)
	QPlaceContent__ReviewType    = QPlaceContent__Type(2)
	QPlaceContent__EditorialType = QPlaceContent__Type(3)
)
