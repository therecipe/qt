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

type QAbstractSliderITF interface {
	QWidgetITF
	QAbstractSliderPTR() *QAbstractSlider
}

func PointerFromQAbstractSlider(ptr QAbstractSliderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSliderPTR().Pointer()
	}
	return nil
}

func QAbstractSliderFromPointer(ptr unsafe.Pointer) *QAbstractSlider {
	var n = new(QAbstractSlider)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractSlider_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractSlider) QAbstractSliderPTR() *QAbstractSlider {
	return ptr
}

//QAbstractSlider::SliderAction
type QAbstractSlider__SliderAction int

var (
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
		return C.QAbstractSlider_HasTracking(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedAppearance() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedAppearance(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedControls() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedControls(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSlider) IsSliderDown() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSlider_IsSliderDown(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSlider) Maximum() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Maximum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) Minimum() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Minimum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QAbstractSlider_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) PageStep() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_PageStep(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) SetInvertedAppearance(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedAppearance(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetInvertedControls(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedControls(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetMaximum(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMaximum(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetMinimum(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMinimum(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetOrientation(v core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetPageStep(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetPageStep(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSingleStep(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSingleStep(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSliderDown(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderDown(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetSliderPosition(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderPosition(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetTracking(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetTracking(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractSlider) SetValue(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetValue(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractSlider) SingleStep() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SingleStep(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) SliderPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SliderPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQAbstractSlider(parent QWidgetITF) *QAbstractSlider {
	return QAbstractSliderFromPointer(unsafe.Pointer(C.QAbstractSlider_NewQAbstractSlider(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QAbstractSlider) ConnectActionTriggered(f func(action int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectActionTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "actionTriggered", f)
	}
}

func (ptr *QAbstractSlider) DisconnectActionTriggered() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectActionTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "actionTriggered")
	}
}

//export callbackQAbstractSliderActionTriggered
func callbackQAbstractSliderActionTriggered(ptrName *C.char, action C.int) {
	qt.GetSignal(C.GoString(ptrName), "actionTriggered").(func(int))(int(action))
}

func (ptr *QAbstractSlider) ConnectRangeChanged(f func(min int, max int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectRangeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rangeChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectRangeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rangeChanged")
	}
}

//export callbackQAbstractSliderRangeChanged
func callbackQAbstractSliderRangeChanged(ptrName *C.char, min C.int, max C.int) {
	qt.GetSignal(C.GoString(ptrName), "rangeChanged").(func(int, int))(int(min), int(max))
}

func (ptr *QAbstractSlider) SetRange(min int, max int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetRange(C.QtObjectPtr(ptr.Pointer()), C.int(min), C.int(max))
	}
}

func (ptr *QAbstractSlider) ConnectSliderMoved(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sliderMoved", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderMoved() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sliderMoved")
	}
}

//export callbackQAbstractSliderSliderMoved
func callbackQAbstractSliderSliderMoved(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "sliderMoved").(func(int))(int(value))
}

func (ptr *QAbstractSlider) ConnectSliderPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sliderPressed", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sliderPressed")
	}
}

//export callbackQAbstractSliderSliderPressed
func callbackQAbstractSliderSliderPressed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sliderPressed").(func())()
}

func (ptr *QAbstractSlider) ConnectSliderReleased(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderReleased(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sliderReleased", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderReleased() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderReleased(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sliderReleased")
	}
}

//export callbackQAbstractSliderSliderReleased
func callbackQAbstractSliderSliderReleased(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sliderReleased").(func())()
}

func (ptr *QAbstractSlider) TriggerAction(action QAbstractSlider__SliderAction) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_TriggerAction(C.QtObjectPtr(ptr.Pointer()), C.int(action))
	}
}

func (ptr *QAbstractSlider) ConnectValueChanged(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQAbstractSliderValueChanged
func callbackQAbstractSliderValueChanged(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(value))
}

func (ptr *QAbstractSlider) DestroyQAbstractSlider() {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_DestroyQAbstractSlider(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
