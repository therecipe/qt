package multimedia

//#include "qaudiobuffer.h"
import "C"
import (
	"unsafe"
)

type QAudioBuffer struct {
	ptr unsafe.Pointer
}

type QAudioBufferITF interface {
	QAudioBufferPTR() *QAudioBuffer
}

func (p *QAudioBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioBuffer(ptr QAudioBufferITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioBufferPTR().Pointer()
	}
	return nil
}

func QAudioBufferFromPointer(ptr unsafe.Pointer) *QAudioBuffer {
	var n = new(QAudioBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAudioBuffer) QAudioBufferPTR() *QAudioBuffer {
	return ptr
}

func NewQAudioBuffer() *QAudioBuffer {
	return QAudioBufferFromPointer(unsafe.Pointer(C.QAudioBuffer_NewQAudioBuffer()))
}

func NewQAudioBuffer3(other QAudioBufferITF) *QAudioBuffer {
	return QAudioBufferFromPointer(unsafe.Pointer(C.QAudioBuffer_NewQAudioBuffer3(C.QtObjectPtr(PointerFromQAudioBuffer(other)))))
}

func (ptr *QAudioBuffer) ByteCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_ByteCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioBuffer) ConstData() {
	if ptr.Pointer() != nil {
		C.QAudioBuffer_ConstData(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioBuffer) Data2() {
	if ptr.Pointer() != nil {
		C.QAudioBuffer_Data2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioBuffer) Data() {
	if ptr.Pointer() != nil {
		C.QAudioBuffer_Data(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioBuffer) FrameCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_FrameCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioBuffer) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAudioBuffer_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAudioBuffer) SampleCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_SampleCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioBuffer) DestroyQAudioBuffer() {
	if ptr.Pointer() != nil {
		C.QAudioBuffer_DestroyQAudioBuffer(C.QtObjectPtr(ptr.Pointer()))
	}
}
