// +build !minimal

package quick

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtQuick_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtQuick_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QQuickAsyncImageProvider struct {
	QQuickImageProvider
}

type QQuickAsyncImageProvider_ITF interface {
	QQuickImageProvider_ITF
	QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider
}

func (ptr *QQuickAsyncImageProvider) QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider {
	return ptr
}

func (ptr *QQuickAsyncImageProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickAsyncImageProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickImageProvider_PTR().SetPointer(p)
	}
}

func PointerFromQQuickAsyncImageProvider(ptr QQuickAsyncImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickAsyncImageProvider_PTR().Pointer()
	}
	return nil
}

func NewQQuickAsyncImageProviderFromPointer(ptr unsafe.Pointer) (n *QQuickAsyncImageProvider) {
	n = new(QQuickAsyncImageProvider)
	n.SetPointer(ptr)
	return
}
func NewQQuickAsyncImageProvider() *QQuickAsyncImageProvider {
	return NewQQuickAsyncImageProviderFromPointer(C.QQuickAsyncImageProvider_NewQQuickAsyncImageProvider())
}

//export callbackQQuickAsyncImageProvider_RequestImageResponse
func callbackQQuickAsyncImageProvider_RequestImageResponse(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, requestedSize unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "requestImageResponse"); signal != nil {
		return PointerFromQQuickImageResponse(signal.(func(string, *core.QSize) *QQuickImageResponse)(cGoUnpackString(id), core.NewQSizeFromPointer(requestedSize)))
	}

	return PointerFromQQuickImageResponse(NewQQuickImageResponse())
}

func (ptr *QQuickAsyncImageProvider) ConnectRequestImageResponse(f func(id string, requestedSize *core.QSize) *QQuickImageResponse) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestImageResponse"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestImageResponse", func(id string, requestedSize *core.QSize) *QQuickImageResponse {
				signal.(func(string, *core.QSize) *QQuickImageResponse)(id, requestedSize)
				return f(id, requestedSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestImageResponse", f)
		}
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectRequestImageResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestImageResponse")
	}
}

func (ptr *QQuickAsyncImageProvider) RequestImageResponse(id string, requestedSize core.QSize_ITF) *QQuickImageResponse {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := NewQQuickImageResponseFromPointer(C.QQuickAsyncImageProvider_RequestImageResponse(ptr.Pointer(), C.struct_QtQuick_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider
func callbackQQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickAsyncImageProvider"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickAsyncImageProviderFromPointer(ptr).DestroyQQuickAsyncImageProviderDefault()
	}
}

func (ptr *QQuickAsyncImageProvider) ConnectDestroyQQuickAsyncImageProvider(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickAsyncImageProvider"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickAsyncImageProvider", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickAsyncImageProvider", f)
		}
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectDestroyQQuickAsyncImageProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickAsyncImageProvider")
	}
}

func (ptr *QQuickAsyncImageProvider) DestroyQQuickAsyncImageProvider() {
	if ptr.Pointer() != nil {
		C.QQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickAsyncImageProvider) DestroyQQuickAsyncImageProviderDefault() {
	if ptr.Pointer() != nil {
		C.QQuickAsyncImageProvider_DestroyQQuickAsyncImageProviderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QQuickFramebufferObject struct {
	QQuickItem
}

type QQuickFramebufferObject_ITF interface {
	QQuickItem_ITF
	QQuickFramebufferObject_PTR() *QQuickFramebufferObject
}

func (ptr *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject {
	return ptr
}

func (ptr *QQuickFramebufferObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickFramebufferObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickItem_PTR().SetPointer(p)
	}
}

func PointerFromQQuickFramebufferObject(ptr QQuickFramebufferObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickFramebufferObject_PTR().Pointer()
	}
	return nil
}

func NewQQuickFramebufferObjectFromPointer(ptr unsafe.Pointer) (n *QQuickFramebufferObject) {
	n = new(QQuickFramebufferObject)
	n.SetPointer(ptr)
	return
}

//export callbackQQuickFramebufferObject_MirrorVerticallyChanged
func callbackQQuickFramebufferObject_MirrorVerticallyChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "mirrorVerticallyChanged"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	}

}

func (ptr *QQuickFramebufferObject) ConnectMirrorVerticallyChanged(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "mirrorVerticallyChanged") {
			C.QQuickFramebufferObject_ConnectMirrorVerticallyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "mirrorVerticallyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mirrorVerticallyChanged", func(vbo bool) {
				signal.(func(bool))(vbo)
				f(vbo)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mirrorVerticallyChanged", f)
		}
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMirrorVerticallyChanged() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectMirrorVerticallyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "mirrorVerticallyChanged")
	}
}

func (ptr *QQuickFramebufferObject) MirrorVerticallyChanged(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MirrorVerticallyChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickFramebufferObject) SetMirrorVertically(enable bool) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_SetMirrorVertically(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_SetTextureFollowsItemSize(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(follows))))
	}
}

//export callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged
func callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "textureFollowsItemSizeChanged"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	}

}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textureFollowsItemSizeChanged") {
			C.QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textureFollowsItemSizeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureFollowsItemSizeChanged", func(vbo bool) {
				signal.(func(bool))(vbo)
				f(vbo)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureFollowsItemSizeChanged", f)
		}
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureFollowsItemSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "textureFollowsItemSizeChanged")
	}
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSizeChanged(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TextureFollowsItemSizeChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickFramebufferObject) MirrorVertically() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickFramebufferObject_MirrorVertically(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickFramebufferObject_TextureFollowsItemSize(ptr.Pointer())) != 0
	}
	return false
}

type QQuickImageProvider struct {
	qml.QQmlImageProviderBase
}

type QQuickImageProvider_ITF interface {
	qml.QQmlImageProviderBase_ITF
	QQuickImageProvider_PTR() *QQuickImageProvider
}

func (ptr *QQuickImageProvider) QQuickImageProvider_PTR() *QQuickImageProvider {
	return ptr
}

func (ptr *QQuickImageProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickImageProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQmlImageProviderBase_PTR().SetPointer(p)
	}
}

func PointerFromQQuickImageProvider(ptr QQuickImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func NewQQuickImageProviderFromPointer(ptr unsafe.Pointer) (n *QQuickImageProvider) {
	n = new(QQuickImageProvider)
	n.SetPointer(ptr)
	return
}

//export callbackQQuickImageProvider_RequestImage
func callbackQQuickImageProvider_RequestImage(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "requestImage"); signal != nil {
		return gui.PointerFromQImage(signal.(func(string, *core.QSize, *core.QSize) *gui.QImage)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return gui.PointerFromQImage(NewQQuickImageProviderFromPointer(ptr).RequestImageDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickImageProvider) ConnectRequestImage(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QImage) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestImage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestImage", func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QImage {
				signal.(func(string, *core.QSize, *core.QSize) *gui.QImage)(id, size, requestedSize)
				return f(id, size, requestedSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestImage", f)
		}
	}
}

func (ptr *QQuickImageProvider) DisconnectRequestImage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestImage")
	}
}

func (ptr *QQuickImageProvider) RequestImage(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := gui.NewQImageFromPointer(C.QQuickImageProvider_RequestImage(ptr.Pointer(), C.struct_QtQuick_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageProvider) RequestImageDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := gui.NewQImageFromPointer(C.QQuickImageProvider_RequestImageDefault(ptr.Pointer(), C.struct_QtQuick_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

//export callbackQQuickImageProvider_RequestPixmap
func callbackQQuickImageProvider_RequestPixmap(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "requestPixmap"); signal != nil {
		return gui.PointerFromQPixmap(signal.(func(string, *core.QSize, *core.QSize) *gui.QPixmap)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return gui.PointerFromQPixmap(NewQQuickImageProviderFromPointer(ptr).RequestPixmapDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickImageProvider) ConnectRequestPixmap(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QPixmap) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestPixmap"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestPixmap", func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QPixmap {
				signal.(func(string, *core.QSize, *core.QSize) *gui.QPixmap)(id, size, requestedSize)
				return f(id, size, requestedSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestPixmap", f)
		}
	}
}

func (ptr *QQuickImageProvider) DisconnectRequestPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestPixmap")
	}
}

func (ptr *QQuickImageProvider) RequestPixmap(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := gui.NewQPixmapFromPointer(C.QQuickImageProvider_RequestPixmap(ptr.Pointer(), C.struct_QtQuick_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageProvider) RequestPixmapDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := gui.NewQPixmapFromPointer(C.QQuickImageProvider_RequestPixmapDefault(ptr.Pointer(), C.struct_QtQuick_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func NewQQuickImageProvider(ty qml.QQmlImageProviderBase__ImageType, flags qml.QQmlImageProviderBase__Flag) *QQuickImageProvider {
	return NewQQuickImageProviderFromPointer(C.QQuickImageProvider_NewQQuickImageProvider(C.longlong(ty), C.longlong(flags)))
}

//export callbackQQuickImageProvider_RequestTexture
func callbackQQuickImageProvider_RequestTexture(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "requestTexture"); signal != nil {
		return PointerFromQQuickTextureFactory(signal.(func(string, *core.QSize, *core.QSize) *QQuickTextureFactory)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return PointerFromQQuickTextureFactory(NewQQuickImageProviderFromPointer(ptr).RequestTextureDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickImageProvider) ConnectRequestTexture(f func(id string, size *core.QSize, requestedSize *core.QSize) *QQuickTextureFactory) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestTexture", func(id string, size *core.QSize, requestedSize *core.QSize) *QQuickTextureFactory {
				signal.(func(string, *core.QSize, *core.QSize) *QQuickTextureFactory)(id, size, requestedSize)
				return f(id, size, requestedSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestTexture", f)
		}
	}
}

func (ptr *QQuickImageProvider) DisconnectRequestTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestTexture")
	}
}

func (ptr *QQuickImageProvider) RequestTexture(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := NewQQuickTextureFactoryFromPointer(C.QQuickImageProvider_RequestTexture(ptr.Pointer(), C.struct_QtQuick_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageProvider) RequestTextureDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := NewQQuickTextureFactoryFromPointer(C.QQuickImageProvider_RequestTextureDefault(ptr.Pointer(), C.struct_QtQuick_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickImageProvider_DestroyQQuickImageProvider
func callbackQQuickImageProvider_DestroyQQuickImageProvider(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickImageProvider"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageProviderFromPointer(ptr).DestroyQQuickImageProviderDefault()
	}
}

func (ptr *QQuickImageProvider) ConnectDestroyQQuickImageProvider(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickImageProvider"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickImageProvider", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickImageProvider", f)
		}
	}
}

func (ptr *QQuickImageProvider) DisconnectDestroyQQuickImageProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickImageProvider")
	}
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProvider() {
	if ptr.Pointer() != nil {
		C.QQuickImageProvider_DestroyQQuickImageProvider(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProviderDefault() {
	if ptr.Pointer() != nil {
		C.QQuickImageProvider_DestroyQQuickImageProviderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickImageProvider_Flags
func callbackQQuickImageProvider_Flags(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func() qml.QQmlImageProviderBase__Flag)())
	}

	return C.longlong(NewQQuickImageProviderFromPointer(ptr).FlagsDefault())
}

func (ptr *QQuickImageProvider) ConnectFlags(f func() qml.QQmlImageProviderBase__Flag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "flags"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "flags", func() qml.QQmlImageProviderBase__Flag {
				signal.(func() qml.QQmlImageProviderBase__Flag)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "flags", f)
		}
	}
}

func (ptr *QQuickImageProvider) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "flags")
	}
}

func (ptr *QQuickImageProvider) Flags() qml.QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__Flag(C.QQuickImageProvider_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickImageProvider) FlagsDefault() qml.QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__Flag(C.QQuickImageProvider_FlagsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickImageProvider_ImageType
func callbackQQuickImageProvider_ImageType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "imageType"); signal != nil {
		return C.longlong(signal.(func() qml.QQmlImageProviderBase__ImageType)())
	}

	return C.longlong(NewQQuickImageProviderFromPointer(ptr).ImageTypeDefault())
}

func (ptr *QQuickImageProvider) ConnectImageType(f func() qml.QQmlImageProviderBase__ImageType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "imageType"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "imageType", func() qml.QQmlImageProviderBase__ImageType {
				signal.(func() qml.QQmlImageProviderBase__ImageType)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "imageType", f)
		}
	}
}

func (ptr *QQuickImageProvider) DisconnectImageType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "imageType")
	}
}

func (ptr *QQuickImageProvider) ImageType() qml.QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__ImageType(C.QQuickImageProvider_ImageType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickImageProvider) ImageTypeDefault() qml.QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__ImageType(C.QQuickImageProvider_ImageTypeDefault(ptr.Pointer()))
	}
	return 0
}

type QQuickImageResponse struct {
	core.QObject
}

type QQuickImageResponse_ITF interface {
	core.QObject_ITF
	QQuickImageResponse_PTR() *QQuickImageResponse
}

func (ptr *QQuickImageResponse) QQuickImageResponse_PTR() *QQuickImageResponse {
	return ptr
}

func (ptr *QQuickImageResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickImageResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickImageResponse(ptr QQuickImageResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageResponse_PTR().Pointer()
	}
	return nil
}

func NewQQuickImageResponseFromPointer(ptr unsafe.Pointer) (n *QQuickImageResponse) {
	n = new(QQuickImageResponse)
	n.SetPointer(ptr)
	return
}
func NewQQuickImageResponse() *QQuickImageResponse {
	tmpValue := NewQQuickImageResponseFromPointer(C.QQuickImageResponse_NewQQuickImageResponse())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQuickImageResponse_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickImageResponse_QQuickImageResponse_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickImageResponse) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickImageResponse_QQuickImageResponse_Tr(sC, cC, C.int(int32(n))))
}

func QQuickImageResponse_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickImageResponse_QQuickImageResponse_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickImageResponse) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickImageResponse_QQuickImageResponse_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQuickImageResponse_Cancel
func callbackQQuickImageResponse_Cancel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cancel"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).CancelDefault()
	}
}

func (ptr *QQuickImageResponse) ConnectCancel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cancel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "cancel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cancel", f)
		}
	}
}

func (ptr *QQuickImageResponse) DisconnectCancel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cancel")
	}
}

func (ptr *QQuickImageResponse) Cancel() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_Cancel(ptr.Pointer())
	}
}

func (ptr *QQuickImageResponse) CancelDefault() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_CancelDefault(ptr.Pointer())
	}
}

//export callbackQQuickImageResponse_Finished
func callbackQQuickImageResponse_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickImageResponse) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QQuickImageResponse_ConnectFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finished", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", f)
		}
	}
}

func (ptr *QQuickImageResponse) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QQuickImageResponse) Finished() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_Finished(ptr.Pointer())
	}
}

//export callbackQQuickImageResponse_DestroyQQuickImageResponse
func callbackQQuickImageResponse_DestroyQQuickImageResponse(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickImageResponse"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).DestroyQQuickImageResponseDefault()
	}
}

func (ptr *QQuickImageResponse) ConnectDestroyQQuickImageResponse(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickImageResponse"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickImageResponse", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickImageResponse", f)
		}
	}
}

func (ptr *QQuickImageResponse) DisconnectDestroyQQuickImageResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickImageResponse")
	}
}

func (ptr *QQuickImageResponse) DestroyQQuickImageResponse() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DestroyQQuickImageResponse(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickImageResponse) DestroyQQuickImageResponseDefault() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DestroyQQuickImageResponseDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickImageResponse_TextureFactory
func callbackQQuickImageResponse_TextureFactory(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "textureFactory"); signal != nil {
		return PointerFromQQuickTextureFactory(signal.(func() *QQuickTextureFactory)())
	}

	return PointerFromQQuickTextureFactory(NewQQuickTextureFactory())
}

func (ptr *QQuickImageResponse) ConnectTextureFactory(f func() *QQuickTextureFactory) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "textureFactory"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureFactory", func() *QQuickTextureFactory {
				signal.(func() *QQuickTextureFactory)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureFactory", f)
		}
	}
}

func (ptr *QQuickImageResponse) DisconnectTextureFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "textureFactory")
	}
}

func (ptr *QQuickImageResponse) TextureFactory() *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickTextureFactoryFromPointer(C.QQuickImageResponse_TextureFactory(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickImageResponse_ErrorString
func callbackQQuickImageResponse_ErrorString(ptr unsafe.Pointer) C.struct_QtQuick_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtQuick_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQQuickImageResponseFromPointer(ptr).ErrorStringDefault()
	return C.struct_QtQuick_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QQuickImageResponse) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QQuickImageResponse) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QQuickImageResponse) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickImageResponse_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickImageResponse) ErrorStringDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickImageResponse_ErrorStringDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickImageResponse_MetaObject
func callbackQQuickImageResponse_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickImageResponseFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickImageResponse) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickImageResponse_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickImageResponse) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickImageResponse___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageResponse) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickImageResponse) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickImageResponse___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickImageResponse) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickImageResponse___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageResponse) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickImageResponse) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickImageResponse___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickImageResponse) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickImageResponse___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageResponse) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickImageResponse) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickImageResponse___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickImageResponse) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickImageResponse___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageResponse) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickImageResponse) __findChildren_newList() unsafe.Pointer {
	return C.QQuickImageResponse___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickImageResponse) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickImageResponse___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageResponse) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickImageResponse) __children_newList() unsafe.Pointer {
	return C.QQuickImageResponse___children_newList(ptr.Pointer())
}

//export callbackQQuickImageResponse_Event
func callbackQQuickImageResponse_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickImageResponseFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickImageResponse) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickImageResponse_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickImageResponse_EventFilter
func callbackQQuickImageResponse_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickImageResponseFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickImageResponse) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickImageResponse_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickImageResponse_ChildEvent
func callbackQQuickImageResponse_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickImageResponse_ConnectNotify
func callbackQQuickImageResponse_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickImageResponseFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickImageResponse) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickImageResponse_CustomEvent
func callbackQQuickImageResponse_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickImageResponse_DeleteLater
func callbackQQuickImageResponse_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickImageResponse) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickImageResponse_Destroyed
func callbackQQuickImageResponse_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickImageResponse_DisconnectNotify
func callbackQQuickImageResponse_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickImageResponseFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickImageResponse) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickImageResponse_ObjectNameChanged
func callbackQQuickImageResponse_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickImageResponse_TimerEvent
func callbackQQuickImageResponse_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickItem struct {
	core.QObject
	qml.QQmlParserStatus
}

type QQuickItem_ITF interface {
	core.QObject_ITF
	qml.QQmlParserStatus_ITF
	QQuickItem_PTR() *QQuickItem
}

func (ptr *QQuickItem) QQuickItem_PTR() *QQuickItem {
	return ptr
}

func (ptr *QQuickItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QQmlParserStatus_PTR().SetPointer(p)
	}
}

func PointerFromQQuickItem(ptr QQuickItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemFromPointer(ptr unsafe.Pointer) (n *QQuickItem) {
	n = new(QQuickItem)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQuickItem__Flag
//QQuickItem::Flag
type QQuickItem__Flag int64

const (
	QQuickItem__ItemClipsChildrenToShape QQuickItem__Flag = QQuickItem__Flag(0x01)
	QQuickItem__ItemAcceptsInputMethod   QQuickItem__Flag = QQuickItem__Flag(0x02)
	QQuickItem__ItemIsFocusScope         QQuickItem__Flag = QQuickItem__Flag(0x04)
	QQuickItem__ItemHasContents          QQuickItem__Flag = QQuickItem__Flag(0x08)
	QQuickItem__ItemAcceptsDrops         QQuickItem__Flag = QQuickItem__Flag(0x10)
)

//go:generate stringer -type=QQuickItem__ItemChange
//QQuickItem::ItemChange
type QQuickItem__ItemChange int64

const (
	QQuickItem__ItemChildAddedChange           QQuickItem__ItemChange = QQuickItem__ItemChange(0)
	QQuickItem__ItemChildRemovedChange         QQuickItem__ItemChange = QQuickItem__ItemChange(1)
	QQuickItem__ItemSceneChange                QQuickItem__ItemChange = QQuickItem__ItemChange(2)
	QQuickItem__ItemVisibleHasChanged          QQuickItem__ItemChange = QQuickItem__ItemChange(3)
	QQuickItem__ItemParentHasChanged           QQuickItem__ItemChange = QQuickItem__ItemChange(4)
	QQuickItem__ItemOpacityHasChanged          QQuickItem__ItemChange = QQuickItem__ItemChange(5)
	QQuickItem__ItemActiveFocusHasChanged      QQuickItem__ItemChange = QQuickItem__ItemChange(6)
	QQuickItem__ItemRotationHasChanged         QQuickItem__ItemChange = QQuickItem__ItemChange(7)
	QQuickItem__ItemAntialiasingHasChanged     QQuickItem__ItemChange = QQuickItem__ItemChange(8)
	QQuickItem__ItemDevicePixelRatioHasChanged QQuickItem__ItemChange = QQuickItem__ItemChange(9)
	QQuickItem__ItemEnabledHasChanged          QQuickItem__ItemChange = QQuickItem__ItemChange(10)
)

//go:generate stringer -type=QQuickItem__TransformOrigin
//QQuickItem::TransformOrigin
type QQuickItem__TransformOrigin int64

const (
	QQuickItem__TopLeft     QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(0)
	QQuickItem__Top         QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(1)
	QQuickItem__TopRight    QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(2)
	QQuickItem__Left        QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(3)
	QQuickItem__Center      QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(4)
	QQuickItem__Right       QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(5)
	QQuickItem__BottomLeft  QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(6)
	QQuickItem__Bottom      QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(7)
	QQuickItem__BottomRight QQuickItem__TransformOrigin = QQuickItem__TransformOrigin(8)
)

func (ptr *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickItem_NextItemInFocusChain(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward)))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQQuickItem(parent QQuickItem_ITF) *QQuickItem {
	tmpValue := NewQQuickItemFromPointer(C.QQuickItem_NewQQuickItem(PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickItem) ChildrenRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QQuickItem_ChildrenRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func QQuickItem_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItem_QQuickItem_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickItem) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItem_QQuickItem_Tr(sC, cC, C.int(int32(n))))
}

func QQuickItem_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItem_QQuickItem_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickItem) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItem_QQuickItem_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQuickItem_ChildMouseEventFilter
func callbackQQuickItem_ChildMouseEventFilter(ptr unsafe.Pointer, item unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "childMouseEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QQuickItem, *core.QEvent) bool)(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).ChildMouseEventFilterDefault(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickItem) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "childMouseEventFilter"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "childMouseEventFilter", func(item *QQuickItem, event *core.QEvent) bool {
				signal.(func(*QQuickItem, *core.QEvent) bool)(item, event)
				return f(item, event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "childMouseEventFilter", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectChildMouseEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "childMouseEventFilter")
	}
}

func (ptr *QQuickItem) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QQuickItem) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_ChildMouseEventFilterDefault(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickItem_Event
func callbackQQuickItem_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QQuickItem) ConnectEvent(f func(ev *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "event"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "event", func(ev *core.QEvent) bool {
				signal.(func(*core.QEvent) bool)(ev)
				return f(ev)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "event", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "event")
	}
}

func (ptr *QQuickItem) Event(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_Event(ptr.Pointer(), core.PointerFromQEvent(ev))) != 0
	}
	return false
}

func (ptr *QQuickItem) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackQQuickItem_ClassBegin
func callbackQQuickItem_ClassBegin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "classBegin"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *QQuickItem) ConnectClassBegin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "classBegin"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "classBegin", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "classBegin", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectClassBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "classBegin")
	}
}

func (ptr *QQuickItem) ClassBegin() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ClassBeginDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ClassBeginDefault(ptr.Pointer())
	}
}

//export callbackQQuickItem_ComponentComplete
func callbackQQuickItem_ComponentComplete(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "componentComplete"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *QQuickItem) ConnectComponentComplete(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "componentComplete"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "componentComplete", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "componentComplete", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectComponentComplete() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "componentComplete")
	}
}

func (ptr *QQuickItem) ComponentComplete() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ComponentComplete(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ComponentCompleteDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ComponentCompleteDefault(ptr.Pointer())
	}
}

//export callbackQQuickItem_ContainmentMaskChanged
func callbackQQuickItem_ContainmentMaskChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "containmentMaskChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItem) ConnectContainmentMaskChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "containmentMaskChanged") {
			C.QQuickItem_ConnectContainmentMaskChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "containmentMaskChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "containmentMaskChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "containmentMaskChanged", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectContainmentMaskChanged() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectContainmentMaskChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "containmentMaskChanged")
	}
}

func (ptr *QQuickItem) ContainmentMaskChanged() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ContainmentMaskChanged(ptr.Pointer())
	}
}

//export callbackQQuickItem_DragEnterEvent
func callbackQQuickItem_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dragEnterEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dragEnterEvent", func(event *gui.QDragEnterEvent) {
				signal.(func(*gui.QDragEnterEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dragEnterEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dragEnterEvent")
	}
}

func (ptr *QQuickItem) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickItem) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQQuickItem_DragLeaveEvent
func callbackQQuickItem_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dragLeaveEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dragLeaveEvent", func(event *gui.QDragLeaveEvent) {
				signal.(func(*gui.QDragLeaveEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dragLeaveEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dragLeaveEvent")
	}
}

func (ptr *QQuickItem) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickItem) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQQuickItem_DragMoveEvent
func callbackQQuickItem_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dragMoveEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dragMoveEvent", func(event *gui.QDragMoveEvent) {
				signal.(func(*gui.QDragMoveEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dragMoveEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dragMoveEvent")
	}
}

func (ptr *QQuickItem) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickItem) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQQuickItem_DropEvent
func callbackQQuickItem_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dropEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dropEvent", func(event *gui.QDropEvent) {
				signal.(func(*gui.QDropEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dropEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dropEvent")
	}
}

func (ptr *QQuickItem) DropEvent(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickItem) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQQuickItem_FocusInEvent
func callbackQQuickItem_FocusInEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewQQuickItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *QQuickItem) ConnectFocusInEvent(f func(vqf *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "focusInEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "focusInEvent", func(vqf *gui.QFocusEvent) {
				signal.(func(*gui.QFocusEvent))(vqf)
				f(vqf)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "focusInEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "focusInEvent")
	}
}

func (ptr *QQuickItem) FocusInEvent(vqf gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(vqf))
	}
}

func (ptr *QQuickItem) FocusInEventDefault(vqf gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(vqf))
	}
}

//export callbackQQuickItem_FocusOutEvent
func callbackQQuickItem_FocusOutEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewQQuickItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *QQuickItem) ConnectFocusOutEvent(f func(vqf *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "focusOutEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "focusOutEvent", func(vqf *gui.QFocusEvent) {
				signal.(func(*gui.QFocusEvent))(vqf)
				f(vqf)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "focusOutEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "focusOutEvent")
	}
}

func (ptr *QQuickItem) FocusOutEvent(vqf gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(vqf))
	}
}

func (ptr *QQuickItem) FocusOutEventDefault(vqf gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(vqf))
	}
}

func (ptr *QQuickItem) ForceActiveFocus() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ForceActiveFocus2(reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus2(ptr.Pointer(), C.longlong(reason))
	}
}

//export callbackQQuickItem_GeometryChanged
func callbackQQuickItem_GeometryChanged(ptr unsafe.Pointer, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickItemFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *QQuickItem) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "geometryChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "geometryChanged", func(newGeometry *core.QRectF, oldGeometry *core.QRectF) {
				signal.(func(*core.QRectF, *core.QRectF))(newGeometry, oldGeometry)
				f(newGeometry, oldGeometry)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometryChanged", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectGeometryChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "geometryChanged")
	}
}

func (ptr *QQuickItem) GeometryChanged(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_GeometryChanged(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickItem) GeometryChangedDefault(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_GeometryChangedDefault(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickItem) GrabMouse() {
	if ptr.Pointer() != nil {
		C.QQuickItem_GrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) GrabTouchPoints(ids []int) {
	if ptr.Pointer() != nil {
		C.QQuickItem_GrabTouchPoints(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQQuickItemFromPointer(NewQQuickItemFromPointer(nil).__grabTouchPoints_ids_newList())
			for _, v := range ids {
				tmpList.__grabTouchPoints_ids_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQQuickItem_HoverEnterEvent
func callbackQQuickItem_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoverEnterEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hoverEnterEvent", func(event *gui.QHoverEvent) {
				signal.(func(*gui.QHoverEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoverEnterEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectHoverEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoverEnterEvent")
	}
}

func (ptr *QQuickItem) HoverEnterEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_HoverEnterEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) HoverEnterEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_HoverEnterEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickItem_HoverLeaveEvent
func callbackQQuickItem_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoverLeaveEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hoverLeaveEvent", func(event *gui.QHoverEvent) {
				signal.(func(*gui.QHoverEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoverLeaveEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectHoverLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoverLeaveEvent")
	}
}

func (ptr *QQuickItem) HoverLeaveEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_HoverLeaveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) HoverLeaveEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_HoverLeaveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickItem_HoverMoveEvent
func callbackQQuickItem_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoverMoveEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hoverMoveEvent", func(event *gui.QHoverEvent) {
				signal.(func(*gui.QHoverEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoverMoveEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectHoverMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoverMoveEvent")
	}
}

func (ptr *QQuickItem) HoverMoveEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_HoverMoveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) HoverMoveEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_HoverMoveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickItem_InputMethodEvent
func callbackQQuickItem_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "inputMethodEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodEvent", func(event *gui.QInputMethodEvent) {
				signal.(func(*gui.QInputMethodEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "inputMethodEvent")
	}
}

func (ptr *QQuickItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQQuickItem_KeyPressEvent
func callbackQQuickItem_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyPressEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", func(event *gui.QKeyEvent) {
				signal.(func(*gui.QKeyEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyPressEvent")
	}
}

func (ptr *QQuickItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQQuickItem_KeyReleaseEvent
func callbackQQuickItem_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyReleaseEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", func(event *gui.QKeyEvent) {
				signal.(func(*gui.QKeyEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyReleaseEvent")
	}
}

func (ptr *QQuickItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQQuickItem_MouseDoubleClickEvent
func callbackQQuickItem_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseDoubleClickEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", func(event *gui.QMouseEvent) {
				signal.(func(*gui.QMouseEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseDoubleClickEvent")
	}
}

func (ptr *QQuickItem) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickItem_MouseMoveEvent
func callbackQQuickItem_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseMoveEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mouseMoveEvent", func(event *gui.QMouseEvent) {
				signal.(func(*gui.QMouseEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseMoveEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseMoveEvent")
	}
}

func (ptr *QQuickItem) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickItem_MousePressEvent
func callbackQQuickItem_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mousePressEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", func(event *gui.QMouseEvent) {
				signal.(func(*gui.QMouseEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mousePressEvent")
	}
}

func (ptr *QQuickItem) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickItem_MouseReleaseEvent
func callbackQQuickItem_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseReleaseEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", func(event *gui.QMouseEvent) {
				signal.(func(*gui.QMouseEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseReleaseEvent")
	}
}

func (ptr *QQuickItem) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickItem_MouseUngrabEvent
func callbackQQuickItem_MouseUngrabEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *QQuickItem) ConnectMouseUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseUngrabEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mouseUngrabEvent", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseUngrabEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectMouseUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseUngrabEvent")
	}
}

func (ptr *QQuickItem) MouseUngrabEvent() {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickItem) MouseUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_MouseUngrabEventDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Polish() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Polish(ptr.Pointer())
	}
}

//export callbackQQuickItem_ReleaseResources
func callbackQQuickItem_ReleaseResources(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickItem) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "releaseResources"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "releaseResources", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "releaseResources", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "releaseResources")
	}
}

func (ptr *QQuickItem) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ReleaseResourcesDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetAntialiasing() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetAntialiasing(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetHeight() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetHeight(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetWidth() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetWidth(ptr.Pointer())
	}
}

func (ptr *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptHoverEvents(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQuickItem) SetAcceptTouchEvents(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptTouchEvents(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQuickItem) SetAcceptedMouseButtons(buttons core.Qt__MouseButton) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptedMouseButtons(ptr.Pointer(), C.longlong(buttons))
	}
}

func (ptr *QQuickItem) SetActiveFocusOnTab(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetActiveFocusOnTab(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetAntialiasing(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAntialiasing(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetBaselineOffset(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetBaselineOffset(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetClip(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetClip(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetContainmentMask(mask core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetContainmentMask(ptr.Pointer(), core.PointerFromQObject(mask))
	}
}

func (ptr *QQuickItem) SetCursor(cursor gui.QCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetCursor(ptr.Pointer(), gui.PointerFromQCursor(cursor))
	}
}

func (ptr *QQuickItem) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFiltersChildMouseEvents(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(filter))))
	}
}

func (ptr *QQuickItem) SetFlag(flag QQuickItem__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlag(ptr.Pointer(), C.longlong(flag), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQuickItem) SetFlags(flags QQuickItem__Flag) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlags(ptr.Pointer(), C.longlong(flags))
	}
}

func (ptr *QQuickItem) SetFocus(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetFocus2(focus bool, reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus2(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(focus))), C.longlong(reason))
	}
}

func (ptr *QQuickItem) SetHeight(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetHeight(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetImplicitHeight(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitHeight(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetImplicitWidth(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitWidth(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetKeepMouseGrab(keep bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepMouseGrab(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(keep))))
	}
}

func (ptr *QQuickItem) SetKeepTouchGrab(keep bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepTouchGrab(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(keep))))
	}
}

func (ptr *QQuickItem) SetOpacity(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetOpacity(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetParentItem(parent QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetParentItem(ptr.Pointer(), PointerFromQQuickItem(parent))
	}
}

func (ptr *QQuickItem) SetRotation(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetRotation(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetScale(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetScale(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetSize(size core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QQuickItem) SetSmooth(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetSmooth(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetState(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QQuickItem_SetState(ptr.Pointer(), C.struct_QtQuick_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QQuickItem) SetTransformOrigin(vqq QQuickItem__TransformOrigin) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetTransformOrigin(ptr.Pointer(), C.longlong(vqq))
	}
}

func (ptr *QQuickItem) SetVisible(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetWidth(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetWidth(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetX(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetX(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetY(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetY(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetZ(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetZ(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) StackAfter(sibling QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_StackAfter(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

func (ptr *QQuickItem) StackBefore(sibling QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_StackBefore(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

//export callbackQQuickItem_TouchEvent
func callbackQQuickItem_TouchEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "touchEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "touchEvent", func(event *gui.QTouchEvent) {
				signal.(func(*gui.QTouchEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "touchEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "touchEvent")
	}
}

func (ptr *QQuickItem) TouchEvent(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickItem) TouchEventDefault(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

//export callbackQQuickItem_TouchUngrabEvent
func callbackQQuickItem_TouchUngrabEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *QQuickItem) ConnectTouchUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "touchUngrabEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "touchUngrabEvent", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "touchUngrabEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectTouchUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "touchUngrabEvent")
	}
}

func (ptr *QQuickItem) TouchUngrabEvent() {
	if ptr.Pointer() != nil {
		C.QQuickItem_TouchUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickItem) TouchUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_TouchUngrabEventDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UngrabMouse() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UngrabTouchPoints() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabTouchPoints(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UnsetCursor() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UnsetCursor(ptr.Pointer())
	}
}

//export callbackQQuickItem_Update
func callbackQQuickItem_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickItem) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "update"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "update", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "update", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "update")
	}
}

func (ptr *QQuickItem) Update() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Update(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UpdateDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UpdateInputMethod(queries core.Qt__InputMethodQuery) {
	if ptr.Pointer() != nil {
		C.QQuickItem_UpdateInputMethod(ptr.Pointer(), C.longlong(queries))
	}
}

//export callbackQQuickItem_UpdatePolish
func callbackQQuickItem_UpdatePolish(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *QQuickItem) ConnectUpdatePolish(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updatePolish"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updatePolish", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updatePolish", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectUpdatePolish() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updatePolish")
	}
}

func (ptr *QQuickItem) UpdatePolish() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UpdatePolish(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UpdatePolishDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UpdatePolishDefault(ptr.Pointer())
	}
}

//export callbackQQuickItem_WheelEvent
func callbackQQuickItem_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "wheelEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "wheelEvent", func(event *gui.QWheelEvent) {
				signal.(func(*gui.QWheelEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "wheelEvent", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "wheelEvent")
	}
}

func (ptr *QQuickItem) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickItem) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQQuickItem_WindowChanged
func callbackQQuickItem_WindowChanged(ptr unsafe.Pointer, window unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowChanged"); signal != nil {
		signal.(func(*QQuickWindow))(NewQQuickWindowFromPointer(window))
	}

}

func (ptr *QQuickItem) ConnectWindowChanged(f func(window *QQuickWindow)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "windowChanged") {
			C.QQuickItem_ConnectWindowChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "windowChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "windowChanged", func(window *QQuickWindow) {
				signal.(func(*QQuickWindow))(window)
				f(window)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "windowChanged", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectWindowChanged() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "windowChanged")
	}
}

func (ptr *QQuickItem) WindowChanged(window QQuickWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_WindowChanged(ptr.Pointer(), PointerFromQQuickWindow(window))
	}
}

//export callbackQQuickItem_DestroyQQuickItem
func callbackQQuickItem_DestroyQQuickItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickItem"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).DestroyQQuickItemDefault()
	}
}

func (ptr *QQuickItem) ConnectDestroyQQuickItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickItem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickItem", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickItem", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectDestroyQQuickItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickItem")
	}
}

func (ptr *QQuickItem) DestroyQQuickItem() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItem(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickItem) DestroyQQuickItemDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickItem) Cursor() *gui.QCursor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQCursorFromPointer(C.QQuickItem_Cursor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QCursor).DestroyQCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) ChildItems() []*QQuickItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQuick_PackedList) []*QQuickItem {
			out := make([]*QQuickItem, int(l.len))
			tmpList := NewQQuickItemFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__childItems_atList(i)
			}
			return out
		}(C.QQuickItem_ChildItems(ptr.Pointer()))
	}
	return make([]*QQuickItem, 0)
}

func (ptr *QQuickItem) ContainmentMask() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItem_ContainmentMask(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapFromGlobal(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QQuickItem_MapFromGlobal(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapFromItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QQuickItem_MapFromItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapFromScene(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QQuickItem_MapFromScene(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapToGlobal(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QQuickItem_MapToGlobal(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapToItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QQuickItem_MapToItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapToScene(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QQuickItem_MapToScene(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) ChildAt(x float64, y float64) *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickItem_ChildAt(ptr.Pointer(), C.double(x), C.double(y)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) ParentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickItem_ParentItem(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) ScopedFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickItem_ScopedFocusItem(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) Flags() QQuickItem__Flag {
	if ptr.Pointer() != nil {
		return QQuickItem__Flag(C.QQuickItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) TransformOrigin() QQuickItem__TransformOrigin {
	if ptr.Pointer() != nil {
		return QQuickItem__TransformOrigin(C.QQuickItem_TransformOrigin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Window() *QQuickWindow {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickWindowFromPointer(C.QQuickItem_Window(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectFromItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QQuickItem_MapRectFromItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectFromScene(rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QQuickItem_MapRectFromScene(ptr.Pointer(), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectToItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QQuickItem_MapRectToItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectToScene(rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QQuickItem_MapRectToScene(ptr.Pointer(), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQQuickItem_TextureProvider
func callbackQQuickItem_TextureProvider(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "textureProvider"); signal != nil {
		return PointerFromQSGTextureProvider(signal.(func() *QSGTextureProvider)())
	}

	return PointerFromQSGTextureProvider(NewQQuickItemFromPointer(ptr).TextureProviderDefault())
}

func (ptr *QQuickItem) ConnectTextureProvider(f func() *QSGTextureProvider) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "textureProvider"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureProvider", func() *QSGTextureProvider {
				signal.(func() *QSGTextureProvider)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureProvider", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "textureProvider")
	}
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProvider(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) TextureProviderDefault() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProviderDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) Size() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QQuickItem_Size(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) State() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickItem_State(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickItem_InputMethodQuery
func callbackQQuickItem_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "inputMethodQuery"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodQuery", func(query core.Qt__InputMethodQuery) *core.QVariant {
				signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(query)
				return f(query)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodQuery", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "inputMethodQuery")
	}
}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QQuickItem_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QQuickItem_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) AcceptedMouseButtons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QQuickItem_AcceptedMouseButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) AcceptHoverEvents() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_AcceptHoverEvents(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) AcceptTouchEvents() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_AcceptTouchEvents(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) ActiveFocusOnTab() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_ActiveFocusOnTab(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) Antialiasing() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_Antialiasing(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) Clip() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_Clip(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickItem_Contains
func callbackQQuickItem_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QQuickItem) ConnectContains(f func(point *core.QPointF) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contains"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contains", func(point *core.QPointF) bool {
				signal.(func(*core.QPointF) bool)(point)
				return f(point)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contains", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectContains() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contains")
	}
}

func (ptr *QQuickItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QQuickItem) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QQuickItem) FiltersChildMouseEvents() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_FiltersChildMouseEvents(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) HasActiveFocus() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_HasActiveFocus(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) HasFocus() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_HasFocus(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) HeightValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_HeightValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsAncestorOf(child QQuickItem_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_IsAncestorOf(ptr.Pointer(), PointerFromQQuickItem(child))) != 0
	}
	return false
}

func (ptr *QQuickItem) IsComponentComplete() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_IsComponentComplete(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_IsEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsFocusScope() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_IsFocusScope(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickItem_IsTextureProvider
func callbackQQuickItem_IsTextureProvider(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isTextureProvider"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).IsTextureProviderDefault())))
}

func (ptr *QQuickItem) ConnectIsTextureProvider(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isTextureProvider"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isTextureProvider", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isTextureProvider", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectIsTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isTextureProvider")
	}
}

func (ptr *QQuickItem) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_IsTextureProvider(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsTextureProviderDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_IsTextureProviderDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepMouseGrab() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_KeepMouseGrab(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepTouchGrab() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_KeepTouchGrab(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) Smooth() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_Smooth(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) WidthValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_WidthValid(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickItem_MetaObject
func callbackQQuickItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickItem) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metaObject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "metaObject", func() *core.QMetaObject {
				signal.(func() *core.QMetaObject)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metaObject", f)
		}
	}
}

func (ptr *QQuickItem) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metaObject")
	}
}

func (ptr *QQuickItem) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) BaselineOffset() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_BaselineOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ImplicitHeight() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ImplicitWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Rotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Rotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Scale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Scale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) __grabTouchPoints_ids_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickItem___grabTouchPoints_ids_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QQuickItem) __grabTouchPoints_ids_setList(i int) {
	if ptr.Pointer() != nil {
		C.QQuickItem___grabTouchPoints_ids_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QQuickItem) __grabTouchPoints_ids_newList() unsafe.Pointer {
	return C.QQuickItem___grabTouchPoints_ids_newList(ptr.Pointer())
}

func (ptr *QQuickItem) __childItems_atList(i int) *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickItem___childItems_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) __childItems_setList(i QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem___childItems_setList(ptr.Pointer(), PointerFromQQuickItem(i))
	}
}

func (ptr *QQuickItem) __childItems_newList() unsafe.Pointer {
	return C.QQuickItem___childItems_newList(ptr.Pointer())
}

func (ptr *QQuickItem) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickItem___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickItem) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickItem___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickItem) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItem___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItem) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickItem___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickItem) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItem___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItem) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickItem___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickItem) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItem___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItem) __findChildren_newList() unsafe.Pointer {
	return C.QQuickItem___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickItem) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItem___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItem) __children_newList() unsafe.Pointer {
	return C.QQuickItem___children_newList(ptr.Pointer())
}

//export callbackQQuickItem_EventFilter
func callbackQQuickItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QQuickItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickItem_ChildEvent
func callbackQQuickItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickItem) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickItem_ConnectNotify
func callbackQQuickItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItem_CustomEvent
func callbackQQuickItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickItem) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickItem) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickItem_DeleteLater
func callbackQQuickItem_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickItem) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickItem_Destroyed
func callbackQQuickItem_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickItem_DisconnectNotify
func callbackQQuickItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItem_ObjectNameChanged
func callbackQQuickItem_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickItem_TimerEvent
func callbackQQuickItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickItem) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResult_ITF interface {
	core.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return ptr
}

func (ptr *QQuickItemGrabResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickItemGrabResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickItemGrabResult(ptr QQuickItemGrabResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemGrabResult_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemGrabResultFromPointer(ptr unsafe.Pointer) (n *QQuickItemGrabResult) {
	n = new(QQuickItemGrabResult)
	n.SetPointer(ptr)
	return
}
func QQuickItemGrabResult_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItemGrabResult_QQuickItemGrabResult_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickItemGrabResult) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItemGrabResult_QQuickItemGrabResult_Tr(sC, cC, C.int(int32(n))))
}

func QQuickItemGrabResult_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItemGrabResult_QQuickItemGrabResult_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickItemGrabResult) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickItemGrabResult_QQuickItemGrabResult_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQuickItemGrabResult_Ready
func callbackQQuickItemGrabResult_Ready(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ready"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "ready") {
			C.QQuickItemGrabResult_ConnectReady(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "ready"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "ready", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ready", f)
		}
	}
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "ready")
	}
}

func (ptr *QQuickItemGrabResult) Ready() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_Ready(ptr.Pointer())
	}
}

func (ptr *QQuickItemGrabResult) Image() *gui.QImage {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQImageFromPointer(C.QQuickItemGrabResult_Image(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQuickItemGrabResult_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		return int8(C.QQuickItemGrabResult_SaveToFile(ptr.Pointer(), C.struct_QtQuick_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_MetaObject
func callbackQQuickItemGrabResult_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickItemGrabResultFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickItemGrabResult) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItemGrabResult_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItemGrabResult) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickItemGrabResult___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickItemGrabResult) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickItemGrabResult___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickItemGrabResult) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItemGrabResult___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItemGrabResult) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickItemGrabResult___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickItemGrabResult) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItemGrabResult___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItemGrabResult) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickItemGrabResult___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickItemGrabResult) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItemGrabResult___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItemGrabResult) __findChildren_newList() unsafe.Pointer {
	return C.QQuickItemGrabResult___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickItemGrabResult) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickItemGrabResult___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickItemGrabResult) __children_newList() unsafe.Pointer {
	return C.QQuickItemGrabResult___children_newList(ptr.Pointer())
}

//export callbackQQuickItemGrabResult_Event
func callbackQQuickItemGrabResult_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemGrabResultFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickItemGrabResult) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItemGrabResult_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_EventFilter
func callbackQQuickItemGrabResult_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemGrabResultFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickItemGrabResult) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickItemGrabResult_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_ChildEvent
func callbackQQuickItemGrabResult_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickItemGrabResult_ConnectNotify
func callbackQQuickItemGrabResult_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItemGrabResult) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItemGrabResult_CustomEvent
func callbackQQuickItemGrabResult_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickItemGrabResult_DeleteLater
func callbackQQuickItemGrabResult_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickItemGrabResult) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickItemGrabResult_Destroyed
func callbackQQuickItemGrabResult_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickItemGrabResult_DisconnectNotify
func callbackQQuickItemGrabResult_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItemGrabResult) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItemGrabResult_ObjectNameChanged
func callbackQQuickItemGrabResult_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickItemGrabResult_TimerEvent
func callbackQQuickItemGrabResult_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickPaintedItem struct {
	QQuickItem
}

type QQuickPaintedItem_ITF interface {
	QQuickItem_ITF
	QQuickPaintedItem_PTR() *QQuickPaintedItem
}

func (ptr *QQuickPaintedItem) QQuickPaintedItem_PTR() *QQuickPaintedItem {
	return ptr
}

func (ptr *QQuickPaintedItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickPaintedItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickItem_PTR().SetPointer(p)
	}
}

func PointerFromQQuickPaintedItem(ptr QQuickPaintedItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickPaintedItem_PTR().Pointer()
	}
	return nil
}

func NewQQuickPaintedItemFromPointer(ptr unsafe.Pointer) (n *QQuickPaintedItem) {
	n = new(QQuickPaintedItem)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQuickPaintedItem__PerformanceHint
//QQuickPaintedItem::PerformanceHint
type QQuickPaintedItem__PerformanceHint int64

const (
	QQuickPaintedItem__FastFBOResizing QQuickPaintedItem__PerformanceHint = QQuickPaintedItem__PerformanceHint(0x1)
)

//go:generate stringer -type=QQuickPaintedItem__RenderTarget
//QQuickPaintedItem::RenderTarget
type QQuickPaintedItem__RenderTarget int64

const (
	QQuickPaintedItem__Image                      QQuickPaintedItem__RenderTarget = QQuickPaintedItem__RenderTarget(0)
	QQuickPaintedItem__FramebufferObject          QQuickPaintedItem__RenderTarget = QQuickPaintedItem__RenderTarget(1)
	QQuickPaintedItem__InvertedYFramebufferObject QQuickPaintedItem__RenderTarget = QQuickPaintedItem__RenderTarget(2)
)

func NewQQuickPaintedItem(parent QQuickItem_ITF) *QQuickPaintedItem {
	tmpValue := NewQQuickPaintedItemFromPointer(C.QQuickPaintedItem_NewQQuickPaintedItem(PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickPaintedItem_ContentsScaleChanged
func callbackQQuickPaintedItem_ContentsScaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsScaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectContentsScaleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsScaleChanged") {
			C.QQuickPaintedItem_ConnectContentsScaleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsScaleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentsScaleChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsScaleChanged", f)
		}
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsScaleChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "contentsScaleChanged")
	}
}

func (ptr *QQuickPaintedItem) ContentsScaleChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsScaleChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_ContentsSizeChanged
func callbackQQuickPaintedItem_ContentsSizeChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectContentsSizeChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsSizeChanged") {
			C.QQuickPaintedItem_ConnectContentsSizeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsSizeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentsSizeChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsSizeChanged", f)
		}
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "contentsSizeChanged")
	}
}

func (ptr *QQuickPaintedItem) ContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsSizeChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_FillColorChanged
func callbackQQuickPaintedItem_FillColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fillColorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectFillColorChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fillColorChanged") {
			C.QQuickPaintedItem_ConnectFillColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fillColorChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fillColorChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fillColorChanged", f)
		}
	}
}

func (ptr *QQuickPaintedItem) DisconnectFillColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectFillColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fillColorChanged")
	}
}

func (ptr *QQuickPaintedItem) FillColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FillColorChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_Paint
func callbackQQuickPaintedItem_Paint(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	}

}

func (ptr *QQuickPaintedItem) ConnectPaint(f func(painter *gui.QPainter)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "paint", func(painter *gui.QPainter) {
				signal.(func(*gui.QPainter))(painter)
				f(painter)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paint", f)
		}
	}
}

func (ptr *QQuickPaintedItem) DisconnectPaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paint")
	}
}

func (ptr *QQuickPaintedItem) Paint(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQQuickPaintedItem_RenderTargetChanged
func callbackQQuickPaintedItem_RenderTargetChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "renderTargetChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectRenderTargetChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "renderTargetChanged") {
			C.QQuickPaintedItem_ConnectRenderTargetChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "renderTargetChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "renderTargetChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "renderTargetChanged", f)
		}
	}
}

func (ptr *QQuickPaintedItem) DisconnectRenderTargetChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectRenderTargetChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "renderTargetChanged")
	}
}

func (ptr *QQuickPaintedItem) RenderTargetChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_RenderTargetChanged(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) SetContentsScale(vqr float64) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsScale(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickPaintedItem) SetContentsSize(vqs core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsSize(ptr.Pointer(), core.PointerFromQSize(vqs))
	}
}

func (ptr *QQuickPaintedItem) SetFillColor(vqc gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetFillColor(ptr.Pointer(), gui.PointerFromQColor(vqc))
	}
}

func (ptr *QQuickPaintedItem) SetMipmap(enable bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetMipmap(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QQuickPaintedItem) SetOpaquePainting(opaque bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetOpaquePainting(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(opaque))))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHint(hint QQuickPaintedItem__PerformanceHint, enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHint(ptr.Pointer(), C.longlong(hint), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHints(hints QQuickPaintedItem__PerformanceHint) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHints(ptr.Pointer(), C.longlong(hints))
	}
}

func (ptr *QQuickPaintedItem) SetRenderTarget(target QQuickPaintedItem__RenderTarget) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetRenderTarget(ptr.Pointer(), C.longlong(target))
	}
}

func (ptr *QQuickPaintedItem) SetTextureSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetTextureSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

//export callbackQQuickPaintedItem_TextureSizeChanged
func callbackQQuickPaintedItem_TextureSizeChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "textureSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectTextureSizeChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textureSizeChanged") {
			C.QQuickPaintedItem_ConnectTextureSizeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textureSizeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureSizeChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureSizeChanged", f)
		}
	}
}

func (ptr *QQuickPaintedItem) DisconnectTextureSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectTextureSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "textureSizeChanged")
	}
}

func (ptr *QQuickPaintedItem) TextureSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TextureSizeChanged(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) Update(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Update(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

//export callbackQQuickPaintedItem_DestroyQQuickPaintedItem
func callbackQQuickPaintedItem_DestroyQQuickPaintedItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickPaintedItem"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DestroyQQuickPaintedItemDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectDestroyQQuickPaintedItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickPaintedItem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickPaintedItem", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickPaintedItem", f)
		}
	}
}

func (ptr *QQuickPaintedItem) DisconnectDestroyQQuickPaintedItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickPaintedItem")
	}
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItem() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DestroyQQuickPaintedItem(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItemDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DestroyQQuickPaintedItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickPaintedItem) FillColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QQuickPaintedItem_FillColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickPaintedItem) PerformanceHints() QQuickPaintedItem__PerformanceHint {
	if ptr.Pointer() != nil {
		return QQuickPaintedItem__PerformanceHint(C.QQuickPaintedItem_PerformanceHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) RenderTarget() QQuickPaintedItem__RenderTarget {
	if ptr.Pointer() != nil {
		return QQuickPaintedItem__RenderTarget(C.QQuickPaintedItem_RenderTarget(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) ContentsSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickPaintedItem_ContentsSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickPaintedItem) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickPaintedItem_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickPaintedItem) Mipmap() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickPaintedItem_Mipmap(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) OpaquePainting() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickPaintedItem_OpaquePainting(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ContentsScale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickPaintedItem_ContentsScale(ptr.Pointer()))
	}
	return 0
}

type QQuickRenderControl struct {
	core.QObject
}

type QQuickRenderControl_ITF interface {
	core.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func (ptr *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return ptr
}

func (ptr *QQuickRenderControl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickRenderControl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickRenderControl(ptr QQuickRenderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickRenderControl_PTR().Pointer()
	}
	return nil
}

func NewQQuickRenderControlFromPointer(ptr unsafe.Pointer) (n *QQuickRenderControl) {
	n = new(QQuickRenderControl)
	n.SetPointer(ptr)
	return
}
func (ptr *QQuickRenderControl) Grab() *gui.QImage {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQImageFromPointer(C.QQuickRenderControl_Grab(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func NewQQuickRenderControl(parent core.QObject_ITF) *QQuickRenderControl {
	tmpValue := NewQQuickRenderControlFromPointer(C.QQuickRenderControl_NewQQuickRenderControl(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQuickRenderControl_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickRenderControl_QQuickRenderControl_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickRenderControl) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickRenderControl_QQuickRenderControl_Tr(sC, cC, C.int(int32(n))))
}

func QQuickRenderControl_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickRenderControl_QQuickRenderControl_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickRenderControl) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickRenderControl_QQuickRenderControl_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQuickRenderControl_RenderWindow
func callbackQQuickRenderControl_RenderWindow(ptr unsafe.Pointer, offset unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "renderWindow"); signal != nil {
		return gui.PointerFromQWindow(signal.(func(*core.QPoint) *gui.QWindow)(core.NewQPointFromPointer(offset)))
	}

	return gui.PointerFromQWindow(NewQQuickRenderControlFromPointer(ptr).RenderWindowDefault(core.NewQPointFromPointer(offset)))
}

func (ptr *QQuickRenderControl) ConnectRenderWindow(f func(offset *core.QPoint) *gui.QWindow) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "renderWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "renderWindow", func(offset *core.QPoint) *gui.QWindow {
				signal.(func(*core.QPoint) *gui.QWindow)(offset)
				return f(offset)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "renderWindow", f)
		}
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "renderWindow")
	}
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindow(ptr.Pointer(), core.PointerFromQPoint(offset)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickRenderControl) RenderWindowDefault(offset core.QPoint_ITF) *gui.QWindow {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindowDefault(ptr.Pointer(), core.PointerFromQPoint(offset)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	tmpValue := gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickRenderControl) RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	tmpValue := gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickRenderControl) Sync() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickRenderControl_Sync(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) Initialize(gl gui.QOpenGLContext_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(gl))
	}
}

func (ptr *QQuickRenderControl) Invalidate() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Invalidate(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PolishItems() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PolishItems(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PrepareThread(targetThread core.QThread_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PrepareThread(ptr.Pointer(), core.PointerFromQThread(targetThread))
	}
}

func (ptr *QQuickRenderControl) Render() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Render(ptr.Pointer())
	}
}

//export callbackQQuickRenderControl_RenderRequested
func callbackQQuickRenderControl_RenderRequested(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "renderRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "renderRequested") {
			C.QQuickRenderControl_ConnectRenderRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "renderRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "renderRequested", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "renderRequested", f)
		}
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectRenderRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "renderRequested")
	}
}

func (ptr *QQuickRenderControl) RenderRequested() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_RenderRequested(ptr.Pointer())
	}
}

//export callbackQQuickRenderControl_SceneChanged
func callbackQQuickRenderControl_SceneChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sceneChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sceneChanged") {
			C.QQuickRenderControl_ConnectSceneChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sceneChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sceneChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sceneChanged", f)
		}
	}
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectSceneChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sceneChanged")
	}
}

func (ptr *QQuickRenderControl) SceneChanged() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_SceneChanged(ptr.Pointer())
	}
}

//export callbackQQuickRenderControl_DestroyQQuickRenderControl
func callbackQQuickRenderControl_DestroyQQuickRenderControl(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickRenderControl"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickRenderControlFromPointer(ptr).DestroyQQuickRenderControlDefault()
	}
}

func (ptr *QQuickRenderControl) ConnectDestroyQQuickRenderControl(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickRenderControl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickRenderControl", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickRenderControl", f)
		}
	}
}

func (ptr *QQuickRenderControl) DisconnectDestroyQQuickRenderControl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickRenderControl")
	}
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControl(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControlDefault() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControlDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickRenderControl_MetaObject
func callbackQQuickRenderControl_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickRenderControlFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickRenderControl) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickRenderControl_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickRenderControl) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickRenderControl___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickRenderControl) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickRenderControl) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickRenderControl___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickRenderControl) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickRenderControl___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickRenderControl) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickRenderControl) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickRenderControl___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickRenderControl) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickRenderControl___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickRenderControl) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickRenderControl) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickRenderControl___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickRenderControl) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickRenderControl___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickRenderControl) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickRenderControl) __findChildren_newList() unsafe.Pointer {
	return C.QQuickRenderControl___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickRenderControl) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickRenderControl___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickRenderControl) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickRenderControl) __children_newList() unsafe.Pointer {
	return C.QQuickRenderControl___children_newList(ptr.Pointer())
}

//export callbackQQuickRenderControl_Event
func callbackQQuickRenderControl_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickRenderControlFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickRenderControl) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickRenderControl_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickRenderControl_EventFilter
func callbackQQuickRenderControl_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickRenderControlFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickRenderControl) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickRenderControl_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickRenderControl_ChildEvent
func callbackQQuickRenderControl_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickRenderControl_ConnectNotify
func callbackQQuickRenderControl_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickRenderControlFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickRenderControl) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickRenderControl_CustomEvent
func callbackQQuickRenderControl_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickRenderControl_DeleteLater
func callbackQQuickRenderControl_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickRenderControlFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickRenderControl) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickRenderControl_Destroyed
func callbackQQuickRenderControl_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickRenderControl_DisconnectNotify
func callbackQQuickRenderControl_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickRenderControlFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickRenderControl) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickRenderControl_ObjectNameChanged
func callbackQQuickRenderControl_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickRenderControl_TimerEvent
func callbackQQuickRenderControl_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocument_ITF interface {
	core.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func (ptr *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument {
	return ptr
}

func (ptr *QQuickTextDocument) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickTextDocument) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickTextDocument(ptr QQuickTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextDocument_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextDocumentFromPointer(ptr unsafe.Pointer) (n *QQuickTextDocument) {
	n = new(QQuickTextDocument)
	n.SetPointer(ptr)
	return
}
func NewQQuickTextDocument(parent QQuickItem_ITF) *QQuickTextDocument {
	tmpValue := NewQQuickTextDocumentFromPointer(C.QQuickTextDocument_NewQQuickTextDocument(PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQuickTextDocument_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextDocument_QQuickTextDocument_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickTextDocument) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextDocument_QQuickTextDocument_Tr(sC, cC, C.int(int32(n))))
}

func QQuickTextDocument_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextDocument_QQuickTextDocument_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickTextDocument) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextDocument_QQuickTextDocument_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickTextDocument) TextDocument() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQTextDocumentFromPointer(C.QQuickTextDocument_TextDocument(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickTextDocument_MetaObject
func callbackQQuickTextDocument_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickTextDocumentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickTextDocument) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextDocument_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextDocument) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickTextDocument___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextDocument) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickTextDocument) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickTextDocument___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickTextDocument) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextDocument___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextDocument) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextDocument) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickTextDocument___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickTextDocument) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextDocument___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextDocument) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextDocument) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickTextDocument___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickTextDocument) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextDocument___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextDocument) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextDocument) __findChildren_newList() unsafe.Pointer {
	return C.QQuickTextDocument___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickTextDocument) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextDocument___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextDocument) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextDocument) __children_newList() unsafe.Pointer {
	return C.QQuickTextDocument___children_newList(ptr.Pointer())
}

//export callbackQQuickTextDocument_Event
func callbackQQuickTextDocument_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextDocumentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickTextDocument) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickTextDocument_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickTextDocument_EventFilter
func callbackQQuickTextDocument_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextDocumentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickTextDocument) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickTextDocument_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickTextDocument_ChildEvent
func callbackQQuickTextDocument_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickTextDocument_ConnectNotify
func callbackQQuickTextDocument_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextDocument) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextDocument_CustomEvent
func callbackQQuickTextDocument_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickTextDocument_DeleteLater
func callbackQQuickTextDocument_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextDocumentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickTextDocument) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickTextDocument_Destroyed
func callbackQQuickTextDocument_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickTextDocument_DisconnectNotify
func callbackQQuickTextDocument_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextDocument) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextDocument_ObjectNameChanged
func callbackQQuickTextDocument_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickTextDocument_TimerEvent
func callbackQQuickTextDocument_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactory_ITF interface {
	core.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func (ptr *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory {
	return ptr
}

func (ptr *QQuickTextureFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickTextureFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQuickTextureFactory(ptr QQuickTextureFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextureFactory_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextureFactoryFromPointer(ptr unsafe.Pointer) (n *QQuickTextureFactory) {
	n = new(QQuickTextureFactory)
	n.SetPointer(ptr)
	return
}
func QQuickTextureFactory_TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {
	tmpValue := NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(gui.PointerFromQImage(image)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickTextureFactory) TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {
	tmpValue := NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(gui.PointerFromQImage(image)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickTextureFactory() *QQuickTextureFactory {
	tmpValue := NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_NewQQuickTextureFactory())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQuickTextureFactory_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextureFactory_QQuickTextureFactory_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickTextureFactory) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextureFactory_QQuickTextureFactory_Tr(sC, cC, C.int(int32(n))))
}

func QQuickTextureFactory_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextureFactory_QQuickTextureFactory_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickTextureFactory) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickTextureFactory_QQuickTextureFactory_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQuickTextureFactory_DestroyQQuickTextureFactory
func callbackQQuickTextureFactory_DestroyQQuickTextureFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickTextureFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DestroyQQuickTextureFactoryDefault()
	}
}

func (ptr *QQuickTextureFactory) ConnectDestroyQQuickTextureFactory(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickTextureFactory"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickTextureFactory", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickTextureFactory", f)
		}
	}
}

func (ptr *QQuickTextureFactory) DisconnectDestroyQQuickTextureFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickTextureFactory")
	}
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactory() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactory(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactoryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickTextureFactory_Image
func callbackQQuickTextureFactory_Image(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "image"); signal != nil {
		return gui.PointerFromQImage(signal.(func() *gui.QImage)())
	}

	return gui.PointerFromQImage(NewQQuickTextureFactoryFromPointer(ptr).ImageDefault())
}

func (ptr *QQuickTextureFactory) ConnectImage(f func() *gui.QImage) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "image"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "image", func() *gui.QImage {
				signal.(func() *gui.QImage)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "image", f)
		}
	}
}

func (ptr *QQuickTextureFactory) DisconnectImage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "image")
	}
}

func (ptr *QQuickTextureFactory) Image() *gui.QImage {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQImageFromPointer(C.QQuickTextureFactory_Image(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextureFactory) ImageDefault() *gui.QImage {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQImageFromPointer(C.QQuickTextureFactory_ImageDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

//export callbackQQuickTextureFactory_CreateTexture
func callbackQQuickTextureFactory_CreateTexture(ptr unsafe.Pointer, window unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createTexture"); signal != nil {
		return PointerFromQSGTexture(signal.(func(*QQuickWindow) *QSGTexture)(NewQQuickWindowFromPointer(window)))
	}

	return PointerFromQSGTexture(NewQSGTexture())
}

func (ptr *QQuickTextureFactory) ConnectCreateTexture(f func(window *QQuickWindow) *QSGTexture) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createTexture", func(window *QQuickWindow) *QSGTexture {
				signal.(func(*QQuickWindow) *QSGTexture)(window)
				return f(window)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createTexture", f)
		}
	}
}

func (ptr *QQuickTextureFactory) DisconnectCreateTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createTexture")
	}
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF) *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QQuickTextureFactory_CreateTexture(ptr.Pointer(), PointerFromQQuickWindow(window)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickTextureFactory_TextureSize
func callbackQQuickTextureFactory_TextureSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QQuickTextureFactory) ConnectTextureSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "textureSize"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureSize", func() *core.QSize {
				signal.(func() *core.QSize)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureSize", f)
		}
	}
}

func (ptr *QQuickTextureFactory) DisconnectTextureSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "textureSize")
	}
}

func (ptr *QQuickTextureFactory) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickTextureFactory_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickTextureFactory_MetaObject
func callbackQQuickTextureFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickTextureFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickTextureFactory) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextureFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickTextureFactory_TextureByteCount
func callbackQQuickTextureFactory_TextureByteCount(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "textureByteCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QQuickTextureFactory) ConnectTextureByteCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "textureByteCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureByteCount", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureByteCount", f)
		}
	}
}

func (ptr *QQuickTextureFactory) DisconnectTextureByteCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "textureByteCount")
	}
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickTextureFactory_TextureByteCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickTextureFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickTextureFactory___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextureFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickTextureFactory) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickTextureFactory___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickTextureFactory) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextureFactory___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextureFactory) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextureFactory) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickTextureFactory___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickTextureFactory) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextureFactory___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextureFactory) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextureFactory) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickTextureFactory___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickTextureFactory) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextureFactory___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextureFactory) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextureFactory) __findChildren_newList() unsafe.Pointer {
	return C.QQuickTextureFactory___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickTextureFactory) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickTextureFactory___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextureFactory) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickTextureFactory) __children_newList() unsafe.Pointer {
	return C.QQuickTextureFactory___children_newList(ptr.Pointer())
}

//export callbackQQuickTextureFactory_Event
func callbackQQuickTextureFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextureFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickTextureFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickTextureFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickTextureFactory_EventFilter
func callbackQQuickTextureFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextureFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickTextureFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickTextureFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickTextureFactory_ChildEvent
func callbackQQuickTextureFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickTextureFactory_ConnectNotify
func callbackQQuickTextureFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextureFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextureFactory_CustomEvent
func callbackQQuickTextureFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickTextureFactory_DeleteLater
func callbackQQuickTextureFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickTextureFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickTextureFactory_Destroyed
func callbackQQuickTextureFactory_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickTextureFactory_DisconnectNotify
func callbackQQuickTextureFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextureFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextureFactory_ObjectNameChanged
func callbackQQuickTextureFactory_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickTextureFactory_TimerEvent
func callbackQQuickTextureFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickView struct {
	QQuickWindow
}

type QQuickView_ITF interface {
	QQuickWindow_ITF
	QQuickView_PTR() *QQuickView
}

func (ptr *QQuickView) QQuickView_PTR() *QQuickView {
	return ptr
}

func (ptr *QQuickView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickWindow_PTR().SetPointer(p)
	}
}

func PointerFromQQuickView(ptr QQuickView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickView_PTR().Pointer()
	}
	return nil
}

func NewQQuickViewFromPointer(ptr unsafe.Pointer) (n *QQuickView) {
	n = new(QQuickView)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQuickView__ResizeMode
//QQuickView::ResizeMode
type QQuickView__ResizeMode int64

const (
	QQuickView__SizeViewToRootObject QQuickView__ResizeMode = QQuickView__ResizeMode(0)
	QQuickView__SizeRootObjectToView QQuickView__ResizeMode = QQuickView__ResizeMode(1)
)

//go:generate stringer -type=QQuickView__Status
//QQuickView::Status
type QQuickView__Status int64

const (
	QQuickView__Null    QQuickView__Status = QQuickView__Status(0)
	QQuickView__Ready   QQuickView__Status = QQuickView__Status(1)
	QQuickView__Loading QQuickView__Status = QQuickView__Status(2)
	QQuickView__Error   QQuickView__Status = QQuickView__Status(3)
)

func NewQQuickView2(engine qml.QQmlEngine_ITF, parent gui.QWindow_ITF) *QQuickView {
	tmpValue := NewQQuickViewFromPointer(C.QQuickView_NewQQuickView2(qml.PointerFromQQmlEngine(engine), gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickView(parent gui.QWindow_ITF) *QQuickView {
	tmpValue := NewQQuickViewFromPointer(C.QQuickView_NewQQuickView(gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickView3(source core.QUrl_ITF, parent gui.QWindow_ITF) *QQuickView {
	tmpValue := NewQQuickViewFromPointer(C.QQuickView_NewQQuickView3(core.PointerFromQUrl(source), gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickView) SetResizeMode(vqq QQuickView__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetResizeMode(ptr.Pointer(), C.longlong(vqq))
	}
}

//export callbackQQuickView_SetSource
func callbackQQuickView_SetSource(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	} else {
		NewQQuickViewFromPointer(ptr).SetSourceDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQuickView) ConnectSetSource(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSource", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSource", f)
		}
	}
}

func (ptr *QQuickView) DisconnectSetSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSource")
	}
}

func (ptr *QQuickView) SetSource(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickView) SetSourceDefault(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetSourceDefault(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQuickView_StatusChanged
func callbackQQuickView_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		signal.(func(QQuickView__Status))(QQuickView__Status(status))
	}

}

func (ptr *QQuickView) ConnectStatusChanged(f func(status QQuickView__Status)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.QQuickView_ConnectStatusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", func(status QQuickView__Status) {
				signal.(func(QQuickView__Status))(status)
				f(status)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", f)
		}
	}
}

func (ptr *QQuickView) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *QQuickView) StatusChanged(status QQuickView__Status) {
	if ptr.Pointer() != nil {
		C.QQuickView_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

//export callbackQQuickView_DestroyQQuickView
func callbackQQuickView_DestroyQQuickView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickView"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).DestroyQQuickViewDefault()
	}
}

func (ptr *QQuickView) ConnectDestroyQQuickView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickView"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickView", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickView", f)
		}
	}
}

func (ptr *QQuickView) DisconnectDestroyQQuickView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickView")
	}
}

func (ptr *QQuickView) DestroyQQuickView() {
	if ptr.Pointer() != nil {
		C.QQuickView_DestroyQQuickView(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickView) DestroyQQuickViewDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_DestroyQQuickViewDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickView) Errors() []*qml.QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQuick_PackedList) []*qml.QQmlError {
			out := make([]*qml.QQmlError, int(l.len))
			tmpList := NewQQuickViewFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__errors_atList(i)
			}
			return out
		}(C.QQuickView_Errors(ptr.Pointer()))
	}
	return make([]*qml.QQmlError, 0)
}

func (ptr *QQuickView) RootContext() *qml.QQmlContext {
	if ptr.Pointer() != nil {
		tmpValue := qml.NewQQmlContextFromPointer(C.QQuickView_RootContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) Engine() *qml.QQmlEngine {
	if ptr.Pointer() != nil {
		tmpValue := qml.NewQQmlEngineFromPointer(C.QQuickView_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) RootObject() *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickView_RootObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) ResizeMode() QQuickView__ResizeMode {
	if ptr.Pointer() != nil {
		return QQuickView__ResizeMode(C.QQuickView_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) Status() QQuickView__Status {
	if ptr.Pointer() != nil {
		return QQuickView__Status(C.QQuickView_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) InitialSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickView_InitialSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) Source() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQuickView_Source(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) __errors_atList(i int) *qml.QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := qml.NewQQmlErrorFromPointer(C.QQuickView___errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*qml.QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) __errors_setList(i qml.QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView___errors_setList(ptr.Pointer(), qml.PointerFromQQmlError(i))
	}
}

func (ptr *QQuickView) __errors_newList() unsafe.Pointer {
	return C.QQuickView___errors_newList(ptr.Pointer())
}

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidget_ITF interface {
	widgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func (ptr *QQuickWidget) QQuickWidget_PTR() *QQuickWidget {
	return ptr
}

func (ptr *QQuickWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQQuickWidget(ptr QQuickWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWidget_PTR().Pointer()
	}
	return nil
}

func NewQQuickWidgetFromPointer(ptr unsafe.Pointer) (n *QQuickWidget) {
	n = new(QQuickWidget)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQuickWidget__ResizeMode
//QQuickWidget::ResizeMode
type QQuickWidget__ResizeMode int64

const (
	QQuickWidget__SizeViewToRootObject QQuickWidget__ResizeMode = QQuickWidget__ResizeMode(0)
	QQuickWidget__SizeRootObjectToView QQuickWidget__ResizeMode = QQuickWidget__ResizeMode(1)
)

//go:generate stringer -type=QQuickWidget__Status
//QQuickWidget::Status
type QQuickWidget__Status int64

const (
	QQuickWidget__Null    QQuickWidget__Status = QQuickWidget__Status(0)
	QQuickWidget__Ready   QQuickWidget__Status = QQuickWidget__Status(1)
	QQuickWidget__Loading QQuickWidget__Status = QQuickWidget__Status(2)
	QQuickWidget__Error   QQuickWidget__Status = QQuickWidget__Status(3)
)

func NewQQuickWidget2(engine qml.QQmlEngine_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	tmpValue := NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget2(qml.PointerFromQQmlEngine(engine), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickWidget(parent widgets.QWidget_ITF) *QQuickWidget {
	tmpValue := NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickWidget3(source core.QUrl_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	tmpValue := NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget3(core.PointerFromQUrl(source), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQuickWidget_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWidget_QQuickWidget_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickWidget) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWidget_QQuickWidget_Tr(sC, cC, C.int(int32(n))))
}

func QQuickWidget_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWidget_QQuickWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickWidget) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWidget_QQuickWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQuickWidget_Event
func callbackQQuickWidget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickWidget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQuickWidget_FocusNextPrevChild
func callbackQQuickWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QQuickWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQQuickWidget_DragEnterEvent
func callbackQQuickWidget_DragEnterEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

//export callbackQQuickWidget_DragLeaveEvent
func callbackQQuickWidget_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackQQuickWidget_DragMoveEvent
func callbackQQuickWidget_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackQQuickWidget_DropEvent
func callbackQQuickWidget_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

//export callbackQQuickWidget_FocusInEvent
func callbackQQuickWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickWidget_FocusOutEvent
func callbackQQuickWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickWidget_HideEvent
func callbackQQuickWidget_HideEvent(ptr unsafe.Pointer, vqh unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(vqh))
	} else {
		NewQQuickWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(vqh))
	}
}

func (ptr *QQuickWidget) HideEventDefault(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

//export callbackQQuickWidget_KeyPressEvent
func callbackQQuickWidget_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWidget_KeyReleaseEvent
func callbackQQuickWidget_KeyReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWidget_MouseDoubleClickEvent
func callbackQQuickWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQQuickWidget_MouseMoveEvent
func callbackQQuickWidget_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQQuickWidget_MousePressEvent
func callbackQQuickWidget_MousePressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQQuickWidget_MouseReleaseEvent
func callbackQQuickWidget_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQQuickWidget_PaintEvent
func callbackQQuickWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQQuickWidget_SceneGraphError
func callbackQQuickWidget_SceneGraphError(ptr unsafe.Pointer, error C.longlong, message C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), cGoUnpackString(message))
	}

}

func (ptr *QQuickWidget) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sceneGraphError") {
			C.QQuickWidget_ConnectSceneGraphError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sceneGraphError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphError", func(error QQuickWindow__SceneGraphError, message string) {
				signal.(func(QQuickWindow__SceneGraphError, string))(error, message)
				f(error, message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphError", f)
		}
	}
}

func (ptr *QQuickWidget) DisconnectSceneGraphError() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sceneGraphError")
	}
}

func (ptr *QQuickWidget) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QQuickWidget_SceneGraphError(ptr.Pointer(), C.longlong(error), C.struct_QtQuick_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *QQuickWidget) SetClearColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetClearColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QQuickWidget) SetFormat(format gui.QSurfaceFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFormat(ptr.Pointer(), gui.PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QQuickWidget) SetResizeMode(vqq QQuickWidget__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetResizeMode(ptr.Pointer(), C.longlong(vqq))
	}
}

//export callbackQQuickWidget_SetSource
func callbackQQuickWidget_SetSource(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	} else {
		NewQQuickWidgetFromPointer(ptr).SetSourceDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQuickWidget) ConnectSetSource(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSource", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSource", f)
		}
	}
}

func (ptr *QQuickWidget) DisconnectSetSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSource")
	}
}

func (ptr *QQuickWidget) SetSource(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickWidget) SetSourceDefault(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetSourceDefault(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQuickWidget_ShowEvent
func callbackQQuickWidget_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(vqs))
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *QQuickWidget) ShowEventDefault(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackQQuickWidget_StatusChanged
func callbackQQuickWidget_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		signal.(func(QQuickWidget__Status))(QQuickWidget__Status(status))
	}

}

func (ptr *QQuickWidget) ConnectStatusChanged(f func(status QQuickWidget__Status)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.QQuickWidget_ConnectStatusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", func(status QQuickWidget__Status) {
				signal.(func(QQuickWidget__Status))(status)
				f(status)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", f)
		}
	}
}

func (ptr *QQuickWidget) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *QQuickWidget) StatusChanged(status QQuickWidget__Status) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

//export callbackQQuickWidget_WheelEvent
func callbackQQuickWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQQuickWidget_DestroyQQuickWidget
func callbackQQuickWidget_DestroyQQuickWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickWidget"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).DestroyQQuickWidgetDefault()
	}
}

func (ptr *QQuickWidget) ConnectDestroyQQuickWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickWidget", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickWidget", f)
		}
	}
}

func (ptr *QQuickWidget) DisconnectDestroyQQuickWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickWidget")
	}
}

func (ptr *QQuickWidget) DestroyQQuickWidget() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DestroyQQuickWidget(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickWidget) DestroyQQuickWidgetDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DestroyQQuickWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickWidget) GrabFramebuffer() *gui.QImage {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQImageFromPointer(C.QQuickWidget_GrabFramebuffer(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) Errors() []*qml.QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQuick_PackedList) []*qml.QQmlError {
			out := make([]*qml.QQmlError, int(l.len))
			tmpList := NewQQuickWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__errors_atList(i)
			}
			return out
		}(C.QQuickWidget_Errors(ptr.Pointer()))
	}
	return make([]*qml.QQmlError, 0)
}

func (ptr *QQuickWidget) RootContext() *qml.QQmlContext {
	if ptr.Pointer() != nil {
		tmpValue := qml.NewQQmlContextFromPointer(C.QQuickWidget_RootContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) Engine() *qml.QQmlEngine {
	if ptr.Pointer() != nil {
		tmpValue := qml.NewQQmlEngineFromPointer(C.QQuickWidget_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) RootObject() *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickWidget_RootObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) ResizeMode() QQuickWidget__ResizeMode {
	if ptr.Pointer() != nil {
		return QQuickWidget__ResizeMode(C.QQuickWidget_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWidget) Status() QQuickWidget__Status {
	if ptr.Pointer() != nil {
		return QQuickWidget__Status(C.QQuickWidget_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWidget) QuickWindow() *QQuickWindow {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickWindowFromPointer(C.QQuickWidget_QuickWindow(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) InitialSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickWidget_InitialSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) Format() *gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQSurfaceFormatFromPointer(C.QQuickWidget_Format(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) Source() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQuickWidget_Source(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_MetaObject
func callbackQQuickWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) __errors_atList(i int) *qml.QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := qml.NewQQmlErrorFromPointer(C.QQuickWidget___errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*qml.QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __errors_setList(i qml.QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___errors_setList(ptr.Pointer(), qml.PointerFromQQmlError(i))
	}
}

func (ptr *QQuickWidget) __errors_newList() unsafe.Pointer {
	return C.QQuickWidget___errors_newList(ptr.Pointer())
}

func (ptr *QQuickWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QQuickWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QQuickWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QQuickWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QQuickWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QQuickWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QQuickWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QQuickWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QQuickWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QQuickWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QQuickWidget) __actions_newList() unsafe.Pointer {
	return C.QQuickWidget___actions_newList(ptr.Pointer())
}

func (ptr *QQuickWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWidget) __findChildren_newList() unsafe.Pointer {
	return C.QQuickWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWidget) __children_newList() unsafe.Pointer {
	return C.QQuickWidget___children_newList(ptr.Pointer())
}

//export callbackQQuickWidget_Close
func callbackQQuickWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QQuickWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickWidget_ActionEvent
func callbackQQuickWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQQuickWidget_ChangeEvent
func callbackQQuickWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_CloseEvent
func callbackQQuickWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQQuickWidget_ContextMenuEvent
func callbackQQuickWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQQuickWidget_CustomContextMenuRequested
func callbackQQuickWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQQuickWidget_EnterEvent
func callbackQQuickWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_Hide
func callbackQQuickWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_InputMethodEvent
func callbackQQuickWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQQuickWidget_LeaveEvent
func callbackQQuickWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_Lower
func callbackQQuickWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_MoveEvent
func callbackQQuickWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQQuickWidget_Raise
func callbackQQuickWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_Repaint
func callbackQQuickWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QQuickWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ResizeEvent
func callbackQQuickWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQQuickWidget_SetDisabled
func callbackQQuickWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QQuickWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQQuickWidget_SetEnabled
func callbackQQuickWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QQuickWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQQuickWidget_SetFocus2
func callbackQQuickWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QQuickWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQQuickWidget_SetHidden
func callbackQQuickWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QQuickWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQQuickWidget_SetStyleSheet
func callbackQQuickWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQQuickWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QQuickWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QQuickWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtQuick_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQQuickWidget_SetVisible
func callbackQQuickWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QQuickWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQQuickWidget_SetWindowModified
func callbackQQuickWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QQuickWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQQuickWidget_SetWindowTitle
func callbackQQuickWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQQuickWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QQuickWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QQuickWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtQuick_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQQuickWidget_Show
func callbackQQuickWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowFullScreen
func callbackQQuickWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowMaximized
func callbackQQuickWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowMinimized
func callbackQQuickWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowNormal
func callbackQQuickWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_TabletEvent
func callbackQQuickWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQQuickWidget_Update
func callbackQQuickWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_UpdateMicroFocus
func callbackQQuickWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QQuickWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_WindowIconChanged
func callbackQQuickWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQQuickWidget_WindowTitleChanged
func callbackQQuickWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQQuickWidget_PaintEngine
func callbackQQuickWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQQuickWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QQuickWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QQuickWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWidget_MinimumSizeHint
func callbackQQuickWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QQuickWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_SizeHint
func callbackQQuickWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QQuickWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_InputMethodQuery
func callbackQQuickWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QQuickWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_HasHeightForWidth
func callbackQQuickWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QQuickWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickWidget_HeightForWidth
func callbackQQuickWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQQuickWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QQuickWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQQuickWidget_Metric
func callbackQQuickWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQQuickWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QQuickWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQQuickWidget_InitPainter
func callbackQQuickWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQQuickWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QQuickWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQQuickWidget_EventFilter
func callbackQQuickWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickWidget_ChildEvent
func callbackQQuickWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWidget_ConnectNotify
func callbackQQuickWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWidget_CustomEvent
func callbackQQuickWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_DeleteLater
func callbackQQuickWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickWidget_Destroyed
func callbackQQuickWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickWidget_DisconnectNotify
func callbackQQuickWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWidget_ObjectNameChanged
func callbackQQuickWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickWidget_TimerEvent
func callbackQQuickWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQuickWindow struct {
	gui.QWindow
}

type QQuickWindow_ITF interface {
	gui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func (ptr *QQuickWindow) QQuickWindow_PTR() *QQuickWindow {
	return ptr
}

func (ptr *QQuickWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *QQuickWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWindow_PTR().SetPointer(p)
	}
}

func PointerFromQQuickWindow(ptr QQuickWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func NewQQuickWindowFromPointer(ptr unsafe.Pointer) (n *QQuickWindow) {
	n = new(QQuickWindow)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQuickWindow__CreateTextureOption
//QQuickWindow::CreateTextureOption
type QQuickWindow__CreateTextureOption int64

const (
	QQuickWindow__TextureHasAlphaChannel QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0001)
	QQuickWindow__TextureHasMipmaps      QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0002)
	QQuickWindow__TextureOwnsGLTexture   QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0004)
	QQuickWindow__TextureCanUseAtlas     QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0008)
	QQuickWindow__TextureIsOpaque        QQuickWindow__CreateTextureOption = QQuickWindow__CreateTextureOption(0x0010)
)

//go:generate stringer -type=QQuickWindow__RenderStage
//QQuickWindow::RenderStage
type QQuickWindow__RenderStage int64

const (
	QQuickWindow__BeforeSynchronizingStage QQuickWindow__RenderStage = QQuickWindow__RenderStage(0)
	QQuickWindow__AfterSynchronizingStage  QQuickWindow__RenderStage = QQuickWindow__RenderStage(1)
	QQuickWindow__BeforeRenderingStage     QQuickWindow__RenderStage = QQuickWindow__RenderStage(2)
	QQuickWindow__AfterRenderingStage      QQuickWindow__RenderStage = QQuickWindow__RenderStage(3)
	QQuickWindow__AfterSwapStage           QQuickWindow__RenderStage = QQuickWindow__RenderStage(4)
	QQuickWindow__NoStage                  QQuickWindow__RenderStage = QQuickWindow__RenderStage(5)
)

//go:generate stringer -type=QQuickWindow__SceneGraphError
//QQuickWindow::SceneGraphError
type QQuickWindow__SceneGraphError int64

const (
	QQuickWindow__ContextNotAvailable QQuickWindow__SceneGraphError = QQuickWindow__SceneGraphError(1)
)

//go:generate stringer -type=QQuickWindow__TextRenderType
//QQuickWindow::TextRenderType
type QQuickWindow__TextRenderType int64

const (
	QQuickWindow__QtTextRendering     QQuickWindow__TextRenderType = QQuickWindow__TextRenderType(0)
	QQuickWindow__NativeTextRendering QQuickWindow__TextRenderType = QQuickWindow__TextRenderType(1)
)

func (ptr *QQuickWindow) GrabWindow() *gui.QImage {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQImageFromPointer(C.QQuickWindow_GrabWindow(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func NewQQuickWindow(parent gui.QWindow_ITF) *QQuickWindow {
	tmpValue := NewQQuickWindowFromPointer(C.QQuickWindow_NewQQuickWindow(gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQuickWindow_TextRenderType() QQuickWindow__TextRenderType {
	return QQuickWindow__TextRenderType(C.QQuickWindow_QQuickWindow_TextRenderType())
}

func (ptr *QQuickWindow) TextRenderType() QQuickWindow__TextRenderType {
	return QQuickWindow__TextRenderType(C.QQuickWindow_QQuickWindow_TextRenderType())
}

func QQuickWindow_SceneGraphBackend() string {
	return cGoUnpackString(C.QQuickWindow_QQuickWindow_SceneGraphBackend())
}

func (ptr *QQuickWindow) SceneGraphBackend() string {
	return cGoUnpackString(C.QQuickWindow_QQuickWindow_SceneGraphBackend())
}

func QQuickWindow_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWindow_QQuickWindow_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickWindow) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWindow_QQuickWindow_Tr(sC, cC, C.int(int32(n))))
}

func QQuickWindow_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWindow_QQuickWindow_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQuickWindow) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QQuickWindow_QQuickWindow_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQuickWindow_Event
func callbackQQuickWindow_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickWindow) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWindow_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func QQuickWindow_HasDefaultAlphaBuffer() bool {
	return int8(C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer()) != 0
}

func (ptr *QQuickWindow) HasDefaultAlphaBuffer() bool {
	return int8(C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer()) != 0
}

//export callbackQQuickWindow_ActiveFocusItemChanged
func callbackQQuickWindow_ActiveFocusItemChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activeFocusItemChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activeFocusItemChanged") {
			C.QQuickWindow_ConnectActiveFocusItemChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activeFocusItemChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activeFocusItemChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeFocusItemChanged", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectActiveFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activeFocusItemChanged")
	}
}

func (ptr *QQuickWindow) ActiveFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ActiveFocusItemChanged(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterAnimating
func callbackQQuickWindow_AfterAnimating(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "afterAnimating"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "afterAnimating") {
			C.QQuickWindow_ConnectAfterAnimating(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "afterAnimating"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "afterAnimating", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "afterAnimating", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterAnimating(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "afterAnimating")
	}
}

func (ptr *QQuickWindow) AfterAnimating() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterAnimating(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterRendering
func callbackQQuickWindow_AfterRendering(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "afterRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "afterRendering") {
			C.QQuickWindow_ConnectAfterRendering(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "afterRendering"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "afterRendering", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "afterRendering", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "afterRendering")
	}
}

func (ptr *QQuickWindow) AfterRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterRendering(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterSynchronizing
func callbackQQuickWindow_AfterSynchronizing(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "afterSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "afterSynchronizing") {
			C.QQuickWindow_ConnectAfterSynchronizing(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "afterSynchronizing"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "afterSynchronizing", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "afterSynchronizing", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "afterSynchronizing")
	}
}

func (ptr *QQuickWindow) AfterSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterSynchronizing(ptr.Pointer())
	}
}

//export callbackQQuickWindow_BeforeRendering
func callbackQQuickWindow_BeforeRendering(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "beforeRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "beforeRendering") {
			C.QQuickWindow_ConnectBeforeRendering(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "beforeRendering"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "beforeRendering", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "beforeRendering", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "beforeRendering")
	}
}

func (ptr *QQuickWindow) BeforeRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_BeforeRendering(ptr.Pointer())
	}
}

//export callbackQQuickWindow_BeforeSynchronizing
func callbackQQuickWindow_BeforeSynchronizing(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "beforeSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "beforeSynchronizing") {
			C.QQuickWindow_ConnectBeforeSynchronizing(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "beforeSynchronizing"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "beforeSynchronizing", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "beforeSynchronizing", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "beforeSynchronizing")
	}
}

func (ptr *QQuickWindow) BeforeSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_BeforeSynchronizing(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ColorChanged
func callbackQQuickWindow_ColorChanged(ptr unsafe.Pointer, vqc unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(vqc))
	}

}

func (ptr *QQuickWindow) ConnectColorChanged(f func(vqc *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorChanged") {
			C.QQuickWindow_ConnectColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", func(vqc *gui.QColor) {
				signal.(func(*gui.QColor))(vqc)
				f(vqc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorChanged")
	}
}

func (ptr *QQuickWindow) ColorChanged(vqc gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(vqc))
	}
}

//export callbackQQuickWindow_ExposeEvent
func callbackQQuickWindow_ExposeEvent(ptr unsafe.Pointer, vqe unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(vqe))
	} else {
		NewQQuickWindowFromPointer(ptr).ExposeEventDefault(gui.NewQExposeEventFromPointer(vqe))
	}
}

func (ptr *QQuickWindow) ExposeEventDefault(vqe gui.QExposeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ExposeEventDefault(ptr.Pointer(), gui.PointerFromQExposeEvent(vqe))
	}
}

//export callbackQQuickWindow_FocusInEvent
func callbackQQuickWindow_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQQuickWindow_FocusOutEvent
func callbackQQuickWindow_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQQuickWindow_FrameSwapped
func callbackQQuickWindow_FrameSwapped(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "frameSwapped"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "frameSwapped") {
			C.QQuickWindow_ConnectFrameSwapped(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "frameSwapped"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "frameSwapped", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "frameSwapped", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectFrameSwapped(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "frameSwapped")
	}
}

func (ptr *QQuickWindow) FrameSwapped() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FrameSwapped(ptr.Pointer())
	}
}

//export callbackQQuickWindow_HideEvent
func callbackQQuickWindow_HideEvent(ptr unsafe.Pointer, vqh unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(vqh))
	} else {
		NewQQuickWindowFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(vqh))
	}
}

func (ptr *QQuickWindow) HideEventDefault(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

//export callbackQQuickWindow_KeyPressEvent
func callbackQQuickWindow_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWindowFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWindow) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWindow_KeyReleaseEvent
func callbackQQuickWindow_KeyReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWindowFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWindow) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWindow_MouseDoubleClickEvent
func callbackQQuickWindow_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickWindow_MouseMoveEvent
func callbackQQuickWindow_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickWindow_MousePressEvent
func callbackQQuickWindow_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickWindow_MouseReleaseEvent
func callbackQQuickWindow_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickWindow_OpenglContextCreated
func callbackQQuickWindow_OpenglContextCreated(ptr unsafe.Pointer, context unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "openglContextCreated"); signal != nil {
		signal.(func(*gui.QOpenGLContext))(gui.NewQOpenGLContextFromPointer(context))
	}

}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context *gui.QOpenGLContext)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "openglContextCreated") {
			C.QQuickWindow_ConnectOpenglContextCreated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "openglContextCreated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "openglContextCreated", func(context *gui.QOpenGLContext) {
				signal.(func(*gui.QOpenGLContext))(context)
				f(context)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "openglContextCreated", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectOpenglContextCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "openglContextCreated")
	}
}

func (ptr *QQuickWindow) OpenglContextCreated(context gui.QOpenGLContext_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_OpenglContextCreated(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

//export callbackQQuickWindow_ReleaseResources
func callbackQQuickWindow_ReleaseResources(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickWindow) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "releaseResources"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "releaseResources", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "releaseResources", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "releaseResources")
	}
}

func (ptr *QQuickWindow) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResourcesDefault(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ResetOpenGLState() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ResetOpenGLState(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ResizeEvent
func callbackQQuickWindow_ResizeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ResizeEventDefault(ev gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

//export callbackQQuickWindow_SceneGraphAboutToStop
func callbackQQuickWindow_SceneGraphAboutToStop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sceneGraphAboutToStop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sceneGraphAboutToStop") {
			C.QQuickWindow_ConnectSceneGraphAboutToStop(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sceneGraphAboutToStop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphAboutToStop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphAboutToStop", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphAboutToStop(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sceneGraphAboutToStop")
	}
}

func (ptr *QQuickWindow) SceneGraphAboutToStop() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphAboutToStop(ptr.Pointer())
	}
}

//export callbackQQuickWindow_SceneGraphError
func callbackQQuickWindow_SceneGraphError(ptr unsafe.Pointer, error C.longlong, message C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), cGoUnpackString(message))
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sceneGraphError") {
			C.QQuickWindow_ConnectSceneGraphError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sceneGraphError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphError", func(error QQuickWindow__SceneGraphError, message string) {
				signal.(func(QQuickWindow__SceneGraphError, string))(error, message)
				f(error, message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphError", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sceneGraphError")
	}
}

func (ptr *QQuickWindow) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QQuickWindow_SceneGraphError(ptr.Pointer(), C.longlong(error), C.struct_QtQuick_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackQQuickWindow_SceneGraphInitialized
func callbackQQuickWindow_SceneGraphInitialized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sceneGraphInitialized"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sceneGraphInitialized") {
			C.QQuickWindow_ConnectSceneGraphInitialized(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sceneGraphInitialized"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphInitialized", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphInitialized", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInitialized(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sceneGraphInitialized")
	}
}

func (ptr *QQuickWindow) SceneGraphInitialized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphInitialized(ptr.Pointer())
	}
}

//export callbackQQuickWindow_SceneGraphInvalidated
func callbackQQuickWindow_SceneGraphInvalidated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sceneGraphInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sceneGraphInvalidated") {
			C.QQuickWindow_ConnectSceneGraphInvalidated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sceneGraphInvalidated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphInvalidated", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphInvalidated", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sceneGraphInvalidated")
	}
}

func (ptr *QQuickWindow) SceneGraphInvalidated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphInvalidated(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ScheduleRenderJob(job core.QRunnable_ITF, stage QQuickWindow__RenderStage) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ScheduleRenderJob(ptr.Pointer(), core.PointerFromQRunnable(job), C.longlong(stage))
	}
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetClearBeforeRendering(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQuickWindow) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	C.QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(C.char(int8(qt.GoBoolToInt(useAlpha))))
}

func (ptr *QQuickWindow) SetDefaultAlphaBuffer(useAlpha bool) {
	C.QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(C.char(int8(qt.GoBoolToInt(useAlpha))))
}

func (ptr *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentOpenGLContext(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(persistent))))
	}
}

func (ptr *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentSceneGraph(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(persistent))))
	}
}

func (ptr *QQuickWindow) SetRenderTarget(fbo gui.QOpenGLFramebufferObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetRenderTarget(ptr.Pointer(), gui.PointerFromQOpenGLFramebufferObject(fbo))
	}
}

func (ptr *QQuickWindow) SetRenderTarget2(fboId uint, size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetRenderTarget2(ptr.Pointer(), C.uint(uint32(fboId)), core.PointerFromQSize(size))
	}
}

func QQuickWindow_SetSceneGraphBackend(api QSGRendererInterface__GraphicsApi) {
	C.QQuickWindow_QQuickWindow_SetSceneGraphBackend(C.longlong(api))
}

func (ptr *QQuickWindow) SetSceneGraphBackend(api QSGRendererInterface__GraphicsApi) {
	C.QQuickWindow_QQuickWindow_SetSceneGraphBackend(C.longlong(api))
}

func QQuickWindow_SetSceneGraphBackend2(backend string) {
	var backendC *C.char
	if backend != "" {
		backendC = C.CString(backend)
		defer C.free(unsafe.Pointer(backendC))
	}
	C.QQuickWindow_QQuickWindow_SetSceneGraphBackend2(C.struct_QtQuick_PackedString{data: backendC, len: C.longlong(len(backend))})
}

func (ptr *QQuickWindow) SetSceneGraphBackend2(backend string) {
	var backendC *C.char
	if backend != "" {
		backendC = C.CString(backend)
		defer C.free(unsafe.Pointer(backendC))
	}
	C.QQuickWindow_QQuickWindow_SetSceneGraphBackend2(C.struct_QtQuick_PackedString{data: backendC, len: C.longlong(len(backend))})
}

func QQuickWindow_SetTextRenderType(renderType QQuickWindow__TextRenderType) {
	C.QQuickWindow_QQuickWindow_SetTextRenderType(C.longlong(renderType))
}

func (ptr *QQuickWindow) SetTextRenderType(renderType QQuickWindow__TextRenderType) {
	C.QQuickWindow_QQuickWindow_SetTextRenderType(C.longlong(renderType))
}

//export callbackQQuickWindow_ShowEvent
func callbackQQuickWindow_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(vqs))
	} else {
		NewQQuickWindowFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *QQuickWindow) ShowEventDefault(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackQQuickWindow_Update
func callbackQQuickWindow_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickWindow) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "update"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "update", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "update", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "update")
	}
}

func (ptr *QQuickWindow) Update() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_WheelEvent
func callbackQQuickWindow_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQQuickWindow_DestroyQQuickWindow
func callbackQQuickWindow_DestroyQQuickWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQuickWindow"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).DestroyQQuickWindowDefault()
	}
}

func (ptr *QQuickWindow) ConnectDestroyQQuickWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQuickWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickWindow", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQuickWindow", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectDestroyQQuickWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQuickWindow")
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindow(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindowDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickWindow_AccessibleRoot
func callbackQQuickWindow_AccessibleRoot(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "accessibleRoot"); signal != nil {
		return gui.PointerFromQAccessibleInterface(signal.(func() *gui.QAccessibleInterface)())
	}

	return gui.PointerFromQAccessibleInterface(NewQQuickWindowFromPointer(ptr).AccessibleRootDefault())
}

func (ptr *QQuickWindow) ConnectAccessibleRoot(f func() *gui.QAccessibleInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "accessibleRoot"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "accessibleRoot", func() *gui.QAccessibleInterface {
				signal.(func() *gui.QAccessibleInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accessibleRoot", f)
		}
	}
}

func (ptr *QQuickWindow) DisconnectAccessibleRoot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "accessibleRoot")
	}
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QQuickWindow_AccessibleRoot(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) AccessibleRootDefault() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QQuickWindow_AccessibleRootDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QQuickWindow_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQOpenGLContextFromPointer(C.QQuickWindow_OpenglContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {
	if ptr.Pointer() != nil {
		return gui.NewQOpenGLFramebufferObjectFromPointer(C.QQuickWindow_RenderTarget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) IncubationController() *qml.QQmlIncubationController {
	if ptr.Pointer() != nil {
		return qml.NewQQmlIncubationControllerFromPointer(C.QQuickWindow_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ActiveFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickWindow_ActiveFocusItem(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickWindow_ContentItem(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQQuickItemFromPointer(C.QQuickWindow_MouseGrabberItem(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) CreateImageNode() *QSGImageNode {
	if ptr.Pointer() != nil {
		return NewQSGImageNodeFromPointer(C.QQuickWindow_CreateImageNode(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) CreateRectangleNode() *QSGRectangleNode {
	if ptr.Pointer() != nil {
		return NewQSGRectangleNodeFromPointer(C.QQuickWindow_CreateRectangleNode(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) RendererInterface() *QSGRendererInterface {
	if ptr.Pointer() != nil {
		return NewQSGRendererInterfaceFromPointer(C.QQuickWindow_RendererInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromId(id uint, size core.QSize_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromId(ptr.Pointer(), C.uint(uint32(id)), core.PointerFromQSize(size), C.longlong(options)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImage_ITF) *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage2(ptr.Pointer(), gui.PointerFromQImage(image)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImage_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.longlong(options)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) RenderTargetSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickWindow_RenderTargetSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWindow_ClearBeforeRendering(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWindow_IsPersistentOpenGLContext(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWindow_IsPersistentSceneGraph(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWindow_IsSceneGraphInitialized(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickWindow_MetaObject
func callbackQQuickWindow_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWindowFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWindow) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWindow_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickWindow_EffectiveDevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWindow) RenderTargetId() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QQuickWindow_RenderTargetId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickWindow) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQuickWindow___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQuickWindow) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQuickWindow___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQuickWindow) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWindow___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWindow) __findChildren_newList2() unsafe.Pointer {
	return C.QQuickWindow___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQuickWindow) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWindow___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWindow) __findChildren_newList3() unsafe.Pointer {
	return C.QQuickWindow___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQuickWindow) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWindow___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWindow) __findChildren_newList() unsafe.Pointer {
	return C.QQuickWindow___findChildren_newList(ptr.Pointer())
}

func (ptr *QQuickWindow) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWindow___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQuickWindow) __children_newList() unsafe.Pointer {
	return C.QQuickWindow___children_newList(ptr.Pointer())
}

//export callbackQQuickWindow_Close
func callbackQQuickWindow_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *QQuickWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWindow_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQuickWindow_ActiveChanged
func callbackQQuickWindow_ActiveChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activeChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQQuickWindow_Alert
func callbackQQuickWindow_Alert(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "alert"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQQuickWindowFromPointer(ptr).AlertDefault(int(int32(msec)))
	}
}

func (ptr *QQuickWindow) AlertDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AlertDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

//export callbackQQuickWindow_ContentOrientationChanged
func callbackQQuickWindow_ContentOrientationChanged(ptr unsafe.Pointer, orientation C.longlong) {
	if signal := qt.GetSignal(ptr, "contentOrientationChanged"); signal != nil {
		signal.(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
	}

}

//export callbackQQuickWindow_FocusObjectChanged
func callbackQQuickWindow_FocusObjectChanged(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusObjectChanged"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

//export callbackQQuickWindow_HeightChanged
func callbackQQuickWindow_HeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "heightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_Hide
func callbackQQuickWindow_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_Lower
func callbackQQuickWindow_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_MaximumHeightChanged
func callbackQQuickWindow_MaximumHeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "maximumHeightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_MaximumWidthChanged
func callbackQQuickWindow_MaximumWidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "maximumWidthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_MinimumHeightChanged
func callbackQQuickWindow_MinimumHeightChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "minimumHeightChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_MinimumWidthChanged
func callbackQQuickWindow_MinimumWidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "minimumWidthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_ModalityChanged
func callbackQQuickWindow_ModalityChanged(ptr unsafe.Pointer, modality C.longlong) {
	if signal := qt.GetSignal(ptr, "modalityChanged"); signal != nil {
		signal.(func(core.Qt__WindowModality))(core.Qt__WindowModality(modality))
	}

}

//export callbackQQuickWindow_MoveEvent
func callbackQQuickWindow_MoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) MoveEventDefault(ev gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

//export callbackQQuickWindow_OpacityChanged
func callbackQQuickWindow_OpacityChanged(ptr unsafe.Pointer, opacity C.double) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		signal.(func(float64))(float64(opacity))
	}

}

//export callbackQQuickWindow_Raise
func callbackQQuickWindow_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickWindow) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_RequestActivate
func callbackQQuickWindow_RequestActivate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestActivate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RequestActivateDefault()
	}
}

func (ptr *QQuickWindow) RequestActivateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestActivateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_RequestUpdate
func callbackQQuickWindow_RequestUpdate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RequestUpdateDefault()
	}
}

func (ptr *QQuickWindow) RequestUpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestUpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ScreenChanged
func callbackQQuickWindow_ScreenChanged(ptr unsafe.Pointer, screen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "screenChanged"); signal != nil {
		signal.(func(*gui.QScreen))(gui.NewQScreenFromPointer(screen))
	}

}

//export callbackQQuickWindow_SetGeometry2
func callbackQQuickWindow_SetGeometry2(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry2"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(rect))
	} else {
		NewQQuickWindowFromPointer(ptr).SetGeometry2Default(core.NewQRectFromPointer(rect))
	}
}

func (ptr *QQuickWindow) SetGeometry2Default(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetGeometry2Default(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

//export callbackQQuickWindow_SetGeometry
func callbackQQuickWindow_SetGeometry(ptr unsafe.Pointer, posx C.int, posy C.int, w C.int, h C.int) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(int, int, int, int))(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetGeometryDefault(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	}
}

func (ptr *QQuickWindow) SetGeometryDefault(posx int, posy int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetGeometryDefault(ptr.Pointer(), C.int(int32(posx)), C.int(int32(posy)), C.int(int32(w)), C.int(int32(h)))
	}
}

//export callbackQQuickWindow_SetHeight
func callbackQQuickWindow_SetHeight(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setHeight"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetHeightDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetHeightDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetHeightDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_SetMaximumHeight
func callbackQQuickWindow_SetMaximumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(ptr, "setMaximumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMaximumHeightDefault(int(int32(h)))
	}
}

func (ptr *QQuickWindow) SetMaximumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackQQuickWindow_SetMaximumWidth
func callbackQQuickWindow_SetMaximumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMaximumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMaximumWidthDefault(int(int32(w)))
	}
}

func (ptr *QQuickWindow) SetMaximumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackQQuickWindow_SetMinimumHeight
func callbackQQuickWindow_SetMinimumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(ptr, "setMinimumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMinimumHeightDefault(int(int32(h)))
	}
}

func (ptr *QQuickWindow) SetMinimumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackQQuickWindow_SetMinimumWidth
func callbackQQuickWindow_SetMinimumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMinimumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMinimumWidthDefault(int(int32(w)))
	}
}

func (ptr *QQuickWindow) SetMinimumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackQQuickWindow_SetTitle
func callbackQQuickWindow_SetTitle(ptr unsafe.Pointer, vqs C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "setTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQQuickWindowFromPointer(ptr).SetTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QQuickWindow) SetTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QQuickWindow_SetTitleDefault(ptr.Pointer(), C.struct_QtQuick_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQQuickWindow_SetVisible
func callbackQQuickWindow_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQQuickWindowFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QQuickWindow) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQQuickWindow_SetWidth
func callbackQQuickWindow_SetWidth(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setWidth"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetWidthDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetWidthDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetWidthDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_SetX
func callbackQQuickWindow_SetX(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setX"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetXDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetXDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetXDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_SetY
func callbackQQuickWindow_SetY(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "setY"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetYDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetYDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetYDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_Show
func callbackQQuickWindow_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowFullScreen
func callbackQQuickWindow_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickWindow) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowMaximized
func callbackQQuickWindow_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowMinimized
func callbackQQuickWindow_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickWindow) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowNormal
func callbackQQuickWindow_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickWindow) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_TabletEvent
func callbackQQuickWindow_TabletEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) TabletEventDefault(ev gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

//export callbackQQuickWindow_TouchEvent
func callbackQQuickWindow_TouchEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) TouchEventDefault(ev gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

//export callbackQQuickWindow_VisibilityChanged
func callbackQQuickWindow_VisibilityChanged(ptr unsafe.Pointer, visibility C.longlong) {
	if signal := qt.GetSignal(ptr, "visibilityChanged"); signal != nil {
		signal.(func(gui.QWindow__Visibility))(gui.QWindow__Visibility(visibility))
	}

}

//export callbackQQuickWindow_VisibleChanged
func callbackQQuickWindow_VisibleChanged(ptr unsafe.Pointer, arg C.char) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		signal.(func(bool))(int8(arg) != 0)
	}

}

//export callbackQQuickWindow_WidthChanged
func callbackQQuickWindow_WidthChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "widthChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_WindowStateChanged
func callbackQQuickWindow_WindowStateChanged(ptr unsafe.Pointer, windowState C.longlong) {
	if signal := qt.GetSignal(ptr, "windowStateChanged"); signal != nil {
		signal.(func(core.Qt__WindowState))(core.Qt__WindowState(windowState))
	}

}

//export callbackQQuickWindow_WindowTitleChanged
func callbackQQuickWindow_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQQuickWindow_XChanged
func callbackQQuickWindow_XChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_YChanged
func callbackQQuickWindow_YChanged(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	}

}

//export callbackQQuickWindow_FocusObject
func callbackQQuickWindow_FocusObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "focusObject"); signal != nil {
		return core.PointerFromQObject(signal.(func() *core.QObject)())
	}

	return core.PointerFromQObject(NewQQuickWindowFromPointer(ptr).FocusObjectDefault())
}

func (ptr *QQuickWindow) FocusObjectDefault() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQuickWindow_FocusObjectDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_Size
func callbackQQuickWindow_Size(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWindowFromPointer(ptr).SizeDefault())
}

func (ptr *QQuickWindow) SizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QQuickWindow_SizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_SurfaceType
func callbackQQuickWindow_SurfaceType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "surfaceType"); signal != nil {
		return C.longlong(signal.(func() gui.QSurface__SurfaceType)())
	}

	return C.longlong(NewQQuickWindowFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *QQuickWindow) SurfaceTypeDefault() gui.QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickWindow_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWindow_Format
func callbackQQuickWindow_Format(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "format"); signal != nil {
		return gui.PointerFromQSurfaceFormat(signal.(func() *gui.QSurfaceFormat)())
	}

	return gui.PointerFromQSurfaceFormat(NewQQuickWindowFromPointer(ptr).FormatDefault())
}

func (ptr *QQuickWindow) FormatDefault() *gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQSurfaceFormatFromPointer(C.QQuickWindow_FormatDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_EventFilter
func callbackQQuickWindow_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickWindow) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQuickWindow_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQuickWindow_ChildEvent
func callbackQQuickWindow_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWindow_ConnectNotify
func callbackQQuickWindow_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWindowFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWindow) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWindow_CustomEvent
func callbackQQuickWindow_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWindow_DeleteLater
func callbackQQuickWindow_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQuickWindow_Destroyed
func callbackQQuickWindow_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQuickWindow_DisconnectNotify
func callbackQQuickWindow_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWindowFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWindow) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWindow_ObjectNameChanged
func callbackQQuickWindow_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQuickWindow_TimerEvent
func callbackQQuickWindow_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSGAbstractRenderer struct {
	core.QObject
}

type QSGAbstractRenderer_ITF interface {
	core.QObject_ITF
	QSGAbstractRenderer_PTR() *QSGAbstractRenderer
}

func (ptr *QSGAbstractRenderer) QSGAbstractRenderer_PTR() *QSGAbstractRenderer {
	return ptr
}

func (ptr *QSGAbstractRenderer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGAbstractRenderer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGAbstractRenderer(ptr QSGAbstractRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGAbstractRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSGAbstractRendererFromPointer(ptr unsafe.Pointer) (n *QSGAbstractRenderer) {
	n = new(QSGAbstractRenderer)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGAbstractRenderer__ClearModeBit
//QSGAbstractRenderer::ClearModeBit
type QSGAbstractRenderer__ClearModeBit int64

const (
	QSGAbstractRenderer__ClearColorBuffer   QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0001)
	QSGAbstractRenderer__ClearDepthBuffer   QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0002)
	QSGAbstractRenderer__ClearStencilBuffer QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0004)
)

func QSGAbstractRenderer_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGAbstractRenderer_QSGAbstractRenderer_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSGAbstractRenderer) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGAbstractRenderer_QSGAbstractRenderer_Tr(sC, cC, C.int(int32(n))))
}

func QSGAbstractRenderer_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGAbstractRenderer_QSGAbstractRenderer_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSGAbstractRenderer) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGAbstractRenderer_QSGAbstractRenderer_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQSGAbstractRenderer_RenderScene
func callbackQSGAbstractRenderer_RenderScene(ptr unsafe.Pointer, fboId C.uint) {
	if signal := qt.GetSignal(ptr, "renderScene"); signal != nil {
		signal.(func(uint))(uint(uint32(fboId)))
	}

}

func (ptr *QSGAbstractRenderer) ConnectRenderScene(f func(fboId uint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "renderScene"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "renderScene", func(fboId uint) {
				signal.(func(uint))(fboId)
				f(fboId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "renderScene", f)
		}
	}
}

func (ptr *QSGAbstractRenderer) DisconnectRenderScene() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "renderScene")
	}
}

func (ptr *QSGAbstractRenderer) RenderScene(fboId uint) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_RenderScene(ptr.Pointer(), C.uint(uint32(fboId)))
	}
}

//export callbackQSGAbstractRenderer_SceneGraphChanged
func callbackQSGAbstractRenderer_SceneGraphChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sceneGraphChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGAbstractRenderer) ConnectSceneGraphChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sceneGraphChanged") {
			C.QSGAbstractRenderer_ConnectSceneGraphChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sceneGraphChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sceneGraphChanged", f)
		}
	}
}

func (ptr *QSGAbstractRenderer) DisconnectSceneGraphChanged() {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectSceneGraphChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sceneGraphChanged")
	}
}

func (ptr *QSGAbstractRenderer) SceneGraphChanged() {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SceneGraphChanged(ptr.Pointer())
	}
}

func (ptr *QSGAbstractRenderer) SetClearColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QSGAbstractRenderer) SetClearMode(mode QSGAbstractRenderer__ClearModeBit) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect2(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrix(matrix gui.QMatrix4x4_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrix(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrixToRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrixToRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect2(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSGAbstractRenderer) ClearColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QSGAbstractRenderer_ClearColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ProjectionMatrix() *gui.QMatrix4x4 {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQMatrix4x4FromPointer(C.QSGAbstractRenderer_ProjectionMatrix(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QMatrix4x4).DestroyQMatrix4x4)
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) DeviceRect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QSGAbstractRenderer_DeviceRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ViewportRect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QSGAbstractRenderer_ViewportRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ClearMode() QSGAbstractRenderer__ClearModeBit {
	if ptr.Pointer() != nil {
		return QSGAbstractRenderer__ClearModeBit(C.QSGAbstractRenderer_ClearMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGAbstractRenderer_MetaObject
func callbackQSGAbstractRenderer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGAbstractRendererFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGAbstractRenderer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGAbstractRenderer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSGAbstractRenderer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSGAbstractRenderer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSGAbstractRenderer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSGAbstractRenderer) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGAbstractRenderer___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGAbstractRenderer) __findChildren_newList2() unsafe.Pointer {
	return C.QSGAbstractRenderer___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSGAbstractRenderer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGAbstractRenderer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGAbstractRenderer) __findChildren_newList3() unsafe.Pointer {
	return C.QSGAbstractRenderer___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSGAbstractRenderer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGAbstractRenderer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGAbstractRenderer) __findChildren_newList() unsafe.Pointer {
	return C.QSGAbstractRenderer___findChildren_newList(ptr.Pointer())
}

func (ptr *QSGAbstractRenderer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGAbstractRenderer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGAbstractRenderer) __children_newList() unsafe.Pointer {
	return C.QSGAbstractRenderer___children_newList(ptr.Pointer())
}

//export callbackQSGAbstractRenderer_Event
func callbackQSGAbstractRenderer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGAbstractRendererFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGAbstractRenderer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGAbstractRenderer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSGAbstractRenderer_EventFilter
func callbackQSGAbstractRenderer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGAbstractRendererFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGAbstractRenderer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGAbstractRenderer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSGAbstractRenderer_ChildEvent
func callbackQSGAbstractRenderer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGAbstractRenderer_ConnectNotify
func callbackQSGAbstractRenderer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGAbstractRenderer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGAbstractRenderer_CustomEvent
func callbackQSGAbstractRenderer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGAbstractRenderer_DeleteLater
func callbackQSGAbstractRenderer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGAbstractRendererFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGAbstractRenderer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSGAbstractRenderer_Destroyed
func callbackQSGAbstractRenderer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSGAbstractRenderer_DisconnectNotify
func callbackQSGAbstractRenderer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGAbstractRenderer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGAbstractRenderer_ObjectNameChanged
func callbackQSGAbstractRenderer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSGAbstractRenderer_TimerEvent
func callbackQSGAbstractRenderer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSGBasicGeometryNode struct {
	QSGNode
}

type QSGBasicGeometryNode_ITF interface {
	QSGNode_ITF
	QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode
}

func (ptr *QSGBasicGeometryNode) QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode {
	return ptr
}

func (ptr *QSGBasicGeometryNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGBasicGeometryNode(ptr QSGBasicGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGBasicGeometryNodeFromPointer(ptr unsafe.Pointer) (n *QSGBasicGeometryNode) {
	n = new(QSGBasicGeometryNode)
	n.SetPointer(ptr)
	return
}
func (ptr *QSGBasicGeometryNode) Geometry2() *QSGGeometry {
	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometry_ITF) {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_SetGeometry(ptr.Pointer(), PointerFromQSGGeometry(geometry))
	}
}

//export callbackQSGBasicGeometryNode_DestroyQSGBasicGeometryNode
func callbackQSGBasicGeometryNode_DestroyQSGBasicGeometryNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGBasicGeometryNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGBasicGeometryNodeFromPointer(ptr).DestroyQSGBasicGeometryNodeDefault()
	}
}

func (ptr *QSGBasicGeometryNode) ConnectDestroyQSGBasicGeometryNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGBasicGeometryNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGBasicGeometryNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGBasicGeometryNode", f)
		}
	}
}

func (ptr *QSGBasicGeometryNode) DisconnectDestroyQSGBasicGeometryNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGBasicGeometryNode")
	}
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNode() {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_DestroyQSGBasicGeometryNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGBasicGeometryNode) Geometry() *QSGGeometry {
	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry(ptr.Pointer()))
	}
	return nil
}

type QSGBatchRenderer struct {
	ptr unsafe.Pointer
}

type QSGBatchRenderer_ITF interface {
	QSGBatchRenderer_PTR() *QSGBatchRenderer
}

func (ptr *QSGBatchRenderer) QSGBatchRenderer_PTR() *QSGBatchRenderer {
	return ptr
}

func (ptr *QSGBatchRenderer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSGBatchRenderer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSGBatchRenderer(ptr QSGBatchRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBatchRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSGBatchRendererFromPointer(ptr unsafe.Pointer) (n *QSGBatchRenderer) {
	n = new(QSGBatchRenderer)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGBatchRenderer) DestroyQSGBatchRenderer() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QSGBatchRenderer__BatchCompatibility
//QSGBatchRenderer::BatchCompatibility
type QSGBatchRenderer__BatchCompatibility int64

const (
	QSGBatchRenderer__BatchBreaksOnCompare QSGBatchRenderer__BatchCompatibility = QSGBatchRenderer__BatchCompatibility(0)
	QSGBatchRenderer__BatchIsCompatible    QSGBatchRenderer__BatchCompatibility = QSGBatchRenderer__BatchCompatibility(1)
)

type QSGClipNode struct {
	QSGBasicGeometryNode
}

type QSGClipNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGClipNode_PTR() *QSGClipNode
}

func (ptr *QSGClipNode) QSGClipNode_PTR() *QSGClipNode {
	return ptr
}

func (ptr *QSGClipNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGClipNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGBasicGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGClipNode(ptr QSGClipNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGClipNode_PTR().Pointer()
	}
	return nil
}

func NewQSGClipNodeFromPointer(ptr unsafe.Pointer) (n *QSGClipNode) {
	n = new(QSGClipNode)
	n.SetPointer(ptr)
	return
}
func NewQSGClipNode() *QSGClipNode {
	return NewQSGClipNodeFromPointer(C.QSGClipNode_NewQSGClipNode())
}

func (ptr *QSGClipNode) SetClipRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGClipNode_SetClipRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGClipNode) SetIsRectangular(rectHint bool) {
	if ptr.Pointer() != nil {
		C.QSGClipNode_SetIsRectangular(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(rectHint))))
	}
}

//export callbackQSGClipNode_DestroyQSGClipNode
func callbackQSGClipNode_DestroyQSGClipNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGClipNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGClipNodeFromPointer(ptr).DestroyQSGClipNodeDefault()
	}
}

func (ptr *QSGClipNode) ConnectDestroyQSGClipNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGClipNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGClipNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGClipNode", f)
		}
	}
}

func (ptr *QSGClipNode) DisconnectDestroyQSGClipNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGClipNode")
	}
}

func (ptr *QSGClipNode) DestroyQSGClipNode() {
	if ptr.Pointer() != nil {
		C.QSGClipNode_DestroyQSGClipNode(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGClipNode) DestroyQSGClipNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGClipNode_DestroyQSGClipNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGClipNode) ClipRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGClipNode_ClipRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGClipNode) IsRectangular() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGClipNode_IsRectangular(ptr.Pointer())) != 0
	}
	return false
}

type QSGDynamicTexture struct {
	QSGTexture
}

type QSGDynamicTexture_ITF interface {
	QSGTexture_ITF
	QSGDynamicTexture_PTR() *QSGDynamicTexture
}

func (ptr *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture {
	return ptr
}

func (ptr *QSGDynamicTexture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGDynamicTexture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGTexture_PTR().SetPointer(p)
	}
}

func PointerFromQSGDynamicTexture(ptr QSGDynamicTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGDynamicTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGDynamicTextureFromPointer(ptr unsafe.Pointer) (n *QSGDynamicTexture) {
	n = new(QSGDynamicTexture)
	n.SetPointer(ptr)
	return
}

//export callbackQSGDynamicTexture_UpdateTexture
func callbackQSGDynamicTexture_UpdateTexture(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "updateTexture"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGDynamicTexture) ConnectUpdateTexture(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateTexture", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateTexture", f)
		}
	}
}

func (ptr *QSGDynamicTexture) DisconnectUpdateTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateTexture")
	}
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGDynamicTexture_UpdateTexture(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_Bind
func callbackQSGDynamicTexture_Bind(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "bind"); signal != nil {
		signal.(func())()
	} else {
		NewQSGDynamicTextureFromPointer(ptr).BindDefault()
	}
}

func (ptr *QSGDynamicTexture) Bind() {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_Bind(ptr.Pointer())
	}
}

func (ptr *QSGDynamicTexture) BindDefault() {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_BindDefault(ptr.Pointer())
	}
}

//export callbackQSGDynamicTexture_TextureSize
func callbackQSGDynamicTexture_TextureSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQSGDynamicTextureFromPointer(ptr).TextureSizeDefault())
}

func (ptr *QSGDynamicTexture) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSGDynamicTexture_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSGDynamicTexture) TextureSizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSGDynamicTexture_TextureSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSGDynamicTexture_HasAlphaChannel
func callbackQSGDynamicTexture_HasAlphaChannel(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasAlphaChannel"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).HasAlphaChannelDefault())))
}

func (ptr *QSGDynamicTexture) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGDynamicTexture_HasAlphaChannel(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) HasAlphaChannelDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGDynamicTexture_HasAlphaChannelDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_HasMipmaps
func callbackQSGDynamicTexture_HasMipmaps(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasMipmaps"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).HasMipmapsDefault())))
}

func (ptr *QSGDynamicTexture) HasMipmaps() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGDynamicTexture_HasMipmaps(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) HasMipmapsDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGDynamicTexture_HasMipmapsDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_TextureId
func callbackQSGDynamicTexture_TextureId(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "textureId"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQSGDynamicTextureFromPointer(ptr).TextureIdDefault()))
}

func (ptr *QSGDynamicTexture) TextureId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGDynamicTexture_TextureId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGDynamicTexture) TextureIdDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGDynamicTexture_TextureIdDefault(ptr.Pointer())))
	}
	return 0
}

type QSGEngine struct {
	core.QObject
}

type QSGEngine_ITF interface {
	core.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func (ptr *QSGEngine) QSGEngine_PTR() *QSGEngine {
	return ptr
}

func (ptr *QSGEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGEngine(ptr QSGEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGEngine_PTR().Pointer()
	}
	return nil
}

func NewQSGEngineFromPointer(ptr unsafe.Pointer) (n *QSGEngine) {
	n = new(QSGEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGEngine__CreateTextureOption
//QSGEngine::CreateTextureOption
type QSGEngine__CreateTextureOption int64

const (
	QSGEngine__TextureHasAlphaChannel QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0001)
	QSGEngine__TextureOwnsGLTexture   QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0004)
	QSGEngine__TextureCanUseAtlas     QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0008)
	QSGEngine__TextureIsOpaque        QSGEngine__CreateTextureOption = QSGEngine__CreateTextureOption(0x0010)
)

func NewQSGEngine(parent core.QObject_ITF) *QSGEngine {
	tmpValue := NewQSGEngineFromPointer(C.QSGEngine_NewQSGEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QSGEngine_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGEngine_QSGEngine_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSGEngine) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGEngine_QSGEngine_Tr(sC, cC, C.int(int32(n))))
}

func QSGEngine_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGEngine_QSGEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSGEngine) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGEngine_QSGEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSGEngine) Initialize(context gui.QOpenGLContext_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

func (ptr *QSGEngine) Invalidate() {
	if ptr.Pointer() != nil {
		C.QSGEngine_Invalidate(ptr.Pointer())
	}
}

//export callbackQSGEngine_DestroyQSGEngine
func callbackQSGEngine_DestroyQSGEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQSGEngineFromPointer(ptr).DestroyQSGEngineDefault()
	}
}

func (ptr *QSGEngine) ConnectDestroyQSGEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGEngine", f)
		}
	}
}

func (ptr *QSGEngine) DisconnectDestroyQSGEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGEngine")
	}
}

func (ptr *QSGEngine) DestroyQSGEngine() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DestroyQSGEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGEngine) DestroyQSGEngineDefault() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DestroyQSGEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGEngine) CreateRenderer() *QSGAbstractRenderer {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGAbstractRendererFromPointer(C.QSGEngine_CreateRenderer(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) CreateImageNode() *QSGImageNode {
	if ptr.Pointer() != nil {
		return NewQSGImageNodeFromPointer(C.QSGEngine_CreateImageNode(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) CreateRectangleNode() *QSGRectangleNode {
	if ptr.Pointer() != nil {
		return NewQSGRectangleNodeFromPointer(C.QSGEngine_CreateRectangleNode(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) RendererInterface() *QSGRendererInterface {
	if ptr.Pointer() != nil {
		return NewQSGRendererInterfaceFromPointer(C.QSGEngine_RendererInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromId(id uint, size core.QSize_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGEngine_CreateTextureFromId(ptr.Pointer(), C.uint(uint32(id)), core.PointerFromQSize(size), C.longlong(options)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromImage(image gui.QImage_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGEngine_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.longlong(options)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSGEngine_MetaObject
func callbackQSGEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSGEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSGEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSGEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSGEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGEngine) __findChildren_newList2() unsafe.Pointer {
	return C.QSGEngine___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSGEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QSGEngine___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSGEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGEngine) __findChildren_newList() unsafe.Pointer {
	return C.QSGEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QSGEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGEngine) __children_newList() unsafe.Pointer {
	return C.QSGEngine___children_newList(ptr.Pointer())
}

//export callbackQSGEngine_Event
func callbackQSGEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSGEngine_EventFilter
func callbackQSGEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSGEngine_ChildEvent
func callbackQSGEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGEngine_ConnectNotify
func callbackQSGEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGEngine_CustomEvent
func callbackQSGEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGEngine_DeleteLater
func callbackQSGEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSGEngine_Destroyed
func callbackQSGEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSGEngine_DisconnectNotify
func callbackQSGEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGEngine_ObjectNameChanged
func callbackQSGEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSGEngine_TimerEvent
func callbackQSGEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSGFlatColorMaterial struct {
	QSGMaterial
}

type QSGFlatColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial
}

func (ptr *QSGFlatColorMaterial) QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial {
	return ptr
}

func (ptr *QSGFlatColorMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGFlatColorMaterial(ptr QSGFlatColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGFlatColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGFlatColorMaterialFromPointer(ptr unsafe.Pointer) (n *QSGFlatColorMaterial) {
	n = new(QSGFlatColorMaterial)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGFlatColorMaterial) DestroyQSGFlatColorMaterial() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQSGFlatColorMaterial() *QSGFlatColorMaterial {
	return NewQSGFlatColorMaterialFromPointer(C.QSGFlatColorMaterial_NewQSGFlatColorMaterial())
}

func (ptr *QSGFlatColorMaterial) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QSGFlatColorMaterial_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QSGFlatColorMaterial) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGFlatColorMaterial_Color(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGFlatColorMaterial_CreateShader
func callbackQSGFlatColorMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGFlatColorMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGFlatColorMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGFlatColorMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) CreateShaderDefault() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGFlatColorMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGFlatColorMaterial_Type
func callbackQSGFlatColorMaterial_Type(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGFlatColorMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGFlatColorMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGFlatColorMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) TypeDefault() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGFlatColorMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}

type QSGGeometry struct {
	ptr unsafe.Pointer
}

type QSGGeometry_ITF interface {
	QSGGeometry_PTR() *QSGGeometry
}

func (ptr *QSGGeometry) QSGGeometry_PTR() *QSGGeometry {
	return ptr
}

func (ptr *QSGGeometry) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSGGeometry) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSGGeometry(ptr QSGGeometry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometry_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryFromPointer(ptr unsafe.Pointer) (n *QSGGeometry) {
	n = new(QSGGeometry)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGGeometry__AttributeType
//QSGGeometry::AttributeType
type QSGGeometry__AttributeType int64

const (
	QSGGeometry__UnknownAttribute   QSGGeometry__AttributeType = QSGGeometry__AttributeType(0)
	QSGGeometry__PositionAttribute  QSGGeometry__AttributeType = QSGGeometry__AttributeType(1)
	QSGGeometry__ColorAttribute     QSGGeometry__AttributeType = QSGGeometry__AttributeType(2)
	QSGGeometry__TexCoordAttribute  QSGGeometry__AttributeType = QSGGeometry__AttributeType(3)
	QSGGeometry__TexCoord1Attribute QSGGeometry__AttributeType = QSGGeometry__AttributeType(4)
	QSGGeometry__TexCoord2Attribute QSGGeometry__AttributeType = QSGGeometry__AttributeType(5)
)

//go:generate stringer -type=QSGGeometry__DataPattern
//QSGGeometry::DataPattern
type QSGGeometry__DataPattern int64

const (
	QSGGeometry__AlwaysUploadPattern QSGGeometry__DataPattern = QSGGeometry__DataPattern(0)
	QSGGeometry__StreamPattern       QSGGeometry__DataPattern = QSGGeometry__DataPattern(1)
	QSGGeometry__DynamicPattern      QSGGeometry__DataPattern = QSGGeometry__DataPattern(2)
	QSGGeometry__StaticPattern       QSGGeometry__DataPattern = QSGGeometry__DataPattern(3)
)

//go:generate stringer -type=QSGGeometry__DrawingMode
//QSGGeometry::DrawingMode
type QSGGeometry__DrawingMode int64

const (
	QSGGeometry__DrawPoints        QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0000)
	QSGGeometry__DrawLines         QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0001)
	QSGGeometry__DrawLineLoop      QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0002)
	QSGGeometry__DrawLineStrip     QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0003)
	QSGGeometry__DrawTriangles     QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0004)
	QSGGeometry__DrawTriangleStrip QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0005)
	QSGGeometry__DrawTriangleFan   QSGGeometry__DrawingMode = QSGGeometry__DrawingMode(0x0006)
)

//go:generate stringer -type=QSGGeometry__Type
//QSGGeometry::Type
type QSGGeometry__Type int64

const (
	QSGGeometry__ByteType          QSGGeometry__Type = QSGGeometry__Type(0x1400)
	QSGGeometry__UnsignedByteType  QSGGeometry__Type = QSGGeometry__Type(0x1401)
	QSGGeometry__ShortType         QSGGeometry__Type = QSGGeometry__Type(0x1402)
	QSGGeometry__UnsignedShortType QSGGeometry__Type = QSGGeometry__Type(0x1403)
	QSGGeometry__IntType           QSGGeometry__Type = QSGGeometry__Type(0x1404)
	QSGGeometry__UnsignedIntType   QSGGeometry__Type = QSGGeometry__Type(0x1405)
	QSGGeometry__FloatType         QSGGeometry__Type = QSGGeometry__Type(0x1406)
)

func (ptr *QSGGeometry) IndexDataAsUShort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSGGeometry_IndexDataAsUShort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexDataAsUInt() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGGeometry_IndexDataAsUInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexData() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return C.QSGGeometry_IndexData(ptr.Pointer())
	}
	return nil
}

func (ptr *QSGGeometry) VertexData() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return C.QSGGeometry_VertexData(ptr.Pointer())
	}
	return nil
}

func (ptr *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_Allocate(ptr.Pointer(), C.int(int32(vertexCount)), C.int(int32(indexCount)))
	}
}

func (ptr *QSGGeometry) MarkIndexDataDirty() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkIndexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) MarkVertexDataDirty() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkVertexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) SetDrawingMode(mode uint) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetDrawingMode(ptr.Pointer(), C.uint(uint32(mode)))
	}
}

func (ptr *QSGGeometry) SetIndexDataPattern(p QSGGeometry__DataPattern) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetIndexDataPattern(ptr.Pointer(), C.longlong(p))
	}
}

func (ptr *QSGGeometry) SetLineWidth(width float32) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetLineWidth(ptr.Pointer(), C.float(width))
	}
}

func (ptr *QSGGeometry) SetVertexDataPattern(p QSGGeometry__DataPattern) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetVertexDataPattern(ptr.Pointer(), C.longlong(p))
	}
}

func QSGGeometry_UpdateColoredRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateColoredRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func (ptr *QSGGeometry) UpdateColoredRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateColoredRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func QSGGeometry_UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func (ptr *QSGGeometry) UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect), core.PointerFromQRectF(textureRect))
}

func (ptr *QSGGeometry) UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect), core.PointerFromQRectF(textureRect))
}

//export callbackQSGGeometry_DestroyQSGGeometry
func callbackQSGGeometry_DestroyQSGGeometry(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGGeometry"); signal != nil {
		signal.(func())()
	} else {
		NewQSGGeometryFromPointer(ptr).DestroyQSGGeometryDefault()
	}
}

func (ptr *QSGGeometry) ConnectDestroyQSGGeometry(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGGeometry"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGGeometry", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGGeometry", f)
		}
	}
}

func (ptr *QSGGeometry) DisconnectDestroyQSGGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGGeometry")
	}
}

func (ptr *QSGGeometry) DestroyQSGGeometry() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometry(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGGeometry) DestroyQSGGeometryDefault() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGGeometry) IndexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_IndexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) VertexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_VertexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexDataAsUShort2() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSGGeometry_IndexDataAsUShort2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexDataAsUInt2() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGGeometry_IndexDataAsUInt2(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexData2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return C.QSGGeometry_IndexData2(ptr.Pointer())
	}
	return nil
}

func (ptr *QSGGeometry) VertexData2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return C.QSGGeometry_VertexData2(ptr.Pointer())
	}
	return nil
}

func (ptr *QSGGeometry) LineWidth() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QSGGeometry_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) AttributeCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_AttributeCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_IndexCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexType() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_IndexType(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) SizeOfIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_SizeOfIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) SizeOfVertex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_SizeOfVertex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) VertexCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_VertexCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) DrawingMode() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGGeometry_DrawingMode(ptr.Pointer())))
	}
	return 0
}

type QSGGeometryNode struct {
	QSGBasicGeometryNode
}

type QSGGeometryNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGGeometryNode_PTR() *QSGGeometryNode
}

func (ptr *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode {
	return ptr
}

func (ptr *QSGGeometryNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGGeometryNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGBasicGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGGeometryNode(ptr QSGGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryNodeFromPointer(ptr unsafe.Pointer) (n *QSGGeometryNode) {
	n = new(QSGGeometryNode)
	n.SetPointer(ptr)
	return
}
func NewQSGGeometryNode() *QSGGeometryNode {
	return NewQSGGeometryNodeFromPointer(C.QSGGeometryNode_NewQSGGeometryNode())
}

func (ptr *QSGGeometryNode) SetMaterial(material QSGMaterial_ITF) {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterial_ITF) {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetOpaqueMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

//export callbackQSGGeometryNode_DestroyQSGGeometryNode
func callbackQSGGeometryNode_DestroyQSGGeometryNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGGeometryNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGGeometryNodeFromPointer(ptr).DestroyQSGGeometryNodeDefault()
	}
}

func (ptr *QSGGeometryNode) ConnectDestroyQSGGeometryNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGGeometryNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGGeometryNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGGeometryNode", f)
		}
	}
}

func (ptr *QSGGeometryNode) DisconnectDestroyQSGGeometryNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGGeometryNode")
	}
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGGeometryNode) Material() *QSGMaterial {
	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_Material(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) OpaqueMaterial() *QSGMaterial {
	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_OpaqueMaterial(ptr.Pointer()))
	}
	return nil
}

type QSGImageNode struct {
	QSGGeometryNode
}

type QSGImageNode_ITF interface {
	QSGGeometryNode_ITF
	QSGImageNode_PTR() *QSGImageNode
}

func (ptr *QSGImageNode) QSGImageNode_PTR() *QSGImageNode {
	return ptr
}

func (ptr *QSGImageNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGImageNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGImageNode(ptr QSGImageNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGImageNode_PTR().Pointer()
	}
	return nil
}

func NewQSGImageNodeFromPointer(ptr unsafe.Pointer) (n *QSGImageNode) {
	n = new(QSGImageNode)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGImageNode__TextureCoordinatesTransformFlag
//QSGImageNode::TextureCoordinatesTransformFlag
type QSGImageNode__TextureCoordinatesTransformFlag int64

const (
	QSGImageNode__NoTransform        QSGImageNode__TextureCoordinatesTransformFlag = QSGImageNode__TextureCoordinatesTransformFlag(0x00)
	QSGImageNode__MirrorHorizontally QSGImageNode__TextureCoordinatesTransformFlag = QSGImageNode__TextureCoordinatesTransformFlag(0x01)
	QSGImageNode__MirrorVertically   QSGImageNode__TextureCoordinatesTransformFlag = QSGImageNode__TextureCoordinatesTransformFlag(0x02)
)

func QSGImageNode_RebuildGeometry(g QSGGeometry_ITF, texture QSGTexture_ITF, rect core.QRectF_ITF, sourceRect core.QRectF_ITF, texCoordMode QSGImageNode__TextureCoordinatesTransformFlag) {
	C.QSGImageNode_QSGImageNode_RebuildGeometry(PointerFromQSGGeometry(g), PointerFromQSGTexture(texture), core.PointerFromQRectF(rect), core.PointerFromQRectF(sourceRect), C.longlong(texCoordMode))
}

func (ptr *QSGImageNode) RebuildGeometry(g QSGGeometry_ITF, texture QSGTexture_ITF, rect core.QRectF_ITF, sourceRect core.QRectF_ITF, texCoordMode QSGImageNode__TextureCoordinatesTransformFlag) {
	C.QSGImageNode_QSGImageNode_RebuildGeometry(PointerFromQSGGeometry(g), PointerFromQSGTexture(texture), core.PointerFromQRectF(rect), core.PointerFromQRectF(sourceRect), C.longlong(texCoordMode))
}

//export callbackQSGImageNode_SetFiltering
func callbackQSGImageNode_SetFiltering(ptr unsafe.Pointer, filtering C.longlong) {
	if signal := qt.GetSignal(ptr, "setFiltering"); signal != nil {
		signal.(func(QSGTexture__Filtering))(QSGTexture__Filtering(filtering))
	}

}

func (ptr *QSGImageNode) ConnectSetFiltering(f func(filtering QSGTexture__Filtering)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFiltering"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFiltering", func(filtering QSGTexture__Filtering) {
				signal.(func(QSGTexture__Filtering))(filtering)
				f(filtering)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFiltering", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSetFiltering() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFiltering")
	}
}

func (ptr *QSGImageNode) SetFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetFiltering(ptr.Pointer(), C.longlong(filtering))
	}
}

//export callbackQSGImageNode_SetMipmapFiltering
func callbackQSGImageNode_SetMipmapFiltering(ptr unsafe.Pointer, filtering C.longlong) {
	if signal := qt.GetSignal(ptr, "setMipmapFiltering"); signal != nil {
		signal.(func(QSGTexture__Filtering))(QSGTexture__Filtering(filtering))
	}

}

func (ptr *QSGImageNode) ConnectSetMipmapFiltering(f func(filtering QSGTexture__Filtering)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMipmapFiltering"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setMipmapFiltering", func(filtering QSGTexture__Filtering) {
				signal.(func(QSGTexture__Filtering))(filtering)
				f(filtering)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMipmapFiltering", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSetMipmapFiltering() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMipmapFiltering")
	}
}

func (ptr *QSGImageNode) SetMipmapFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetMipmapFiltering(ptr.Pointer(), C.longlong(filtering))
	}
}

//export callbackQSGImageNode_SetOwnsTexture
func callbackQSGImageNode_SetOwnsTexture(ptr unsafe.Pointer, owns C.char) {
	if signal := qt.GetSignal(ptr, "setOwnsTexture"); signal != nil {
		signal.(func(bool))(int8(owns) != 0)
	}

}

func (ptr *QSGImageNode) ConnectSetOwnsTexture(f func(owns bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOwnsTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setOwnsTexture", func(owns bool) {
				signal.(func(bool))(owns)
				f(owns)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOwnsTexture", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSetOwnsTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOwnsTexture")
	}
}

func (ptr *QSGImageNode) SetOwnsTexture(owns bool) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetOwnsTexture(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(owns))))
	}
}

//export callbackQSGImageNode_SetRect
func callbackQSGImageNode_SetRect(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setRect"); signal != nil {
		signal.(func(*core.QRectF))(core.NewQRectFFromPointer(rect))
	}

}

func (ptr *QSGImageNode) ConnectSetRect(f func(rect *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRect", func(rect *core.QRectF) {
				signal.(func(*core.QRectF))(rect)
				f(rect)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRect", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSetRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRect")
	}
}

func (ptr *QSGImageNode) SetRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGImageNode) SetRect2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

//export callbackQSGImageNode_SetSourceRect
func callbackQSGImageNode_SetSourceRect(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setSourceRect"); signal != nil {
		signal.(func(*core.QRectF))(core.NewQRectFFromPointer(rect))
	}

}

func (ptr *QSGImageNode) ConnectSetSourceRect(f func(rect *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSourceRect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSourceRect", func(rect *core.QRectF) {
				signal.(func(*core.QRectF))(rect)
				f(rect)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSourceRect", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSetSourceRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSourceRect")
	}
}

func (ptr *QSGImageNode) SetSourceRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetSourceRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGImageNode) SetSourceRect2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetSourceRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

//export callbackQSGImageNode_SetTexture
func callbackQSGImageNode_SetTexture(ptr unsafe.Pointer, texture unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setTexture"); signal != nil {
		signal.(func(*QSGTexture))(NewQSGTextureFromPointer(texture))
	}

}

func (ptr *QSGImageNode) ConnectSetTexture(f func(texture *QSGTexture)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTexture", func(texture *QSGTexture) {
				signal.(func(*QSGTexture))(texture)
				f(texture)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTexture", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSetTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTexture")
	}
}

func (ptr *QSGImageNode) SetTexture(texture QSGTexture_ITF) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetTexture(ptr.Pointer(), PointerFromQSGTexture(texture))
	}
}

//export callbackQSGImageNode_SetTextureCoordinatesTransform
func callbackQSGImageNode_SetTextureCoordinatesTransform(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "setTextureCoordinatesTransform"); signal != nil {
		signal.(func(QSGImageNode__TextureCoordinatesTransformFlag))(QSGImageNode__TextureCoordinatesTransformFlag(mode))
	}

}

func (ptr *QSGImageNode) ConnectSetTextureCoordinatesTransform(f func(mode QSGImageNode__TextureCoordinatesTransformFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTextureCoordinatesTransform"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTextureCoordinatesTransform", func(mode QSGImageNode__TextureCoordinatesTransformFlag) {
				signal.(func(QSGImageNode__TextureCoordinatesTransformFlag))(mode)
				f(mode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTextureCoordinatesTransform", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSetTextureCoordinatesTransform() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTextureCoordinatesTransform")
	}
}

func (ptr *QSGImageNode) SetTextureCoordinatesTransform(mode QSGImageNode__TextureCoordinatesTransformFlag) {
	if ptr.Pointer() != nil {
		C.QSGImageNode_SetTextureCoordinatesTransform(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQSGImageNode_DestroyQSGImageNode
func callbackQSGImageNode_DestroyQSGImageNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGImageNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGImageNodeFromPointer(ptr).DestroyQSGImageNodeDefault()
	}
}

func (ptr *QSGImageNode) ConnectDestroyQSGImageNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGImageNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGImageNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGImageNode", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectDestroyQSGImageNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGImageNode")
	}
}

func (ptr *QSGImageNode) DestroyQSGImageNode() {
	if ptr.Pointer() != nil {
		C.QSGImageNode_DestroyQSGImageNode(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGImageNode) DestroyQSGImageNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGImageNode_DestroyQSGImageNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSGImageNode_Rect
func callbackQSGImageNode_Rect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "rect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(core.NewQRectF())
}

func (ptr *QSGImageNode) ConnectRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rect", func() *core.QRectF {
				signal.(func() *core.QRectF)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rect", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rect")
	}
}

func (ptr *QSGImageNode) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGImageNode_Rect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQSGImageNode_SourceRect
func callbackQSGImageNode_SourceRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sourceRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(core.NewQRectF())
}

func (ptr *QSGImageNode) ConnectSourceRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sourceRect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sourceRect", func() *core.QRectF {
				signal.(func() *core.QRectF)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sourceRect", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectSourceRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sourceRect")
	}
}

func (ptr *QSGImageNode) SourceRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGImageNode_SourceRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQSGImageNode_TextureCoordinatesTransform
func callbackQSGImageNode_TextureCoordinatesTransform(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "textureCoordinatesTransform"); signal != nil {
		return C.longlong(signal.(func() QSGImageNode__TextureCoordinatesTransformFlag)())
	}

	return C.longlong(0)
}

func (ptr *QSGImageNode) ConnectTextureCoordinatesTransform(f func() QSGImageNode__TextureCoordinatesTransformFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "textureCoordinatesTransform"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureCoordinatesTransform", func() QSGImageNode__TextureCoordinatesTransformFlag {
				signal.(func() QSGImageNode__TextureCoordinatesTransformFlag)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureCoordinatesTransform", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectTextureCoordinatesTransform() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "textureCoordinatesTransform")
	}
}

func (ptr *QSGImageNode) TextureCoordinatesTransform() QSGImageNode__TextureCoordinatesTransformFlag {
	if ptr.Pointer() != nil {
		return QSGImageNode__TextureCoordinatesTransformFlag(C.QSGImageNode_TextureCoordinatesTransform(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGImageNode_Texture
func callbackQSGImageNode_Texture(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "texture"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGTexture())
}

func (ptr *QSGImageNode) ConnectTexture(f func() *QSGTexture) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "texture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "texture", func() *QSGTexture {
				signal.(func() *QSGTexture)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "texture", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "texture")
	}
}

func (ptr *QSGImageNode) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGImageNode_Texture(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSGImageNode_Filtering
func callbackQSGImageNode_Filtering(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "filtering"); signal != nil {
		return C.longlong(signal.(func() QSGTexture__Filtering)())
	}

	return C.longlong(0)
}

func (ptr *QSGImageNode) ConnectFiltering(f func() QSGTexture__Filtering) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filtering"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "filtering", func() QSGTexture__Filtering {
				signal.(func() QSGTexture__Filtering)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filtering", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectFiltering() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filtering")
	}
}

func (ptr *QSGImageNode) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGImageNode_Filtering(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGImageNode_MipmapFiltering
func callbackQSGImageNode_MipmapFiltering(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "mipmapFiltering"); signal != nil {
		return C.longlong(signal.(func() QSGTexture__Filtering)())
	}

	return C.longlong(0)
}

func (ptr *QSGImageNode) ConnectMipmapFiltering(f func() QSGTexture__Filtering) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mipmapFiltering"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mipmapFiltering", func() QSGTexture__Filtering {
				signal.(func() QSGTexture__Filtering)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mipmapFiltering", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectMipmapFiltering() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mipmapFiltering")
	}
}

func (ptr *QSGImageNode) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGImageNode_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGImageNode_OwnsTexture
func callbackQSGImageNode_OwnsTexture(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "ownsTexture"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGImageNode) ConnectOwnsTexture(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ownsTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "ownsTexture", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ownsTexture", f)
		}
	}
}

func (ptr *QSGImageNode) DisconnectOwnsTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ownsTexture")
	}
}

func (ptr *QSGImageNode) OwnsTexture() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGImageNode_OwnsTexture(ptr.Pointer())) != 0
	}
	return false
}

type QSGMaterial struct {
	ptr unsafe.Pointer
}

type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (ptr *QSGMaterial) QSGMaterial_PTR() *QSGMaterial {
	return ptr
}

func (ptr *QSGMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSGMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSGMaterial(ptr QSGMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialFromPointer(ptr unsafe.Pointer) (n *QSGMaterial) {
	n = new(QSGMaterial)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGMaterial) DestroyQSGMaterial() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QSGMaterial__Flag
//QSGMaterial::Flag
type QSGMaterial__Flag int64

const (
	QSGMaterial__Blending                          QSGMaterial__Flag = QSGMaterial__Flag(0x0001)
	QSGMaterial__RequiresDeterminant               QSGMaterial__Flag = QSGMaterial__Flag(0x0002)
	QSGMaterial__RequiresFullMatrixExceptTranslate QSGMaterial__Flag = QSGMaterial__Flag(0x0004 | QSGMaterial__RequiresDeterminant)
	QSGMaterial__RequiresFullMatrix                QSGMaterial__Flag = QSGMaterial__Flag(0x0008 | QSGMaterial__RequiresFullMatrixExceptTranslate)
	QSGMaterial__CustomCompileStep                 QSGMaterial__Flag = QSGMaterial__Flag(0x0010)
)

func (ptr *QSGMaterial) SetFlag(flags QSGMaterial__Flag, on bool) {
	if ptr.Pointer() != nil {
		C.QSGMaterial_SetFlag(ptr.Pointer(), C.longlong(flags), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QSGMaterial) Flags() QSGMaterial__Flag {
	if ptr.Pointer() != nil {
		return QSGMaterial__Flag(C.QSGMaterial_Flags(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGMaterial_CreateShader
func callbackQSGMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(nil)
}

func (ptr *QSGMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createShader"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createShader", func() *QSGMaterialShader {
				signal.(func() *QSGMaterialShader)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createShader", f)
		}
	}
}

func (ptr *QSGMaterial) DisconnectCreateShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createShader")
	}
}

func (ptr *QSGMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGMaterial_Type
func callbackQSGMaterial_Type(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(nil)
}

func (ptr *QSGMaterial) ConnectType(f func() *QSGMaterialType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "type", func() *QSGMaterialType {
				signal.(func() *QSGMaterialType)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", f)
		}
	}
}

func (ptr *QSGMaterial) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QSGMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGMaterial_Type(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGMaterial_Compare
func callbackQSGMaterial_Compare(ptr unsafe.Pointer, other unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "compare"); signal != nil {
		return C.int(int32(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other))))
	}

	return C.int(int32(NewQSGMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other))))
}

func (ptr *QSGMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "compare"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "compare", func(other *QSGMaterial) int {
				signal.(func(*QSGMaterial) int)(other)
				return f(other)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "compare", f)
		}
	}
}

func (ptr *QSGMaterial) DisconnectCompare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "compare")
	}
}

func (ptr *QSGMaterial) Compare(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

func (ptr *QSGMaterial) CompareDefault(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

type QSGMaterialShader struct {
	ptr unsafe.Pointer
}

type QSGMaterialShader_ITF interface {
	QSGMaterialShader_PTR() *QSGMaterialShader
}

func (ptr *QSGMaterialShader) QSGMaterialShader_PTR() *QSGMaterialShader {
	return ptr
}

func (ptr *QSGMaterialShader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSGMaterialShader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSGMaterialShader(ptr QSGMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialShaderFromPointer(ptr unsafe.Pointer) (n *QSGMaterialShader) {
	n = new(QSGMaterialShader)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGMaterialShader) DestroyQSGMaterialShader() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQOpenGLShaderProgramFromPointer(C.QSGMaterialShader_Program(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSGMaterialShader_Activate
func callbackQSGMaterialShader_Activate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).ActivateDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectActivate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activate", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activate", f)
		}
	}
}

func (ptr *QSGMaterialShader) DisconnectActivate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activate")
	}
}

func (ptr *QSGMaterialShader) Activate() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Activate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ActivateDefault() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_ActivateDefault(ptr.Pointer())
	}
}

//export callbackQSGMaterialShader_Compile
func callbackQSGMaterialShader_Compile(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "compile"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).CompileDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectCompile(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "compile"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "compile", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "compile", f)
		}
	}
}

func (ptr *QSGMaterialShader) DisconnectCompile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "compile")
	}
}

func (ptr *QSGMaterialShader) Compile() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Compile(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) CompileDefault() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_CompileDefault(ptr.Pointer())
	}
}

//export callbackQSGMaterialShader_Deactivate
func callbackQSGMaterialShader_Deactivate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deactivate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).DeactivateDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectDeactivate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deactivate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deactivate", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deactivate", f)
		}
	}
}

func (ptr *QSGMaterialShader) DisconnectDeactivate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deactivate")
	}
}

func (ptr *QSGMaterialShader) Deactivate() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Deactivate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) DeactivateDefault() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_DeactivateDefault(ptr.Pointer())
	}
}

//export callbackQSGMaterialShader_Initialize
func callbackQSGMaterialShader_Initialize(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initialize"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).InitializeDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectInitialize(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "initialize"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initialize", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initialize", f)
		}
	}
}

func (ptr *QSGMaterialShader) DisconnectInitialize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "initialize")
	}
}

func (ptr *QSGMaterialShader) Initialize() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Initialize(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) InitializeDefault() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_InitializeDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) SetShaderSourceFile(ty gui.QOpenGLShader__ShaderTypeBit, sourceFile string) {
	if ptr.Pointer() != nil {
		var sourceFileC *C.char
		if sourceFile != "" {
			sourceFileC = C.CString(sourceFile)
			defer C.free(unsafe.Pointer(sourceFileC))
		}
		C.QSGMaterialShader_SetShaderSourceFile(ptr.Pointer(), C.longlong(ty), C.struct_QtQuick_PackedString{data: sourceFileC, len: C.longlong(len(sourceFile))})
	}
}

func (ptr *QSGMaterialShader) SetShaderSourceFiles(ty gui.QOpenGLShader__ShaderTypeBit, sourceFiles []string) {
	if ptr.Pointer() != nil {
		sourceFilesC := C.CString(strings.Join(sourceFiles, "|"))
		defer C.free(unsafe.Pointer(sourceFilesC))
		C.QSGMaterialShader_SetShaderSourceFiles(ptr.Pointer(), C.longlong(ty), C.struct_QtQuick_PackedString{data: sourceFilesC, len: C.longlong(len(strings.Join(sourceFiles, "|")))})
	}
}

//export callbackQSGMaterialShader_FragmentShader
func callbackQSGMaterialShader_FragmentShader(ptr unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(ptr, "fragmentShader"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSGMaterialShaderFromPointer(ptr).FragmentShaderDefault())
}

func (ptr *QSGMaterialShader) ConnectFragmentShader(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fragmentShader"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fragmentShader", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fragmentShader", f)
		}
	}
}

func (ptr *QSGMaterialShader) DisconnectFragmentShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fragmentShader")
	}
}

func (ptr *QSGMaterialShader) FragmentShader() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSGMaterialShader_FragmentShader(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGMaterialShader) FragmentShaderDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSGMaterialShader_FragmentShaderDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQSGMaterialShader_VertexShader
func callbackQSGMaterialShader_VertexShader(ptr unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(ptr, "vertexShader"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSGMaterialShaderFromPointer(ptr).VertexShaderDefault())
}

func (ptr *QSGMaterialShader) ConnectVertexShader(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "vertexShader"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "vertexShader", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "vertexShader", f)
		}
	}
}

func (ptr *QSGMaterialShader) DisconnectVertexShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "vertexShader")
	}
}

func (ptr *QSGMaterialShader) VertexShader() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSGMaterialShader_VertexShader(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGMaterialShader) VertexShaderDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSGMaterialShader_VertexShaderDefault(ptr.Pointer()))
	}
	return ""
}

type QSGMaterialType struct {
	ptr unsafe.Pointer
}

type QSGMaterialType_ITF interface {
	QSGMaterialType_PTR() *QSGMaterialType
}

func (ptr *QSGMaterialType) QSGMaterialType_PTR() *QSGMaterialType {
	return ptr
}

func (ptr *QSGMaterialType) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSGMaterialType) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSGMaterialType(ptr QSGMaterialType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialType_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialTypeFromPointer(ptr unsafe.Pointer) (n *QSGMaterialType) {
	n = new(QSGMaterialType)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGMaterialType) DestroyQSGMaterialType() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QSGNode struct {
	ptr unsafe.Pointer
}

type QSGNode_ITF interface {
	QSGNode_PTR() *QSGNode
}

func (ptr *QSGNode) QSGNode_PTR() *QSGNode {
	return ptr
}

func (ptr *QSGNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSGNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSGNode(ptr QSGNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func NewQSGNodeFromPointer(ptr unsafe.Pointer) (n *QSGNode) {
	n = new(QSGNode)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGNode__DirtyStateBit
//QSGNode::DirtyStateBit
type QSGNode__DirtyStateBit int64

const (
	QSGNode__DirtySubtreeBlocked QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0080)
	QSGNode__DirtyMatrix         QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0100)
	QSGNode__DirtyNodeAdded      QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0400)
	QSGNode__DirtyNodeRemoved    QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x0800)
	QSGNode__DirtyGeometry       QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x1000)
	QSGNode__DirtyMaterial       QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x2000)
	QSGNode__DirtyOpacity        QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x4000)
	QSGNode__DirtyForceUpdate    QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(0x8000)
	QSGNode__DirtyUsePreprocess  QSGNode__DirtyStateBit = QSGNode__DirtyStateBit(QSGNode__UsePreprocess)
)

//go:generate stringer -type=QSGNode__Flag
//QSGNode::Flag
type QSGNode__Flag int64

const (
	QSGNode__OwnedByParent      QSGNode__Flag = QSGNode__Flag(0x0001)
	QSGNode__UsePreprocess      QSGNode__Flag = QSGNode__Flag(0x0002)
	QSGNode__OwnsGeometry       QSGNode__Flag = QSGNode__Flag(0x00010000)
	QSGNode__OwnsMaterial       QSGNode__Flag = QSGNode__Flag(0x00020000)
	QSGNode__OwnsOpaqueMaterial QSGNode__Flag = QSGNode__Flag(0x00040000)
	QSGNode__IsVisitableNode    QSGNode__Flag = QSGNode__Flag(0x01000000)
	QSGNode__InternalReserved   QSGNode__Flag = QSGNode__Flag(0x01000000)
)

//go:generate stringer -type=QSGNode__NodeType
//QSGNode::NodeType
type QSGNode__NodeType int64

const (
	QSGNode__BasicNodeType     QSGNode__NodeType = QSGNode__NodeType(0)
	QSGNode__GeometryNodeType  QSGNode__NodeType = QSGNode__NodeType(1)
	QSGNode__TransformNodeType QSGNode__NodeType = QSGNode__NodeType(2)
	QSGNode__ClipNodeType      QSGNode__NodeType = QSGNode__NodeType(3)
	QSGNode__OpacityNodeType   QSGNode__NodeType = QSGNode__NodeType(4)
	QSGNode__RootNodeType      QSGNode__NodeType = QSGNode__NodeType(5)
	QSGNode__RenderNodeType    QSGNode__NodeType = QSGNode__NodeType(6)
)

func NewQSGNode() *QSGNode {
	return NewQSGNodeFromPointer(C.QSGNode_NewQSGNode())
}

func (ptr *QSGNode) AppendChildNode(node QSGNode_ITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_AppendChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) InsertChildNodeAfter(node QSGNode_ITF, after QSGNode_ITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeAfter(ptr.Pointer(), PointerFromQSGNode(node), PointerFromQSGNode(after))
	}
}

func (ptr *QSGNode) InsertChildNodeBefore(node QSGNode_ITF, before QSGNode_ITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeBefore(ptr.Pointer(), PointerFromQSGNode(node), PointerFromQSGNode(before))
	}
}

func (ptr *QSGNode) MarkDirty(bits QSGNode__DirtyStateBit) {
	if ptr.Pointer() != nil {
		C.QSGNode_MarkDirty(ptr.Pointer(), C.longlong(bits))
	}
}

func (ptr *QSGNode) PrependChildNode(node QSGNode_ITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_PrependChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

//export callbackQSGNode_Preprocess
func callbackQSGNode_Preprocess(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "preprocess"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "preprocess", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preprocess", f)
		}
	}
}

func (ptr *QSGNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "preprocess")
	}
}

func (ptr *QSGNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGNode_PreprocessDefault(ptr.Pointer())
	}
}

func (ptr *QSGNode) RemoveAllChildNodes() {
	if ptr.Pointer() != nil {
		C.QSGNode_RemoveAllChildNodes(ptr.Pointer())
	}
}

func (ptr *QSGNode) RemoveChildNode(node QSGNode_ITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_RemoveChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) SetFlag(ff QSGNode__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QSGNode_SetFlag(ptr.Pointer(), C.longlong(ff), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QSGNode) SetFlags(ff QSGNode__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QSGNode_SetFlags(ptr.Pointer(), C.longlong(ff), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQSGNode_DestroyQSGNode
func callbackQSGNode_DestroyQSGNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGNodeFromPointer(ptr).DestroyQSGNodeDefault()
	}
}

func (ptr *QSGNode) ConnectDestroyQSGNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGNode", f)
		}
	}
}

func (ptr *QSGNode) DisconnectDestroyQSGNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGNode")
	}
}

func (ptr *QSGNode) DestroyQSGNode() {
	if ptr.Pointer() != nil {
		C.QSGNode_DestroyQSGNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGNode) DestroyQSGNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGNode_DestroyQSGNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGNode) ChildAtIndex(i int) *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_ChildAtIndex(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QSGNode) FirstChild() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_FirstChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) LastChild() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_LastChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) NextSibling() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_NextSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) Parent() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) PreviousSibling() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_PreviousSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) Flags() QSGNode__Flag {
	if ptr.Pointer() != nil {
		return QSGNode__Flag(C.QSGNode_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGNode) Type() QSGNode__NodeType {
	if ptr.Pointer() != nil {
		return QSGNode__NodeType(C.QSGNode_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGNode_IsSubtreeBlocked
func callbackQSGNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isSubtreeBlocked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isSubtreeBlocked", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isSubtreeBlocked", f)
		}
	}
}

func (ptr *QSGNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isSubtreeBlocked")
	}
}

func (ptr *QSGNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGNode_IsSubtreeBlocked(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGNode_IsSubtreeBlockedDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGNode) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGNode_ChildCount(ptr.Pointer())))
	}
	return 0
}

type QSGOpacityNode struct {
	QSGNode
}

type QSGOpacityNode_ITF interface {
	QSGNode_ITF
	QSGOpacityNode_PTR() *QSGOpacityNode
}

func (ptr *QSGOpacityNode) QSGOpacityNode_PTR() *QSGOpacityNode {
	return ptr
}

func (ptr *QSGOpacityNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpacityNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpacityNode(ptr QSGOpacityNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpacityNode_PTR().Pointer()
	}
	return nil
}

func NewQSGOpacityNodeFromPointer(ptr unsafe.Pointer) (n *QSGOpacityNode) {
	n = new(QSGOpacityNode)
	n.SetPointer(ptr)
	return
}
func NewQSGOpacityNode() *QSGOpacityNode {
	return NewQSGOpacityNodeFromPointer(C.QSGOpacityNode_NewQSGOpacityNode())
}

func (ptr *QSGOpacityNode) SetOpacity(opacity float64) {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

//export callbackQSGOpacityNode_DestroyQSGOpacityNode
func callbackQSGOpacityNode_DestroyQSGOpacityNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGOpacityNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGOpacityNodeFromPointer(ptr).DestroyQSGOpacityNodeDefault()
	}
}

func (ptr *QSGOpacityNode) ConnectDestroyQSGOpacityNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGOpacityNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGOpacityNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGOpacityNode", f)
		}
	}
}

func (ptr *QSGOpacityNode) DisconnectDestroyQSGOpacityNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGOpacityNode")
	}
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNode() {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_DestroyQSGOpacityNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_DestroyQSGOpacityNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGOpacityNode) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QSGOpacityNode_Opacity(ptr.Pointer()))
	}
	return 0
}

type QSGOpaqueTextureMaterial struct {
	QSGMaterial
}

type QSGOpaqueTextureMaterial_ITF interface {
	QSGMaterial_ITF
	QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial
}

func (ptr *QSGOpaqueTextureMaterial) QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial {
	return ptr
}

func (ptr *QSGOpaqueTextureMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGOpaqueTextureMaterial(ptr QSGOpaqueTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGOpaqueTextureMaterialFromPointer(ptr unsafe.Pointer) (n *QSGOpaqueTextureMaterial) {
	n = new(QSGOpaqueTextureMaterial)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGOpaqueTextureMaterial) DestroyQSGOpaqueTextureMaterial() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQSGOpaqueTextureMaterial() *QSGOpaqueTextureMaterial {
	return NewQSGOpaqueTextureMaterialFromPointer(C.QSGOpaqueTextureMaterial_NewQSGOpaqueTextureMaterial())
}

func (ptr *QSGOpaqueTextureMaterial) SetAnisotropyLevel(level QSGTexture__AnisotropyLevel) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetAnisotropyLevel(ptr.Pointer(), C.longlong(level))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetFiltering(ptr.Pointer(), C.longlong(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetHorizontalWrapMode(mode QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetHorizontalWrapMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetMipmapFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetMipmapFiltering(ptr.Pointer(), C.longlong(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetTexture(texture QSGTexture_ITF) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetTexture(ptr.Pointer(), PointerFromQSGTexture(texture))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetVerticalWrapMode(mode QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetVerticalWrapMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGOpaqueTextureMaterial_Texture(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) AnisotropyLevel() QSGTexture__AnisotropyLevel {
	if ptr.Pointer() != nil {
		return QSGTexture__AnisotropyLevel(C.QSGOpaqueTextureMaterial_AnisotropyLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) HorizontalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) VerticalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) M_anisotropy_level() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGOpaqueTextureMaterial_M_anisotropy_level(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetM_anisotropy_level(vui uint) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_anisotropy_level(ptr.Pointer(), C.uint(uint32(vui)))
	}
}

func (ptr *QSGOpaqueTextureMaterial) M_filtering() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGOpaqueTextureMaterial_M_filtering(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetM_filtering(vui uint) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_filtering(ptr.Pointer(), C.uint(uint32(vui)))
	}
}

func (ptr *QSGOpaqueTextureMaterial) M_horizontal_wrap() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGOpaqueTextureMaterial_M_horizontal_wrap(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetM_horizontal_wrap(vui uint) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_horizontal_wrap(ptr.Pointer(), C.uint(uint32(vui)))
	}
}

func (ptr *QSGOpaqueTextureMaterial) M_mipmap_filtering() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGOpaqueTextureMaterial_M_mipmap_filtering(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetM_mipmap_filtering(vui uint) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_mipmap_filtering(ptr.Pointer(), C.uint(uint32(vui)))
	}
}

func (ptr *QSGOpaqueTextureMaterial) M_reserved() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGOpaqueTextureMaterial_M_reserved(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetM_reserved(vui uint) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_reserved(ptr.Pointer(), C.uint(uint32(vui)))
	}
}

func (ptr *QSGOpaqueTextureMaterial) M_texture() *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGOpaqueTextureMaterial_M_texture(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) SetM_texture(vqs QSGTexture_ITF) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_texture(ptr.Pointer(), PointerFromQSGTexture(vqs))
	}
}

func (ptr *QSGOpaqueTextureMaterial) M_vertical_wrap() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGOpaqueTextureMaterial_M_vertical_wrap(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetM_vertical_wrap(vui uint) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_vertical_wrap(ptr.Pointer(), C.uint(uint32(vui)))
	}
}

//export callbackQSGOpaqueTextureMaterial_CreateShader
func callbackQSGOpaqueTextureMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGOpaqueTextureMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGOpaqueTextureMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGOpaqueTextureMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) CreateShaderDefault() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGOpaqueTextureMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGOpaqueTextureMaterial_Type
func callbackQSGOpaqueTextureMaterial_Type(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGOpaqueTextureMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGOpaqueTextureMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGOpaqueTextureMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) TypeDefault() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGOpaqueTextureMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}

type QSGRectangleNode struct {
	QSGGeometryNode
}

type QSGRectangleNode_ITF interface {
	QSGGeometryNode_ITF
	QSGRectangleNode_PTR() *QSGRectangleNode
}

func (ptr *QSGRectangleNode) QSGRectangleNode_PTR() *QSGRectangleNode {
	return ptr
}

func (ptr *QSGRectangleNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGRectangleNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGRectangleNode(ptr QSGRectangleNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGRectangleNode_PTR().Pointer()
	}
	return nil
}

func NewQSGRectangleNodeFromPointer(ptr unsafe.Pointer) (n *QSGRectangleNode) {
	n = new(QSGRectangleNode)
	n.SetPointer(ptr)
	return
}

//export callbackQSGRectangleNode_SetColor
func callbackQSGRectangleNode_SetColor(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setColor"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QSGRectangleNode) ConnectSetColor(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setColor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setColor", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setColor", f)
		}
	}
}

func (ptr *QSGRectangleNode) DisconnectSetColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setColor")
	}
}

func (ptr *QSGRectangleNode) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QSGRectangleNode_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQSGRectangleNode_SetRect
func callbackQSGRectangleNode_SetRect(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setRect"); signal != nil {
		signal.(func(*core.QRectF))(core.NewQRectFFromPointer(rect))
	}

}

func (ptr *QSGRectangleNode) ConnectSetRect(f func(rect *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRect", func(rect *core.QRectF) {
				signal.(func(*core.QRectF))(rect)
				f(rect)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRect", f)
		}
	}
}

func (ptr *QSGRectangleNode) DisconnectSetRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRect")
	}
}

func (ptr *QSGRectangleNode) SetRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGRectangleNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGRectangleNode) SetRect2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QSGRectangleNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

//export callbackQSGRectangleNode_DestroyQSGRectangleNode
func callbackQSGRectangleNode_DestroyQSGRectangleNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGRectangleNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGRectangleNodeFromPointer(ptr).DestroyQSGRectangleNodeDefault()
	}
}

func (ptr *QSGRectangleNode) ConnectDestroyQSGRectangleNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGRectangleNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGRectangleNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGRectangleNode", f)
		}
	}
}

func (ptr *QSGRectangleNode) DisconnectDestroyQSGRectangleNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGRectangleNode")
	}
}

func (ptr *QSGRectangleNode) DestroyQSGRectangleNode() {
	if ptr.Pointer() != nil {
		C.QSGRectangleNode_DestroyQSGRectangleNode(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGRectangleNode) DestroyQSGRectangleNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGRectangleNode_DestroyQSGRectangleNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSGRectangleNode_Color
func callbackQSGRectangleNode_Color(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "color"); signal != nil {
		return gui.PointerFromQColor(signal.(func() *gui.QColor)())
	}

	return gui.PointerFromQColor(gui.NewQColor())
}

func (ptr *QSGRectangleNode) ConnectColor(f func() *gui.QColor) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "color"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "color", func() *gui.QColor {
				signal.(func() *gui.QColor)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "color", f)
		}
	}
}

func (ptr *QSGRectangleNode) DisconnectColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "color")
	}
}

func (ptr *QSGRectangleNode) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QSGRectangleNode_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQSGRectangleNode_Rect
func callbackQSGRectangleNode_Rect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "rect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(core.NewQRectF())
}

func (ptr *QSGRectangleNode) ConnectRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rect", func() *core.QRectF {
				signal.(func() *core.QRectF)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rect", f)
		}
	}
}

func (ptr *QSGRectangleNode) DisconnectRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rect")
	}
}

func (ptr *QSGRectangleNode) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGRectangleNode_Rect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

type QSGRenderNode struct {
	QSGNode
}

type QSGRenderNode_ITF interface {
	QSGNode_ITF
	QSGRenderNode_PTR() *QSGRenderNode
}

func (ptr *QSGRenderNode) QSGRenderNode_PTR() *QSGRenderNode {
	return ptr
}

func (ptr *QSGRenderNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGRenderNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGRenderNode(ptr QSGRenderNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGRenderNode_PTR().Pointer()
	}
	return nil
}

func NewQSGRenderNodeFromPointer(ptr unsafe.Pointer) (n *QSGRenderNode) {
	n = new(QSGRenderNode)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGRenderNode__RenderingFlag
//QSGRenderNode::RenderingFlag
type QSGRenderNode__RenderingFlag int64

const (
	QSGRenderNode__BoundedRectRendering QSGRenderNode__RenderingFlag = QSGRenderNode__RenderingFlag(0x01)
	QSGRenderNode__DepthAwareRendering  QSGRenderNode__RenderingFlag = QSGRenderNode__RenderingFlag(0x02)
	QSGRenderNode__OpaqueRendering      QSGRenderNode__RenderingFlag = QSGRenderNode__RenderingFlag(0x04)
)

//go:generate stringer -type=QSGRenderNode__StateFlag
//QSGRenderNode::StateFlag
type QSGRenderNode__StateFlag int64

const (
	QSGRenderNode__DepthState        QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x01)
	QSGRenderNode__StencilState      QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x02)
	QSGRenderNode__ScissorState      QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x04)
	QSGRenderNode__ColorState        QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x08)
	QSGRenderNode__BlendState        QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x10)
	QSGRenderNode__CullState         QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x20)
	QSGRenderNode__ViewportState     QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x40)
	QSGRenderNode__RenderTargetState QSGRenderNode__StateFlag = QSGRenderNode__StateFlag(0x80)
)

//export callbackQSGRenderNode_ReleaseResources
func callbackQSGRenderNode_ReleaseResources(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQSGRenderNodeFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QSGRenderNode) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "releaseResources"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "releaseResources", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "releaseResources", f)
		}
	}
}

func (ptr *QSGRenderNode) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "releaseResources")
	}
}

func (ptr *QSGRenderNode) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QSGRenderNode_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QSGRenderNode) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.QSGRenderNode_ReleaseResourcesDefault(ptr.Pointer())
	}
}

//export callbackQSGRenderNode_DestroyQSGRenderNode
func callbackQSGRenderNode_DestroyQSGRenderNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGRenderNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGRenderNodeFromPointer(ptr).DestroyQSGRenderNodeDefault()
	}
}

func (ptr *QSGRenderNode) ConnectDestroyQSGRenderNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGRenderNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGRenderNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGRenderNode", f)
		}
	}
}

func (ptr *QSGRenderNode) DisconnectDestroyQSGRenderNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGRenderNode")
	}
}

func (ptr *QSGRenderNode) DestroyQSGRenderNode() {
	if ptr.Pointer() != nil {
		C.QSGRenderNode_DestroyQSGRenderNode(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGRenderNode) DestroyQSGRenderNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGRenderNode_DestroyQSGRenderNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSGRenderNode_Rect
func callbackQSGRenderNode_Rect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "rect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQSGRenderNodeFromPointer(ptr).RectDefault())
}

func (ptr *QSGRenderNode) ConnectRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rect", func() *core.QRectF {
				signal.(func() *core.QRectF)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rect", f)
		}
	}
}

func (ptr *QSGRenderNode) DisconnectRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rect")
	}
}

func (ptr *QSGRenderNode) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGRenderNode_Rect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGRenderNode) RectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGRenderNode_RectDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQSGRenderNode_Flags
func callbackQSGRenderNode_Flags(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func() QSGRenderNode__RenderingFlag)())
	}

	return C.longlong(NewQSGRenderNodeFromPointer(ptr).FlagsDefault())
}

func (ptr *QSGRenderNode) ConnectFlags(f func() QSGRenderNode__RenderingFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "flags"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "flags", func() QSGRenderNode__RenderingFlag {
				signal.(func() QSGRenderNode__RenderingFlag)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "flags", f)
		}
	}
}

func (ptr *QSGRenderNode) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "flags")
	}
}

func (ptr *QSGRenderNode) Flags() QSGRenderNode__RenderingFlag {
	if ptr.Pointer() != nil {
		return QSGRenderNode__RenderingFlag(C.QSGRenderNode_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGRenderNode) FlagsDefault() QSGRenderNode__RenderingFlag {
	if ptr.Pointer() != nil {
		return QSGRenderNode__RenderingFlag(C.QSGRenderNode_FlagsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGRenderNode_ChangedStates
func callbackQSGRenderNode_ChangedStates(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "changedStates"); signal != nil {
		return C.longlong(signal.(func() QSGRenderNode__StateFlag)())
	}

	return C.longlong(NewQSGRenderNodeFromPointer(ptr).ChangedStatesDefault())
}

func (ptr *QSGRenderNode) ConnectChangedStates(f func() QSGRenderNode__StateFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "changedStates"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "changedStates", func() QSGRenderNode__StateFlag {
				signal.(func() QSGRenderNode__StateFlag)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changedStates", f)
		}
	}
}

func (ptr *QSGRenderNode) DisconnectChangedStates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "changedStates")
	}
}

func (ptr *QSGRenderNode) ChangedStates() QSGRenderNode__StateFlag {
	if ptr.Pointer() != nil {
		return QSGRenderNode__StateFlag(C.QSGRenderNode_ChangedStates(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGRenderNode) ChangedStatesDefault() QSGRenderNode__StateFlag {
	if ptr.Pointer() != nil {
		return QSGRenderNode__StateFlag(C.QSGRenderNode_ChangedStatesDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGRenderNode) Matrix() *gui.QMatrix4x4 {
	if ptr.Pointer() != nil {
		return gui.NewQMatrix4x4FromPointer(C.QSGRenderNode_Matrix(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGRenderNode) ClipList() *QSGClipNode {
	if ptr.Pointer() != nil {
		return NewQSGClipNodeFromPointer(C.QSGRenderNode_ClipList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGRenderNode) InheritedOpacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QSGRenderNode_InheritedOpacity(ptr.Pointer()))
	}
	return 0
}

type QSGRendererInterface struct {
	ptr unsafe.Pointer
}

type QSGRendererInterface_ITF interface {
	QSGRendererInterface_PTR() *QSGRendererInterface
}

func (ptr *QSGRendererInterface) QSGRendererInterface_PTR() *QSGRendererInterface {
	return ptr
}

func (ptr *QSGRendererInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSGRendererInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSGRendererInterface(ptr QSGRendererInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGRendererInterface_PTR().Pointer()
	}
	return nil
}

func NewQSGRendererInterfaceFromPointer(ptr unsafe.Pointer) (n *QSGRendererInterface) {
	n = new(QSGRendererInterface)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGRendererInterface__GraphicsApi
//QSGRendererInterface::GraphicsApi
type QSGRendererInterface__GraphicsApi int64

const (
	QSGRendererInterface__Unknown    QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(0)
	QSGRendererInterface__Software   QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(1)
	QSGRendererInterface__OpenGL     QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(2)
	QSGRendererInterface__Direct3D12 QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(3)
	QSGRendererInterface__OpenVG     QSGRendererInterface__GraphicsApi = QSGRendererInterface__GraphicsApi(4)
)

//go:generate stringer -type=QSGRendererInterface__Resource
//QSGRendererInterface::Resource
type QSGRendererInterface__Resource int64

const (
	QSGRendererInterface__DeviceResource       QSGRendererInterface__Resource = QSGRendererInterface__Resource(0)
	QSGRendererInterface__CommandQueueResource QSGRendererInterface__Resource = QSGRendererInterface__Resource(1)
	QSGRendererInterface__CommandListResource  QSGRendererInterface__Resource = QSGRendererInterface__Resource(2)
	QSGRendererInterface__PainterResource      QSGRendererInterface__Resource = QSGRendererInterface__Resource(3)
)

//go:generate stringer -type=QSGRendererInterface__ShaderCompilationType
//QSGRendererInterface::ShaderCompilationType
type QSGRendererInterface__ShaderCompilationType int64

const (
	QSGRendererInterface__RuntimeCompilation QSGRendererInterface__ShaderCompilationType = QSGRendererInterface__ShaderCompilationType(0x01)
	QSGRendererInterface__OfflineCompilation QSGRendererInterface__ShaderCompilationType = QSGRendererInterface__ShaderCompilationType(0x02)
)

//go:generate stringer -type=QSGRendererInterface__ShaderSourceType
//QSGRendererInterface::ShaderSourceType
type QSGRendererInterface__ShaderSourceType int64

const (
	QSGRendererInterface__ShaderSourceString QSGRendererInterface__ShaderSourceType = QSGRendererInterface__ShaderSourceType(0x01)
	QSGRendererInterface__ShaderSourceFile   QSGRendererInterface__ShaderSourceType = QSGRendererInterface__ShaderSourceType(0x02)
	QSGRendererInterface__ShaderByteCode     QSGRendererInterface__ShaderSourceType = QSGRendererInterface__ShaderSourceType(0x04)
)

//go:generate stringer -type=QSGRendererInterface__ShaderType
//QSGRendererInterface::ShaderType
type QSGRendererInterface__ShaderType int64

const (
	QSGRendererInterface__UnknownShadingLanguage QSGRendererInterface__ShaderType = QSGRendererInterface__ShaderType(0)
	QSGRendererInterface__GLSL                   QSGRendererInterface__ShaderType = QSGRendererInterface__ShaderType(1)
	QSGRendererInterface__HLSL                   QSGRendererInterface__ShaderType = QSGRendererInterface__ShaderType(2)
)

//export callbackQSGRendererInterface_DestroyQSGRendererInterface
func callbackQSGRendererInterface_DestroyQSGRendererInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGRendererInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQSGRendererInterfaceFromPointer(ptr).DestroyQSGRendererInterfaceDefault()
	}
}

func (ptr *QSGRendererInterface) ConnectDestroyQSGRendererInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGRendererInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGRendererInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGRendererInterface", f)
		}
	}
}

func (ptr *QSGRendererInterface) DisconnectDestroyQSGRendererInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGRendererInterface")
	}
}

func (ptr *QSGRendererInterface) DestroyQSGRendererInterface() {
	if ptr.Pointer() != nil {
		C.QSGRendererInterface_DestroyQSGRendererInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGRendererInterface) DestroyQSGRendererInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QSGRendererInterface_DestroyQSGRendererInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGRendererInterface_GraphicsApi
func callbackQSGRendererInterface_GraphicsApi(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "graphicsApi"); signal != nil {
		return C.longlong(signal.(func() QSGRendererInterface__GraphicsApi)())
	}

	return C.longlong(0)
}

func (ptr *QSGRendererInterface) ConnectGraphicsApi(f func() QSGRendererInterface__GraphicsApi) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "graphicsApi"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "graphicsApi", func() QSGRendererInterface__GraphicsApi {
				signal.(func() QSGRendererInterface__GraphicsApi)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "graphicsApi", f)
		}
	}
}

func (ptr *QSGRendererInterface) DisconnectGraphicsApi() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "graphicsApi")
	}
}

func (ptr *QSGRendererInterface) GraphicsApi() QSGRendererInterface__GraphicsApi {
	if ptr.Pointer() != nil {
		return QSGRendererInterface__GraphicsApi(C.QSGRendererInterface_GraphicsApi(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGRendererInterface_ShaderCompilationType
func callbackQSGRendererInterface_ShaderCompilationType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "shaderCompilationType"); signal != nil {
		return C.longlong(signal.(func() QSGRendererInterface__ShaderCompilationType)())
	}

	return C.longlong(0)
}

func (ptr *QSGRendererInterface) ConnectShaderCompilationType(f func() QSGRendererInterface__ShaderCompilationType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "shaderCompilationType"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "shaderCompilationType", func() QSGRendererInterface__ShaderCompilationType {
				signal.(func() QSGRendererInterface__ShaderCompilationType)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shaderCompilationType", f)
		}
	}
}

func (ptr *QSGRendererInterface) DisconnectShaderCompilationType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "shaderCompilationType")
	}
}

func (ptr *QSGRendererInterface) ShaderCompilationType() QSGRendererInterface__ShaderCompilationType {
	if ptr.Pointer() != nil {
		return QSGRendererInterface__ShaderCompilationType(C.QSGRendererInterface_ShaderCompilationType(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGRendererInterface_ShaderSourceType
func callbackQSGRendererInterface_ShaderSourceType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "shaderSourceType"); signal != nil {
		return C.longlong(signal.(func() QSGRendererInterface__ShaderSourceType)())
	}

	return C.longlong(0)
}

func (ptr *QSGRendererInterface) ConnectShaderSourceType(f func() QSGRendererInterface__ShaderSourceType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "shaderSourceType"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "shaderSourceType", func() QSGRendererInterface__ShaderSourceType {
				signal.(func() QSGRendererInterface__ShaderSourceType)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shaderSourceType", f)
		}
	}
}

func (ptr *QSGRendererInterface) DisconnectShaderSourceType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "shaderSourceType")
	}
}

func (ptr *QSGRendererInterface) ShaderSourceType() QSGRendererInterface__ShaderSourceType {
	if ptr.Pointer() != nil {
		return QSGRendererInterface__ShaderSourceType(C.QSGRendererInterface_ShaderSourceType(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGRendererInterface_ShaderType
func callbackQSGRendererInterface_ShaderType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "shaderType"); signal != nil {
		return C.longlong(signal.(func() QSGRendererInterface__ShaderType)())
	}

	return C.longlong(0)
}

func (ptr *QSGRendererInterface) ConnectShaderType(f func() QSGRendererInterface__ShaderType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "shaderType"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "shaderType", func() QSGRendererInterface__ShaderType {
				signal.(func() QSGRendererInterface__ShaderType)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shaderType", f)
		}
	}
}

func (ptr *QSGRendererInterface) DisconnectShaderType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "shaderType")
	}
}

func (ptr *QSGRendererInterface) ShaderType() QSGRendererInterface__ShaderType {
	if ptr.Pointer() != nil {
		return QSGRendererInterface__ShaderType(C.QSGRendererInterface_ShaderType(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGRendererInterface_GetResource
func callbackQSGRendererInterface_GetResource(ptr unsafe.Pointer, window unsafe.Pointer, resource C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "getResource"); signal != nil {
		return signal.(func(*QQuickWindow, QSGRendererInterface__Resource) unsafe.Pointer)(NewQQuickWindowFromPointer(window), QSGRendererInterface__Resource(resource))
	}

	return NewQSGRendererInterfaceFromPointer(ptr).GetResourceDefault(NewQQuickWindowFromPointer(window), QSGRendererInterface__Resource(resource))
}

func (ptr *QSGRendererInterface) ConnectGetResource(f func(window *QQuickWindow, resource QSGRendererInterface__Resource) unsafe.Pointer) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getResource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getResource", func(window *QQuickWindow, resource QSGRendererInterface__Resource) unsafe.Pointer {
				signal.(func(*QQuickWindow, QSGRendererInterface__Resource) unsafe.Pointer)(window, resource)
				return f(window, resource)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getResource", f)
		}
	}
}

func (ptr *QSGRendererInterface) DisconnectGetResource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getResource")
	}
}

func (ptr *QSGRendererInterface) GetResource(window QQuickWindow_ITF, resource QSGRendererInterface__Resource) unsafe.Pointer {
	if ptr.Pointer() != nil {
		return C.QSGRendererInterface_GetResource(ptr.Pointer(), PointerFromQQuickWindow(window), C.longlong(resource))
	}
	return nil
}

func (ptr *QSGRendererInterface) GetResourceDefault(window QQuickWindow_ITF, resource QSGRendererInterface__Resource) unsafe.Pointer {
	if ptr.Pointer() != nil {
		return C.QSGRendererInterface_GetResourceDefault(ptr.Pointer(), PointerFromQQuickWindow(window), C.longlong(resource))
	}
	return nil
}

//export callbackQSGRendererInterface_GetResource2
func callbackQSGRendererInterface_GetResource2(ptr unsafe.Pointer, window unsafe.Pointer, resource C.struct_QtQuick_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "getResource2"); signal != nil {
		return signal.(func(*QQuickWindow, string) unsafe.Pointer)(NewQQuickWindowFromPointer(window), cGoUnpackString(resource))
	}

	return NewQSGRendererInterfaceFromPointer(ptr).GetResource2Default(NewQQuickWindowFromPointer(window), cGoUnpackString(resource))
}

func (ptr *QSGRendererInterface) ConnectGetResource2(f func(window *QQuickWindow, resource string) unsafe.Pointer) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getResource2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getResource2", func(window *QQuickWindow, resource string) unsafe.Pointer {
				signal.(func(*QQuickWindow, string) unsafe.Pointer)(window, resource)
				return f(window, resource)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getResource2", f)
		}
	}
}

func (ptr *QSGRendererInterface) DisconnectGetResource2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getResource2")
	}
}

func (ptr *QSGRendererInterface) GetResource2(window QQuickWindow_ITF, resource string) unsafe.Pointer {
	if ptr.Pointer() != nil {
		var resourceC *C.char
		if resource != "" {
			resourceC = C.CString(resource)
			defer C.free(unsafe.Pointer(resourceC))
		}
		return C.QSGRendererInterface_GetResource2(ptr.Pointer(), PointerFromQQuickWindow(window), resourceC)
	}
	return nil
}

func (ptr *QSGRendererInterface) GetResource2Default(window QQuickWindow_ITF, resource string) unsafe.Pointer {
	if ptr.Pointer() != nil {
		var resourceC *C.char
		if resource != "" {
			resourceC = C.CString(resource)
			defer C.free(unsafe.Pointer(resourceC))
		}
		return C.QSGRendererInterface_GetResource2Default(ptr.Pointer(), PointerFromQQuickWindow(window), resourceC)
	}
	return nil
}

type QSGSimpleMaterial struct {
	QSGMaterial
}

type QSGSimpleMaterial_ITF interface {
	QSGMaterial_ITF
	QSGSimpleMaterial_PTR() *QSGSimpleMaterial
}

func (ptr *QSGSimpleMaterial) QSGSimpleMaterial_PTR() *QSGSimpleMaterial {
	return ptr
}

func (ptr *QSGSimpleMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleMaterial(ptr QSGSimpleMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialFromPointer(ptr unsafe.Pointer) (n *QSGSimpleMaterial) {
	n = new(QSGSimpleMaterial)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGSimpleMaterial) DestroyQSGSimpleMaterial() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QSGSimpleMaterialShader struct {
	QSGMaterialShader
}

type QSGSimpleMaterialShader_ITF interface {
	QSGMaterialShader_ITF
	QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader
}

func (ptr *QSGSimpleMaterialShader) QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader {
	return ptr
}

func (ptr *QSGSimpleMaterialShader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleMaterialShader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterialShader_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleMaterialShader(ptr QSGSimpleMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialShaderFromPointer(ptr unsafe.Pointer) (n *QSGSimpleMaterialShader) {
	n = new(QSGSimpleMaterialShader)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGSimpleMaterialShader) DestroyQSGSimpleMaterialShader() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QSGSimpleRectNode struct {
	QSGGeometryNode
}

type QSGSimpleRectNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleRectNode_PTR() *QSGSimpleRectNode
}

func (ptr *QSGSimpleRectNode) QSGSimpleRectNode_PTR() *QSGSimpleRectNode {
	return ptr
}

func (ptr *QSGSimpleRectNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleRectNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleRectNode(ptr QSGSimpleRectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleRectNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleRectNodeFromPointer(ptr unsafe.Pointer) (n *QSGSimpleRectNode) {
	n = new(QSGSimpleRectNode)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGSimpleRectNode) DestroyQSGSimpleRectNode() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQSGSimpleRectNode2() *QSGSimpleRectNode {
	return NewQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode2())
}

func NewQSGSimpleRectNode(rect core.QRectF_ITF, color gui.QColor_ITF) *QSGSimpleRectNode {
	return NewQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode(core.PointerFromQRectF(rect), gui.PointerFromQColor(color)))
}

func (ptr *QSGSimpleRectNode) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QSGSimpleRectNode) SetRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGSimpleRectNode) SetRect2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleRectNode) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QSGSimpleRectNode_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QSGSimpleRectNode) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGSimpleRectNode_Rect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

type QSGSimpleTextureNode struct {
	QSGGeometryNode
}

type QSGSimpleTextureNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode
}

func (ptr *QSGSimpleTextureNode) QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode {
	return ptr
}

func (ptr *QSGSimpleTextureNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGGeometryNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGSimpleTextureNode(ptr QSGSimpleTextureNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleTextureNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleTextureNodeFromPointer(ptr unsafe.Pointer) (n *QSGSimpleTextureNode) {
	n = new(QSGSimpleTextureNode)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGSimpleTextureNode__TextureCoordinatesTransformFlag
//QSGSimpleTextureNode::TextureCoordinatesTransformFlag
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag int64

const (
	QSGSimpleTextureNode__NoTransform        QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x00)
	QSGSimpleTextureNode__MirrorHorizontally QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x01)
	QSGSimpleTextureNode__MirrorVertically   QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x02)
)

func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {
	return NewQSGSimpleTextureNodeFromPointer(C.QSGSimpleTextureNode_NewQSGSimpleTextureNode())
}

func (ptr *QSGSimpleTextureNode) SetFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetFiltering(ptr.Pointer(), C.longlong(filtering))
	}
}

func (ptr *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetOwnsTexture(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(owns))))
	}
}

func (ptr *QSGSimpleTextureNode) SetRect(r core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(r))
	}
}

func (ptr *QSGSimpleTextureNode) SetRect2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleTextureNode) SetSourceRect(r core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetSourceRect(ptr.Pointer(), core.PointerFromQRectF(r))
	}
}

func (ptr *QSGSimpleTextureNode) SetSourceRect2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetSourceRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleTextureNode) SetTexture(texture QSGTexture_ITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTexture(ptr.Pointer(), PointerFromQSGTexture(texture))
	}
}

func (ptr *QSGSimpleTextureNode) SetTextureCoordinatesTransform(mode QSGSimpleTextureNode__TextureCoordinatesTransformFlag) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTextureCoordinatesTransform(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQSGSimpleTextureNode_DestroyQSGSimpleTextureNode
func callbackQSGSimpleTextureNode_DestroyQSGSimpleTextureNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGSimpleTextureNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleTextureNodeFromPointer(ptr).DestroyQSGSimpleTextureNodeDefault()
	}
}

func (ptr *QSGSimpleTextureNode) ConnectDestroyQSGSimpleTextureNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGSimpleTextureNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGSimpleTextureNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGSimpleTextureNode", f)
		}
	}
}

func (ptr *QSGSimpleTextureNode) DisconnectDestroyQSGSimpleTextureNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGSimpleTextureNode")
	}
}

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNode() {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_DestroyQSGSimpleTextureNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGSimpleTextureNode) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGSimpleTextureNode_Rect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) SourceRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGSimpleTextureNode_SourceRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) TextureCoordinatesTransform() QSGSimpleTextureNode__TextureCoordinatesTransformFlag {
	if ptr.Pointer() != nil {
		return QSGSimpleTextureNode__TextureCoordinatesTransformFlag(C.QSGSimpleTextureNode_TextureCoordinatesTransform(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGSimpleTextureNode_Texture(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGSimpleTextureNode_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) OwnsTexture() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGSimpleTextureNode_OwnsTexture(ptr.Pointer())) != 0
	}
	return false
}

type QSGTexture struct {
	core.QObject
}

type QSGTexture_ITF interface {
	core.QObject_ITF
	QSGTexture_PTR() *QSGTexture
}

func (ptr *QSGTexture) QSGTexture_PTR() *QSGTexture {
	return ptr
}

func (ptr *QSGTexture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTexture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGTexture(ptr QSGTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureFromPointer(ptr unsafe.Pointer) (n *QSGTexture) {
	n = new(QSGTexture)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSGTexture__AnisotropyLevel
//QSGTexture::AnisotropyLevel
type QSGTexture__AnisotropyLevel int64

const (
	QSGTexture__AnisotropyNone QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(0)
	QSGTexture__Anisotropy2x   QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(1)
	QSGTexture__Anisotropy4x   QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(2)
	QSGTexture__Anisotropy8x   QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(3)
	QSGTexture__Anisotropy16x  QSGTexture__AnisotropyLevel = QSGTexture__AnisotropyLevel(4)
)

//go:generate stringer -type=QSGTexture__Filtering
//QSGTexture::Filtering
type QSGTexture__Filtering int64

const (
	QSGTexture__None    QSGTexture__Filtering = QSGTexture__Filtering(0)
	QSGTexture__Nearest QSGTexture__Filtering = QSGTexture__Filtering(1)
	QSGTexture__Linear  QSGTexture__Filtering = QSGTexture__Filtering(2)
)

//go:generate stringer -type=QSGTexture__WrapMode
//QSGTexture::WrapMode
type QSGTexture__WrapMode int64

const (
	QSGTexture__Repeat         QSGTexture__WrapMode = QSGTexture__WrapMode(0)
	QSGTexture__ClampToEdge    QSGTexture__WrapMode = QSGTexture__WrapMode(1)
	QSGTexture__MirroredRepeat QSGTexture__WrapMode = QSGTexture__WrapMode(2)
)

func NewQSGTexture() *QSGTexture {
	tmpValue := NewQSGTextureFromPointer(C.QSGTexture_NewQSGTexture())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QSGTexture_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTexture_QSGTexture_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSGTexture) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTexture_QSGTexture_Tr(sC, cC, C.int(int32(n))))
}

func QSGTexture_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTexture_QSGTexture_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSGTexture) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTexture_QSGTexture_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQSGTexture_Bind
func callbackQSGTexture_Bind(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "bind"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGTexture) ConnectBind(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "bind"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "bind", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bind", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectBind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "bind")
	}
}

func (ptr *QSGTexture) Bind() {
	if ptr.Pointer() != nil {
		C.QSGTexture_Bind(ptr.Pointer())
	}
}

func (ptr *QSGTexture) SetAnisotropyLevel(level QSGTexture__AnisotropyLevel) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetAnisotropyLevel(ptr.Pointer(), C.longlong(level))
	}
}

func (ptr *QSGTexture) SetFiltering(filter QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetFiltering(ptr.Pointer(), C.longlong(filter))
	}
}

func (ptr *QSGTexture) SetHorizontalWrapMode(hwrap QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetHorizontalWrapMode(ptr.Pointer(), C.longlong(hwrap))
	}
}

func (ptr *QSGTexture) SetMipmapFiltering(filter QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetMipmapFiltering(ptr.Pointer(), C.longlong(filter))
	}
}

func (ptr *QSGTexture) SetVerticalWrapMode(vwrap QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetVerticalWrapMode(ptr.Pointer(), C.longlong(vwrap))
	}
}

func (ptr *QSGTexture) UpdateBindOptions(force bool) {
	if ptr.Pointer() != nil {
		C.QSGTexture_UpdateBindOptions(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(force))))
	}
}

//export callbackQSGTexture_DestroyQSGTexture
func callbackQSGTexture_DestroyQSGTexture(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGTexture"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTextureFromPointer(ptr).DestroyQSGTextureDefault()
	}
}

func (ptr *QSGTexture) ConnectDestroyQSGTexture(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGTexture", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGTexture", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectDestroyQSGTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGTexture")
	}
}

func (ptr *QSGTexture) DestroyQSGTexture() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DestroyQSGTexture(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGTexture) DestroyQSGTextureDefault() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DestroyQSGTextureDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSGTexture) ConvertToNormalizedSourceRect(rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGTexture_ConvertToNormalizedSourceRect(ptr.Pointer(), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQSGTexture_NormalizedTextureSubRect
func callbackQSGTexture_NormalizedTextureSubRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "normalizedTextureSubRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQSGTextureFromPointer(ptr).NormalizedTextureSubRectDefault())
}

func (ptr *QSGTexture) ConnectNormalizedTextureSubRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "normalizedTextureSubRect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "normalizedTextureSubRect", func() *core.QRectF {
				signal.(func() *core.QRectF)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "normalizedTextureSubRect", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectNormalizedTextureSubRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "normalizedTextureSubRect")
	}
}

func (ptr *QSGTexture) NormalizedTextureSubRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGTexture_NormalizedTextureSubRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) NormalizedTextureSubRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSGTexture_NormalizedTextureSubRectDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQSGTexture_RemovedFromAtlas
func callbackQSGTexture_RemovedFromAtlas(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "removedFromAtlas"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGTextureFromPointer(ptr).RemovedFromAtlasDefault())
}

func (ptr *QSGTexture) ConnectRemovedFromAtlas(f func() *QSGTexture) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removedFromAtlas"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removedFromAtlas", func() *QSGTexture {
				signal.(func() *QSGTexture)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removedFromAtlas", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectRemovedFromAtlas() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removedFromAtlas")
	}
}

func (ptr *QSGTexture) RemovedFromAtlas() *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlas(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) RemovedFromAtlasDefault() *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlasDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) AnisotropyLevel() QSGTexture__AnisotropyLevel {
	if ptr.Pointer() != nil {
		return QSGTexture__AnisotropyLevel(C.QSGTexture_AnisotropyLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) HorizontalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) VerticalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGTexture_TextureSize
func callbackQSGTexture_TextureSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QSGTexture) ConnectTextureSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "textureSize"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureSize", func() *core.QSize {
				signal.(func() *core.QSize)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureSize", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectTextureSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "textureSize")
	}
}

func (ptr *QSGTexture) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSGTexture_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSGTexture_HasAlphaChannel
func callbackQSGTexture_HasAlphaChannel(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasAlphaChannel"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGTexture) ConnectHasAlphaChannel(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasAlphaChannel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasAlphaChannel", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasAlphaChannel", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectHasAlphaChannel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasAlphaChannel")
	}
}

func (ptr *QSGTexture) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTexture_HasAlphaChannel(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSGTexture_HasMipmaps
func callbackQSGTexture_HasMipmaps(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasMipmaps"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGTexture) ConnectHasMipmaps(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasMipmaps"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasMipmaps", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasMipmaps", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectHasMipmaps() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasMipmaps")
	}
}

func (ptr *QSGTexture) HasMipmaps() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTexture_HasMipmaps(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSGTexture_IsAtlasTexture
func callbackQSGTexture_IsAtlasTexture(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isAtlasTexture"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).IsAtlasTextureDefault())))
}

func (ptr *QSGTexture) ConnectIsAtlasTexture(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isAtlasTexture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isAtlasTexture", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isAtlasTexture", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectIsAtlasTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isAtlasTexture")
	}
}

func (ptr *QSGTexture) IsAtlasTexture() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTexture_IsAtlasTexture(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGTexture) IsAtlasTextureDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTexture_IsAtlasTextureDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSGTexture_MetaObject
func callbackQSGTexture_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGTextureFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGTexture) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTexture_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGTexture_TextureId
func callbackQSGTexture_TextureId(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "textureId"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QSGTexture) ConnectTextureId(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "textureId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureId", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureId", f)
		}
	}
}

func (ptr *QSGTexture) DisconnectTextureId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "textureId")
	}
}

func (ptr *QSGTexture) TextureId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGTexture_TextureId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGTexture) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSGTexture___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSGTexture) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSGTexture___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSGTexture) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTexture___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTexture) __findChildren_newList2() unsafe.Pointer {
	return C.QSGTexture___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSGTexture) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTexture___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTexture) __findChildren_newList3() unsafe.Pointer {
	return C.QSGTexture___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSGTexture) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTexture___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTexture) __findChildren_newList() unsafe.Pointer {
	return C.QSGTexture___findChildren_newList(ptr.Pointer())
}

func (ptr *QSGTexture) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTexture___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTexture) __children_newList() unsafe.Pointer {
	return C.QSGTexture___children_newList(ptr.Pointer())
}

//export callbackQSGTexture_Event
func callbackQSGTexture_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGTexture) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTexture_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSGTexture_EventFilter
func callbackQSGTexture_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGTexture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTexture_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSGTexture_ChildEvent
func callbackQSGTexture_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGTexture) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGTexture_ConnectNotify
func callbackQSGTexture_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTexture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTexture_CustomEvent
func callbackQSGTexture_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGTexture) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGTexture_DeleteLater
func callbackQSGTexture_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTextureFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGTexture) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSGTexture_Destroyed
func callbackQSGTexture_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSGTexture_DisconnectNotify
func callbackQSGTexture_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTexture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTexture_ObjectNameChanged
func callbackQSGTexture_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSGTexture_TimerEvent
func callbackQSGTexture_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGTexture) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSGTextureMaterial struct {
	QSGOpaqueTextureMaterial
}

type QSGTextureMaterial_ITF interface {
	QSGOpaqueTextureMaterial_ITF
	QSGTextureMaterial_PTR() *QSGTextureMaterial
}

func (ptr *QSGTextureMaterial) QSGTextureMaterial_PTR() *QSGTextureMaterial {
	return ptr
}

func (ptr *QSGTextureMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTextureMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGOpaqueTextureMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGTextureMaterial(ptr QSGTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureMaterialFromPointer(ptr unsafe.Pointer) (n *QSGTextureMaterial) {
	n = new(QSGTextureMaterial)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGTextureMaterial) DestroyQSGTextureMaterial() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProvider_ITF interface {
	core.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func (ptr *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider {
	return ptr
}

func (ptr *QSGTextureProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTextureProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSGTextureProvider(ptr QSGTextureProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureProvider_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureProviderFromPointer(ptr unsafe.Pointer) (n *QSGTextureProvider) {
	n = new(QSGTextureProvider)
	n.SetPointer(ptr)
	return
}
func QSGTextureProvider_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTextureProvider_QSGTextureProvider_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSGTextureProvider) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTextureProvider_QSGTextureProvider_Tr(sC, cC, C.int(int32(n))))
}

func QSGTextureProvider_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTextureProvider_QSGTextureProvider_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSGTextureProvider) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSGTextureProvider_QSGTextureProvider_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQSGTextureProvider_TextureChanged
func callbackQSGTextureProvider_TextureChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "textureChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGTextureProvider) ConnectTextureChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textureChanged") {
			C.QSGTextureProvider_ConnectTextureChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textureChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textureChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textureChanged", f)
		}
	}
}

func (ptr *QSGTextureProvider) DisconnectTextureChanged() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectTextureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "textureChanged")
	}
}

func (ptr *QSGTextureProvider) TextureChanged() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TextureChanged(ptr.Pointer())
	}
}

//export callbackQSGTextureProvider_Texture
func callbackQSGTextureProvider_Texture(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "texture"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGTexture())
}

func (ptr *QSGTextureProvider) ConnectTexture(f func() *QSGTexture) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "texture"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "texture", func() *QSGTexture {
				signal.(func() *QSGTexture)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "texture", f)
		}
	}
}

func (ptr *QSGTextureProvider) DisconnectTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "texture")
	}
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		tmpValue := NewQSGTextureFromPointer(C.QSGTextureProvider_Texture(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSGTextureProvider_MetaObject
func callbackQSGTextureProvider_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGTextureProviderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGTextureProvider) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTextureProvider_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureProvider) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSGTextureProvider___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSGTextureProvider) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSGTextureProvider) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSGTextureProvider___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSGTextureProvider) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTextureProvider___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTextureProvider) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTextureProvider) __findChildren_newList2() unsafe.Pointer {
	return C.QSGTextureProvider___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSGTextureProvider) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTextureProvider___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTextureProvider) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTextureProvider) __findChildren_newList3() unsafe.Pointer {
	return C.QSGTextureProvider___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSGTextureProvider) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTextureProvider___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTextureProvider) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTextureProvider) __findChildren_newList() unsafe.Pointer {
	return C.QSGTextureProvider___findChildren_newList(ptr.Pointer())
}

func (ptr *QSGTextureProvider) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSGTextureProvider___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTextureProvider) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSGTextureProvider) __children_newList() unsafe.Pointer {
	return C.QSGTextureProvider___children_newList(ptr.Pointer())
}

//export callbackQSGTextureProvider_Event
func callbackQSGTextureProvider_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureProviderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGTextureProvider) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTextureProvider_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSGTextureProvider_EventFilter
func callbackQSGTextureProvider_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureProviderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGTextureProvider) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSGTextureProvider_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSGTextureProvider_ChildEvent
func callbackQSGTextureProvider_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGTextureProvider_ConnectNotify
func callbackQSGTextureProvider_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTextureProvider) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTextureProvider_CustomEvent
func callbackQSGTextureProvider_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGTextureProvider_DeleteLater
func callbackQSGTextureProvider_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTextureProviderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGTextureProvider) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSGTextureProvider_Destroyed
func callbackQSGTextureProvider_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSGTextureProvider_DisconnectNotify
func callbackQSGTextureProvider_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureProviderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTextureProvider) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTextureProvider_ObjectNameChanged
func callbackQSGTextureProvider_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSGTextureProvider_TimerEvent
func callbackQSGTextureProvider_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSGTransformNode struct {
	QSGNode
}

type QSGTransformNode_ITF interface {
	QSGNode_ITF
	QSGTransformNode_PTR() *QSGTransformNode
}

func (ptr *QSGTransformNode) QSGTransformNode_PTR() *QSGTransformNode {
	return ptr
}

func (ptr *QSGTransformNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGTransformNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGNode_PTR().SetPointer(p)
	}
}

func PointerFromQSGTransformNode(ptr QSGTransformNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTransformNode_PTR().Pointer()
	}
	return nil
}

func NewQSGTransformNodeFromPointer(ptr unsafe.Pointer) (n *QSGTransformNode) {
	n = new(QSGTransformNode)
	n.SetPointer(ptr)
	return
}
func NewQSGTransformNode() *QSGTransformNode {
	return NewQSGTransformNodeFromPointer(C.QSGTransformNode_NewQSGTransformNode())
}

func (ptr *QSGTransformNode) SetMatrix(matrix gui.QMatrix4x4_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_SetMatrix(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

//export callbackQSGTransformNode_DestroyQSGTransformNode
func callbackQSGTransformNode_DestroyQSGTransformNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSGTransformNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTransformNodeFromPointer(ptr).DestroyQSGTransformNodeDefault()
	}
}

func (ptr *QSGTransformNode) ConnectDestroyQSGTransformNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSGTransformNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSGTransformNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSGTransformNode", f)
		}
	}
}

func (ptr *QSGTransformNode) DisconnectDestroyQSGTransformNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSGTransformNode")
	}
}

func (ptr *QSGTransformNode) DestroyQSGTransformNode() {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_DestroyQSGTransformNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGTransformNode) DestroyQSGTransformNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_DestroyQSGTransformNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGTransformNode) Matrix() *gui.QMatrix4x4 {
	if ptr.Pointer() != nil {
		return gui.NewQMatrix4x4FromPointer(C.QSGTransformNode_Matrix(ptr.Pointer()))
	}
	return nil
}

type QSGVertexColorMaterial struct {
	QSGMaterial
}

type QSGVertexColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial
}

func (ptr *QSGVertexColorMaterial) QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial {
	return ptr
}

func (ptr *QSGVertexColorMaterial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (ptr *QSGVertexColorMaterial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSGMaterial_PTR().SetPointer(p)
	}
}

func PointerFromQSGVertexColorMaterial(ptr QSGVertexColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGVertexColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGVertexColorMaterialFromPointer(ptr unsafe.Pointer) (n *QSGVertexColorMaterial) {
	n = new(QSGVertexColorMaterial)
	n.SetPointer(ptr)
	return
}

func (ptr *QSGVertexColorMaterial) DestroyQSGVertexColorMaterial() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQSGVertexColorMaterial() *QSGVertexColorMaterial {
	return NewQSGVertexColorMaterialFromPointer(C.QSGVertexColorMaterial_NewQSGVertexColorMaterial())
}

//export callbackQSGVertexColorMaterial_CreateShader
func callbackQSGVertexColorMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGVertexColorMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGVertexColorMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGVertexColorMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGVertexColorMaterial) CreateShaderDefault() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGVertexColorMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGVertexColorMaterial_Type
func callbackQSGVertexColorMaterial_Type(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGVertexColorMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGVertexColorMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGVertexColorMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGVertexColorMaterial) TypeDefault() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGVertexColorMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}
