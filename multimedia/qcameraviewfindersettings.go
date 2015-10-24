package multimedia

//#include "qcameraviewfindersettings.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraViewfinderSettings struct {
	ptr unsafe.Pointer
}

type QCameraViewfinderSettingsITF interface {
	QCameraViewfinderSettingsPTR() *QCameraViewfinderSettings
}

func (p *QCameraViewfinderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraViewfinderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraViewfinderSettings(ptr QCameraViewfinderSettingsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsPTR().Pointer()
	}
	return nil
}

func QCameraViewfinderSettingsFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettings {
	var n = new(QCameraViewfinderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCameraViewfinderSettings) QCameraViewfinderSettingsPTR() *QCameraViewfinderSettings {
	return ptr
}

func NewQCameraViewfinderSettings() *QCameraViewfinderSettings {
	return QCameraViewfinderSettingsFromPointer(unsafe.Pointer(C.QCameraViewfinderSettings_NewQCameraViewfinderSettings()))
}

func NewQCameraViewfinderSettings2(other QCameraViewfinderSettingsITF) *QCameraViewfinderSettings {
	return QCameraViewfinderSettingsFromPointer(unsafe.Pointer(C.QCameraViewfinderSettings_NewQCameraViewfinderSettings2(C.QtObjectPtr(PointerFromQCameraViewfinderSettings(other)))))
}

func (ptr *QCameraViewfinderSettings) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QCameraViewfinderSettings_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraViewfinderSettings) PixelFormat() QVideoFrame__PixelFormat {
	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraViewfinderSettings_PixelFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraViewfinderSettings) SetPixelAspectRatio(ratio core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelAspectRatio(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(ratio)))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelAspectRatio2(horizontal int, vertical int) {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelAspectRatio2(C.QtObjectPtr(ptr.Pointer()), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelFormat(format QVideoFrame__PixelFormat) {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelFormat(C.QtObjectPtr(ptr.Pointer()), C.int(format))
	}
}

func (ptr *QCameraViewfinderSettings) SetResolution(resolution core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetResolution(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(resolution)))
	}
}

func (ptr *QCameraViewfinderSettings) SetResolution2(width int, height int) {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetResolution2(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height))
	}
}

func (ptr *QCameraViewfinderSettings) Swap(other QCameraViewfinderSettingsITF) {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCameraViewfinderSettings(other)))
	}
}

func (ptr *QCameraViewfinderSettings) DestroyQCameraViewfinderSettings() {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(C.QtObjectPtr(ptr.Pointer()))
	}
}
