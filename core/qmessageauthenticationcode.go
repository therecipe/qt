package core

//#include "qmessageauthenticationcode.h"
import "C"
import (
	"unsafe"
)

type QMessageAuthenticationCode struct {
	ptr unsafe.Pointer
}

type QMessageAuthenticationCodeITF interface {
	QMessageAuthenticationCodePTR() *QMessageAuthenticationCode
}

func (p *QMessageAuthenticationCode) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMessageAuthenticationCode) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMessageAuthenticationCode(ptr QMessageAuthenticationCodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageAuthenticationCodePTR().Pointer()
	}
	return nil
}

func QMessageAuthenticationCodeFromPointer(ptr unsafe.Pointer) *QMessageAuthenticationCode {
	var n = new(QMessageAuthenticationCode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMessageAuthenticationCode) QMessageAuthenticationCodePTR() *QMessageAuthenticationCode {
	return ptr
}

func NewQMessageAuthenticationCode(method QCryptographicHash__Algorithm, key QByteArrayITF) *QMessageAuthenticationCode {
	return QMessageAuthenticationCodeFromPointer(unsafe.Pointer(C.QMessageAuthenticationCode_NewQMessageAuthenticationCode(C.int(method), C.QtObjectPtr(PointerFromQByteArray(key)))))
}

func (ptr *QMessageAuthenticationCode) AddData2(device QIODeviceITF) bool {
	if ptr.Pointer() != nil {
		return C.QMessageAuthenticationCode_AddData2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIODevice(device))) != 0
	}
	return false
}

func (ptr *QMessageAuthenticationCode) AddData3(data QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_AddData3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(data)))
	}
}

func (ptr *QMessageAuthenticationCode) AddData(data string, length int) {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_AddData(C.QtObjectPtr(ptr.Pointer()), C.CString(data), C.int(length))
	}
}

func (ptr *QMessageAuthenticationCode) Reset() {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMessageAuthenticationCode) SetKey(key QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_SetKey(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(key)))
	}
}

func (ptr *QMessageAuthenticationCode) DestroyQMessageAuthenticationCode() {
	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(C.QtObjectPtr(ptr.Pointer()))
	}
}
