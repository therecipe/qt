package widgets

//#include "qstyleoptiontabbarbase.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionTabBarBase struct {
	QStyleOption
}

type QStyleOptionTabBarBase_ITF interface {
	QStyleOption_ITF
	QStyleOptionTabBarBase_PTR() *QStyleOptionTabBarBase
}

func PointerFromQStyleOptionTabBarBase(ptr QStyleOptionTabBarBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionTabBarBase_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionTabBarBaseFromPointer(ptr unsafe.Pointer) *QStyleOptionTabBarBase {
	var n = new(QStyleOptionTabBarBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionTabBarBase) QStyleOptionTabBarBase_PTR() *QStyleOptionTabBarBase {
	return ptr
}

//QStyleOptionTabBarBase::StyleOptionType
type QStyleOptionTabBarBase__StyleOptionType int64

var (
	QStyleOptionTabBarBase__Type = QStyleOptionTabBarBase__StyleOptionType(QStyleOption__SO_TabBarBase)
)

//QStyleOptionTabBarBase::StyleOptionVersion
type QStyleOptionTabBarBase__StyleOptionVersion int64

var (
	QStyleOptionTabBarBase__Version = QStyleOptionTabBarBase__StyleOptionVersion(2)
)

func NewQStyleOptionTabBarBase() *QStyleOptionTabBarBase {
	return NewQStyleOptionTabBarBaseFromPointer(C.QStyleOptionTabBarBase_NewQStyleOptionTabBarBase())
}

func NewQStyleOptionTabBarBase2(other QStyleOptionTabBarBase_ITF) *QStyleOptionTabBarBase {
	return NewQStyleOptionTabBarBaseFromPointer(C.QStyleOptionTabBarBase_NewQStyleOptionTabBarBase2(PointerFromQStyleOptionTabBarBase(other)))
}
