package multimedia

//#include "qmediaresource.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QMediaResource struct {
	ptr unsafe.Pointer
}

type QMediaResource_ITF interface {
	QMediaResource_PTR() *QMediaResource
}

func (p *QMediaResource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaResource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaResource(ptr QMediaResource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaResource_PTR().Pointer()
	}
	return nil
}

func NewQMediaResourceFromPointer(ptr unsafe.Pointer) *QMediaResource {
	var n = new(QMediaResource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaResource) QMediaResource_PTR() *QMediaResource {
	return ptr
}

func NewQMediaResource() *QMediaResource {
	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource())
}

func NewQMediaResource4(other QMediaResource_ITF) *QMediaResource {
	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource4(PointerFromQMediaResource(other)))
}

func NewQMediaResource3(request network.QNetworkRequest_ITF, mimeType string) *QMediaResource {
	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource3(network.PointerFromQNetworkRequest(request), C.CString(mimeType)))
}

func NewQMediaResource2(url core.QUrl_ITF, mimeType string) *QMediaResource {
	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource2(core.PointerFromQUrl(url), C.CString(mimeType)))
}

func (ptr *QMediaResource) AudioBitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_AudioBitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) AudioCodec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_AudioCodec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) ChannelCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QMediaResource_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaResource) Language() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_Language(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) MimeType() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_MimeType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) SampleRate() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) SetAudioBitRate(rate int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QMediaResource) SetAudioCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QMediaResource) SetChannelCount(channels int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QMediaResource) SetLanguage(language string) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetLanguage(ptr.Pointer(), C.CString(language))
	}
}

func (ptr *QMediaResource) SetResolution(resolution core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QMediaResource) SetResolution2(width int, height int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QMediaResource) SetSampleRate(sampleRate int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetSampleRate(ptr.Pointer(), C.int(sampleRate))
	}
}

func (ptr *QMediaResource) SetVideoBitRate(rate int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QMediaResource) SetVideoCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QMediaResource) VideoBitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_VideoBitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) VideoCodec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_VideoCodec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) DestroyQMediaResource() {
	if ptr.Pointer() != nil {
		C.QMediaResource_DestroyQMediaResource(ptr.Pointer())
	}
}
