package network

//#include "network.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSslEllipticCurve struct {
	ptr unsafe.Pointer
}

type QSslEllipticCurve_ITF interface {
	QSslEllipticCurve_PTR() *QSslEllipticCurve
}

func (p *QSslEllipticCurve) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslEllipticCurve) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslEllipticCurve(ptr QSslEllipticCurve_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslEllipticCurve_PTR().Pointer()
	}
	return nil
}

func NewQSslEllipticCurveFromPointer(ptr unsafe.Pointer) *QSslEllipticCurve {
	var n = new(QSslEllipticCurve)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslEllipticCurve) QSslEllipticCurve_PTR() *QSslEllipticCurve {
	return ptr
}

func NewQSslEllipticCurve() *QSslEllipticCurve {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslEllipticCurve::QSslEllipticCurve")
		}
	}()

	return NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_NewQSslEllipticCurve())
}

func (ptr *QSslEllipticCurve) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslEllipticCurve::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) IsTlsNamedCurve() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslEllipticCurve::isTlsNamedCurve")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsTlsNamedCurve(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) LongName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslEllipticCurve::longName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_LongName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslEllipticCurve) ShortName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslEllipticCurve::shortName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_ShortName(ptr.Pointer()))
	}
	return ""
}
