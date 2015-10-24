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

type QAudioDeviceInfoITF interface {
	QAudioDeviceInfoPTR() *QAudioDeviceInfo
}

func (p *QAudioDeviceInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioDeviceInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioDeviceInfo(ptr QAudioDeviceInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDeviceInfoPTR().Pointer()
	}
	return nil
}

func QAudioDeviceInfoFromPointer(ptr unsafe.Pointer) *QAudioDeviceInfo {
	var n = new(QAudioDeviceInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAudioDeviceInfo) QAudioDeviceInfoPTR() *QAudioDeviceInfo {
	return ptr
}

func NewQAudioDeviceInfo() *QAudioDeviceInfo {
	return QAudioDeviceInfoFromPointer(unsafe.Pointer(C.QAudioDeviceInfo_NewQAudioDeviceInfo()))
}

func NewQAudioDeviceInfo2(other QAudioDeviceInfoITF) *QAudioDeviceInfo {
	return QAudioDeviceInfoFromPointer(unsafe.Pointer(C.QAudioDeviceInfo_NewQAudioDeviceInfo2(C.QtObjectPtr(PointerFromQAudioDeviceInfo(other)))))
}

func (ptr *QAudioDeviceInfo) DeviceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDeviceInfo_DeviceName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAudioDeviceInfo) IsFormatSupported(settings QAudioFormatITF) bool {
	if ptr.Pointer() != nil {
		return C.QAudioDeviceInfo_IsFormatSupported(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAudioFormat(settings))) != 0
	}
	return false
}

func (ptr *QAudioDeviceInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QAudioDeviceInfo_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAudioDeviceInfo) SupportedCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioDeviceInfo_SupportedCodecs(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QAudioDeviceInfo) DestroyQAudioDeviceInfo() {
	if ptr.Pointer() != nil {
		C.QAudioDeviceInfo_DestroyQAudioDeviceInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
