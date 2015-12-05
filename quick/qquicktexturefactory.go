package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactory_ITF interface {
	core.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func PointerFromQQuickTextureFactory(ptr QQuickTextureFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextureFactory_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextureFactoryFromPointer(ptr unsafe.Pointer) *QQuickTextureFactory {
	var n = new(QQuickTextureFactory)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickTextureFactory_") {
		n.SetObjectName("QQuickTextureFactory_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory {
	return ptr
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF) *QSGTexture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickTextureFactory::createTexture")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickTextureFactory_CreateTexture(ptr.Pointer(), PointerFromQQuickWindow(window)))
	}
	return nil
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickTextureFactory::textureByteCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QQuickTextureFactory_TextureByteCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactory() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickTextureFactory::~QQuickTextureFactory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
