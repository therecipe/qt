package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QStyleHintReturnVariant struct {
	QStyleHintReturn
}

type QStyleHintReturnVariant_ITF interface {
	QStyleHintReturn_ITF
	QStyleHintReturnVariant_PTR() *QStyleHintReturnVariant
}

func PointerFromQStyleHintReturnVariant(ptr QStyleHintReturnVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintReturnVariant_PTR().Pointer()
	}
	return nil
}

func NewQStyleHintReturnVariantFromPointer(ptr unsafe.Pointer) *QStyleHintReturnVariant {
	var n = new(QStyleHintReturnVariant)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleHintReturnVariant) QStyleHintReturnVariant_PTR() *QStyleHintReturnVariant {
	return ptr
}

//QStyleHintReturnVariant::StyleOptionType
type QStyleHintReturnVariant__StyleOptionType int64

var (
	QStyleHintReturnVariant__Type = QStyleHintReturnVariant__StyleOptionType(QStyleHintReturn__SH_Variant)
)

//QStyleHintReturnVariant::StyleOptionVersion
type QStyleHintReturnVariant__StyleOptionVersion int64

var (
	QStyleHintReturnVariant__Version = QStyleHintReturnVariant__StyleOptionVersion(1)
)

func NewQStyleHintReturnVariant() *QStyleHintReturnVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHintReturnVariant::QStyleHintReturnVariant")
		}
	}()

	return NewQStyleHintReturnVariantFromPointer(C.QStyleHintReturnVariant_NewQStyleHintReturnVariant())
}

func (ptr *QStyleHintReturnVariant) DestroyQStyleHintReturnVariant() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHintReturnVariant::~QStyleHintReturnVariant")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHintReturnVariant_DestroyQStyleHintReturnVariant(ptr.Pointer())
	}
}
