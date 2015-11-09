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

type QVideoEncoderSettings_ITF interface {
	QVideoEncoderSettings_PTR() *QVideoEncoderSettings
}

func (p *QVideoEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoEncoderSettings(ptr QVideoEncoderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoEncoderSettings_PTR().Pointer()
	}
	return nil
}

func NewQVideoEncoderSettingsFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettings {
	var n = new(QVideoEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoEncoderSettings) QVideoEncoderSettings_PTR() *QVideoEncoderSettings {
	return ptr
}

func (ptr *QVideoEncoderSettings) SetFrameRate(rate float64) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func NewQVideoEncoderSettings() *QVideoEncoderSettings {
	return NewQVideoEncoderSettingsFromPointer(C.QVideoEncoderSettings_NewQVideoEncoderSettings())
}

func NewQVideoEncoderSettings2(other QVideoEncoderSettings_ITF) *QVideoEncoderSettings {
	return NewQVideoEncoderSettingsFromPointer(C.QVideoEncoderSettings_NewQVideoEncoderSettings2(PointerFromQVideoEncoderSettings(other)))
}

func (ptr *QVideoEncoderSettings) BitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoEncoderSettings_BitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoEncoderSettings) Codec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoEncoderSettings_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QVideoEncoderSettings) EncodingOption(option string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoEncoderSettings_EncodingOption(ptr.Pointer(), C.CString(option)))
	}
	return nil
}

func (ptr *QVideoEncoderSettings) FrameRate() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QVideoEncoderSettings_FrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoEncoderSettings) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QVideoEncoderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoEncoderSettings) SetBitRate(value int) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetBitRate(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QVideoEncoderSettings) SetCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QVideoEncoderSettings) SetEncodingOption(option string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetEncodingOption(ptr.Pointer(), C.CString(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QVideoEncoderSettings) SetResolution(resolution core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QVideoEncoderSettings) SetResolution2(width int, height int) {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QVideoEncoderSettings) DestroyQVideoEncoderSettings() {
	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_DestroyQVideoEncoderSettings(ptr.Pointer())
	}
}
