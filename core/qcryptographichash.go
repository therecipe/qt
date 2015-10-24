package core

//#include "qcryptographichash.h"
import "C"
import (
	"unsafe"
)

type QCryptographicHash struct {
	ptr unsafe.Pointer
}

type QCryptographicHashITF interface {
	QCryptographicHashPTR() *QCryptographicHash
}

func (p *QCryptographicHash) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCryptographicHash) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCryptographicHash(ptr QCryptographicHashITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCryptographicHashPTR().Pointer()
	}
	return nil
}

func QCryptographicHashFromPointer(ptr unsafe.Pointer) *QCryptographicHash {
	var n = new(QCryptographicHash)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCryptographicHash) QCryptographicHashPTR() *QCryptographicHash {
	return ptr
}

//QCryptographicHash::Algorithm
type QCryptographicHash__Algorithm int

var (
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
	return QCryptographicHashFromPointer(unsafe.Pointer(C.QCryptographicHash_NewQCryptographicHash(C.int(method))))
}

func (ptr *QCryptographicHash) AddData2(device QIODeviceITF) bool {
	if ptr.Pointer() != nil {
		return C.QCryptographicHash_AddData2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIODevice(device))) != 0
	}
	return false
}

func (ptr *QCryptographicHash) AddData3(data QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QCryptographicHash_AddData3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(data)))
	}
}

func (ptr *QCryptographicHash) AddData(data string, length int) {
	if ptr.Pointer() != nil {
		C.QCryptographicHash_AddData(C.QtObjectPtr(ptr.Pointer()), C.CString(data), C.int(length))
	}
}

func (ptr *QCryptographicHash) Reset() {
	if ptr.Pointer() != nil {
		C.QCryptographicHash_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCryptographicHash) DestroyQCryptographicHash() {
	if ptr.Pointer() != nil {
		C.QCryptographicHash_DestroyQCryptographicHash(C.QtObjectPtr(ptr.Pointer()))
	}
}
