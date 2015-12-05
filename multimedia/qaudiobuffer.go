package multimedia

//#include "multimedia.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::QAudioBuffer")
		}
	}()

	return NewQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer())
}

func NewQAudioBuffer3(other QAudioBuffer_ITF) *QAudioBuffer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::QAudioBuffer")
		}
	}()

	return NewQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer3(PointerFromQAudioBuffer(other)))
}

func (ptr *QAudioBuffer) ByteCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::byteCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_ByteCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) ConstData() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::constData")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_ConstData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) Data2() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::data")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_Data2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) Data() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::data")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) FrameCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::frameCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_FrameCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAudioBuffer_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioBuffer) SampleCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::sampleCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_SampleCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) DestroyQAudioBuffer() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioBuffer::~QAudioBuffer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioBuffer_DestroyQAudioBuffer(ptr.Pointer())
	}
}
