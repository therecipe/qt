package quick

//#include "qquickimageprovider.h"
import "C"
import (
	"github.com/therecipe/qt/qml"
	"unsafe"
)

type QQuickImageProvider struct {
	qml.QQmlImageProviderBase
}

type QQuickImageProviderITF interface {
	qml.QQmlImageProviderBaseITF
	QQuickImageProviderPTR() *QQuickImageProvider
}

func PointerFromQQuickImageProvider(ptr QQuickImageProviderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageProviderPTR().Pointer()
	}
	return nil
}

func QQuickImageProviderFromPointer(ptr unsafe.Pointer) *QQuickImageProvider {
	var n = new(QQuickImageProvider)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuickImageProvider) QQuickImageProviderPTR() *QQuickImageProvider {
	return ptr
}

func NewQQuickImageProvider(ty qml.QQmlImageProviderBase__ImageType, flags qml.QQmlImageProviderBase__Flag) *QQuickImageProvider {
	return QQuickImageProviderFromPointer(unsafe.Pointer(C.QQuickImageProvider_NewQQuickImageProvider(C.int(ty), C.int(flags))))
}

func (ptr *QQuickImageProvider) Flags() qml.QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__Flag(C.QQuickImageProvider_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickImageProvider) ImageType() qml.QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__ImageType(C.QQuickImageProvider_ImageType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProvider() {
	if ptr.Pointer() != nil {
		C.QQuickImageProvider_DestroyQQuickImageProvider(C.QtObjectPtr(ptr.Pointer()))
	}
}
