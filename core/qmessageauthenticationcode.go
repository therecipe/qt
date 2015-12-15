package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QMessageAuthenticationCode::QMessageAuthenticationCode")

	return NewQMessageAuthenticationCodeFromPointer(C.QMessageAuthenticationCode_NewQMessageAuthenticationCode(C.int(method), PointerFromQByteArray(key)))
}

func (ptr *QMessageAuthenticationCode) AddData2(device QIODevice_ITF) bool {
	defer qt.Recovering("QMessageAuthenticationCode::addData")

	if ptr.Pointer() != nil {
		return C.QMessageAuthenticationCode_AddData2(ptr.Pointer(), PointerFromQIODevice(device)) != 0
	}
	return false
}

func (ptr *QMessageAuthenticationCode) AddData3(data QByteArray_ITF) {
	defer qt.Recovering("QMessageAuthenticationCode::addData")

	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_AddData3(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QMessageAuthenticationCode) AddData(data string, length int) {
	defer qt.Recovering("QMessageAuthenticationCode::addData")

	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_AddData(ptr.Pointer(), C.CString(data), C.int(length))
	}
}

func QMessageAuthenticationCode_Hash(message QByteArray_ITF, key QByteArray_ITF, method QCryptographicHash__Algorithm) *QByteArray {
	defer qt.Recovering("QMessageAuthenticationCode::hash")

	return NewQByteArrayFromPointer(C.QMessageAuthenticationCode_QMessageAuthenticationCode_Hash(PointerFromQByteArray(message), PointerFromQByteArray(key), C.int(method)))
}

func (ptr *QMessageAuthenticationCode) Reset() {
	defer qt.Recovering("QMessageAuthenticationCode::reset")

	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_Reset(ptr.Pointer())
	}
}

func (ptr *QMessageAuthenticationCode) Result() *QByteArray {
	defer qt.Recovering("QMessageAuthenticationCode::result")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QMessageAuthenticationCode_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMessageAuthenticationCode) SetKey(key QByteArray_ITF) {
	defer qt.Recovering("QMessageAuthenticationCode::setKey")

	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_SetKey(ptr.Pointer(), PointerFromQByteArray(key))
	}
}

func (ptr *QMessageAuthenticationCode) DestroyQMessageAuthenticationCode() {
	defer qt.Recovering("QMessageAuthenticationCode::~QMessageAuthenticationCode")

	if ptr.Pointer() != nil {
		C.QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(ptr.Pointer())
	}
}
