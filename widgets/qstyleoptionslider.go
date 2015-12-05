package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QStyleOptionSlider struct {
	QStyleOptionComplex
}

type QStyleOptionSlider_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSlider_PTR() *QStyleOptionSlider
}

func PointerFromQStyleOptionSlider(ptr QStyleOptionSlider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionSlider_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionSliderFromPointer(ptr unsafe.Pointer) *QStyleOptionSlider {
	var n = new(QStyleOptionSlider)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionSlider) QStyleOptionSlider_PTR() *QStyleOptionSlider {
	return ptr
}

//QStyleOptionSlider::StyleOptionType
type QStyleOptionSlider__StyleOptionType int64

var (
	QStyleOptionSlider__Type = QStyleOptionSlider__StyleOptionType(QStyleOption__SO_Slider)
)

//QStyleOptionSlider::StyleOptionVersion
type QStyleOptionSlider__StyleOptionVersion int64

var (
	QStyleOptionSlider__Version = QStyleOptionSlider__StyleOptionVersion(1)
)

func NewQStyleOptionSlider() *QStyleOptionSlider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionSlider::QStyleOptionSlider")
		}
	}()

	return NewQStyleOptionSliderFromPointer(C.QStyleOptionSlider_NewQStyleOptionSlider())
}

func NewQStyleOptionSlider2(other QStyleOptionSlider_ITF) *QStyleOptionSlider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionSlider::QStyleOptionSlider")
		}
	}()

	return NewQStyleOptionSliderFromPointer(C.QStyleOptionSlider_NewQStyleOptionSlider2(PointerFromQStyleOptionSlider(other)))
}
