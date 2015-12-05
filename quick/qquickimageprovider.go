package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt/qml"
	"log"
	"unsafe"
)

type QQuickImageProvider struct {
	qml.QQmlImageProviderBase
}

type QQuickImageProvider_ITF interface {
	qml.QQmlImageProviderBase_ITF
	QQuickImageProvider_PTR() *QQuickImageProvider
}

func PointerFromQQuickImageProvider(ptr QQuickImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func NewQQuickImageProviderFromPointer(ptr unsafe.Pointer) *QQuickImageProvider {
	var n = new(QQuickImageProvider)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuickImageProvider) QQuickImageProvider_PTR() *QQuickImageProvider {
	return ptr
}

func NewQQuickImageProvider(ty qml.QQmlImageProviderBase__ImageType, flags qml.QQmlImageProviderBase__Flag) *QQuickImageProvider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickImageProvider::QQuickImageProvider")
		}
	}()

	return NewQQuickImageProviderFromPointer(C.QQuickImageProvider_NewQQuickImageProvider(C.int(ty), C.int(flags)))
}

func (ptr *QQuickImageProvider) Flags() qml.QQmlImageProviderBase__Flag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickImageProvider::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__Flag(C.QQuickImageProvider_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickImageProvider) ImageType() qml.QQmlImageProviderBase__ImageType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickImageProvider::imageType")
		}
	}()

	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__ImageType(C.QQuickImageProvider_ImageType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProvider() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickImageProvider::~QQuickImageProvider")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickImageProvider_DestroyQQuickImageProvider(ptr.Pointer())
	}
}
