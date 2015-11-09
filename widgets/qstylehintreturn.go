package widgets

//#include "qstylehintreturn.h"
import "C"
import (
	"unsafe"
)

type QStyleHintReturn struct {
	ptr unsafe.Pointer
}

type QStyleHintReturn_ITF interface {
	QStyleHintReturn_PTR() *QStyleHintReturn
}

func (p *QStyleHintReturn) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStyleHintReturn) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStyleHintReturn(ptr QStyleHintReturn_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintReturn_PTR().Pointer()
	}
	return nil
}

func NewQStyleHintReturnFromPointer(ptr unsafe.Pointer) *QStyleHintReturn {
	var n = new(QStyleHintReturn)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleHintReturn) QStyleHintReturn_PTR() *QStyleHintReturn {
	return ptr
}

//QStyleHintReturn::HintReturnType
type QStyleHintReturn__HintReturnType int64

var (
	QStyleHintReturn__SH_Default = QStyleHintReturn__HintReturnType(0xf000)
	QStyleHintReturn__SH_Mask    = QStyleHintReturn__HintReturnType(C.QStyleHintReturn_SH_Mask_Type())
	QStyleHintReturn__SH_Variant = QStyleHintReturn__HintReturnType(C.QStyleHintReturn_SH_Variant_Type())
)

//QStyleHintReturn::StyleOptionType
type QStyleHintReturn__StyleOptionType int64

var (
	QStyleHintReturn__Type = QStyleHintReturn__StyleOptionType(QStyleHintReturn__SH_Default)
)

//QStyleHintReturn::StyleOptionVersion
type QStyleHintReturn__StyleOptionVersion int64

var (
	QStyleHintReturn__Version = QStyleHintReturn__StyleOptionVersion(1)
)

func NewQStyleHintReturn(version int, ty int) *QStyleHintReturn {
	return NewQStyleHintReturnFromPointer(C.QStyleHintReturn_NewQStyleHintReturn(C.int(version), C.int(ty)))
}
