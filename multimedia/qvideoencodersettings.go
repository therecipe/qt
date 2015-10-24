package multimedia

//#include "qvideoencodersettings.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVideoEncoderSettings struct {
	ptr unsafe.Pointer
}

type QVideoEncoderSettingsITF interface {
	QVideoEncoderSettingsPTR() *QVideoEncoderSettings
}

func (p *QVideoEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoEncoderSettings(ptr QVideoEncoderSettingsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoEncoderSettingsPTR().Pointer()
	}
	return nil
}

func QVideoEncoderSettingsFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettings {
	var n = new(QVideoEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoEncoderSettings) QVideoEncoderSettingsPTR() *QVideoEncoderSettings {
	return ptr
}

func NewQVideoEncoderSettings() *QVideoEncoderSettings {
	return QVideoEncoderSettingsFromPointer(unsafe.Pointer(C.QVideoEncoderSettings_NewQVideoEncoderSettings()))
}

func NewQVideoEncoderSettings2(other QVideoEncoderSettingsITF) *QVideoEncoderSettings {
	return QVideoEncoderSettingsFromPointer(unsafe.Pointer(C.QVideoEncoderSettings_NewQVideoEncoderSettings2(C.QtObjectPtr(PointerFromQVideoEncoderSettings(other)))))
}

func (ptr *QVideoEncoderSettings) BitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoEncoderSettings_BitRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoEncoderSettings) Codec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoEncoderSettings_Codec(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QVideoEncoderSettings) EncodingOption(option string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoEncoderSettings_EncodingOption(C.QtObjectPtr(ptr.Pointer()), C.CString(option)))
	}
	return ""
}

func (ptr *QVideoEncoderSettings) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QVideoEncoderSettings_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoEncoderSettings) SetBitRate(value int) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetBitRate(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QVideoEncoderSettings) SetCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetCodec(C.QtObjectPtr(ptr.Pointer()), C.CString(codec))
	}
}

func (ptr *QVideoEncoderSettings) SetEncodingOption(option string, value string) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetEncodingOption(C.QtObjectPtr(ptr.Pointer()), C.CString(option), C.CString(value))
	}
}

func (ptr *QVideoEncoderSettings) SetResolution(resolution core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetResolution(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(resolution)))
	}
}

func (ptr *QVideoEncoderSettings) SetResolution2(width int, height int) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetResolution2(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height))
	}
}

func (ptr *QVideoEncoderSettings) DestroyQVideoEncoderSettings() {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_DestroyQVideoEncoderSettings(C.QtObjectPtr(ptr.Pointer()))
	}
}
