package multimedia

//#include "multimedia.h"
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
	for len(n.ObjectName()) < len("QAudioEncoderSettingsControl_") {
		n.SetObjectName("QAudioEncoderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioEncoderSettingsControl) QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl {
	return ptr
}

func (ptr *QAudioEncoderSettingsControl) CodecDescription(codec string) string {
	defer qt.Recovering("QAudioEncoderSettingsControl::codecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettingsControl_CodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QAudioEncoderSettingsControl) SetAudioSettings(settings QAudioEncoderSettings_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::setAudioSettings")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_SetAudioSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(settings))
	}
}

func (ptr *QAudioEncoderSettingsControl) SupportedAudioCodecs() []string {
	defer qt.Recovering("QAudioEncoderSettingsControl::supportedAudioCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioEncoderSettingsControl_SupportedAudioCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAudioEncoderSettingsControl) DestroyQAudioEncoderSettingsControl() {
	defer qt.Recovering("QAudioEncoderSettingsControl::~QAudioEncoderSettingsControl")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
