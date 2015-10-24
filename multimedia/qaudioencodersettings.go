package multimedia

//#include "qaudioencodersettings.h"
import "C"
import (
	"unsafe"
)

type QAudioEncoderSettings struct {
	ptr unsafe.Pointer
}

type QAudioEncoderSettingsITF interface {
	QAudioEncoderSettingsPTR() *QAudioEncoderSettings
}

func (p *QAudioEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioEncoderSettings(ptr QAudioEncoderSettingsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioEncoderSettingsPTR().Pointer()
	}
	return nil
}

func QAudioEncoderSettingsFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettings {
	var n = new(QAudioEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAudioEncoderSettings) QAudioEncoderSettingsPTR() *QAudioEncoderSettings {
	return ptr
}

func NewQAudioEncoderSettings() *QAudioEncoderSettings {
	return QAudioEncoderSettingsFromPointer(unsafe.Pointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings()))
}

func NewQAudioEncoderSettings2(other QAudioEncoderSettingsITF) *QAudioEncoderSettings {
	return QAudioEncoderSettingsFromPointer(unsafe.Pointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings2(C.QtObjectPtr(PointerFromQAudioEncoderSettings(other)))))
}

func (ptr *QAudioEncoderSettings) BitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_BitRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) ChannelCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_ChannelCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) Codec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettings_Codec(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAudioEncoderSettings) EncodingOption(option string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettings_EncodingOption(C.QtObjectPtr(ptr.Pointer()), C.CString(option)))
	}
	return ""
}

func (ptr *QAudioEncoderSettings) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QAudioEncoderSettings_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAudioEncoderSettings) SampleRate() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_SampleRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) SetBitRate(rate int) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetBitRate(C.QtObjectPtr(ptr.Pointer()), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) SetChannelCount(channels int) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetChannelCount(C.QtObjectPtr(ptr.Pointer()), C.int(channels))
	}
}

func (ptr *QAudioEncoderSettings) SetCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetCodec(C.QtObjectPtr(ptr.Pointer()), C.CString(codec))
	}
}

func (ptr *QAudioEncoderSettings) SetEncodingOption(option string, value string) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetEncodingOption(C.QtObjectPtr(ptr.Pointer()), C.CString(option), C.CString(value))
	}
}

func (ptr *QAudioEncoderSettings) SetSampleRate(rate int) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetSampleRate(C.QtObjectPtr(ptr.Pointer()), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) DestroyQAudioEncoderSettings() {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_DestroyQAudioEncoderSettings(C.QtObjectPtr(ptr.Pointer()))
	}
}
