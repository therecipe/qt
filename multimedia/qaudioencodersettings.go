package multimedia

//#include "qaudioencodersettings.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioEncoderSettings struct {
	ptr unsafe.Pointer
}

type QAudioEncoderSettings_ITF interface {
	QAudioEncoderSettings_PTR() *QAudioEncoderSettings
}

func (p *QAudioEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioEncoderSettings(ptr QAudioEncoderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioEncoderSettings_PTR().Pointer()
	}
	return nil
}

func NewQAudioEncoderSettingsFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettings {
	var n = new(QAudioEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAudioEncoderSettings) QAudioEncoderSettings_PTR() *QAudioEncoderSettings {
	return ptr
}

func NewQAudioEncoderSettings() *QAudioEncoderSettings {
	return NewQAudioEncoderSettingsFromPointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings())
}

func NewQAudioEncoderSettings2(other QAudioEncoderSettings_ITF) *QAudioEncoderSettings {
	return NewQAudioEncoderSettingsFromPointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings2(PointerFromQAudioEncoderSettings(other)))
}

func (ptr *QAudioEncoderSettings) BitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_BitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) ChannelCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) Codec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettings_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioEncoderSettings) EncodingOption(option string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAudioEncoderSettings_EncodingOption(ptr.Pointer(), C.CString(option)))
	}
	return nil
}

func (ptr *QAudioEncoderSettings) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QAudioEncoderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioEncoderSettings) SampleRate() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) SetBitRate(rate int) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) SetChannelCount(channels int) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QAudioEncoderSettings) SetCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QAudioEncoderSettings) SetEncodingOption(option string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetEncodingOption(ptr.Pointer(), C.CString(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAudioEncoderSettings) SetSampleRate(rate int) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetSampleRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) DestroyQAudioEncoderSettings() {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_DestroyQAudioEncoderSettings(ptr.Pointer())
	}
}
