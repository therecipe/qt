package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QMediaGaplessPlaybackControl struct {
	QMediaControl
}

type QMediaGaplessPlaybackControl_ITF interface {
	QMediaControl_ITF
	QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl
}

func PointerFromQMediaGaplessPlaybackControl(ptr QMediaGaplessPlaybackControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaGaplessPlaybackControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaGaplessPlaybackControlFromPointer(ptr unsafe.Pointer) *QMediaGaplessPlaybackControl {
	var n = new(QMediaGaplessPlaybackControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaGaplessPlaybackControl_") {
		n.SetObjectName("QMediaGaplessPlaybackControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaGaplessPlaybackControl) QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl {
	return ptr
}

func (ptr *QMediaGaplessPlaybackControl) ConnectAdvancedToNextMedia(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::advancedToNextMedia")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "advancedToNextMedia", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectAdvancedToNextMedia() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::advancedToNextMedia")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "advancedToNextMedia")
	}
}

//export callbackQMediaGaplessPlaybackControlAdvancedToNextMedia
func callbackQMediaGaplessPlaybackControlAdvancedToNextMedia(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::advancedToNextMedia")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "advancedToNextMedia").(func())()
}

func (ptr *QMediaGaplessPlaybackControl) CrossfadeTime() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::crossfadeTime")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMediaGaplessPlaybackControl_CrossfadeTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaGaplessPlaybackControl) IsCrossfadeSupported() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::isCrossfadeSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaGaplessPlaybackControl_IsCrossfadeSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaGaplessPlaybackControl) SetCrossfadeTime(crossfadeTime float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::setCrossfadeTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_SetCrossfadeTime(ptr.Pointer(), C.double(crossfadeTime))
	}
}

func (ptr *QMediaGaplessPlaybackControl) SetNextMedia(media QMediaContent_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::setNextMedia")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_SetNextMedia(ptr.Pointer(), PointerFromQMediaContent(media))
	}
}

func (ptr *QMediaGaplessPlaybackControl) DestroyQMediaGaplessPlaybackControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaGaplessPlaybackControl::~QMediaGaplessPlaybackControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
