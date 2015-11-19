package widgets

//#include "qabstractslider.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractSlider struct {
	QWidget
}

type QAbstractSlider_ITF interface {
	QWidget_ITF
	QAbstractSlider_PTR() *QAbstractSlider
}

func PointerFromQAbstractSlider(ptr QAbstractSlider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSlider_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSliderFromPointer(ptr unsafe.Pointer) *QAbstractSlider {
	var n = new(QAbstractSlider)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractSlider_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractSlider) QAbstractSlider_PTR() *QAbstractSlider {
	return ptr
}

//QAbstractSlider::SliderAction
type QAbstractSlider__SliderAction int64

const (
	QAbstractSlider__SliderNoAction      = QAbstractSlider__SliderAction(0)
	QAbstractSlider__SliderSingleStepAdd = QAbstractSlider__SliderAction(1)
	QAbstractSlider__SliderSingleStepSub = QAbstractSlider__SliderAction(2)
	QAbstractSlider__SliderPageStepAdd   = QAbstractSlider__SliderAction(3)
	QAbstractSlider__SliderPageStepSub   = QAbstractSlider__SliderAction(4)
	QAbstractSlider__SliderToMinimum     = QAbstractSlider__SliderAction(5)
	QAbstractSlider__SliderToMaximum     = QAbstractSlider__SliderAction(6)
	QAbstractSlider__SliderMove          = QAbstractSlider__SliderAction(7)
)

func (ptr *QAbstractSlider) HasTracking() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSlider_HasTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedAppearance() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedAppearance(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedControls() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedControls(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) IsSliderDown() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSlider_IsSliderDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) Maximum() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Minimum() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QAbstractSlider_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) PageStep() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_PageStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) SetInvertedAppearance(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedAppearance(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetInvertedControls(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedControls(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetMaximum(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMaximum(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetMinimum(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMinimum(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetOrientation(v core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetPageStep(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetPageStep(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSingleStep(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSingleStep(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSliderDown(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderDown(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetSliderPosition(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetTracking(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractSlider) SetValue(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetValue(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SingleStep() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) SliderPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SliderPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Value(ptr.Pointer()))
	}
	return 0
}

func NewQAbstractSlider(parent QWidget_ITF) *QAbstractSlider {
	return NewQAbstractSliderFromPointer(C.QAbstractSlider_NewQAbstractSlider(PointerFromQWidget(parent)))
}

func (ptr *QAbstractSlider) ConnectActionTriggered(f func(action int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectActionTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actionTriggered", f)
	}
}

func (ptr *QAbstractSlider) DisconnectActionTriggered() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectActionTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actionTriggered")
	}
}

//export callbackQAbstractSliderActionTriggered
func callbackQAbstractSliderActionTriggered(ptrName *C.char, action C.int) {
	qt.GetSignal(C.GoString(ptrName), "actionTriggered").(func(int))(int(action))
}

func (ptr *QAbstractSlider) ConnectRangeChanged(f func(min int, max int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rangeChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rangeChanged")
	}
}

//export callbackQAbstractSliderRangeChanged
func callbackQAbstractSliderRangeChanged(ptrName *C.char, min C.int, max C.int) {
	qt.GetSignal(C.GoString(ptrName), "rangeChanged").(func(int, int))(int(min), int(max))
}

func (ptr *QAbstractSlider) SetRange(min int, max int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetRange(ptr.Pointer(), C.int(min), C.int(max))
	}
}

func (ptr *QAbstractSlider) ConnectSliderMoved(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderMoved", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderMoved() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderMoved")
	}
}

//export callbackQAbstractSliderSliderMoved
func callbackQAbstractSliderSliderMoved(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "sliderMoved").(func(int))(int(value))
}

func (ptr *QAbstractSlider) ConnectSliderPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderPressed", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderPressed")
	}
}

//export callbackQAbstractSliderSliderPressed
func callbackQAbstractSliderSliderPressed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sliderPressed").(func())()
}

func (ptr *QAbstractSlider) ConnectSliderReleased(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderReleased", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderReleased() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderReleased")
	}
}

//export callbackQAbstractSliderSliderReleased
func callbackQAbstractSliderSliderReleased(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sliderReleased").(func())()
}

func (ptr *QAbstractSlider) TriggerAction(action QAbstractSlider__SliderAction) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_TriggerAction(ptr.Pointer(), C.int(action))
	}
}

func (ptr *QAbstractSlider) ConnectValueChanged(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQAbstractSliderValueChanged
func callbackQAbstractSliderValueChanged(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(value))
}

func (ptr *QAbstractSlider) DestroyQAbstractSlider() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DestroyQAbstractSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
