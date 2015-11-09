package multimedia

//#include "qaudiobuffer.h"
import "C"
import (
	"unsafe"
)

type QAudioBuffer struct {
	ptr unsafe.Pointer
}

type QAudioBuffer_ITF interface {
	QAudioBuffer_PTR() *QAudioBuffer
}

func (p *QAudioBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioBuffer(ptr QAudioBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioBuffer_PTR().Pointer()
	}
	return nil
}

func NewQAudioBufferFromPointer(ptr unsafe.Pointer) *QAudioBuffer {
	var n = new(QAudioBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAudioBuffer) QAudioBuffer_PTR() *QAudioBuffer {
	return ptr
}

func NewQAudioBuffer() *QAudioBuffer {
	return NewQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer())
}

func NewQAudioBuffer3(other QAudioBuffer_ITF) *QAudioBuffer {
	return NewQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer3(PointerFromQAudioBuffer(other)))
}

func (ptr *QAudioBuffer) ByteCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_ByteCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) ConstData() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_ConstData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) Data2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_Data2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) Data() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) FrameCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_FrameCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAudioBuffer_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioBuffer) SampleCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_SampleCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) DestroyQAudioBuffer() {
	if ptr.Pointer() != nil {
		C.QAudioBuffer_DestroyQAudioBuffer(ptr.Pointer())
	}
}
