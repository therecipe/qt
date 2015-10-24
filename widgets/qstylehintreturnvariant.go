package widgets

//#include "qstylehintreturnvariant.h"
import "C"
import (
	"unsafe"
)

type QStyleHintReturnVariant struct {
	QStyleHintReturn
}

type QStyleHintReturnVariantITF interface {
	QStyleHintReturnITF
	QStyleHintReturnVariantPTR() *QStyleHintReturnVariant
}

func PointerFromQStyleHintReturnVariant(ptr QStyleHintReturnVariantITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintReturnVariantPTR().Pointer()
	}
	return nil
}

func QStyleHintReturnVariantFromPointer(ptr unsafe.Pointer) *QStyleHintReturnVariant {
	var n = new(QStyleHintReturnVariant)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleHintReturnVariant) QStyleHintReturnVariantPTR() *QStyleHintReturnVariant {
	return ptr
}

//QStyleHintReturnVariant::StyleOptionType
type QStyleHintReturnVariant__StyleOptionType int

var (
	QStyleHintReturnVariant__Type = QStyleHintReturnVariant__StyleOptionType(QStyleHintReturn__SH_Variant)
)

//QStyleHintReturnVariant::StyleOptionVersion
type QStyleHintReturnVariant__StyleOptionVersion int

var (
	QStyleHintReturnVariant__Version = QStyleHintReturnVariant__StyleOptionVersion(1)
)

func NewQStyleHintReturnVariant() *QStyleHintReturnVariant {
	return QStyleHintReturnVariantFromPointer(unsafe.Pointer(C.QStyleHintReturnVariant_NewQStyleHintReturnVariant()))
}

func (ptr *QStyleHintReturnVariant) DestroyQStyleHintReturnVariant() {
	if ptr.Pointer() != nil {
		C.QStyleHintReturnVariant_DestroyQStyleHintReturnVariant(C.QtObjectPtr(ptr.Pointer()))
	}
}
