package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::QMediaResource")
		}
	}()

	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource())
}

func NewQMediaResource4(other QMediaResource_ITF) *QMediaResource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::QMediaResource")
		}
	}()

	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource4(PointerFromQMediaResource(other)))
}

func NewQMediaResource3(request network.QNetworkRequest_ITF, mimeType string) *QMediaResource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::QMediaResource")
		}
	}()

	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource3(network.PointerFromQNetworkRequest(request), C.CString(mimeType)))
}

func NewQMediaResource2(url core.QUrl_ITF, mimeType string) *QMediaResource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::QMediaResource")
		}
	}()

	return NewQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource2(core.PointerFromQUrl(url), C.CString(mimeType)))
}

func (ptr *QMediaResource) AudioBitRate() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::audioBitRate")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_AudioBitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) AudioCodec() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::audioCodec")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_AudioCodec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) ChannelCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::channelCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaResource_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaResource) Language() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::language")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_Language(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) MimeType() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::mimeType")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_MimeType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) SampleRate() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::sampleRate")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) SetAudioBitRate(rate int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setAudioBitRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QMediaResource) SetAudioCodec(codec string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setAudioCodec")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QMediaResource) SetChannelCount(channels int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setChannelCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QMediaResource) SetLanguage(language string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setLanguage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetLanguage(ptr.Pointer(), C.CString(language))
	}
}

func (ptr *QMediaResource) SetResolution(resolution core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setResolution")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QMediaResource) SetResolution2(width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setResolution")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QMediaResource) SetSampleRate(sampleRate int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setSampleRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetSampleRate(ptr.Pointer(), C.int(sampleRate))
	}
}

func (ptr *QMediaResource) SetVideoBitRate(rate int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setVideoBitRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QMediaResource) SetVideoCodec(codec string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::setVideoCodec")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QMediaResource) VideoBitRate() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::videoBitRate")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_VideoBitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) VideoCodec() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::videoCodec")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_VideoCodec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) DestroyQMediaResource() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaResource::~QMediaResource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaResource_DestroyQMediaResource(ptr.Pointer())
	}
}
