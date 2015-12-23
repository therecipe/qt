package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraViewfinderSettingsControl struct {
	QMediaControl
}

type QCameraViewfinderSettingsControl_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl
}

func PointerFromQCameraViewfinderSettingsControl(ptr QCameraViewfinderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsControlFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl {
	var n = new(QCameraViewfinderSettingsControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinderSettingsControl_") {
		n.SetObjectName("QCameraViewfinderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraViewfinderSettingsControl) QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl {
	return ptr
}

//QCameraViewfinderSettingsControl::ViewfinderParameter
type QCameraViewfinderSettingsControl__ViewfinderParameter int64

const (
	QCameraViewfinderSettingsControl__Resolution       = QCameraViewfinderSettingsControl__ViewfinderParameter(0)
	QCameraViewfinderSettingsControl__PixelAspectRatio = QCameraViewfinderSettingsControl__ViewfinderParameter(1)
	QCameraViewfinderSettingsControl__MinimumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(2)
	QCameraViewfinderSettingsControl__MaximumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(3)
	QCameraViewfinderSettingsControl__PixelFormat      = QCameraViewfinderSettingsControl__ViewfinderParameter(4)
	QCameraViewfinderSettingsControl__UserParameter    = QCameraViewfinderSettingsControl__ViewfinderParameter(1000)
)

func (ptr *QCameraViewfinderSettingsControl) IsViewfinderParameterSupported(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) bool {
	defer qt.Recovering("QCameraViewfinderSettingsControl::isViewfinderParameterSupported")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraViewfinderSettingsControl) SetViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::setViewfinderParameter")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_SetViewfinderParameter(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) *core.QVariant {
	defer qt.Recovering("QCameraViewfinderSettingsControl::viewfinderParameter")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraViewfinderSettingsControl_ViewfinderParameter(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraViewfinderSettingsControl) DestroyQCameraViewfinderSettingsControl() {
	defer qt.Recovering("QCameraViewfinderSettingsControl::~QCameraViewfinderSettingsControl")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraViewfinderSettingsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraViewfinderSettingsControlTimerEvent
func callbackQCameraViewfinderSettingsControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinderSettingsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraViewfinderSettingsControlChildEvent
func callbackQCameraViewfinderSettingsControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinderSettingsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraViewfinderSettingsControlCustomEvent
func callbackQCameraViewfinderSettingsControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
