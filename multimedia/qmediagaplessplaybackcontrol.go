package multimedia

//#include "qmediagaplessplaybackcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaGaplessPlaybackControl struct {
	QMediaControl
}

type QMediaGaplessPlaybackControlITF interface {
	QMediaControlITF
	QMediaGaplessPlaybackControlPTR() *QMediaGaplessPlaybackControl
}

func PointerFromQMediaGaplessPlaybackControl(ptr QMediaGaplessPlaybackControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaGaplessPlaybackControlPTR().Pointer()
	}
	return nil
}

func QMediaGaplessPlaybackControlFromPointer(ptr unsafe.Pointer) *QMediaGaplessPlaybackControl {
	var n = new(QMediaGaplessPlaybackControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaGaplessPlaybackControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaGaplessPlaybackControl) QMediaGaplessPlaybackControlPTR() *QMediaGaplessPlaybackControl {
	return ptr
}

func (ptr *QMediaGaplessPlaybackControl) ConnectAdvancedToNextMedia(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "advancedToNextMedia", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectAdvancedToNextMedia() {
	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "advancedToNextMedia")
	}
}

//export callbackQMediaGaplessPlaybackControlAdvancedToNextMedia
func callbackQMediaGaplessPlaybackControlAdvancedToNextMedia(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "advancedToNextMedia").(func())()
}

func (ptr *QMediaGaplessPlaybackControl) IsCrossfadeSupported() bool {
	if ptr.Pointer() != nil {
		return C.QMediaGaplessPlaybackControl_IsCrossfadeSupported(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaGaplessPlaybackControl) SetNextMedia(media QMediaContentITF) {
	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_SetNextMedia(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaContent(media)))
	}
}

func (ptr *QMediaGaplessPlaybackControl) DestroyQMediaGaplessPlaybackControl() {
	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
