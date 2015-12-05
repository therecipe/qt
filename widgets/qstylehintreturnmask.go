package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QStyleHintReturnMask struct {
	QStyleHintReturn
}

type QStyleHintReturnMask_ITF interface {
	QStyleHintReturn_ITF
	QStyleHintReturnMask_PTR() *QStyleHintReturnMask
}

func PointerFromQStyleHintReturnMask(ptr QStyleHintReturnMask_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintReturnMask_PTR().Pointer()
	}
	return nil
}

func NewQStyleHintReturnMaskFromPointer(ptr unsafe.Pointer) *QStyleHintReturnMask {
	var n = new(QStyleHintReturnMask)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleHintReturnMask) QStyleHintReturnMask_PTR() *QStyleHintReturnMask {
	return ptr
}

//QStyleHintReturnMask::StyleOptionType
type QStyleHintReturnMask__StyleOptionType int64

var (
	QStyleHintReturnMask__Type = QStyleHintReturnMask__StyleOptionType(QStyleHintReturn__SH_Mask)
)

//QStyleHintReturnMask::StyleOptionVersion
type QStyleHintReturnMask__StyleOptionVersion int64

var (
	QStyleHintReturnMask__Version = QStyleHintReturnMask__StyleOptionVersion(1)
)

func NewQStyleHintReturnMask() *QStyleHintReturnMask {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHintReturnMask::QStyleHintReturnMask")
		}
	}()

	return NewQStyleHintReturnMaskFromPointer(C.QStyleHintReturnMask_NewQStyleHintReturnMask())
}

func (ptr *QStyleHintReturnMask) DestroyQStyleHintReturnMask() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHintReturnMask::~QStyleHintReturnMask")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHintReturnMask_DestroyQStyleHintReturnMask(ptr.Pointer())
	}
}
