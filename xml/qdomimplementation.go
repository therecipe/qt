package xml

//#include "qdomimplementation.h"
import "C"
import (
	"unsafe"
)

type QDomImplementation struct {
	ptr unsafe.Pointer
}

type QDomImplementationITF interface {
	QDomImplementationPTR() *QDomImplementation
}

func (p *QDomImplementation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomImplementation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomImplementation(ptr QDomImplementationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomImplementationPTR().Pointer()
	}
	return nil
}

func QDomImplementationFromPointer(ptr unsafe.Pointer) *QDomImplementation {
	var n = new(QDomImplementation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomImplementation) QDomImplementationPTR() *QDomImplementation {
	return ptr
}

//QDomImplementation::InvalidDataPolicy
type QDomImplementation__InvalidDataPolicy int

var (
	QDomImplementation__AcceptInvalidChars = QDomImplementation__InvalidDataPolicy(0)
	QDomImplementation__DropInvalidChars   = QDomImplementation__InvalidDataPolicy(1)
	QDomImplementation__ReturnNullNode     = QDomImplementation__InvalidDataPolicy(2)
)

func NewQDomImplementation() *QDomImplementation {
	return QDomImplementationFromPointer(unsafe.Pointer(C.QDomImplementation_NewQDomImplementation()))
}

func NewQDomImplementation2(x QDomImplementationITF) *QDomImplementation {
	return QDomImplementationFromPointer(unsafe.Pointer(C.QDomImplementation_NewQDomImplementation2(C.QtObjectPtr(PointerFromQDomImplementation(x)))))
}

func (ptr *QDomImplementation) HasFeature(feature string, version string) bool {
	if ptr.Pointer() != nil {
		return C.QDomImplementation_HasFeature(C.QtObjectPtr(ptr.Pointer()), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func QDomImplementation_InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QDomImplementation_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QDomImplementation_SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.int(policy))
}

func (ptr *QDomImplementation) DestroyQDomImplementation() {
	if ptr.Pointer() != nil {
		C.QDomImplementation_DestroyQDomImplementation(C.QtObjectPtr(ptr.Pointer()))
	}
}
