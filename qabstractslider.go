package qt

//#include "qabstractslider.h"
import "C"

type qabstractslider struct {
	qwidget
}

type QAbstractSlider interface {
	QWidget
	HasTracking() bool
	InvertedAppearance() bool
	InvertedControls() bool
	IsSliderDown() bool
	Maximum() int
	Minimum() int
	Orientation() Orientation
	PageStep() int
	SetInvertedAppearance_Bool(invertedAppearance bool)
	SetInvertedControls_Bool(invertedControls bool)
	SetMaximum_Int(maximum int)
	SetMinimum_Int(minimum int)
	SetPageStep_Int(pageStep int)
	SetSingleStep_Int(singleStep int)
	SetSliderDown_Bool(sliderDown bool)
	SetSliderPosition_Int(sliderPosition int)
	SetTracking_Bool(enable bool)
	SingleStep() int
	SliderPosition() int
	Value() int
	ConnectSlotSetOrientation()
	DisconnectSlotSetOrientation()
	SlotSetOrientation_Orientation(orientation Orientation)
	ConnectSlotSetRange()
	DisconnectSlotSetRange()
	SlotSetRange_Int_Int(min int, max int)
	ConnectSlotSetValue()
	DisconnectSlotSetValue()
	SlotSetValue_Int(value int)
	ConnectSignalActionTriggered(f func())
	DisconnectSignalActionTriggered()
	SignalActionTriggered() func()
	ConnectSignalRangeChanged(f func())
	DisconnectSignalRangeChanged()
	SignalRangeChanged() func()
	ConnectSignalSliderMoved(f func())
	DisconnectSignalSliderMoved()
	SignalSliderMoved() func()
	ConnectSignalSliderPressed(f func())
	DisconnectSignalSliderPressed()
	SignalSliderPressed() func()
	ConnectSignalSliderReleased(f func())
	DisconnectSignalSliderReleased()
	SignalSliderReleased() func()
	ConnectSignalValueChanged(f func())
	DisconnectSignalValueChanged()
	SignalValueChanged() func()
}

func (p *qabstractslider) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qabstractslider) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQAbstractSlider_QWidget(parent QWidget) QAbstractSlider {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qabstractslider = new(qabstractslider)
	qabstractslider.SetPointer(C.QAbstractSlider_New_QWidget(parentPtr))
	qabstractslider.SetObjectName_String("QAbstractSlider_" + randomIdentifier())
	return qabstractslider
}

func (p *qabstractslider) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QAbstractSlider_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qabstractslider) HasTracking() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QAbstractSlider_HasTracking(p.Pointer()) != 0
	}
}

func (p *qabstractslider) InvertedAppearance() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QAbstractSlider_InvertedAppearance(p.Pointer()) != 0
	}
}

func (p *qabstractslider) InvertedControls() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QAbstractSlider_InvertedControls(p.Pointer()) != 0
	}
}

func (p *qabstractslider) IsSliderDown() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QAbstractSlider_IsSliderDown(p.Pointer()) != 0
	}
}

func (p *qabstractslider) Maximum() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QAbstractSlider_Maximum(p.Pointer()))
	}
}

func (p *qabstractslider) Minimum() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QAbstractSlider_Minimum(p.Pointer()))
	}
}

func (p *qabstractslider) Orientation() Orientation {
	if p.Pointer() == nil {
		return 0
	} else {
		return Orientation(C.QAbstractSlider_Orientation(p.Pointer()))
	}
}

func (p *qabstractslider) PageStep() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QAbstractSlider_PageStep(p.Pointer()))
	}
}

func (p *qabstractslider) SetInvertedAppearance_Bool(invertedAppearance bool) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetInvertedAppearance_Bool(p.Pointer(), goBoolToCInt(invertedAppearance))
	}
}

func (p *qabstractslider) SetInvertedControls_Bool(invertedControls bool) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetInvertedControls_Bool(p.Pointer(), goBoolToCInt(invertedControls))
	}
}

func (p *qabstractslider) SetMaximum_Int(maximum int) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetMaximum_Int(p.Pointer(), C.int(maximum))
	}
}

func (p *qabstractslider) SetMinimum_Int(minimum int) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetMinimum_Int(p.Pointer(), C.int(minimum))
	}
}

func (p *qabstractslider) SetPageStep_Int(pageStep int) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetPageStep_Int(p.Pointer(), C.int(pageStep))
	}
}

func (p *qabstractslider) SetSingleStep_Int(singleStep int) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetSingleStep_Int(p.Pointer(), C.int(singleStep))
	}
}

func (p *qabstractslider) SetSliderDown_Bool(sliderDown bool) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetSliderDown_Bool(p.Pointer(), goBoolToCInt(sliderDown))
	}
}

func (p *qabstractslider) SetSliderPosition_Int(sliderPosition int) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetSliderPosition_Int(p.Pointer(), C.int(sliderPosition))
	}
}

func (p *qabstractslider) SetTracking_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetTracking_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qabstractslider) SingleStep() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QAbstractSlider_SingleStep(p.Pointer()))
	}
}

func (p *qabstractslider) SliderPosition() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QAbstractSlider_SliderPosition(p.Pointer()))
	}
}

func (p *qabstractslider) Value() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QAbstractSlider_Value(p.Pointer()))
	}
}

func (p *qabstractslider) ConnectSlotSetOrientation() {
	C.QAbstractSlider_ConnectSlotSetOrientation(p.Pointer())
}

func (p *qabstractslider) DisconnectSlotSetOrientation() {
	C.QAbstractSlider_DisconnectSlotSetOrientation(p.Pointer())
}

func (p *qabstractslider) SlotSetOrientation_Orientation(orientation Orientation) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetOrientation_Orientation(p.Pointer(), C.int(orientation))
	}
}

func (p *qabstractslider) ConnectSlotSetRange() {
	C.QAbstractSlider_ConnectSlotSetRange(p.Pointer())
}

func (p *qabstractslider) DisconnectSlotSetRange() {
	C.QAbstractSlider_DisconnectSlotSetRange(p.Pointer())
}

func (p *qabstractslider) SlotSetRange_Int_Int(min int, max int) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetRange_Int_Int(p.Pointer(), C.int(min), C.int(max))
	}
}

func (p *qabstractslider) ConnectSlotSetValue() {
	C.QAbstractSlider_ConnectSlotSetValue(p.Pointer())
}

func (p *qabstractslider) DisconnectSlotSetValue() {
	C.QAbstractSlider_DisconnectSlotSetValue(p.Pointer())
}

func (p *qabstractslider) SlotSetValue_Int(value int) {
	if p.Pointer() != nil {
		C.QAbstractSlider_SetValue_Int(p.Pointer(), C.int(value))
	}
}

func (p *qabstractslider) ConnectSignalActionTriggered(f func()) {
	C.QAbstractSlider_ConnectSignalActionTriggered(p.Pointer())
	connectSignal(p.ObjectName(), "actionTriggered", f)
}

func (p *qabstractslider) DisconnectSignalActionTriggered() {
	C.QAbstractSlider_DisconnectSignalActionTriggered(p.Pointer())
	disconnectSignal(p.ObjectName(), "actionTriggered")
}

func (p *qabstractslider) SignalActionTriggered() func() {
	return getSignal(p.ObjectName(), "actionTriggered")
}

func (p *qabstractslider) ConnectSignalRangeChanged(f func()) {
	C.QAbstractSlider_ConnectSignalRangeChanged(p.Pointer())
	connectSignal(p.ObjectName(), "rangeChanged", f)
}

func (p *qabstractslider) DisconnectSignalRangeChanged() {
	C.QAbstractSlider_DisconnectSignalRangeChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "rangeChanged")
}

func (p *qabstractslider) SignalRangeChanged() func() {
	return getSignal(p.ObjectName(), "rangeChanged")
}

func (p *qabstractslider) ConnectSignalSliderMoved(f func()) {
	C.QAbstractSlider_ConnectSignalSliderMoved(p.Pointer())
	connectSignal(p.ObjectName(), "sliderMoved", f)
}

func (p *qabstractslider) DisconnectSignalSliderMoved() {
	C.QAbstractSlider_DisconnectSignalSliderMoved(p.Pointer())
	disconnectSignal(p.ObjectName(), "sliderMoved")
}

func (p *qabstractslider) SignalSliderMoved() func() {
	return getSignal(p.ObjectName(), "sliderMoved")
}

func (p *qabstractslider) ConnectSignalSliderPressed(f func()) {
	C.QAbstractSlider_ConnectSignalSliderPressed(p.Pointer())
	connectSignal(p.ObjectName(), "sliderPressed", f)
}

func (p *qabstractslider) DisconnectSignalSliderPressed() {
	C.QAbstractSlider_DisconnectSignalSliderPressed(p.Pointer())
	disconnectSignal(p.ObjectName(), "sliderPressed")
}

func (p *qabstractslider) SignalSliderPressed() func() {
	return getSignal(p.ObjectName(), "sliderPressed")
}

func (p *qabstractslider) ConnectSignalSliderReleased(f func()) {
	C.QAbstractSlider_ConnectSignalSliderReleased(p.Pointer())
	connectSignal(p.ObjectName(), "sliderReleased", f)
}

func (p *qabstractslider) DisconnectSignalSliderReleased() {
	C.QAbstractSlider_DisconnectSignalSliderReleased(p.Pointer())
	disconnectSignal(p.ObjectName(), "sliderReleased")
}

func (p *qabstractslider) SignalSliderReleased() func() {
	return getSignal(p.ObjectName(), "sliderReleased")
}

func (p *qabstractslider) ConnectSignalValueChanged(f func()) {
	C.QAbstractSlider_ConnectSignalValueChanged(p.Pointer())
	connectSignal(p.ObjectName(), "valueChanged", f)
}

func (p *qabstractslider) DisconnectSignalValueChanged() {
	C.QAbstractSlider_DisconnectSignalValueChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "valueChanged")
}

func (p *qabstractslider) SignalValueChanged() func() {
	return getSignal(p.ObjectName(), "valueChanged")
}

//export callbackQAbstractSlider
func callbackQAbstractSlider(ptr C.QtObjectPtr, msg *C.char) {
	var qabstractslider = new(qabstractslider)
	qabstractslider.SetPointer(ptr)
	getSignal(qabstractslider.ObjectName(), C.GoString(msg))()
}
