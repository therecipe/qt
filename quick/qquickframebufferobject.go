package quick

//#include "qquickframebufferobject.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QQuickFramebufferObject struct {
	QQuickItem
}

type QQuickFramebufferObjectITF interface {
	QQuickItemITF
	QQuickFramebufferObjectPTR() *QQuickFramebufferObject
}

func PointerFromQQuickFramebufferObject(ptr QQuickFramebufferObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickFramebufferObjectPTR().Pointer()
	}
	return nil
}

func QQuickFramebufferObjectFromPointer(ptr unsafe.Pointer) *QQuickFramebufferObject {
	var n = new(QQuickFramebufferObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickFramebufferObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickFramebufferObject) QQuickFramebufferObjectPTR() *QQuickFramebufferObject {
	return ptr
}

func (ptr *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_SetTextureFollowsItemSize(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(follows)))
	}
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_TextureFollowsItemSize(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProvider(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ReleaseResources(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(v bool)) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureFollowsItemSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged")
	}
}

//export callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged
func callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(ptrName *C.char, v C.int) {
	qt.GetSignal(C.GoString(ptrName), "textureFollowsItemSizeChanged").(func(bool))(int(v) != 0)
}

func (ptr *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		return QSGTextureProviderFromPointer(unsafe.Pointer(C.QQuickFramebufferObject_TextureProvider(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
