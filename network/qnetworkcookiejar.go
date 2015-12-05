package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNetworkCookieJar struct {
	core.QObject
}

type QNetworkCookieJar_ITF interface {
	core.QObject_ITF
	QNetworkCookieJar_PTR() *QNetworkCookieJar
}

func PointerFromQNetworkCookieJar(ptr QNetworkCookieJar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookieJar_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieJarFromPointer(ptr unsafe.Pointer) *QNetworkCookieJar {
	var n = new(QNetworkCookieJar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkCookieJar_") {
		n.SetObjectName("QNetworkCookieJar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkCookieJar) QNetworkCookieJar_PTR() *QNetworkCookieJar {
	return ptr
}

func NewQNetworkCookieJar(parent core.QObject_ITF) *QNetworkCookieJar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCookieJar::QNetworkCookieJar")
		}
	}()

	return NewQNetworkCookieJarFromPointer(C.QNetworkCookieJar_NewQNetworkCookieJar(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCookieJar::deleteCookie")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCookieJar::insertCookie")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCookieJar::updateCookie")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkCookieJar::~QNetworkCookieJar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
