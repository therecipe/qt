package multimedia

//#include "qimageencodercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QImageEncoderControl struct {
	QMediaControl
}

type QImageEncoderControlITF interface {
	QMediaControlITF
	QImageEncoderControlPTR() *QImageEncoderControl
}

func PointerFromQImageEncoderControl(ptr QImageEncoderControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageEncoderControlPTR().Pointer()
	}
	return nil
}

func QImageEncoderControlFromPointer(ptr unsafe.Pointer) *QImageEncoderControl {
	var n = new(QImageEncoderControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QImageEncoderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QImageEncoderControl) QImageEncoderControlPTR() *QImageEncoderControl {
	return ptr
}

func (ptr *QImageEncoderControl) ImageCodecDescription(codec string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderControl_ImageCodecDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(codec)))
	}
	return ""
}

func (ptr *QImageEncoderControl) SetImageSettings(settings QImageEncoderSettingsITF) {
	if ptr.Pointer() != nil {
		C.QImageEncoderControl_SetImageSettings(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImageEncoderSettings(settings)))
	}
}

func (ptr *QImageEncoderControl) SupportedImageCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImageEncoderControl_SupportedImageCodecs(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QImageEncoderControl) DestroyQImageEncoderControl() {
	if ptr.Pointer() != nil {
		C.QImageEncoderControl_DestroyQImageEncoderControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
