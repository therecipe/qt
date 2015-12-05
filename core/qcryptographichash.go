package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QCryptographicHash struct {
	ptr unsafe.Pointer
}

type QCryptographicHash_ITF interface {
	QCryptographicHash_PTR() *QCryptographicHash
}

func (p *QCryptographicHash) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCryptographicHash) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCryptographicHash(ptr QCryptographicHash_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCryptographicHash_PTR().Pointer()
	}
	return nil
}

func NewQCryptographicHashFromPointer(ptr unsafe.Pointer) *QCryptographicHash {
	var n = new(QCryptographicHash)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCryptographicHash) QCryptographicHash_PTR() *QCryptographicHash {
	return ptr
}

//QCryptographicHash::Algorithm
type QCryptographicHash__Algorithm int64

const (
	QCryptographicHash__Md4      = QCryptographicHash__Algorithm(0)
	QCryptographicHash__Md5      = QCryptographicHash__Algorithm(1)
	QCryptographicHash__Sha1     = QCryptographicHash__Algorithm(2)
	QCryptographicHash__Sha224   = QCryptographicHash__Algorithm(3)
	QCryptographicHash__Sha256   = QCryptographicHash__Algorithm(4)
	QCryptographicHash__Sha384   = QCryptographicHash__Algorithm(5)
	QCryptographicHash__Sha512   = QCryptographicHash__Algorithm(6)
	QCryptographicHash__Sha3_224 = QCryptographicHash__Algorithm(7)
	QCryptographicHash__Sha3_256 = QCryptographicHash__Algorithm(8)
	QCryptographicHash__Sha3_384 = QCryptographicHash__Algorithm(9)
	QCryptographicHash__Sha3_512 = QCryptographicHash__Algorithm(10)
)

func NewQCryptographicHash(method QCryptographicHash__Algorithm) *QCryptographicHash {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::QCryptographicHash")
		}
	}()

	return NewQCryptographicHashFromPointer(C.QCryptographicHash_NewQCryptographicHash(C.int(method)))
}

func (ptr *QCryptographicHash) AddData2(device QIODevice_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::addData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCryptographicHash_AddData2(ptr.Pointer(), PointerFromQIODevice(device)) != 0
	}
	return false
}

func (ptr *QCryptographicHash) AddData3(data QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::addData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCryptographicHash_AddData3(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QCryptographicHash) AddData(data string, length int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::addData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCryptographicHash_AddData(ptr.Pointer(), C.CString(data), C.int(length))
	}
}

func QCryptographicHash_Hash(data QByteArray_ITF, method QCryptographicHash__Algorithm) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::hash")
		}
	}()

	return NewQByteArrayFromPointer(C.QCryptographicHash_QCryptographicHash_Hash(PointerFromQByteArray(data), C.int(method)))
}

func (ptr *QCryptographicHash) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCryptographicHash_Reset(ptr.Pointer())
	}
}

func (ptr *QCryptographicHash) Result() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::result")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QCryptographicHash_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCryptographicHash) DestroyQCryptographicHash() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCryptographicHash::~QCryptographicHash")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCryptographicHash_DestroyQCryptographicHash(ptr.Pointer())
	}
}
