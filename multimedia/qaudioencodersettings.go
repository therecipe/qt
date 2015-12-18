package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QAudioEncoderSettings::QAudioEncoderSettings")

	return NewQAudioEncoderSettingsFromPointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings())
}

func NewQAudioEncoderSettings2(other QAudioEncoderSettings_ITF) *QAudioEncoderSettings {
	defer qt.Recovering("QAudioEncoderSettings::QAudioEncoderSettings")

	return NewQAudioEncoderSettingsFromPointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings2(PointerFromQAudioEncoderSettings(other)))
}

func (ptr *QAudioEncoderSettings) BitRate() int {
	defer qt.Recovering("QAudioEncoderSettings::bitRate")

	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_BitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) ChannelCount() int {
	defer qt.Recovering("QAudioEncoderSettings::channelCount")

	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) Codec() string {
	defer qt.Recovering("QAudioEncoderSettings::codec")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettings_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioEncoderSettings) EncodingMode() QMultimedia__EncodingMode {
	defer qt.Recovering("QAudioEncoderSettings::encodingMode")

	if ptr.Pointer() != nil {
		return QMultimedia__EncodingMode(C.QAudioEncoderSettings_EncodingMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) EncodingOption(option string) *core.QVariant {
	defer qt.Recovering("QAudioEncoderSettings::encodingOption")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAudioEncoderSettings_EncodingOption(ptr.Pointer(), C.CString(option)))
	}
	return nil
}

func (ptr *QAudioEncoderSettings) IsNull() bool {
	defer qt.Recovering("QAudioEncoderSettings::isNull")

	if ptr.Pointer() != nil {
		return C.QAudioEncoderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioEncoderSettings) Quality() QMultimedia__EncodingQuality {
	defer qt.Recovering("QAudioEncoderSettings::quality")

	if ptr.Pointer() != nil {
		return QMultimedia__EncodingQuality(C.QAudioEncoderSettings_Quality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) SampleRate() int {
	defer qt.Recovering("QAudioEncoderSettings::sampleRate")

	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) SetBitRate(rate int) {
	defer qt.Recovering("QAudioEncoderSettings::setBitRate")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) SetChannelCount(channels int) {
	defer qt.Recovering("QAudioEncoderSettings::setChannelCount")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QAudioEncoderSettings) SetCodec(codec string) {
	defer qt.Recovering("QAudioEncoderSettings::setCodec")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QAudioEncoderSettings) SetEncodingMode(mode QMultimedia__EncodingMode) {
	defer qt.Recovering("QAudioEncoderSettings::setEncodingMode")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetEncodingMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAudioEncoderSettings) SetEncodingOption(option string, value core.QVariant_ITF) {
	defer qt.Recovering("QAudioEncoderSettings::setEncodingOption")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetEncodingOption(ptr.Pointer(), C.CString(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAudioEncoderSettings) SetQuality(quality QMultimedia__EncodingQuality) {
	defer qt.Recovering("QAudioEncoderSettings::setQuality")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetQuality(ptr.Pointer(), C.int(quality))
	}
}

func (ptr *QAudioEncoderSettings) SetSampleRate(rate int) {
	defer qt.Recovering("QAudioEncoderSettings::setSampleRate")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetSampleRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) DestroyQAudioEncoderSettings() {
	defer qt.Recovering("QAudioEncoderSettings::~QAudioEncoderSettings")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_DestroyQAudioEncoderSettings(ptr.Pointer())
	}
}
