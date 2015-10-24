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

type QSliderITF interface {
	QAbstractSliderITF
	QSliderPTR() *QSlider
}

func PointerFromQSlider(ptr QSliderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSliderPTR().Pointer()
	}
	return nil
}

func QSliderFromPointer(ptr unsafe.Pointer) *QSlider {
	var n = new(QSlider)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSlider_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSlider) QSliderPTR() *QSlider {
	return ptr
}

//QSlider::TickPosition
type QSlider__TickPosition int

var (
	QSlider__NoTicks        = QSlider__TickPosition(0)
	QSlider__TicksAbove     = QSlider__TickPosition(1)
	QSlider__TicksLeft      = QSlider__TickPosition(QSlider__TicksAbove)
	QSlider__TicksBelow     = QSlider__TickPosition(2)
	QSlider__TicksRight     = QSlider__TickPosition(QSlider__TicksBelow)
	QSlider__TicksBothSides = QSlider__TickPosition(3)
)

func (ptr *QSlider) SetTickInterval(ti int) {
	if ptr.Pointer() != nil {
		C.QSlider_SetTickInterval(C.QtObjectPtr(ptr.Pointer()), C.int(ti))
	}
}

func (ptr *QSlider) SetTickPosition(position QSlider__TickPosition) {
	if ptr.Pointer() != nil {
		C.QSlider_SetTickPosition(C.QtObjectPtr(ptr.Pointer()), C.int(position))
	}
}

func (ptr *QSlider) TickInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QSlider_TickInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSlider) TickPosition() QSlider__TickPosition {
	if ptr.Pointer() != nil {
		return QSlider__TickPosition(C.QSlider_TickPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQSlider(parent QWidgetITF) *QSlider {
	return QSliderFromPointer(unsafe.Pointer(C.QSlider_NewQSlider(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQSlider2(orientation core.Qt__Orientation, parent QWidgetITF) *QSlider {
	return QSliderFromPointer(unsafe.Pointer(C.QSlider_NewQSlider2(C.int(orientation), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QSlider) Event(event core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QSlider_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QSlider) DestroyQSlider() {
	if ptr.Pointer() != nil {
		C.QSlider_DestroyQSlider(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
