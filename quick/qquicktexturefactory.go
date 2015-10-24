package quick

//#include "qquicktexturefactory.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactoryITF interface {
	core.QObjectITF
	QQuickTextureFactoryPTR() *QQuickTextureFactory
}

func PointerFromQQuickTextureFactory(ptr QQuickTextureFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextureFactoryPTR().Pointer()
	}
	return nil
}

func QQuickTextureFactoryFromPointer(ptr unsafe.Pointer) *QQuickTextureFactory {
	var n = new(QQuickTextureFactory)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickTextureFactory_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickTextureFactory) QQuickTextureFactoryPTR() *QQuickTextureFactory {
	return ptr
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindowITF) *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QQuickTextureFactory_CreateTexture(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQuickWindow(window)))))
	}
	return nil
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {
	if ptr.Pointer() != nil {
		return int(C.QQuickTextureFactory_TextureByteCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactory() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactory(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
