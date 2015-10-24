package quick

//#include "qsgtextureprovider.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProviderITF interface {
	core.QObjectITF
	QSGTextureProviderPTR() *QSGTextureProvider
}

func PointerFromQSGTextureProvider(ptr QSGTextureProviderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureProviderPTR().Pointer()
	}
	return nil
}

func QSGTextureProviderFromPointer(ptr unsafe.Pointer) *QSGTextureProvider {
	var n = new(QSGTextureProvider)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSGTextureProvider_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGTextureProvider) QSGTextureProviderPTR() *QSGTextureProvider {
	return ptr
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QSGTextureProvider_Texture(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGTextureProvider) ConnectTextureChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectTextureChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textureChanged", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTextureChanged() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectTextureChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textureChanged")
	}
}

//export callbackQSGTextureProviderTextureChanged
func callbackQSGTextureProviderTextureChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textureChanged").(func())()
}
