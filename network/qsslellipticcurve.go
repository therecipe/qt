package network

//#include "qsslellipticcurve.h"
import "C"
import (
	"unsafe"
)

type QSslEllipticCurve struct {
	ptr unsafe.Pointer
}

type QSslEllipticCurveITF interface {
	QSslEllipticCurvePTR() *QSslEllipticCurve
}

func (p *QSslEllipticCurve) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslEllipticCurve) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslEllipticCurve(ptr QSslEllipticCurveITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslEllipticCurvePTR().Pointer()
	}
	return nil
}

func QSslEllipticCurveFromPointer(ptr unsafe.Pointer) *QSslEllipticCurve {
	var n = new(QSslEllipticCurve)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslEllipticCurve) QSslEllipticCurvePTR() *QSslEllipticCurve {
	return ptr
}

func NewQSslEllipticCurve() *QSslEllipticCurve {
	return QSslEllipticCurveFromPointer(unsafe.Pointer(C.QSslEllipticCurve_NewQSslEllipticCurve()))
}

func (ptr *QSslEllipticCurve) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) IsTlsNamedCurve() bool {
	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsTlsNamedCurve(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) LongName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_LongName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslEllipticCurve) ShortName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_ShortName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
