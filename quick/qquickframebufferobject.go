package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
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
		n.SetObjectName("QQuickFramebufferObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject {
	return ptr
}

func (ptr *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {
	defer qt.Recovering("QQuickFramebufferObject::setTextureFollowsItemSize")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_SetTextureFollowsItemSize(ptr.Pointer(), C.int(qt.GoBoolToInt(follows)))
	}
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	defer qt.Recovering("QQuickFramebufferObject::textureFollowsItemSize")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_TextureFollowsItemSize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) IsTextureProvider() bool {
	defer qt.Recovering("QQuickFramebufferObject::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ConnectReleaseResources(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::releaseResources")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "releaseResources", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectReleaseResources() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::releaseResources")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "releaseResources")
	}
}

//export callbackQQuickFramebufferObjectReleaseResources
func callbackQQuickFramebufferObjectReleaseResources(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::releaseResources")

	var signal = qt.GetSignal(C.GoString(ptrName), "releaseResources")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(v bool)) {
	defer qt.Recovering("connect QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureFollowsItemSizeChanged() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged")
	}
}

//export callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged
func callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(ptrName *C.char, v C.int) {
	defer qt.Recovering("callback QQuickFramebufferObject::textureFollowsItemSizeChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "textureFollowsItemSizeChanged")
	if signal != nil {
		signal.(func(bool))(int(v) != 0)
	}

}

func (ptr *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickFramebufferObject::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickFramebufferObject_TextureProvider(ptr.Pointer()))
	}
	return nil
}
