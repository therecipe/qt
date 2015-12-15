package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QImageEncoderControl struct {
	QMediaControl
}

type QImageEncoderControl_ITF interface {
	QMediaControl_ITF
	QImageEncoderControl_PTR() *QImageEncoderControl
}

func PointerFromQImageEncoderControl(ptr QImageEncoderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageEncoderControl_PTR().Pointer()
	}
	return nil
}

func NewQImageEncoderControlFromPointer(ptr unsafe.Pointer) *QImageEncoderControl {
	var n = new(QImageEncoderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QImageEncoderControl_") {
		n.SetObjectName("QImageEncoderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QImageEncoderControl) QImageEncoderControl_PTR() *QImageEncoderControl {
	return ptr
}

func (ptr *QImageEncoderControl) ImageCodecDescription(codec string) string {
	defer qt.Recovering("QImageEncoderControl::imageCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderControl_ImageCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QImageEncoderControl) SetImageSettings(settings QImageEncoderSettings_ITF) {
	defer qt.Recovering("QImageEncoderControl::setImageSettings")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_SetImageSettings(ptr.Pointer(), PointerFromQImageEncoderSettings(settings))
	}
}

func (ptr *QImageEncoderControl) SupportedImageCodecs() []string {
	defer qt.Recovering("QImageEncoderControl::supportedImageCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImageEncoderControl_SupportedImageCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QImageEncoderControl) DestroyQImageEncoderControl() {
	defer qt.Recovering("QImageEncoderControl::~QImageEncoderControl")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_DestroyQImageEncoderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
