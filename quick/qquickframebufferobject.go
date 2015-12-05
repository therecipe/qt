package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QQuickFramebufferObject struct {
	QQuickItem
}

type QQuickFramebufferObject_ITF interface {
	QQuickItem_ITF
	QQuickFramebufferObject_PTR() *QQuickFramebufferObject
}

func PointerFromQQuickFramebufferObject(ptr QQuickFramebufferObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickFramebufferObject_PTR().Pointer()
	}
	return nil
}

func NewQQuickFramebufferObjectFromPointer(ptr unsafe.Pointer) *QQuickFramebufferObject {
	var n = new(QQuickFramebufferObject)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickFramebufferObject_") {
		n.SetObjectName("QQuickFramebufferObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject {
	return ptr
}

func (ptr *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::setTextureFollowsItemSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_SetTextureFollowsItemSize(ptr.Pointer(), C.int(qt.GoBoolToInt(follows)))
	}
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::textureFollowsItemSize")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_TextureFollowsItemSize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) IsTextureProvider() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::isTextureProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ReleaseResources() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::releaseResources")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(v bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::textureFollowsItemSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureFollowsItemSizeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::textureFollowsItemSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged")
	}
}

//export callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged
func callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(ptrName *C.char, v C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::textureFollowsItemSizeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "textureFollowsItemSizeChanged").(func(bool))(int(v) != 0)
}

func (ptr *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickFramebufferObject::textureProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickFramebufferObject_TextureProvider(ptr.Pointer()))
	}
	return nil
}
