package multimedia

//#include "qimageencodersettings.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QImageEncoderSettings struct {
	ptr unsafe.Pointer
}

type QImageEncoderSettingsITF interface {
	QImageEncoderSettingsPTR() *QImageEncoderSettings
}

func (p *QImageEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageEncoderSettings(ptr QImageEncoderSettingsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageEncoderSettingsPTR().Pointer()
	}
	return nil
}

func QImageEncoderSettingsFromPointer(ptr unsafe.Pointer) *QImageEncoderSettings {
	var n = new(QImageEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageEncoderSettings) QImageEncoderSettingsPTR() *QImageEncoderSettings {
	return ptr
}

func NewQImageEncoderSettings() *QImageEncoderSettings {
	return QImageEncoderSettingsFromPointer(unsafe.Pointer(C.QImageEncoderSettings_NewQImageEncoderSettings()))
}

func NewQImageEncoderSettings2(other QImageEncoderSettingsITF) *QImageEncoderSettings {
	return QImageEncoderSettingsFromPointer(unsafe.Pointer(C.QImageEncoderSettings_NewQImageEncoderSettings2(C.QtObjectPtr(PointerFromQImageEncoderSettings(other)))))
}

func (ptr *QImageEncoderSettings) Codec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderSettings_Codec(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QImageEncoderSettings) EncodingOption(option string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderSettings_EncodingOption(C.QtObjectPtr(ptr.Pointer()), C.CString(option)))
	}
	return ""
}

func (ptr *QImageEncoderSettings) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QImageEncoderSettings_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageEncoderSettings) SetCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetCodec(C.QtObjectPtr(ptr.Pointer()), C.CString(codec))
	}
}

func (ptr *QImageEncoderSettings) SetEncodingOption(option string, value string) {
	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetEncodingOption(C.QtObjectPtr(ptr.Pointer()), C.CString(option), C.CString(value))
	}
}

func (ptr *QImageEncoderSettings) SetResolution(resolution core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetResolution(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(resolution)))
	}
}

func (ptr *QImageEncoderSettings) SetResolution2(width int, height int) {
	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetResolution2(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height))
	}
}

func (ptr *QImageEncoderSettings) DestroyQImageEncoderSettings() {
	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_DestroyQImageEncoderSettings(C.QtObjectPtr(ptr.Pointer()))
	}
}
