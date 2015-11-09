package widgets

//#include "qslider.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSlider struct {
	QAbstractSlider
}

type QSlider_ITF interface {
	QAbstractSlider_ITF
	QSlider_PTR() *QSlider
}

func PointerFromQSlider(ptr QSlider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSlider_PTR().Pointer()
	}
	return nil
}

func NewQSliderFromPointer(ptr unsafe.Pointer) *QSlider {
	var n = new(QSlider)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSlider_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSlider) QSlider_PTR() *QSlider {
	return ptr
}

//QSlider::TickPosition
type QSlider__TickPosition int64

const (
	QSlider__NoTicks        = QSlider__TickPosition(0)
	QSlider__TicksAbove     = QSlider__TickPosition(1)
	QSlider__TicksLeft      = QSlider__TickPosition(QSlider__TicksAbove)
	QSlider__TicksBelow     = QSlider__TickPosition(2)
	QSlider__TicksRight     = QSlider__TickPosition(QSlider__TicksBelow)
	QSlider__TicksBothSides = QSlider__TickPosition(3)
)

func (ptr *QSlider) SetTickInterval(ti int) {
	if ptr.Pointer() != nil {
		C.QSlider_SetTickInterval(ptr.Pointer(), C.int(ti))
	}
}

func (ptr *QSlider) SetTickPosition(position QSlider__TickPosition) {
	if ptr.Pointer() != nil {
		C.QSlider_SetTickPosition(ptr.Pointer(), C.int(position))
	}
}

func (ptr *QSlider) TickInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QSlider_TickInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSlider) TickPosition() QSlider__TickPosition {
	if ptr.Pointer() != nil {
		return QSlider__TickPosition(C.QSlider_TickPosition(ptr.Pointer()))
	}
	return 0
}

func NewQSlider(parent QWidget_ITF) *QSlider {
	return NewQSliderFromPointer(C.QSlider_NewQSlider(PointerFromQWidget(parent)))
}

func NewQSlider2(orientation core.Qt__Orientation, parent QWidget_ITF) *QSlider {
	return NewQSliderFromPointer(C.QSlider_NewQSlider2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QSlider) Event(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSlider_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSlider) DestroyQSlider() {
	if ptr.Pointer() != nil {
		C.QSlider_DestroyQSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
