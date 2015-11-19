package multimedia

//#include "qaudioencodersettingscontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QAudioEncoderSettingsControl struct {
	QMediaControl
}

type QAudioEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl
}

func PointerFromQAudioEncoderSettingsControl(ptr QAudioEncoderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioEncoderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettingsControl {
	var n = new(QAudioEncoderSettingsControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioEncoderSettingsControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioEncoderSettingsControl) QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl {
	return ptr
}

func (ptr *QAudioEncoderSettingsControl) CodecDescription(codec string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettingsControl_CodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QAudioEncoderSettingsControl) SetAudioSettings(settings QAudioEncoderSettings_ITF) {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_SetAudioSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(settings))
	}
}

func (ptr *QAudioEncoderSettingsControl) SupportedAudioCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioEncoderSettingsControl_SupportedAudioCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QAudioEncoderSettingsControl) DestroyQAudioEncoderSettingsControl() {
	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
