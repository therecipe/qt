package core

//#include "qmessageauthenticationcode.h"
import "C"
import (
	"unsafe"
)

type QMessageAuthenticationCode struct {
	ptr unsafe.Pointer
}

type QMessageAuthenticationCode_ITF interface {
	QMessageAuthenticationCode_PTR() *QMessageAuthenticationCode
}

func (p *QMessageAuthenticationCode) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMessageAuthenticationCode) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMessageAuthenticationCode(ptr QMessageAuthenticationCode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageAuthenticationCode_PTR().Pointer()
	}
	return nil
}

func NewQMessageAuthenticationCodeFromPointer(ptr unsafe.Pointer) *QMessageAuthenticationCode {
	var n = new(QMessageAuthenticationCode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMessageAuthenticationCode) QMessageAuthenticationCode_PTR() *QMessageAuthenticationCode {
	return ptr
}

func NewQMessageAuthenticationCode(method QCryptographicHash__Algorithm, key QByteArray_ITF) *QMessageAuthenticationCode {
	return NewQMessageAuthenticationCodeFromPointer(C.QMessageAuthenticationCode_NewQMessageAuthenticationCode(C.int(method), PointerFromQByteArray(key)))
}

func (ptr *QMessageAuthenticationCode) AddData2(device QIODevice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMessageAuthenticationCode_AddData2(ptr.Pointer(), PointerFromQIODevice(device)) != 0
	}
	return false
}

func (ptr *QMessageAuthenticationCode) AddData3(data QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_AddData3(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QMessageAuthenticationCode) AddData(data string, length int) {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_AddData(ptr.Pointer(), C.CString(data), C.int(length))
	}
}

func QMessageAuthenticationCode_Hash(message QByteArray_ITF, key QByteArray_ITF, method QCryptographicHash__Algorithm) *QByteArray {
	return NewQByteArrayFromPointer(C.QMessageAuthenticationCode_QMessageAuthenticationCode_Hash(PointerFromQByteArray(message), PointerFromQByteArray(key), C.int(method)))
}

func (ptr *QMessageAuthenticationCode) Reset() {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_Reset(ptr.Pointer())
	}
}

func (ptr *QMessageAuthenticationCode) Result() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QMessageAuthenticationCode_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMessageAuthenticationCode) SetKey(key QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_SetKey(ptr.Pointer(), PointerFromQByteArray(key))
	}
}

func (ptr *QMessageAuthenticationCode) DestroyQMessageAuthenticationCode() {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(ptr.Pointer())
	}
}
