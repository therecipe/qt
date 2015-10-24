package dbus

//#include "qdbussignature.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusSignature struct {
	ptr unsafe.Pointer
}

type QDBusSignatureITF interface {
	QDBusSignaturePTR() *QDBusSignature
}

func (p *QDBusSignature) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusSignature) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusSignature(ptr QDBusSignatureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusSignaturePTR().Pointer()
	}
	return nil
}

func QDBusSignatureFromPointer(ptr unsafe.Pointer) *QDBusSignature {
	var n = new(QDBusSignature)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusSignature) QDBusSignaturePTR() *QDBusSignature {
	return ptr
}

func NewQDBusSignature() *QDBusSignature {
	return QDBusSignatureFromPointer(unsafe.Pointer(C.QDBusSignature_NewQDBusSignature()))
}

func NewQDBusSignature3(signature core.QLatin1StringITF) *QDBusSignature {
	return QDBusSignatureFromPointer(unsafe.Pointer(C.QDBusSignature_NewQDBusSignature3(C.QtObjectPtr(core.PointerFromQLatin1String(signature)))))
}

func NewQDBusSignature4(signature string) *QDBusSignature {
	return QDBusSignatureFromPointer(unsafe.Pointer(C.QDBusSignature_NewQDBusSignature4(C.CString(signature))))
}

func NewQDBusSignature2(signature string) *QDBusSignature {
	return QDBusSignatureFromPointer(unsafe.Pointer(C.QDBusSignature_NewQDBusSignature2(C.CString(signature))))
}

func (ptr *QDBusSignature) SetSignature(signature string) {
	if ptr.Pointer() != nil {
		C.QDBusSignature_SetSignature(C.QtObjectPtr(ptr.Pointer()), C.CString(signature))
	}
}

func (ptr *QDBusSignature) Signature() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusSignature_Signature(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
