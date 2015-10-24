package network

//#include "qnetworkcookiejar.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkCookieJar struct {
	core.QObject
}

type QNetworkCookieJarITF interface {
	core.QObjectITF
	QNetworkCookieJarPTR() *QNetworkCookieJar
}

func PointerFromQNetworkCookieJar(ptr QNetworkCookieJarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookieJarPTR().Pointer()
	}
	return nil
}

func QNetworkCookieJarFromPointer(ptr unsafe.Pointer) *QNetworkCookieJar {
	var n = new(QNetworkCookieJar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNetworkCookieJar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkCookieJar) QNetworkCookieJarPTR() *QNetworkCookieJar {
	return ptr
}

func NewQNetworkCookieJar(parent core.QObjectITF) *QNetworkCookieJar {
	return QNetworkCookieJarFromPointer(unsafe.Pointer(C.QNetworkCookieJar_NewQNetworkCookieJar(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookieITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookie(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookieITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookie(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookieITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookie(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCookie(cookie))) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {
	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
