package widgets

//#include "qstylehintreturnmask.h"
import "C"
import (
	"unsafe"
)

type QStyleHintReturnMask struct {
	QStyleHintReturn
}

type QStyleHintReturnMaskITF interface {
	QStyleHintReturnITF
	QStyleHintReturnMaskPTR() *QStyleHintReturnMask
}

func PointerFromQStyleHintReturnMask(ptr QStyleHintReturnMaskITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintReturnMaskPTR().Pointer()
	}
	return nil
}

func QStyleHintReturnMaskFromPointer(ptr unsafe.Pointer) *QStyleHintReturnMask {
	var n = new(QStyleHintReturnMask)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleHintReturnMask) QStyleHintReturnMaskPTR() *QStyleHintReturnMask {
	return ptr
}

//QStyleHintReturnMask::StyleOptionType
type QStyleHintReturnMask__StyleOptionType int

var (
	QStyleHintReturnMask__Type = QStyleHintReturnMask__StyleOptionType(QStyleHintReturn__SH_Mask)
)

//QStyleHintReturnMask::StyleOptionVersion
type QStyleHintReturnMask__StyleOptionVersion int

var (
	QStyleHintReturnMask__Version = QStyleHintReturnMask__StyleOptionVersion(1)
)

func NewQStyleHintReturnMask() *QStyleHintReturnMask {
	return QStyleHintReturnMaskFromPointer(unsafe.Pointer(C.QStyleHintReturnMask_NewQStyleHintReturnMask()))
}

func (ptr *QStyleHintReturnMask) DestroyQStyleHintReturnMask() {
	if ptr.Pointer() != nil {
		C.QStyleHintReturnMask_DestroyQStyleHintReturnMask(C.QtObjectPtr(ptr.Pointer()))
	}
}
