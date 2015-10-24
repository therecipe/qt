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

type QMediaResourceITF interface {
	QMediaResourcePTR() *QMediaResource
}

func (p *QMediaResource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaResource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaResource(ptr QMediaResourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaResourcePTR().Pointer()
	}
	return nil
}

func QMediaResourceFromPointer(ptr unsafe.Pointer) *QMediaResource {
	var n = new(QMediaResource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaResource) QMediaResourcePTR() *QMediaResource {
	return ptr
}

func NewQMediaResource() *QMediaResource {
	return QMediaResourceFromPointer(unsafe.Pointer(C.QMediaResource_NewQMediaResource()))
}

func NewQMediaResource4(other QMediaResourceITF) *QMediaResource {
	return QMediaResourceFromPointer(unsafe.Pointer(C.QMediaResource_NewQMediaResource4(C.QtObjectPtr(PointerFromQMediaResource(other)))))
}

func NewQMediaResource3(request network.QNetworkRequestITF, mimeType string) *QMediaResource {
	return QMediaResourceFromPointer(unsafe.Pointer(C.QMediaResource_NewQMediaResource3(C.QtObjectPtr(network.PointerFromQNetworkRequest(request)), C.CString(mimeType))))
}

func NewQMediaResource2(url string, mimeType string) *QMediaResource {
	return QMediaResourceFromPointer(unsafe.Pointer(C.QMediaResource_NewQMediaResource2(C.CString(url), C.CString(mimeType))))
}

func (ptr *QMediaResource) AudioBitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_AudioBitRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaResource) AudioCodec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_AudioCodec(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaResource) ChannelCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_ChannelCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaResource) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QMediaResource_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaResource) Language() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_Language(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaResource) MimeType() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_MimeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaResource) SampleRate() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_SampleRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaResource) SetAudioBitRate(rate int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioBitRate(C.QtObjectPtr(ptr.Pointer()), C.int(rate))
	}
}

func (ptr *QMediaResource) SetAudioCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioCodec(C.QtObjectPtr(ptr.Pointer()), C.CString(codec))
	}
}

func (ptr *QMediaResource) SetChannelCount(channels int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetChannelCount(C.QtObjectPtr(ptr.Pointer()), C.int(channels))
	}
}

func (ptr *QMediaResource) SetLanguage(language string) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetLanguage(C.QtObjectPtr(ptr.Pointer()), C.CString(language))
	}
}

func (ptr *QMediaResource) SetResolution(resolution core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(resolution)))
	}
}

func (ptr *QMediaResource) SetResolution2(width int, height int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution2(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height))
	}
}

func (ptr *QMediaResource) SetSampleRate(sampleRate int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetSampleRate(C.QtObjectPtr(ptr.Pointer()), C.int(sampleRate))
	}
}

func (ptr *QMediaResource) SetVideoBitRate(rate int) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoBitRate(C.QtObjectPtr(ptr.Pointer()), C.int(rate))
	}
}

func (ptr *QMediaResource) SetVideoCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoCodec(C.QtObjectPtr(ptr.Pointer()), C.CString(codec))
	}
}

func (ptr *QMediaResource) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaResource) VideoBitRate() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaResource_VideoBitRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaResource) VideoCodec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_VideoCodec(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaResource) DestroyQMediaResource() {
	if ptr.Pointer() != nil {
		C.QMediaResource_DestroyQMediaResource(C.QtObjectPtr(ptr.Pointer()))
	}
}
