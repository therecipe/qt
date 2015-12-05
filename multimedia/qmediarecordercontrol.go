package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QMediaRecorderControl struct {
	QMediaControl
}

type QMediaRecorderControl_ITF interface {
	QMediaControl_ITF
	QMediaRecorderControl_PTR() *QMediaRecorderControl
}

func PointerFromQMediaRecorderControl(ptr QMediaRecorderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaRecorderControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaRecorderControlFromPointer(ptr unsafe.Pointer) *QMediaRecorderControl {
	var n = new(QMediaRecorderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaRecorderControl_") {
		n.SetObjectName("QMediaRecorderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaRecorderControl) QMediaRecorderControl_PTR() *QMediaRecorderControl {
	return ptr
}

func (ptr *QMediaRecorderControl) ApplySettings() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::applySettings")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ApplySettings(ptr.Pointer())
	}
}

func (ptr *QMediaRecorderControl) ConnectError(f func(error int, errorString string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaRecorderControlError
func callbackQMediaRecorderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::error")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QMediaRecorderControl) IsMuted() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::isMuted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) ConnectMutedChanged(f func(muted bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectMutedChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderControlMutedChanged
func callbackQMediaRecorderControlMutedChanged(ptrName *C.char, muted C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::mutedChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QMediaRecorderControl) SetMuted(muted bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::setMuted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorderControl) SetOutputLocation(location core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::setOutputLocation")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_SetOutputLocation(ptr.Pointer(), core.PointerFromQUrl(location)) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) SetState(state QMediaRecorder__State) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::setState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaRecorderControl) SetVolume(gain float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::setVolume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetVolume(ptr.Pointer(), C.double(gain))
	}
}

func (ptr *QMediaRecorderControl) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderControlStateChanged
func callbackQMediaRecorderControlStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaRecorder__State))(QMediaRecorder__State(state))
}

func (ptr *QMediaRecorderControl) Status() QMediaRecorder__Status {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::status")
		}
	}()

	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorderControl_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::statusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::statusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderControlStatusChanged
func callbackQMediaRecorderControlStatusChanged(ptrName *C.char, status C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::statusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
}

func (ptr *QMediaRecorderControl) Volume() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::volume")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMediaRecorderControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) DestroyQMediaRecorderControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaRecorderControl::~QMediaRecorderControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DestroyQMediaRecorderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
