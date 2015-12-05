package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QCameraViewfinderSettings struct {
	ptr unsafe.Pointer
}

type QCameraViewfinderSettings_ITF interface {
	QCameraViewfinderSettings_PTR() *QCameraViewfinderSettings
}

func (p *QCameraViewfinderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraViewfinderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraViewfinderSettings(ptr QCameraViewfinderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettings_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettings {
	var n = new(QCameraViewfinderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCameraViewfinderSettings) QCameraViewfinderSettings_PTR() *QCameraViewfinderSettings {
	return ptr
}

func NewQCameraViewfinderSettings() *QCameraViewfinderSettings {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::QCameraViewfinderSettings")
		}
	}()

	return NewQCameraViewfinderSettingsFromPointer(C.QCameraViewfinderSettings_NewQCameraViewfinderSettings())
}

func NewQCameraViewfinderSettings2(other QCameraViewfinderSettings_ITF) *QCameraViewfinderSettings {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::QCameraViewfinderSettings")
		}
	}()

	return NewQCameraViewfinderSettingsFromPointer(C.QCameraViewfinderSettings_NewQCameraViewfinderSettings2(PointerFromQCameraViewfinderSettings(other)))
}

func (ptr *QCameraViewfinderSettings) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraViewfinderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraViewfinderSettings) MaximumFrameRate() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::maximumFrameRate")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraViewfinderSettings_MaximumFrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraViewfinderSettings) MinimumFrameRate() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::minimumFrameRate")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraViewfinderSettings_MinimumFrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraViewfinderSettings) PixelFormat() QVideoFrame__PixelFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::pixelFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraViewfinderSettings_PixelFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraViewfinderSettings) SetMaximumFrameRate(rate float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::setMaximumFrameRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetMaximumFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QCameraViewfinderSettings) SetMinimumFrameRate(rate float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::setMinimumFrameRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetMinimumFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelAspectRatio(ratio core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::setPixelAspectRatio")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelAspectRatio(ptr.Pointer(), core.PointerFromQSize(ratio))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelAspectRatio2(horizontal int, vertical int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::setPixelAspectRatio")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelAspectRatio2(ptr.Pointer(), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelFormat(format QVideoFrame__PixelFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::setPixelFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraViewfinderSettings) SetResolution(resolution core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::setResolution")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QCameraViewfinderSettings) SetResolution2(width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::setResolution")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QCameraViewfinderSettings) Swap(other QCameraViewfinderSettings_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_Swap(ptr.Pointer(), PointerFromQCameraViewfinderSettings(other))
	}
}

func (ptr *QCameraViewfinderSettings) DestroyQCameraViewfinderSettings() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettings::~QCameraViewfinderSettings")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(ptr.Pointer())
	}
}
