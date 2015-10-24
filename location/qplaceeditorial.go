package location

//#include "qplaceeditorial.h"
import "C"
import (
	"unsafe"
)

type QPlaceEditorial struct {
	QPlaceContent
}

type QPlaceEditorialITF interface {
	QPlaceContentITF
	QPlaceEditorialPTR() *QPlaceEditorial
}

func PointerFromQPlaceEditorial(ptr QPlaceEditorialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceEditorialPTR().Pointer()
	}
	return nil
}

func QPlaceEditorialFromPointer(ptr unsafe.Pointer) *QPlaceEditorial {
	var n = new(QPlaceEditorial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceEditorial) QPlaceEditorialPTR() *QPlaceEditorial {
	return ptr
}
