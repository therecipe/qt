package widgets

//#include "qstylehintreturn.h"
import "C"
import (
	"unsafe"
)

type QStyleHintReturn struct {
	ptr unsafe.Pointer
}

type QStyleHintReturnITF interface {
	QStyleHintReturnPTR() *QStyleHintReturn
}

func (p *QStyleHintReturn) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStyleHintReturn) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStyleHintReturn(ptr QStyleHintReturnITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintReturnPTR().Pointer()
	}
	return nil
}

func QStyleHintReturnFromPointer(ptr unsafe.Pointer) *QStyleHintReturn {
	var n = new(QStyleHintReturn)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleHintReturn) QStyleHintReturnPTR() *QStyleHintReturn {
	return ptr
}

//QStyleHintReturn::HintReturnType
type QStyleHintReturn__HintReturnType int

var (
	QStyleHintReturn__SH_Default = QStyleHintReturn__HintReturnType(0xf000)
	QStyleHintReturn__SH_Mask    = QStyleHintReturn__HintReturnType(C.QStyleHintReturn_SH_Mask_Type())
	QStyleHintReturn__SH_Variant = QStyleHintReturn__HintReturnType(C.QStyleHintReturn_SH_Variant_Type())
)

//QStyleHintReturn::StyleOptionType
type QStyleHintReturn__StyleOptionType int

var (
	QStyleHintReturn__Type = QStyleHintReturn__StyleOptionType(QStyleHintReturn__SH_Default)
)

//QStyleHintReturn::StyleOptionVersion
type QStyleHintReturn__StyleOptionVersion int

var (
	QStyleHintReturn__Version = QStyleHintReturn__StyleOptionVersion(1)
)

func NewQStyleHintReturn(version int, ty int) *QStyleHintReturn {
	return QStyleHintReturnFromPointer(unsafe.Pointer(C.QStyleHintReturn_NewQStyleHintReturn(C.int(version), C.int(ty))))
}
