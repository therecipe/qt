package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDomImplementation struct {
	ptr unsafe.Pointer
}

type QDomImplementation_ITF interface {
	QDomImplementation_PTR() *QDomImplementation
}

func (p *QDomImplementation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomImplementation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomImplementation(ptr QDomImplementation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomImplementation_PTR().Pointer()
	}
	return nil
}

func NewQDomImplementationFromPointer(ptr unsafe.Pointer) *QDomImplementation {
	var n = new(QDomImplementation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomImplementation) QDomImplementation_PTR() *QDomImplementation {
	return ptr
}

//QDomImplementation::InvalidDataPolicy
type QDomImplementation__InvalidDataPolicy int64

const (
	QDomImplementation__AcceptInvalidChars = QDomImplementation__InvalidDataPolicy(0)
	QDomImplementation__DropInvalidChars   = QDomImplementation__InvalidDataPolicy(1)
	QDomImplementation__ReturnNullNode     = QDomImplementation__InvalidDataPolicy(2)
)

func NewQDomImplementation() *QDomImplementation {
	defer qt.Recovering("QDomImplementation::QDomImplementation")

	return NewQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation())
}

func NewQDomImplementation2(x QDomImplementation_ITF) *QDomImplementation {
	defer qt.Recovering("QDomImplementation::QDomImplementation")

	return NewQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation2(PointerFromQDomImplementation(x)))
}

func (ptr *QDomImplementation) HasFeature(feature string, version string) bool {
	defer qt.Recovering("QDomImplementation::hasFeature")

	if ptr.Pointer() != nil {
		return C.QDomImplementation_HasFeature(ptr.Pointer(), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func QDomImplementation_InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	defer qt.Recovering("QDomImplementation::invalidDataPolicy")

	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) IsNull() bool {
	defer qt.Recovering("QDomImplementation::isNull")

	if ptr.Pointer() != nil {
		return C.QDomImplementation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func QDomImplementation_SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	defer qt.Recovering("QDomImplementation::setInvalidDataPolicy")

	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.int(policy))
}

func (ptr *QDomImplementation) DestroyQDomImplementation() {
	defer qt.Recovering("QDomImplementation::~QDomImplementation")

	if ptr.Pointer() != nil {
		C.QDomImplementation_DestroyQDomImplementation(ptr.Pointer())
	}
}
