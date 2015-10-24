package widgets

//#include "qstyleoptionslider.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionSlider struct {
	QStyleOptionComplex
}

type QStyleOptionSliderITF interface {
	QStyleOptionComplexITF
	QStyleOptionSliderPTR() *QStyleOptionSlider
}

func PointerFromQStyleOptionSlider(ptr QStyleOptionSliderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionSliderPTR().Pointer()
	}
	return nil
}

func QStyleOptionSliderFromPointer(ptr unsafe.Pointer) *QStyleOptionSlider {
	var n = new(QStyleOptionSlider)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionSlider) QStyleOptionSliderPTR() *QStyleOptionSlider {
	return ptr
}

//QStyleOptionSlider::StyleOptionType
type QStyleOptionSlider__StyleOptionType int

var (
	QStyleOptionSlider__Type = QStyleOptionSlider__StyleOptionType(QStyleOption__SO_Slider)
)

//QStyleOptionSlider::StyleOptionVersion
type QStyleOptionSlider__StyleOptionVersion int

var (
	QStyleOptionSlider__Version = QStyleOptionSlider__StyleOptionVersion(1)
)

func NewQStyleOptionSlider() *QStyleOptionSlider {
	return QStyleOptionSliderFromPointer(unsafe.Pointer(C.QStyleOptionSlider_NewQStyleOptionSlider()))
}

func NewQStyleOptionSlider2(other QStyleOptionSliderITF) *QStyleOptionSlider {
	return QStyleOptionSliderFromPointer(unsafe.Pointer(C.QStyleOptionSlider_NewQStyleOptionSlider2(C.QtObjectPtr(PointerFromQStyleOptionSlider(other)))))
}
