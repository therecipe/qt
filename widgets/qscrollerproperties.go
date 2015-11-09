package widgets

//#include "qscrollerproperties.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScrollerProperties struct {
	ptr unsafe.Pointer
}

type QScrollerProperties_ITF interface {
	QScrollerProperties_PTR() *QScrollerProperties
}

func (p *QScrollerProperties) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScrollerProperties) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScrollerProperties(ptr QScrollerProperties_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollerProperties_PTR().Pointer()
	}
	return nil
}

func NewQScrollerPropertiesFromPointer(ptr unsafe.Pointer) *QScrollerProperties {
	var n = new(QScrollerProperties)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScrollerProperties) QScrollerProperties_PTR() *QScrollerProperties {
	return ptr
}

//QScrollerProperties::FrameRates
type QScrollerProperties__FrameRates int64

const (
	QScrollerProperties__Standard = QScrollerProperties__FrameRates(0)
	QScrollerProperties__Fps60    = QScrollerProperties__FrameRates(1)
	QScrollerProperties__Fps30    = QScrollerProperties__FrameRates(2)
	QScrollerProperties__Fps20    = QScrollerProperties__FrameRates(3)
)

//QScrollerProperties::OvershootPolicy
type QScrollerProperties__OvershootPolicy int64

const (
	QScrollerProperties__OvershootWhenScrollable = QScrollerProperties__OvershootPolicy(0)
	QScrollerProperties__OvershootAlwaysOff      = QScrollerProperties__OvershootPolicy(1)
	QScrollerProperties__OvershootAlwaysOn       = QScrollerProperties__OvershootPolicy(2)
)

//QScrollerProperties::ScrollMetric
type QScrollerProperties__ScrollMetric int64

const (
	QScrollerProperties__MousePressEventDelay           = QScrollerProperties__ScrollMetric(0)
	QScrollerProperties__DragStartDistance              = QScrollerProperties__ScrollMetric(1)
	QScrollerProperties__DragVelocitySmoothingFactor    = QScrollerProperties__ScrollMetric(2)
	QScrollerProperties__AxisLockThreshold              = QScrollerProperties__ScrollMetric(3)
	QScrollerProperties__ScrollingCurve                 = QScrollerProperties__ScrollMetric(4)
	QScrollerProperties__DecelerationFactor             = QScrollerProperties__ScrollMetric(5)
	QScrollerProperties__MinimumVelocity                = QScrollerProperties__ScrollMetric(6)
	QScrollerProperties__MaximumVelocity                = QScrollerProperties__ScrollMetric(7)
	QScrollerProperties__MaximumClickThroughVelocity    = QScrollerProperties__ScrollMetric(8)
	QScrollerProperties__AcceleratingFlickMaximumTime   = QScrollerProperties__ScrollMetric(9)
	QScrollerProperties__AcceleratingFlickSpeedupFactor = QScrollerProperties__ScrollMetric(10)
	QScrollerProperties__SnapPositionRatio              = QScrollerProperties__ScrollMetric(11)
	QScrollerProperties__SnapTime                       = QScrollerProperties__ScrollMetric(12)
	QScrollerProperties__OvershootDragResistanceFactor  = QScrollerProperties__ScrollMetric(13)
	QScrollerProperties__OvershootDragDistanceFactor    = QScrollerProperties__ScrollMetric(14)
	QScrollerProperties__OvershootScrollDistanceFactor  = QScrollerProperties__ScrollMetric(15)
	QScrollerProperties__OvershootScrollTime            = QScrollerProperties__ScrollMetric(16)
	QScrollerProperties__HorizontalOvershootPolicy      = QScrollerProperties__ScrollMetric(17)
	QScrollerProperties__VerticalOvershootPolicy        = QScrollerProperties__ScrollMetric(18)
	QScrollerProperties__FrameRate                      = QScrollerProperties__ScrollMetric(19)
	QScrollerProperties__ScrollMetricCount              = QScrollerProperties__ScrollMetric(20)
)

func NewQScrollerProperties() *QScrollerProperties {
	return NewQScrollerPropertiesFromPointer(C.QScrollerProperties_NewQScrollerProperties())
}

func NewQScrollerProperties2(sp QScrollerProperties_ITF) *QScrollerProperties {
	return NewQScrollerPropertiesFromPointer(C.QScrollerProperties_NewQScrollerProperties2(PointerFromQScrollerProperties(sp)))
}

func (ptr *QScrollerProperties) ScrollMetric(metric QScrollerProperties__ScrollMetric) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScrollerProperties_ScrollMetric(ptr.Pointer(), C.int(metric)))
	}
	return nil
}

func QScrollerProperties_SetDefaultScrollerProperties(sp QScrollerProperties_ITF) {
	C.QScrollerProperties_QScrollerProperties_SetDefaultScrollerProperties(PointerFromQScrollerProperties(sp))
}

func (ptr *QScrollerProperties) SetScrollMetric(metric QScrollerProperties__ScrollMetric, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QScrollerProperties_SetScrollMetric(ptr.Pointer(), C.int(metric), core.PointerFromQVariant(value))
	}
}

func QScrollerProperties_UnsetDefaultScrollerProperties() {
	C.QScrollerProperties_QScrollerProperties_UnsetDefaultScrollerProperties()
}

func (ptr *QScrollerProperties) DestroyQScrollerProperties() {
	if ptr.Pointer() != nil {
		C.QScrollerProperties_DestroyQScrollerProperties(ptr.Pointer())
	}
}
