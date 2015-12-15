package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QVideoEncoderSettingsControl struct {
	QMediaControl
}

type QVideoEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl
}

func PointerFromQVideoEncoderSettingsControl(ptr QVideoEncoderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoEncoderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettingsControl {
	var n = new(QVideoEncoderSettingsControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoEncoderSettingsControl_") {
		n.SetObjectName("QVideoEncoderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoEncoderSettingsControl) QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl {
	return ptr
}

func (ptr *QVideoEncoderSettingsControl) SetVideoSettings(settings QVideoEncoderSettings_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::setVideoSettings")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_SetVideoSettings(ptr.Pointer(), PointerFromQVideoEncoderSettings(settings))
	}
}

func (ptr *QVideoEncoderSettingsControl) SupportedVideoCodecs() []string {
	defer qt.Recovering("QVideoEncoderSettingsControl::supportedVideoCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QVideoEncoderSettingsControl_SupportedVideoCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QVideoEncoderSettingsControl) VideoCodecDescription(codec string) string {
	defer qt.Recovering("QVideoEncoderSettingsControl::videoCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoEncoderSettingsControl_VideoCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QVideoEncoderSettingsControl) DestroyQVideoEncoderSettingsControl() {
	defer qt.Recovering("QVideoEncoderSettingsControl::~QVideoEncoderSettingsControl")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
