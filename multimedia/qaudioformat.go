package multimedia

//#include "qaudioformat.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioFormat struct {
	ptr unsafe.Pointer
}

type QAudioFormat_ITF interface {
	QAudioFormat_PTR() *QAudioFormat
}

func (p *QAudioFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioFormat(ptr QAudioFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioFormat_PTR().Pointer()
	}
	return nil
}

func NewQAudioFormatFromPointer(ptr unsafe.Pointer) *QAudioFormat {
	var n = new(QAudioFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAudioFormat) QAudioFormat_PTR() *QAudioFormat {
	return ptr
}

//QAudioFormat::Endian
type QAudioFormat__Endian int64

const (
	QAudioFormat__BigEndian    = QAudioFormat__Endian(core.QSysInfo__BigEndian)
	QAudioFormat__LittleEndian = QAudioFormat__Endian(core.QSysInfo__LittleEndian)
)

//QAudioFormat::SampleType
type QAudioFormat__SampleType int64

const (
	QAudioFormat__Unknown     = QAudioFormat__SampleType(0)
	QAudioFormat__SignedInt   = QAudioFormat__SampleType(1)
	QAudioFormat__UnSignedInt = QAudioFormat__SampleType(2)
	QAudioFormat__Float       = QAudioFormat__SampleType(3)
)

func NewQAudioFormat() *QAudioFormat {
	return NewQAudioFormatFromPointer(C.QAudioFormat_NewQAudioFormat())
}

func NewQAudioFormat2(other QAudioFormat_ITF) *QAudioFormat {
	return NewQAudioFormatFromPointer(C.QAudioFormat_NewQAudioFormat2(PointerFromQAudioFormat(other)))
}

func (ptr *QAudioFormat) ByteOrder() QAudioFormat__Endian {
	if ptr.Pointer() != nil {
		return QAudioFormat__Endian(C.QAudioFormat_ByteOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) BytesPerFrame() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_BytesPerFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) ChannelCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) Codec() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioFormat_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAudioFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioFormat) SampleRate() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) SampleSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_SampleSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) SampleType() QAudioFormat__SampleType {
	if ptr.Pointer() != nil {
		return QAudioFormat__SampleType(C.QAudioFormat_SampleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) SetByteOrder(byteOrder QAudioFormat__Endian) {
	if ptr.Pointer() != nil {
		C.QAudioFormat_SetByteOrder(ptr.Pointer(), C.int(byteOrder))
	}
}

func (ptr *QAudioFormat) SetChannelCount(channels int) {
	if ptr.Pointer() != nil {
		C.QAudioFormat_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QAudioFormat) SetCodec(codec string) {
	if ptr.Pointer() != nil {
		C.QAudioFormat_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QAudioFormat) SetSampleRate(samplerate int) {
	if ptr.Pointer() != nil {
		C.QAudioFormat_SetSampleRate(ptr.Pointer(), C.int(samplerate))
	}
}

func (ptr *QAudioFormat) SetSampleSize(sampleSize int) {
	if ptr.Pointer() != nil {
		C.QAudioFormat_SetSampleSize(ptr.Pointer(), C.int(sampleSize))
	}
}

func (ptr *QAudioFormat) SetSampleType(sampleType QAudioFormat__SampleType) {
	if ptr.Pointer() != nil {
		C.QAudioFormat_SetSampleType(ptr.Pointer(), C.int(sampleType))
	}
}

func (ptr *QAudioFormat) DestroyQAudioFormat() {
	if ptr.Pointer() != nil {
		C.QAudioFormat_DestroyQAudioFormat(ptr.Pointer())
	}
}
