package widgets

//#include "qstyleoptiontabbarbase.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionTabBarBase struct {
	QStyleOption
}

type QStyleOptionTabBarBaseITF interface {
	QStyleOptionITF
	QStyleOptionTabBarBasePTR() *QStyleOptionTabBarBase
}

func PointerFromQStyleOptionTabBarBase(ptr QStyleOptionTabBarBaseITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionTabBarBasePTR().Pointer()
	}
	return nil
}

func QStyleOptionTabBarBaseFromPointer(ptr unsafe.Pointer) *QStyleOptionTabBarBase {
	var n = new(QStyleOptionTabBarBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionTabBarBase) QStyleOptionTabBarBasePTR() *QStyleOptionTabBarBase {
	return ptr
}

//QStyleOptionTabBarBase::StyleOptionType
type QStyleOptionTabBarBase__StyleOptionType int

var (
	QStyleOptionTabBarBase__Type = QStyleOptionTabBarBase__StyleOptionType(QStyleOption__SO_TabBarBase)
)

//QStyleOptionTabBarBase::StyleOptionVersion
type QStyleOptionTabBarBase__StyleOptionVersion int

var (
	QStyleOptionTabBarBase__Version = QStyleOptionTabBarBase__StyleOptionVersion(2)
)

func NewQStyleOptionTabBarBase() *QStyleOptionTabBarBase {
	return QStyleOptionTabBarBaseFromPointer(unsafe.Pointer(C.QStyleOptionTabBarBase_NewQStyleOptionTabBarBase()))
}

func NewQStyleOptionTabBarBase2(other QStyleOptionTabBarBaseITF) *QStyleOptionTabBarBase {
	return QStyleOptionTabBarBaseFromPointer(unsafe.Pointer(C.QStyleOptionTabBarBase_NewQStyleOptionTabBarBase2(C.QtObjectPtr(PointerFromQStyleOptionTabBarBase(other)))))
}
