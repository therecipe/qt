package multimedia

//#include "qaudiodeviceinfo.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QAudioDeviceInfo struct {
	ptr unsafe.Pointer
}

type QAudioDeviceInfo_ITF interface {
	QAudioDeviceInfo_PTR() *QAudioDeviceInfo
}

func (p *QAudioDeviceInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioDeviceInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioDeviceInfo(ptr QAudioDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDeviceInfo_PTR().Pointer()
	}
	return nil
}

func NewQAudioDeviceInfoFromPointer(ptr unsafe.Pointer) *QAudioDeviceInfo {
	var n = new(QAudioDeviceInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAudioDeviceInfo) QAudioDeviceInfo_PTR() *QAudioDeviceInfo {
	return ptr
}

func NewQAudioDeviceInfo() *QAudioDeviceInfo {
	return NewQAudioDeviceInfoFromPointer(C.QAudioDeviceInfo_NewQAudioDeviceInfo())
}

func NewQAudioDeviceInfo2(other QAudioDeviceInfo_ITF) *QAudioDeviceInfo {
	return NewQAudioDeviceInfoFromPointer(C.QAudioDeviceInfo_NewQAudioDeviceInfo2(PointerFromQAudioDeviceInfo(other)))
}

func (ptr *QAudioDeviceInfo) DeviceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDeviceInfo_DeviceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDeviceInfo) IsFormatSupported(settings QAudioFormat_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAudioDeviceInfo_IsFormatSupported(ptr.Pointer(), PointerFromQAudioFormat(settings)) != 0
	}
	return false
}

func (ptr *QAudioDeviceInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QAudioDeviceInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDeviceInfo) SupportedCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioDeviceInfo_SupportedCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QAudioDeviceInfo) DestroyQAudioDeviceInfo() {
	if ptr.Pointer() != nil {
		C.QAudioDeviceInfo_DestroyQAudioDeviceInfo(ptr.Pointer())
	}
}
