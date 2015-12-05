package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	for len(n.ObjectName()) < len("QAbstractSlider_") {
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::hasTracking")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_HasTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedAppearance() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::invertedAppearance")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedAppearance(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedControls() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::invertedControls")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedControls(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) IsSliderDown() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::isSliderDown")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_IsSliderDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) Maximum() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::maximum")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Minimum() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::minimum")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QAbstractSlider_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) PageStep() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::pageStep")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_PageStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) SetInvertedAppearance(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setInvertedAppearance")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedAppearance(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetInvertedControls(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setInvertedControls")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedControls(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetMaximum(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setMaximum")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMaximum(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetMinimum(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setMinimum")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMinimum(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetOrientation(v core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetPageStep(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setPageStep")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetPageStep(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSingleStep(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setSingleStep")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSingleStep(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSliderDown(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setSliderDown")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderDown(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetSliderPosition(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setSliderPosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetTracking(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setTracking")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractSlider) SetValue(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setValue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetValue(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SingleStep() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::singleStep")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) SliderPosition() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SliderPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Value() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::value")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Value(ptr.Pointer()))
	}
	return 0
}

func NewQAbstractSlider(parent QWidget_ITF) *QAbstractSlider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::QAbstractSlider")
		}
	}()

	return NewQAbstractSliderFromPointer(C.QAbstractSlider_NewQAbstractSlider(PointerFromQWidget(parent)))
}

func (ptr *QAbstractSlider) ConnectActionTriggered(f func(action int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::actionTriggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectActionTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actionTriggered", f)
	}
}

func (ptr *QAbstractSlider) DisconnectActionTriggered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::actionTriggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectActionTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actionTriggered")
	}
}

//export callbackQAbstractSliderActionTriggered
func callbackQAbstractSliderActionTriggered(ptrName *C.char, action C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::actionTriggered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "actionTriggered").(func(int))(int(action))
}

func (ptr *QAbstractSlider) ConnectRangeChanged(f func(min int, max int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::rangeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rangeChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectRangeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::rangeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rangeChanged")
	}
}

//export callbackQAbstractSliderRangeChanged
func callbackQAbstractSliderRangeChanged(ptrName *C.char, min C.int, max C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::rangeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "rangeChanged").(func(int, int))(int(min), int(max))
}

func (ptr *QAbstractSlider) SetRange(min int, max int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::setRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetRange(ptr.Pointer(), C.int(min), C.int(max))
	}
}

func (ptr *QAbstractSlider) ConnectSliderMoved(f func(value int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderMoved", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderMoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderMoved")
	}
}

//export callbackQAbstractSliderSliderMoved
func callbackQAbstractSliderSliderMoved(ptrName *C.char, value C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderMoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sliderMoved").(func(int))(int(value))
}

func (ptr *QAbstractSlider) ConnectSliderPressed(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderPressed", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderPressed")
	}
}

//export callbackQAbstractSliderSliderPressed
func callbackQAbstractSliderSliderPressed(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderPressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sliderPressed").(func())()
}

func (ptr *QAbstractSlider) ConnectSliderReleased(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderReleased")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderReleased", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderReleased() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderReleased")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderReleased")
	}
}

//export callbackQAbstractSliderSliderReleased
func callbackQAbstractSliderSliderReleased(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::sliderReleased")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sliderReleased").(func())()
}

func (ptr *QAbstractSlider) TriggerAction(action QAbstractSlider__SliderAction) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::triggerAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_TriggerAction(ptr.Pointer(), C.int(action))
	}
}

func (ptr *QAbstractSlider) ConnectValueChanged(f func(value int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::valueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectValueChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::valueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQAbstractSliderValueChanged
func callbackQAbstractSliderValueChanged(ptrName *C.char, value C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::valueChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(value))
}

func (ptr *QAbstractSlider) DestroyQAbstractSlider() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSlider::~QAbstractSlider")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DestroyQAbstractSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
