// +build !minimal

package quick

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "quick.h"
import "C"
import (
	"fmt"
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
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QQuickAsyncImageProvider struct {
	QQuickImageProvider
}

type QQuickAsyncImageProvider_ITF interface {
	QQuickImageProvider_ITF
	QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider
}

func (p *QQuickAsyncImageProvider) QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider {
	return p
}

func (p *QQuickAsyncImageProvider) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func (p *QQuickAsyncImageProvider) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickImageProvider_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickAsyncImageProvider(ptr QQuickAsyncImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickAsyncImageProvider_PTR().Pointer()
	}
	return nil
}

func NewQQuickAsyncImageProviderFromPointer(ptr unsafe.Pointer) *QQuickAsyncImageProvider {
	var n = new(QQuickAsyncImageProvider)
	n.SetPointer(ptr)
	return n
}
func NewQQuickAsyncImageProvider() *QQuickAsyncImageProvider {
	return NewQQuickAsyncImageProviderFromPointer(C.QQuickAsyncImageProvider_NewQQuickAsyncImageProvider())
}

//export callbackQQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider
func callbackQQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickAsyncImageProvider::~QQuickAsyncImageProvider"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickAsyncImageProviderFromPointer(ptr).DestroyQQuickAsyncImageProviderDefault()
	}
}

func (ptr *QQuickAsyncImageProvider) ConnectDestroyQQuickAsyncImageProvider(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::~QQuickAsyncImageProvider", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectDestroyQQuickAsyncImageProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::~QQuickAsyncImageProvider")
	}
}

func (ptr *QQuickAsyncImageProvider) DestroyQQuickAsyncImageProvider() {
	if ptr.Pointer() != nil {
		C.QQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickAsyncImageProvider) DestroyQQuickAsyncImageProviderDefault() {
	if ptr.Pointer() != nil {
		C.QQuickAsyncImageProvider_DestroyQQuickAsyncImageProviderDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickAsyncImageProvider_RequestImageResponse
func callbackQQuickAsyncImageProvider_RequestImageResponse(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, requestedSize unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickAsyncImageProvider::requestImageResponse"); signal != nil {
		return PointerFromQQuickImageResponse(signal.(func(string, *core.QSize) *QQuickImageResponse)(cGoUnpackString(id), core.NewQSizeFromPointer(requestedSize)))
	}

	return PointerFromQQuickImageResponse(NewQQuickImageResponse())
}

func (ptr *QQuickAsyncImageProvider) ConnectRequestImageResponse(f func(id string, requestedSize *core.QSize) *QQuickImageResponse) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestImageResponse", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectRequestImageResponse(id string, requestedSize core.QSize_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestImageResponse")
	}
}

func (ptr *QQuickAsyncImageProvider) RequestImageResponse(id string, requestedSize core.QSize_ITF) *QQuickImageResponse {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = NewQQuickImageResponseFromPointer(C.QQuickAsyncImageProvider_RequestImageResponse(ptr.Pointer(), idC, core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickAsyncImageProvider_Flags
func callbackQQuickAsyncImageProvider_Flags(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickAsyncImageProvider::flags"); signal != nil {
		return C.longlong(signal.(func() qml.QQmlImageProviderBase__Flag)())
	}

	return C.longlong(NewQQuickAsyncImageProviderFromPointer(ptr).FlagsDefault())
}

func (ptr *QQuickAsyncImageProvider) ConnectFlags(f func() qml.QQmlImageProviderBase__Flag) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::flags", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::flags")
	}
}

func (ptr *QQuickAsyncImageProvider) Flags() qml.QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__Flag(C.QQuickAsyncImageProvider_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickAsyncImageProvider) FlagsDefault() qml.QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__Flag(C.QQuickAsyncImageProvider_FlagsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickAsyncImageProvider_ImageType
func callbackQQuickAsyncImageProvider_ImageType(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickAsyncImageProvider::imageType"); signal != nil {
		return C.longlong(signal.(func() qml.QQmlImageProviderBase__ImageType)())
	}

	return C.longlong(NewQQuickAsyncImageProviderFromPointer(ptr).ImageTypeDefault())
}

func (ptr *QQuickAsyncImageProvider) ConnectImageType(f func() qml.QQmlImageProviderBase__ImageType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::imageType", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectImageType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::imageType")
	}
}

func (ptr *QQuickAsyncImageProvider) ImageType() qml.QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__ImageType(C.QQuickAsyncImageProvider_ImageType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickAsyncImageProvider) ImageTypeDefault() qml.QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__ImageType(C.QQuickAsyncImageProvider_ImageTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickAsyncImageProvider_RequestImage
func callbackQQuickAsyncImageProvider_RequestImage(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickAsyncImageProvider::requestImage"); signal != nil {
		return gui.PointerFromQImage(signal.(func(string, *core.QSize, *core.QSize) *gui.QImage)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return gui.PointerFromQImage(NewQQuickAsyncImageProviderFromPointer(ptr).RequestImageDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickAsyncImageProvider) ConnectRequestImage(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QImage) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestImage", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectRequestImage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestImage")
	}
}

func (ptr *QQuickAsyncImageProvider) RequestImage(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQImageFromPointer(C.QQuickAsyncImageProvider_RequestImage(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickAsyncImageProvider) RequestImageDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQImageFromPointer(C.QQuickAsyncImageProvider_RequestImageDefault(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

//export callbackQQuickAsyncImageProvider_RequestPixmap
func callbackQQuickAsyncImageProvider_RequestPixmap(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickAsyncImageProvider::requestPixmap"); signal != nil {
		return gui.PointerFromQPixmap(signal.(func(string, *core.QSize, *core.QSize) *gui.QPixmap)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return gui.PointerFromQPixmap(NewQQuickAsyncImageProviderFromPointer(ptr).RequestPixmapDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickAsyncImageProvider) ConnectRequestPixmap(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QPixmap) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestPixmap", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectRequestPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestPixmap")
	}
}

func (ptr *QQuickAsyncImageProvider) RequestPixmap(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQPixmapFromPointer(C.QQuickAsyncImageProvider_RequestPixmap(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickAsyncImageProvider) RequestPixmapDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQPixmapFromPointer(C.QQuickAsyncImageProvider_RequestPixmapDefault(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQQuickAsyncImageProvider_RequestTexture
func callbackQQuickAsyncImageProvider_RequestTexture(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickAsyncImageProvider::requestTexture"); signal != nil {
		return PointerFromQQuickTextureFactory(signal.(func(string, *core.QSize, *core.QSize) *QQuickTextureFactory)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return PointerFromQQuickTextureFactory(NewQQuickAsyncImageProviderFromPointer(ptr).RequestTextureDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickAsyncImageProvider) ConnectRequestTexture(f func(id string, size *core.QSize, requestedSize *core.QSize) *QQuickTextureFactory) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestTexture", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectRequestTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickAsyncImageProvider::requestTexture")
	}
}

func (ptr *QQuickAsyncImageProvider) RequestTexture(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickAsyncImageProvider_RequestTexture(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickAsyncImageProvider) RequestTextureDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickAsyncImageProvider_RequestTextureDefault(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QQuickFramebufferObject struct {
	QQuickItem
}

type QQuickFramebufferObject_ITF interface {
	QQuickItem_ITF
	QQuickFramebufferObject_PTR() *QQuickFramebufferObject
}

func (p *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject {
	return p
}

func (p *QQuickFramebufferObject) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (p *QQuickFramebufferObject) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickItem_PTR().SetPointer(ptr)
	}
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
	return n
}

func (ptr *QQuickFramebufferObject) DestroyQQuickFramebufferObject() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QQuickFramebufferObject) MirrorVertically() bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_MirrorVertically(ptr.Pointer()) != 0
	}
	return false
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

func (ptr *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_TextureFollowsItemSize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) DisconnectCreateRenderer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::createRenderer")
	}
}

//export callbackQQuickFramebufferObject_IsTextureProvider
func callbackQQuickFramebufferObject_IsTextureProvider(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::isTextureProvider"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).IsTextureProviderDefault())))
}

func (ptr *QQuickFramebufferObject) ConnectIsTextureProvider(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::isTextureProvider", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectIsTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::isTextureProvider")
	}
}

func (ptr *QQuickFramebufferObject) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) IsTextureProviderDefault() bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProviderDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_MirrorVerticallyChanged
func callbackQQuickFramebufferObject_MirrorVerticallyChanged(ptr unsafe.Pointer, vbo C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::mirrorVerticallyChanged"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	}

}

func (ptr *QQuickFramebufferObject) ConnectMirrorVerticallyChanged(f func(vbo bool)) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectMirrorVerticallyChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mirrorVerticallyChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMirrorVerticallyChanged() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectMirrorVerticallyChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mirrorVerticallyChanged")
	}
}

func (ptr *QQuickFramebufferObject) MirrorVerticallyChanged(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MirrorVerticallyChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQQuickFramebufferObject_ReleaseResources
func callbackQQuickFramebufferObject_ReleaseResources(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::releaseResources", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::releaseResources")
	}
}

func (ptr *QQuickFramebufferObject) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ReleaseResourcesDefault(ptr.Pointer())
	}
}

//export callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged
func callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged(ptr unsafe.Pointer, vbo C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::textureFollowsItemSizeChanged"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	}

}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(vbo bool)) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::textureFollowsItemSizeChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureFollowsItemSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::textureFollowsItemSizeChanged")
	}
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSizeChanged(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TextureFollowsItemSizeChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQQuickFramebufferObject_TextureProvider
func callbackQQuickFramebufferObject_TextureProvider(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::textureProvider"); signal != nil {
		return PointerFromQSGTextureProvider(signal.(func() *QSGTextureProvider)())
	}

	return PointerFromQSGTextureProvider(NewQQuickFramebufferObjectFromPointer(ptr).TextureProviderDefault())
}

func (ptr *QQuickFramebufferObject) ConnectTextureProvider(f func() *QSGTextureProvider) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::textureProvider", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::textureProvider")
	}
}

func (ptr *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureProviderFromPointer(C.QQuickFramebufferObject_TextureProvider(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickFramebufferObject) TextureProviderDefault() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureProviderFromPointer(C.QQuickFramebufferObject_TextureProviderDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickFramebufferObject_ChildMouseEventFilter
func callbackQQuickFramebufferObject_ChildMouseEventFilter(ptr unsafe.Pointer, item unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::childMouseEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QQuickItem, *core.QEvent) bool)(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).ChildMouseEventFilterDefault(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickFramebufferObject) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::childMouseEventFilter", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectChildMouseEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::childMouseEventFilter")
	}
}

func (ptr *QQuickFramebufferObject) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_ChildMouseEventFilterDefault(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_ClassBegin
func callbackQQuickFramebufferObject_ClassBegin(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::classBegin"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectClassBegin(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::classBegin", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectClassBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::classBegin")
	}
}

func (ptr *QQuickFramebufferObject) ClassBegin() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) ClassBeginDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ClassBeginDefault(ptr.Pointer())
	}
}

//export callbackQQuickFramebufferObject_ComponentComplete
func callbackQQuickFramebufferObject_ComponentComplete(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::componentComplete"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectComponentComplete(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::componentComplete", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectComponentComplete() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::componentComplete")
	}
}

func (ptr *QQuickFramebufferObject) ComponentComplete() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ComponentComplete(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) ComponentCompleteDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ComponentCompleteDefault(ptr.Pointer())
	}
}

//export callbackQQuickFramebufferObject_Contains
func callbackQQuickFramebufferObject_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QQuickFramebufferObject) ConnectContains(f func(point *core.QPointF) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::contains", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectContains() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::contains")
	}
}

func (ptr *QQuickFramebufferObject) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_DragEnterEvent
func callbackQQuickFramebufferObject_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dragEnterEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dragEnterEvent")
	}
}

func (ptr *QQuickFramebufferObject) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQQuickFramebufferObject_DragLeaveEvent
func callbackQQuickFramebufferObject_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dragLeaveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dragLeaveEvent")
	}
}

func (ptr *QQuickFramebufferObject) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQQuickFramebufferObject_DragMoveEvent
func callbackQQuickFramebufferObject_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dragMoveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dragMoveEvent")
	}
}

func (ptr *QQuickFramebufferObject) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQQuickFramebufferObject_DropEvent
func callbackQQuickFramebufferObject_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dropEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::dropEvent")
	}
}

func (ptr *QQuickFramebufferObject) DropEvent(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQQuickFramebufferObject_Event
func callbackQQuickFramebufferObject_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QQuickFramebufferObject) ConnectEvent(f func(ev *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::event", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::event")
	}
}

func (ptr *QQuickFramebufferObject) Event(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_FocusInEvent
func callbackQQuickFramebufferObject_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::focusInEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::focusInEvent")
	}
}

func (ptr *QQuickFramebufferObject) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickFramebufferObject_FocusOutEvent
func callbackQQuickFramebufferObject_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::focusOutEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::focusOutEvent")
	}
}

func (ptr *QQuickFramebufferObject) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickFramebufferObject_GeometryChanged
func callbackQQuickFramebufferObject_GeometryChanged(ptr unsafe.Pointer, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *QQuickFramebufferObject) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::geometryChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectGeometryChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::geometryChanged")
	}
}

func (ptr *QQuickFramebufferObject) GeometryChanged(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_GeometryChanged(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickFramebufferObject) GeometryChangedDefault(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_GeometryChangedDefault(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

//export callbackQQuickFramebufferObject_HoverEnterEvent
func callbackQQuickFramebufferObject_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::hoverEnterEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectHoverEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::hoverEnterEvent")
	}
}

func (ptr *QQuickFramebufferObject) HoverEnterEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverEnterEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverEnterEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverEnterEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickFramebufferObject_HoverLeaveEvent
func callbackQQuickFramebufferObject_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::hoverLeaveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectHoverLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::hoverLeaveEvent")
	}
}

func (ptr *QQuickFramebufferObject) HoverLeaveEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverLeaveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverLeaveEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverLeaveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickFramebufferObject_HoverMoveEvent
func callbackQQuickFramebufferObject_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::hoverMoveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectHoverMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::hoverMoveEvent")
	}
}

func (ptr *QQuickFramebufferObject) HoverMoveEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverMoveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverMoveEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverMoveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickFramebufferObject_InputMethodEvent
func callbackQQuickFramebufferObject_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::inputMethodEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::inputMethodEvent")
	}
}

func (ptr *QQuickFramebufferObject) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQQuickFramebufferObject_InputMethodQuery
func callbackQQuickFramebufferObject_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickFramebufferObjectFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickFramebufferObject) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::inputMethodQuery", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::inputMethodQuery")
	}
}

func (ptr *QQuickFramebufferObject) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickFramebufferObject_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickFramebufferObject) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickFramebufferObject_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickFramebufferObject) DisconnectItemChange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::itemChange")
	}
}

//export callbackQQuickFramebufferObject_KeyPressEvent
func callbackQQuickFramebufferObject_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::keyPressEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::keyPressEvent")
	}
}

func (ptr *QQuickFramebufferObject) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQQuickFramebufferObject_KeyReleaseEvent
func callbackQQuickFramebufferObject_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::keyReleaseEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::keyReleaseEvent")
	}
}

func (ptr *QQuickFramebufferObject) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQQuickFramebufferObject_MouseDoubleClickEvent
func callbackQQuickFramebufferObject_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseDoubleClickEvent")
	}
}

func (ptr *QQuickFramebufferObject) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickFramebufferObject_MouseMoveEvent
func callbackQQuickFramebufferObject_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseMoveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseMoveEvent")
	}
}

func (ptr *QQuickFramebufferObject) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickFramebufferObject_MousePressEvent
func callbackQQuickFramebufferObject_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mousePressEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mousePressEvent")
	}
}

func (ptr *QQuickFramebufferObject) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickFramebufferObject_MouseReleaseEvent
func callbackQQuickFramebufferObject_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseReleaseEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseReleaseEvent")
	}
}

func (ptr *QQuickFramebufferObject) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickFramebufferObject_MouseUngrabEvent
func callbackQQuickFramebufferObject_MouseUngrabEvent(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectMouseUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseUngrabEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::mouseUngrabEvent")
	}
}

func (ptr *QQuickFramebufferObject) MouseUngrabEvent() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) MouseUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseUngrabEventDefault(ptr.Pointer())
	}
}

//export callbackQQuickFramebufferObject_TouchEvent
func callbackQQuickFramebufferObject_TouchEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::touchEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::touchEvent")
	}
}

func (ptr *QQuickFramebufferObject) TouchEvent(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) TouchEventDefault(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

//export callbackQQuickFramebufferObject_TouchUngrabEvent
func callbackQQuickFramebufferObject_TouchUngrabEvent(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectTouchUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::touchUngrabEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTouchUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::touchUngrabEvent")
	}
}

func (ptr *QQuickFramebufferObject) TouchUngrabEvent() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) TouchUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchUngrabEventDefault(ptr.Pointer())
	}
}

//export callbackQQuickFramebufferObject_Update
func callbackQQuickFramebufferObject_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::update", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::update")
	}
}

func (ptr *QQuickFramebufferObject) Update() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_Update(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_UpdateDefault(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) DisconnectUpdatePaintNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::updatePaintNode")
	}
}

//export callbackQQuickFramebufferObject_UpdatePolish
func callbackQQuickFramebufferObject_UpdatePolish(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectUpdatePolish(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::updatePolish", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectUpdatePolish() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::updatePolish")
	}
}

func (ptr *QQuickFramebufferObject) UpdatePolish() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_UpdatePolish(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) UpdatePolishDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_UpdatePolishDefault(ptr.Pointer())
	}
}

//export callbackQQuickFramebufferObject_WheelEvent
func callbackQQuickFramebufferObject_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::wheelEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::wheelEvent")
	}
}

func (ptr *QQuickFramebufferObject) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQQuickFramebufferObject_TimerEvent
func callbackQQuickFramebufferObject_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::timerEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::timerEvent")
	}
}

func (ptr *QQuickFramebufferObject) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickFramebufferObject_ChildEvent
func callbackQQuickFramebufferObject_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::childEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::childEvent")
	}
}

func (ptr *QQuickFramebufferObject) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickFramebufferObject_ConnectNotify
func callbackQQuickFramebufferObject_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickFramebufferObject) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::connectNotify", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::connectNotify")
	}
}

func (ptr *QQuickFramebufferObject) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickFramebufferObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickFramebufferObject_CustomEvent
func callbackQQuickFramebufferObject_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::customEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::customEvent")
	}
}

func (ptr *QQuickFramebufferObject) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickFramebufferObject_DeleteLater
func callbackQQuickFramebufferObject_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::deleteLater", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::deleteLater")
	}
}

func (ptr *QQuickFramebufferObject) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickFramebufferObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickFramebufferObject_DisconnectNotify
func callbackQQuickFramebufferObject_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickFramebufferObject) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::disconnectNotify", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::disconnectNotify")
	}
}

func (ptr *QQuickFramebufferObject) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickFramebufferObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickFramebufferObject_EventFilter
func callbackQQuickFramebufferObject_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickFramebufferObject) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::eventFilter", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::eventFilter")
	}
}

func (ptr *QQuickFramebufferObject) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_MetaObject
func callbackQQuickFramebufferObject_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickFramebufferObject::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickFramebufferObjectFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickFramebufferObject) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::metaObject", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickFramebufferObject::metaObject")
	}
}

func (ptr *QQuickFramebufferObject) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickFramebufferObject_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickFramebufferObject) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickFramebufferObject_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickImageProvider struct {
	qml.QQmlImageProviderBase
}

type QQuickImageProvider_ITF interface {
	qml.QQmlImageProviderBase_ITF
	QQuickImageProvider_PTR() *QQuickImageProvider
}

func (p *QQuickImageProvider) QQuickImageProvider_PTR() *QQuickImageProvider {
	return p
}

func (p *QQuickImageProvider) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func (p *QQuickImageProvider) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQmlImageProviderBase_PTR().SetPointer(ptr)
	}
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
func NewQQuickImageProvider(ty qml.QQmlImageProviderBase__ImageType, flags qml.QQmlImageProviderBase__Flag) *QQuickImageProvider {
	return NewQQuickImageProviderFromPointer(C.QQuickImageProvider_NewQQuickImageProvider(C.longlong(ty), C.longlong(flags)))
}

//export callbackQQuickImageProvider_Flags
func callbackQQuickImageProvider_Flags(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageProvider::flags"); signal != nil {
		return C.longlong(signal.(func() qml.QQmlImageProviderBase__Flag)())
	}

	return C.longlong(NewQQuickImageProviderFromPointer(ptr).FlagsDefault())
}

func (ptr *QQuickImageProvider) ConnectFlags(f func() qml.QQmlImageProviderBase__Flag) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::flags", f)
	}
}

func (ptr *QQuickImageProvider) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::flags")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageProvider::imageType"); signal != nil {
		return C.longlong(signal.(func() qml.QQmlImageProviderBase__ImageType)())
	}

	return C.longlong(NewQQuickImageProviderFromPointer(ptr).ImageTypeDefault())
}

func (ptr *QQuickImageProvider) ConnectImageType(f func() qml.QQmlImageProviderBase__ImageType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::imageType", f)
	}
}

func (ptr *QQuickImageProvider) DisconnectImageType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::imageType")
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

//export callbackQQuickImageProvider_RequestImage
func callbackQQuickImageProvider_RequestImage(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageProvider::requestImage"); signal != nil {
		return gui.PointerFromQImage(signal.(func(string, *core.QSize, *core.QSize) *gui.QImage)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return gui.PointerFromQImage(NewQQuickImageProviderFromPointer(ptr).RequestImageDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickImageProvider) ConnectRequestImage(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QImage) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::requestImage", f)
	}
}

func (ptr *QQuickImageProvider) DisconnectRequestImage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::requestImage")
	}
}

func (ptr *QQuickImageProvider) RequestImage(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQImageFromPointer(C.QQuickImageProvider_RequestImage(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageProvider) RequestImageDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQImageFromPointer(C.QQuickImageProvider_RequestImageDefault(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

//export callbackQQuickImageProvider_RequestPixmap
func callbackQQuickImageProvider_RequestPixmap(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageProvider::requestPixmap"); signal != nil {
		return gui.PointerFromQPixmap(signal.(func(string, *core.QSize, *core.QSize) *gui.QPixmap)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return gui.PointerFromQPixmap(NewQQuickImageProviderFromPointer(ptr).RequestPixmapDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickImageProvider) ConnectRequestPixmap(f func(id string, size *core.QSize, requestedSize *core.QSize) *gui.QPixmap) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::requestPixmap", f)
	}
}

func (ptr *QQuickImageProvider) DisconnectRequestPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::requestPixmap")
	}
}

func (ptr *QQuickImageProvider) RequestPixmap(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQPixmapFromPointer(C.QQuickImageProvider_RequestPixmap(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageProvider) RequestPixmapDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = gui.NewQPixmapFromPointer(C.QQuickImageProvider_RequestPixmapDefault(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQQuickImageProvider_RequestTexture
func callbackQQuickImageProvider_RequestTexture(ptr unsafe.Pointer, id C.struct_QtQuick_PackedString, size unsafe.Pointer, requestedSize unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageProvider::requestTexture"); signal != nil {
		return PointerFromQQuickTextureFactory(signal.(func(string, *core.QSize, *core.QSize) *QQuickTextureFactory)(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
	}

	return PointerFromQQuickTextureFactory(NewQQuickImageProviderFromPointer(ptr).RequestTextureDefault(cGoUnpackString(id), core.NewQSizeFromPointer(size), core.NewQSizeFromPointer(requestedSize)))
}

func (ptr *QQuickImageProvider) ConnectRequestTexture(f func(id string, size *core.QSize, requestedSize *core.QSize) *QQuickTextureFactory) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::requestTexture", f)
	}
}

func (ptr *QQuickImageProvider) DisconnectRequestTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::requestTexture")
	}
}

func (ptr *QQuickImageProvider) RequestTexture(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickImageProvider_RequestTexture(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickImageProvider) RequestTextureDefault(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickImageProvider_RequestTextureDefault(ptr.Pointer(), idC, core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickImageProvider_DestroyQQuickImageProvider
func callbackQQuickImageProvider_DestroyQQuickImageProvider(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageProvider::~QQuickImageProvider"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageProviderFromPointer(ptr).DestroyQQuickImageProviderDefault()
	}
}

func (ptr *QQuickImageProvider) ConnectDestroyQQuickImageProvider(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::~QQuickImageProvider", f)
	}
}

func (ptr *QQuickImageProvider) DisconnectDestroyQQuickImageProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageProvider::~QQuickImageProvider")
	}
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProvider() {
	if ptr.Pointer() != nil {
		C.QQuickImageProvider_DestroyQQuickImageProvider(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProviderDefault() {
	if ptr.Pointer() != nil {
		C.QQuickImageProvider_DestroyQQuickImageProviderDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QQuickImageResponse struct {
	core.QObject
}

type QQuickImageResponse_ITF interface {
	core.QObject_ITF
	QQuickImageResponse_PTR() *QQuickImageResponse
}

func (p *QQuickImageResponse) QQuickImageResponse_PTR() *QQuickImageResponse {
	return p
}

func (p *QQuickImageResponse) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickImageResponse) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickImageResponse(ptr QQuickImageResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageResponse_PTR().Pointer()
	}
	return nil
}

func NewQQuickImageResponseFromPointer(ptr unsafe.Pointer) *QQuickImageResponse {
	var n = new(QQuickImageResponse)
	n.SetPointer(ptr)
	return n
}
func NewQQuickImageResponse() *QQuickImageResponse {
	var tmpValue = NewQQuickImageResponseFromPointer(C.QQuickImageResponse_NewQQuickImageResponse())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickImageResponse_Cancel
func callbackQQuickImageResponse_Cancel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::cancel"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).CancelDefault()
	}
}

func (ptr *QQuickImageResponse) ConnectCancel(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::cancel", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectCancel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::cancel")
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

//export callbackQQuickImageResponse_ErrorString
func callbackQQuickImageResponse_ErrorString(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQQuickImageResponseFromPointer(ptr).ErrorStringDefault())
}

func (ptr *QQuickImageResponse) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::errorString", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::errorString")
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

//export callbackQQuickImageResponse_Finished
func callbackQQuickImageResponse_Finished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickImageResponse) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::finished", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::finished")
	}
}

func (ptr *QQuickImageResponse) Finished() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_Finished(ptr.Pointer())
	}
}

//export callbackQQuickImageResponse_TextureFactory
func callbackQQuickImageResponse_TextureFactory(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::textureFactory"); signal != nil {
		return PointerFromQQuickTextureFactory(signal.(func() *QQuickTextureFactory)())
	}

	return PointerFromQQuickTextureFactory(NewQQuickTextureFactory())
}

func (ptr *QQuickImageResponse) ConnectTextureFactory(f func() *QQuickTextureFactory) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::textureFactory", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectTextureFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::textureFactory")
	}
}

func (ptr *QQuickImageResponse) TextureFactory() *QQuickTextureFactory {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickImageResponse_TextureFactory(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickImageResponse_DestroyQQuickImageResponse
func callbackQQuickImageResponse_DestroyQQuickImageResponse(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::~QQuickImageResponse"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).DestroyQQuickImageResponseDefault()
	}
}

func (ptr *QQuickImageResponse) ConnectDestroyQQuickImageResponse(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::~QQuickImageResponse", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectDestroyQQuickImageResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::~QQuickImageResponse")
	}
}

func (ptr *QQuickImageResponse) DestroyQQuickImageResponse() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DestroyQQuickImageResponse(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickImageResponse) DestroyQQuickImageResponseDefault() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DestroyQQuickImageResponseDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickImageResponse_TimerEvent
func callbackQQuickImageResponse_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::timerEvent", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::timerEvent")
	}
}

func (ptr *QQuickImageResponse) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickImageResponse) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickImageResponse_ChildEvent
func callbackQQuickImageResponse_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::childEvent", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::childEvent")
	}
}

func (ptr *QQuickImageResponse) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickImageResponse) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickImageResponse_ConnectNotify
func callbackQQuickImageResponse_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickImageResponseFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickImageResponse) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::connectNotify", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::connectNotify")
	}
}

func (ptr *QQuickImageResponse) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickImageResponse) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickImageResponse_CustomEvent
func callbackQQuickImageResponse_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::customEvent", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::customEvent")
	}
}

func (ptr *QQuickImageResponse) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickImageResponse) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickImageResponse_DeleteLater
func callbackQQuickImageResponse_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickImageResponse) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::deleteLater", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::deleteLater")
	}
}

func (ptr *QQuickImageResponse) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickImageResponse) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickImageResponse_DisconnectNotify
func callbackQQuickImageResponse_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickImageResponseFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickImageResponse) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::disconnectNotify", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::disconnectNotify")
	}
}

func (ptr *QQuickImageResponse) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickImageResponse) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickImageResponse_Event
func callbackQQuickImageResponse_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickImageResponseFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickImageResponse) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::event", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::event")
	}
}

func (ptr *QQuickImageResponse) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickImageResponse) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickImageResponse_EventFilter
func callbackQQuickImageResponse_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickImageResponseFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickImageResponse) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::eventFilter", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::eventFilter")
	}
}

func (ptr *QQuickImageResponse) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickImageResponse) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickImageResponse_MetaObject
func callbackQQuickImageResponse_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickImageResponse::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickImageResponseFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickImageResponse) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::metaObject", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickImageResponse::metaObject")
	}
}

func (ptr *QQuickImageResponse) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickImageResponse_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickImageResponse) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickImageResponse_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickItem struct {
	core.QObject
	qml.QQmlParserStatus
}

type QQuickItem_ITF interface {
	core.QObject_ITF
	qml.QQmlParserStatus_ITF
	QQuickItem_PTR() *QQuickItem
}

func (p *QQuickItem) QQuickItem_PTR() *QQuickItem {
	return p
}

func (p *QQuickItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
		p.QQmlParserStatus_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickItem(ptr QQuickItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemFromPointer(ptr unsafe.Pointer) *QQuickItem {
	var n = new(QQuickItem)
	n.SetPointer(ptr)
	return n
}
func NewQQuickItem(parent QQuickItem_ITF) *QQuickItem {
	var tmpValue = NewQQuickItemFromPointer(C.QQuickItem_NewQQuickItem(PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickItem) ActiveFocusOnTab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_ActiveFocusOnTab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Antialiasing() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) BaselineOffset() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_BaselineOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ChildrenRect() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QQuickItem_ChildrenRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) Clip() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Clip(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasActiveFocus() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_HasActiveFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_HasFocus(ptr.Pointer()) != 0
	}
	return false
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

func (ptr *QQuickItem) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickItem_IsTextureProvider
func callbackQQuickItem_IsTextureProvider(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::isTextureProvider"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).IsTextureProviderDefault())))
}

func (ptr *QQuickItem) ConnectIsTextureProvider(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::isTextureProvider", f)
	}
}

func (ptr *QQuickItem) DisconnectIsTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::isTextureProvider")
	}
}

func (ptr *QQuickItem) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsTextureProviderDefault() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProviderDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ParentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickItem_ParentItem(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
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

func (ptr *QQuickItem) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
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

func (ptr *QQuickItem) SetSmooth(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetSmooth(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickItem) SetState(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QQuickItem_SetState(ptr.Pointer(), vqsC)
	}
}

func (ptr *QQuickItem) SetTransformOrigin(vtr QQuickItem__TransformOrigin) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetTransformOrigin(ptr.Pointer(), C.longlong(vtr))
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

func (ptr *QQuickItem) Smooth() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Smooth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) State() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQuickItem_State(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickItem_TextureProvider
func callbackQQuickItem_TextureProvider(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::textureProvider"); signal != nil {
		return PointerFromQSGTextureProvider(signal.(func() *QSGTextureProvider)())
	}

	return PointerFromQSGTextureProvider(NewQQuickItemFromPointer(ptr).TextureProviderDefault())
}

func (ptr *QQuickItem) ConnectTextureProvider(f func() *QSGTextureProvider) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::textureProvider", f)
	}
}

func (ptr *QQuickItem) DisconnectTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::textureProvider")
	}
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProvider(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) TextureProviderDefault() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProviderDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) TransformOrigin() QQuickItem__TransformOrigin {
	if ptr.Pointer() != nil {
		return QQuickItem__TransformOrigin(C.QQuickItem_TransformOrigin(ptr.Pointer()))
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

func (ptr *QQuickItem) AcceptHoverEvents() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_AcceptHoverEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) AcceptedMouseButtons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QQuickItem_AcceptedMouseButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ChildAt(x float64, y float64) *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickItem_ChildAt(ptr.Pointer(), C.double(x), C.double(y)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) ChildItems() []*QQuickItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQuick_PackedList) []*QQuickItem {
			var out = make([]*QQuickItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQQuickItemFromPointer(l.data).childItems_atList(i)
			}
			return out
		}(C.QQuickItem_ChildItems(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickItem_ChildMouseEventFilter
func callbackQQuickItem_ChildMouseEventFilter(ptr unsafe.Pointer, item unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::childMouseEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QQuickItem, *core.QEvent) bool)(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).ChildMouseEventFilterDefault(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickItem) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::childMouseEventFilter", f)
	}
}

func (ptr *QQuickItem) DisconnectChildMouseEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::childMouseEventFilter")
	}
}

func (ptr *QQuickItem) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItem) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_ChildMouseEventFilterDefault(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickItem_ClassBegin
func callbackQQuickItem_ClassBegin(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::classBegin"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *QQuickItem) ConnectClassBegin(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::classBegin", f)
	}
}

func (ptr *QQuickItem) DisconnectClassBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::classBegin")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::componentComplete"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *QQuickItem) ConnectComponentComplete(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::componentComplete", f)
	}
}

func (ptr *QQuickItem) DisconnectComponentComplete() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::componentComplete")
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

//export callbackQQuickItem_Contains
func callbackQQuickItem_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QQuickItem) ConnectContains(f func(point *core.QPointF) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::contains", f)
	}
}

func (ptr *QQuickItem) DisconnectContains() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::contains")
	}
}

func (ptr *QQuickItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickItem) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickItem) Cursor() *gui.QCursor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQCursorFromPointer(C.QQuickItem_Cursor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QCursor).DestroyQCursor)
		return tmpValue
	}
	return nil
}

//export callbackQQuickItem_DragEnterEvent
func callbackQQuickItem_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dragEnterEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dragEnterEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dragLeaveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dragLeaveEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dragMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dragMoveEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dropEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::dropEvent")
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

//export callbackQQuickItem_Event
func callbackQQuickItem_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QQuickItem) ConnectEvent(f func(ev *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::event", f)
	}
}

func (ptr *QQuickItem) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::event")
	}
}

func (ptr *QQuickItem) Event(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QQuickItem) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QQuickItem) FiltersChildMouseEvents() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_FiltersChildMouseEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Flags() QQuickItem__Flag {
	if ptr.Pointer() != nil {
		return QQuickItem__Flag(C.QQuickItem_Flags(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickItem_FocusInEvent
func callbackQQuickItem_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::focusInEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::focusInEvent")
	}
}

func (ptr *QQuickItem) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickItem_FocusOutEvent
func callbackQQuickItem_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::focusOutEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::focusOutEvent")
	}
}

func (ptr *QQuickItem) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickItemFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *QQuickItem) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::geometryChanged", f)
	}
}

func (ptr *QQuickItem) DisconnectGeometryChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::geometryChanged")
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

func (ptr *QQuickItem) HeightValid() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_HeightValid(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickItem_HoverEnterEvent
func callbackQQuickItem_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::hoverEnterEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::hoverEnterEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::hoverLeaveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::hoverLeaveEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::hoverMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::hoverMoveEvent")
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

func (ptr *QQuickItem) ImplicitWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickItem_InputMethodEvent
func callbackQQuickItem_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::inputMethodEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::inputMethodEvent")
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

//export callbackQQuickItem_InputMethodQuery
func callbackQQuickItem_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::inputMethodQuery", f)
	}
}

func (ptr *QQuickItem) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::inputMethodQuery")
	}
}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickItem_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickItem_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) IsAncestorOf(child QQuickItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsAncestorOf(ptr.Pointer(), PointerFromQQuickItem(child)) != 0
	}
	return false
}

func (ptr *QQuickItem) IsComponentComplete() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsComponentComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsFocusScope() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsFocusScope(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) DisconnectItemChange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::itemChange")
	}
}

func (ptr *QQuickItem) KeepMouseGrab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepMouseGrab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepTouchGrab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepTouchGrab(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickItem_KeyPressEvent
func callbackQQuickItem_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::keyPressEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::keyPressEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::keyReleaseEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::keyReleaseEvent")
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

func (ptr *QQuickItem) MapFromGlobal(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QQuickItem_MapFromGlobal(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapFromItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QQuickItem_MapFromItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapFromScene(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QQuickItem_MapFromScene(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectFromItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QQuickItem_MapRectFromItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectFromScene(rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QQuickItem_MapRectFromScene(ptr.Pointer(), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectToItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QQuickItem_MapRectToItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapRectToScene(rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QQuickItem_MapRectToScene(ptr.Pointer(), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapToGlobal(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QQuickItem_MapToGlobal(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapToItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QQuickItem_MapToItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) MapToScene(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QQuickItem_MapToScene(ptr.Pointer(), core.PointerFromQPointF(point)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

//export callbackQQuickItem_MouseDoubleClickEvent
func callbackQQuickItem_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseDoubleClickEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseMoveEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mousePressEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mousePressEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseReleaseEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseReleaseEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *QQuickItem) ConnectMouseUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseUngrabEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::mouseUngrabEvent")
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

func (ptr *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickItem_NextItemInFocusChain(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward)))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) Polish() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Polish(ptr.Pointer())
	}
}

//export callbackQQuickItem_ReleaseResources
func callbackQQuickItem_ReleaseResources(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickItem) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::releaseResources", f)
	}
}

func (ptr *QQuickItem) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::releaseResources")
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

func (ptr *QQuickItem) ScopedFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickItem_ScopedFocusItem(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptHoverEvents(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQuickItem) SetAcceptedMouseButtons(buttons core.Qt__MouseButton) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptedMouseButtons(ptr.Pointer(), C.longlong(buttons))
	}
}

func (ptr *QQuickItem) SetCursor(cursor gui.QCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetCursor(ptr.Pointer(), gui.PointerFromQCursor(cursor))
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::touchEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::touchEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *QQuickItem) ConnectTouchUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::touchUngrabEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTouchUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::touchUngrabEvent")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::update"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItem) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::update", f)
	}
}

func (ptr *QQuickItem) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::update")
	}
}

func (ptr *QQuickItem) Update() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Update(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UpdateInputMethod(queries core.Qt__InputMethodQuery) {
	if ptr.Pointer() != nil {
		C.QQuickItem_UpdateInputMethod(ptr.Pointer(), C.longlong(queries))
	}
}

func (ptr *QQuickItem) DisconnectUpdatePaintNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::updatePaintNode")
	}
}

//export callbackQQuickItem_UpdatePolish
func callbackQQuickItem_UpdatePolish(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *QQuickItem) ConnectUpdatePolish(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::updatePolish", f)
	}
}

func (ptr *QQuickItem) DisconnectUpdatePolish() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::updatePolish")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::wheelEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::wheelEvent")
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

func (ptr *QQuickItem) WidthValid() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_WidthValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Window() *QQuickWindow {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickWindowFromPointer(C.QQuickItem_Window(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickItem_WindowChanged
func callbackQQuickItem_WindowChanged(ptr unsafe.Pointer, window unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::windowChanged"); signal != nil {
		signal.(func(*QQuickWindow))(NewQQuickWindowFromPointer(window))
	}

}

func (ptr *QQuickItem) ConnectWindowChanged(f func(window *QQuickWindow)) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectWindowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::windowChanged", f)
	}
}

func (ptr *QQuickItem) DisconnectWindowChanged() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::windowChanged")
	}
}

func (ptr *QQuickItem) WindowChanged(window QQuickWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_WindowChanged(ptr.Pointer(), PointerFromQQuickWindow(window))
	}
}

//export callbackQQuickItem_DestroyQQuickItem
func callbackQQuickItem_DestroyQQuickItem(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::~QQuickItem"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).DestroyQQuickItemDefault()
	}
}

func (ptr *QQuickItem) ConnectDestroyQQuickItem(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::~QQuickItem", f)
	}
}

func (ptr *QQuickItem) DisconnectDestroyQQuickItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::~QQuickItem")
	}
}

func (ptr *QQuickItem) DestroyQQuickItem() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItem(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItem) DestroyQQuickItemDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItemDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItem) childItems_atList(i int) *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickItem_childItems_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickItem_TimerEvent
func callbackQQuickItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::timerEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::timerEvent")
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

//export callbackQQuickItem_ChildEvent
func callbackQQuickItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::childEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::childEvent")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::connectNotify", f)
	}
}

func (ptr *QQuickItem) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::connectNotify")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::customEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::customEvent")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickItem) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::deleteLater", f)
	}
}

func (ptr *QQuickItem) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::deleteLater")
	}
}

func (ptr *QQuickItem) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickItem_DisconnectNotify
func callbackQQuickItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::disconnectNotify", f)
	}
}

func (ptr *QQuickItem) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::disconnectNotify")
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

//export callbackQQuickItem_EventFilter
func callbackQQuickItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::eventFilter", f)
	}
}

func (ptr *QQuickItem) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::eventFilter")
	}
}

func (ptr *QQuickItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickItem_MetaObject
func callbackQQuickItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItem::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickItem) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::metaObject", f)
	}
}

func (ptr *QQuickItem) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItem::metaObject")
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

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResult_ITF interface {
	core.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func (p *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return p
}

func (p *QQuickItemGrabResult) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickItemGrabResult) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickItemGrabResult(ptr QQuickItemGrabResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemGrabResult_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemGrabResultFromPointer(ptr unsafe.Pointer) *QQuickItemGrabResult {
	var n = new(QQuickItemGrabResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuickItemGrabResult) DestroyQQuickItemGrabResult() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QQuickItemGrabResult) Image() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QQuickItemGrabResult_Image(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickItemGrabResult) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQuickItemGrabResult_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQuickItemGrabResult_Ready
func callbackQQuickItemGrabResult_Ready(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::ready"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectReady(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::ready", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectReady(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::ready")
	}
}

func (ptr *QQuickItemGrabResult) Ready() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_Ready(ptr.Pointer())
	}
}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		return C.QQuickItemGrabResult_SaveToFile(ptr.Pointer(), fileNameC) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_TimerEvent
func callbackQQuickItemGrabResult_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::timerEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::timerEvent")
	}
}

func (ptr *QQuickItemGrabResult) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickItemGrabResult_ChildEvent
func callbackQQuickItemGrabResult_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::childEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::childEvent")
	}
}

func (ptr *QQuickItemGrabResult) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickItemGrabResult_ConnectNotify
func callbackQQuickItemGrabResult_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItemGrabResult) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::connectNotify", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::connectNotify")
	}
}

func (ptr *QQuickItemGrabResult) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItemGrabResult) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItemGrabResult_CustomEvent
func callbackQQuickItemGrabResult_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::customEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::customEvent")
	}
}

func (ptr *QQuickItemGrabResult) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickItemGrabResult_DeleteLater
func callbackQQuickItemGrabResult_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickItemGrabResult) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::deleteLater", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::deleteLater")
	}
}

func (ptr *QQuickItemGrabResult) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItemGrabResult) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickItemGrabResult_DisconnectNotify
func callbackQQuickItemGrabResult_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItemGrabResult) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::disconnectNotify", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::disconnectNotify")
	}
}

func (ptr *QQuickItemGrabResult) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItemGrabResult) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItemGrabResult_Event
func callbackQQuickItemGrabResult_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemGrabResultFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickItemGrabResult) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::event", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::event")
	}
}

func (ptr *QQuickItemGrabResult) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickItemGrabResult) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_EventFilter
func callbackQQuickItemGrabResult_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickItemGrabResultFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickItemGrabResult) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::eventFilter", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::eventFilter")
	}
}

func (ptr *QQuickItemGrabResult) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItemGrabResult) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_MetaObject
func callbackQQuickItemGrabResult_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickItemGrabResult::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickItemGrabResultFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickItemGrabResult) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::metaObject", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickItemGrabResult::metaObject")
	}
}

func (ptr *QQuickItemGrabResult) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItemGrabResult_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItemGrabResult) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItemGrabResult_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickPaintedItem struct {
	QQuickItem
}

type QQuickPaintedItem_ITF interface {
	QQuickItem_ITF
	QQuickPaintedItem_PTR() *QQuickPaintedItem
}

func (p *QQuickPaintedItem) QQuickPaintedItem_PTR() *QQuickPaintedItem {
	return p
}

func (p *QQuickPaintedItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (p *QQuickPaintedItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickItem_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickPaintedItem(ptr QQuickPaintedItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickPaintedItem_PTR().Pointer()
	}
	return nil
}

func NewQQuickPaintedItemFromPointer(ptr unsafe.Pointer) *QQuickPaintedItem {
	var n = new(QQuickPaintedItem)
	n.SetPointer(ptr)
	return n
}
func (ptr *QQuickPaintedItem) ContentsScale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickPaintedItem_ContentsScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) ContentsSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickPaintedItem_ContentsSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickPaintedItem) FillColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QQuickPaintedItem_FillColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickPaintedItem) DisconnectItemChange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::itemChange")
	}
}

func (ptr *QQuickPaintedItem) RenderTarget() QQuickPaintedItem__RenderTarget {
	if ptr.Pointer() != nil {
		return QQuickPaintedItem__RenderTarget(C.QQuickPaintedItem_RenderTarget(ptr.Pointer()))
	}
	return 0
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

func (ptr *QQuickPaintedItem) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickPaintedItem_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func NewQQuickPaintedItem(parent QQuickItem_ITF) *QQuickPaintedItem {
	var tmpValue = NewQQuickPaintedItemFromPointer(C.QQuickPaintedItem_NewQQuickPaintedItem(PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickPaintedItem) Antialiasing() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_ContentsScaleChanged
func callbackQQuickPaintedItem_ContentsScaleChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::contentsScaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectContentsScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsScaleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::contentsScaleChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsScaleChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::contentsScaleChanged")
	}
}

func (ptr *QQuickPaintedItem) ContentsScaleChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsScaleChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_ContentsSizeChanged
func callbackQQuickPaintedItem_ContentsSizeChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::contentsSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectContentsSizeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::contentsSizeChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::contentsSizeChanged")
	}
}

func (ptr *QQuickPaintedItem) ContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsSizeChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_FillColorChanged
func callbackQQuickPaintedItem_FillColorChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::fillColorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectFillColorChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectFillColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::fillColorChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFillColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectFillColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::fillColorChanged")
	}
}

func (ptr *QQuickPaintedItem) FillColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FillColorChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_IsTextureProvider
func callbackQQuickPaintedItem_IsTextureProvider(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::isTextureProvider"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).IsTextureProviderDefault())))
}

func (ptr *QQuickPaintedItem) ConnectIsTextureProvider(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::isTextureProvider", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectIsTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::isTextureProvider")
	}
}

func (ptr *QQuickPaintedItem) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) IsTextureProviderDefault() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_IsTextureProviderDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) Mipmap() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Mipmap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) OpaquePainting() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_OpaquePainting(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_Paint
func callbackQQuickPaintedItem_Paint(ptr unsafe.Pointer, painter unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::paint"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	}

}

func (ptr *QQuickPaintedItem) ConnectPaint(f func(painter *gui.QPainter)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::paint", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectPaint(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::paint")
	}
}

func (ptr *QQuickPaintedItem) Paint(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QQuickPaintedItem) PerformanceHints() QQuickPaintedItem__PerformanceHint {
	if ptr.Pointer() != nil {
		return QQuickPaintedItem__PerformanceHint(C.QQuickPaintedItem_PerformanceHints(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickPaintedItem_ReleaseResources
func callbackQQuickPaintedItem_ReleaseResources(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::releaseResources", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::releaseResources")
	}
}

func (ptr *QQuickPaintedItem) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ReleaseResourcesDefault(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_RenderTargetChanged
func callbackQQuickPaintedItem_RenderTargetChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::renderTargetChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectRenderTargetChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectRenderTargetChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::renderTargetChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectRenderTargetChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectRenderTargetChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::renderTargetChanged")
	}
}

func (ptr *QQuickPaintedItem) RenderTargetChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_RenderTargetChanged(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) SetAntialiasing(enable bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetAntialiasing(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
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

//export callbackQQuickPaintedItem_TextureProvider
func callbackQQuickPaintedItem_TextureProvider(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::textureProvider"); signal != nil {
		return PointerFromQSGTextureProvider(signal.(func() *QSGTextureProvider)())
	}

	return PointerFromQSGTextureProvider(NewQQuickPaintedItemFromPointer(ptr).TextureProviderDefault())
}

func (ptr *QQuickPaintedItem) ConnectTextureProvider(f func() *QSGTextureProvider) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::textureProvider", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTextureProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::textureProvider")
	}
}

func (ptr *QQuickPaintedItem) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureProviderFromPointer(C.QQuickPaintedItem_TextureProvider(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickPaintedItem) TextureProviderDefault() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureProviderFromPointer(C.QQuickPaintedItem_TextureProviderDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickPaintedItem_TextureSizeChanged
func callbackQQuickPaintedItem_TextureSizeChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::textureSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectTextureSizeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectTextureSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::textureSizeChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTextureSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectTextureSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::textureSizeChanged")
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

func (ptr *QQuickPaintedItem) DisconnectUpdatePaintNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::updatePaintNode")
	}
}

//export callbackQQuickPaintedItem_DestroyQQuickPaintedItem
func callbackQQuickPaintedItem_DestroyQQuickPaintedItem(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::~QQuickPaintedItem"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DestroyQQuickPaintedItemDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectDestroyQQuickPaintedItem(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::~QQuickPaintedItem", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDestroyQQuickPaintedItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::~QQuickPaintedItem")
	}
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItem() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DestroyQQuickPaintedItem(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItemDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DestroyQQuickPaintedItemDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickPaintedItem_ChildMouseEventFilter
func callbackQQuickPaintedItem_ChildMouseEventFilter(ptr unsafe.Pointer, item unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::childMouseEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QQuickItem, *core.QEvent) bool)(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).ChildMouseEventFilterDefault(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickPaintedItem) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::childMouseEventFilter", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectChildMouseEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::childMouseEventFilter")
	}
}

func (ptr *QQuickPaintedItem) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_ChildMouseEventFilterDefault(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_ClassBegin
func callbackQQuickPaintedItem_ClassBegin(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::classBegin"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectClassBegin(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::classBegin", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectClassBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::classBegin")
	}
}

func (ptr *QQuickPaintedItem) ClassBegin() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ClassBeginDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ClassBeginDefault(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_ComponentComplete
func callbackQQuickPaintedItem_ComponentComplete(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::componentComplete"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectComponentComplete(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::componentComplete", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectComponentComplete() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::componentComplete")
	}
}

func (ptr *QQuickPaintedItem) ComponentComplete() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ComponentComplete(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ComponentCompleteDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ComponentCompleteDefault(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_Contains
func callbackQQuickPaintedItem_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QQuickPaintedItem) ConnectContains(f func(point *core.QPointF) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::contains", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContains() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::contains")
	}
}

func (ptr *QQuickPaintedItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_DragEnterEvent
func callbackQQuickPaintedItem_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dragEnterEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dragEnterEvent")
	}
}

func (ptr *QQuickPaintedItem) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQQuickPaintedItem_DragLeaveEvent
func callbackQQuickPaintedItem_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dragLeaveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dragLeaveEvent")
	}
}

func (ptr *QQuickPaintedItem) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQQuickPaintedItem_DragMoveEvent
func callbackQQuickPaintedItem_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dragMoveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dragMoveEvent")
	}
}

func (ptr *QQuickPaintedItem) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQQuickPaintedItem_DropEvent
func callbackQQuickPaintedItem_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dropEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::dropEvent")
	}
}

func (ptr *QQuickPaintedItem) DropEvent(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQQuickPaintedItem_Event
func callbackQQuickPaintedItem_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QQuickPaintedItem) ConnectEvent(f func(ev *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::event", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::event")
	}
}

func (ptr *QQuickPaintedItem) Event(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_FocusInEvent
func callbackQQuickPaintedItem_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::focusInEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::focusInEvent")
	}
}

func (ptr *QQuickPaintedItem) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickPaintedItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickPaintedItem_FocusOutEvent
func callbackQQuickPaintedItem_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::focusOutEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::focusOutEvent")
	}
}

func (ptr *QQuickPaintedItem) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickPaintedItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickPaintedItem_GeometryChanged
func callbackQQuickPaintedItem_GeometryChanged(ptr unsafe.Pointer, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *QQuickPaintedItem) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::geometryChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectGeometryChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::geometryChanged")
	}
}

func (ptr *QQuickPaintedItem) GeometryChanged(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_GeometryChanged(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickPaintedItem) GeometryChangedDefault(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_GeometryChangedDefault(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

//export callbackQQuickPaintedItem_HoverEnterEvent
func callbackQQuickPaintedItem_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::hoverEnterEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectHoverEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::hoverEnterEvent")
	}
}

func (ptr *QQuickPaintedItem) HoverEnterEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverEnterEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) HoverEnterEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverEnterEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickPaintedItem_HoverLeaveEvent
func callbackQQuickPaintedItem_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::hoverLeaveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectHoverLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::hoverLeaveEvent")
	}
}

func (ptr *QQuickPaintedItem) HoverLeaveEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverLeaveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) HoverLeaveEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverLeaveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickPaintedItem_HoverMoveEvent
func callbackQQuickPaintedItem_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::hoverMoveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectHoverMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::hoverMoveEvent")
	}
}

func (ptr *QQuickPaintedItem) HoverMoveEvent(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverMoveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) HoverMoveEventDefault(event gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverMoveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

//export callbackQQuickPaintedItem_InputMethodEvent
func callbackQQuickPaintedItem_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::inputMethodEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::inputMethodEvent")
	}
}

func (ptr *QQuickPaintedItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickPaintedItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQQuickPaintedItem_InputMethodQuery
func callbackQQuickPaintedItem_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickPaintedItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickPaintedItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::inputMethodQuery", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::inputMethodQuery")
	}
}

func (ptr *QQuickPaintedItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickPaintedItem_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickPaintedItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickPaintedItem_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQQuickPaintedItem_KeyPressEvent
func callbackQQuickPaintedItem_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::keyPressEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::keyPressEvent")
	}
}

func (ptr *QQuickPaintedItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickPaintedItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQQuickPaintedItem_KeyReleaseEvent
func callbackQQuickPaintedItem_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::keyReleaseEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::keyReleaseEvent")
	}
}

func (ptr *QQuickPaintedItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickPaintedItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQQuickPaintedItem_MouseDoubleClickEvent
func callbackQQuickPaintedItem_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseDoubleClickEvent")
	}
}

func (ptr *QQuickPaintedItem) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickPaintedItem_MouseMoveEvent
func callbackQQuickPaintedItem_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseMoveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseMoveEvent")
	}
}

func (ptr *QQuickPaintedItem) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickPaintedItem_MousePressEvent
func callbackQQuickPaintedItem_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mousePressEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mousePressEvent")
	}
}

func (ptr *QQuickPaintedItem) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickPaintedItem_MouseReleaseEvent
func callbackQQuickPaintedItem_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseReleaseEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseReleaseEvent")
	}
}

func (ptr *QQuickPaintedItem) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickPaintedItem_MouseUngrabEvent
func callbackQQuickPaintedItem_MouseUngrabEvent(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseUngrabEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::mouseUngrabEvent")
	}
}

func (ptr *QQuickPaintedItem) MouseUngrabEvent() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) MouseUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseUngrabEventDefault(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_TouchEvent
func callbackQQuickPaintedItem_TouchEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::touchEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::touchEvent")
	}
}

func (ptr *QQuickPaintedItem) TouchEvent(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickPaintedItem) TouchEventDefault(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

//export callbackQQuickPaintedItem_TouchUngrabEvent
func callbackQQuickPaintedItem_TouchUngrabEvent(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectTouchUngrabEvent(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::touchUngrabEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTouchUngrabEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::touchUngrabEvent")
	}
}

func (ptr *QQuickPaintedItem) TouchUngrabEvent() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) TouchUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchUngrabEventDefault(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_UpdatePolish
func callbackQQuickPaintedItem_UpdatePolish(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectUpdatePolish(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::updatePolish", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectUpdatePolish() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::updatePolish")
	}
}

func (ptr *QQuickPaintedItem) UpdatePolish() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_UpdatePolish(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) UpdatePolishDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_UpdatePolishDefault(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_WheelEvent
func callbackQQuickPaintedItem_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::wheelEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::wheelEvent")
	}
}

func (ptr *QQuickPaintedItem) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickPaintedItem) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQQuickPaintedItem_TimerEvent
func callbackQQuickPaintedItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::timerEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::timerEvent")
	}
}

func (ptr *QQuickPaintedItem) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickPaintedItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickPaintedItem_ChildEvent
func callbackQQuickPaintedItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::childEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::childEvent")
	}
}

func (ptr *QQuickPaintedItem) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickPaintedItem_ConnectNotify
func callbackQQuickPaintedItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickPaintedItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::connectNotify", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::connectNotify")
	}
}

func (ptr *QQuickPaintedItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickPaintedItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickPaintedItem_CustomEvent
func callbackQQuickPaintedItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::customEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::customEvent")
	}
}

func (ptr *QQuickPaintedItem) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickPaintedItem) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickPaintedItem_DeleteLater
func callbackQQuickPaintedItem_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::deleteLater", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::deleteLater")
	}
}

func (ptr *QQuickPaintedItem) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickPaintedItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickPaintedItem_DisconnectNotify
func callbackQQuickPaintedItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickPaintedItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::disconnectNotify", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::disconnectNotify")
	}
}

func (ptr *QQuickPaintedItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickPaintedItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickPaintedItem_EventFilter
func callbackQQuickPaintedItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickPaintedItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::eventFilter", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::eventFilter")
	}
}

func (ptr *QQuickPaintedItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_MetaObject
func callbackQQuickPaintedItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickPaintedItem::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickPaintedItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickPaintedItem) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::metaObject", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickPaintedItem::metaObject")
	}
}

func (ptr *QQuickPaintedItem) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickPaintedItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickPaintedItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickRenderControl struct {
	core.QObject
}

type QQuickRenderControl_ITF interface {
	core.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func (p *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return p
}

func (p *QQuickRenderControl) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickRenderControl) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickRenderControl(ptr QQuickRenderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickRenderControl_PTR().Pointer()
	}
	return nil
}

func NewQQuickRenderControlFromPointer(ptr unsafe.Pointer) *QQuickRenderControl {
	var n = new(QQuickRenderControl)
	n.SetPointer(ptr)
	return n
}
func NewQQuickRenderControl(parent core.QObject_ITF) *QQuickRenderControl {
	var tmpValue = NewQQuickRenderControlFromPointer(C.QQuickRenderControl_NewQQuickRenderControl(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickRenderControl) Grab() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QQuickRenderControl_Grab(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::renderRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectRenderRequested(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::renderRequested", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectRenderRequested(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::renderRequested")
	}
}

func (ptr *QQuickRenderControl) RenderRequested() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_RenderRequested(ptr.Pointer())
	}
}

//export callbackQQuickRenderControl_RenderWindow
func callbackQQuickRenderControl_RenderWindow(ptr unsafe.Pointer, offset unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::renderWindow"); signal != nil {
		return gui.PointerFromQWindow(signal.(func(*core.QPoint) *gui.QWindow)(core.NewQPointFromPointer(offset)))
	}

	return gui.PointerFromQWindow(NewQQuickRenderControlFromPointer(ptr).RenderWindowDefault(core.NewQPointFromPointer(offset)))
}

func (ptr *QQuickRenderControl) ConnectRenderWindow(f func(offset *core.QPoint) *gui.QWindow) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::renderWindow", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::renderWindow")
	}
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindow(ptr.Pointer(), core.PointerFromQPoint(offset)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickRenderControl) RenderWindowDefault(offset core.QPoint_ITF) *gui.QWindow {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindowDefault(ptr.Pointer(), core.PointerFromQPoint(offset)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	var tmpValue = gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickRenderControl) RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	var tmpValue = gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickRenderControl_SceneChanged
func callbackQQuickRenderControl_SceneChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::sceneChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectSceneChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::sceneChanged", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectSceneChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::sceneChanged")
	}
}

func (ptr *QQuickRenderControl) SceneChanged() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_SceneChanged(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) Sync() bool {
	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Sync(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControl(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickRenderControl_TimerEvent
func callbackQQuickRenderControl_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::timerEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::timerEvent")
	}
}

func (ptr *QQuickRenderControl) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickRenderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickRenderControl_ChildEvent
func callbackQQuickRenderControl_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::childEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::childEvent")
	}
}

func (ptr *QQuickRenderControl) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickRenderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickRenderControl_ConnectNotify
func callbackQQuickRenderControl_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickRenderControlFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickRenderControl) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::connectNotify", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::connectNotify")
	}
}

func (ptr *QQuickRenderControl) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickRenderControl) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickRenderControl_CustomEvent
func callbackQQuickRenderControl_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::customEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::customEvent")
	}
}

func (ptr *QQuickRenderControl) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickRenderControl) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickRenderControl_DeleteLater
func callbackQQuickRenderControl_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickRenderControlFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickRenderControl) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::deleteLater", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::deleteLater")
	}
}

func (ptr *QQuickRenderControl) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickRenderControl) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickRenderControl_DisconnectNotify
func callbackQQuickRenderControl_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickRenderControlFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickRenderControl) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::disconnectNotify", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::disconnectNotify")
	}
}

func (ptr *QQuickRenderControl) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickRenderControl) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickRenderControl_Event
func callbackQQuickRenderControl_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickRenderControlFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickRenderControl) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::event", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::event")
	}
}

func (ptr *QQuickRenderControl) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickRenderControl_EventFilter
func callbackQQuickRenderControl_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickRenderControlFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickRenderControl) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::eventFilter", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::eventFilter")
	}
}

func (ptr *QQuickRenderControl) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickRenderControl_MetaObject
func callbackQQuickRenderControl_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickRenderControl::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickRenderControlFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickRenderControl) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::metaObject", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickRenderControl::metaObject")
	}
}

func (ptr *QQuickRenderControl) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickRenderControl_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickRenderControl) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickRenderControl_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocument_ITF interface {
	core.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func (p *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument {
	return p
}

func (p *QQuickTextDocument) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickTextDocument) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickTextDocument(ptr QQuickTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextDocument_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextDocumentFromPointer(ptr unsafe.Pointer) *QQuickTextDocument {
	var n = new(QQuickTextDocument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuickTextDocument) DestroyQQuickTextDocument() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func NewQQuickTextDocument(parent QQuickItem_ITF) *QQuickTextDocument {
	var tmpValue = NewQQuickTextDocumentFromPointer(C.QQuickTextDocument_NewQQuickTextDocument(PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickTextDocument) TextDocument() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQTextDocumentFromPointer(C.QQuickTextDocument_TextDocument(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickTextDocument_TimerEvent
func callbackQQuickTextDocument_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::timerEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::timerEvent")
	}
}

func (ptr *QQuickTextDocument) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextDocument) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickTextDocument_ChildEvent
func callbackQQuickTextDocument_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::childEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::childEvent")
	}
}

func (ptr *QQuickTextDocument) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextDocument) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickTextDocument_ConnectNotify
func callbackQQuickTextDocument_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextDocument) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::connectNotify", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::connectNotify")
	}
}

func (ptr *QQuickTextDocument) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextDocument) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextDocument_CustomEvent
func callbackQQuickTextDocument_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::customEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::customEvent")
	}
}

func (ptr *QQuickTextDocument) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickTextDocument) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickTextDocument_DeleteLater
func callbackQQuickTextDocument_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextDocumentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickTextDocument) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::deleteLater", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::deleteLater")
	}
}

func (ptr *QQuickTextDocument) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickTextDocument) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickTextDocument_DisconnectNotify
func callbackQQuickTextDocument_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextDocument) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::disconnectNotify", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::disconnectNotify")
	}
}

func (ptr *QQuickTextDocument) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextDocument) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextDocument_Event
func callbackQQuickTextDocument_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextDocumentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickTextDocument) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::event", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::event")
	}
}

func (ptr *QQuickTextDocument) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickTextDocument) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickTextDocument_EventFilter
func callbackQQuickTextDocument_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextDocumentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickTextDocument) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::eventFilter", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::eventFilter")
	}
}

func (ptr *QQuickTextDocument) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickTextDocument) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickTextDocument_MetaObject
func callbackQQuickTextDocument_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextDocument::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickTextDocumentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickTextDocument) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::metaObject", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextDocument::metaObject")
	}
}

func (ptr *QQuickTextDocument) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextDocument_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextDocument) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextDocument_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactory_ITF interface {
	core.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func (p *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory {
	return p
}

func (p *QQuickTextureFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickTextureFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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
	return n
}

//export callbackQQuickTextureFactory_Image
func callbackQQuickTextureFactory_Image(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::image"); signal != nil {
		return gui.PointerFromQImage(signal.(func() *gui.QImage)())
	}

	return gui.PointerFromQImage(NewQQuickTextureFactoryFromPointer(ptr).ImageDefault())
}

func (ptr *QQuickTextureFactory) ConnectImage(f func() *gui.QImage) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::image", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectImage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::image")
	}
}

func (ptr *QQuickTextureFactory) Image() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QQuickTextureFactory_Image(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickTextureFactory) ImageDefault() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QQuickTextureFactory_ImageDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func NewQQuickTextureFactory() *QQuickTextureFactory {
	var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_NewQQuickTextureFactory())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickTextureFactory_CreateTexture
func callbackQQuickTextureFactory_CreateTexture(ptr unsafe.Pointer, window unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::createTexture"); signal != nil {
		return PointerFromQSGTexture(signal.(func(*QQuickWindow) *QSGTexture)(NewQQuickWindowFromPointer(window)))
	}

	return PointerFromQSGTexture(NewQSGTexture())
}

func (ptr *QQuickTextureFactory) ConnectCreateTexture(f func(window *QQuickWindow) *QSGTexture) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::createTexture", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectCreateTexture(window QQuickWindow_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::createTexture")
	}
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF) *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QQuickTextureFactory_CreateTexture(ptr.Pointer(), PointerFromQQuickWindow(window)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickTextureFactory_TextureByteCount
func callbackQQuickTextureFactory_TextureByteCount(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::textureByteCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QQuickTextureFactory) ConnectTextureByteCount(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::textureByteCount", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectTextureByteCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::textureByteCount")
	}
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickTextureFactory_TextureByteCount(ptr.Pointer())))
	}
	return 0
}

func QQuickTextureFactory_TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {
	var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(gui.PointerFromQImage(image)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickTextureFactory) TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {
	var tmpValue = NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(gui.PointerFromQImage(image)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickTextureFactory_TextureSize
func callbackQQuickTextureFactory_TextureSize(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QQuickTextureFactory) ConnectTextureSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::textureSize", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectTextureSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::textureSize")
	}
}

func (ptr *QQuickTextureFactory) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickTextureFactory_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickTextureFactory_DestroyQQuickTextureFactory
func callbackQQuickTextureFactory_DestroyQQuickTextureFactory(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::~QQuickTextureFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DestroyQQuickTextureFactoryDefault()
	}
}

func (ptr *QQuickTextureFactory) ConnectDestroyQQuickTextureFactory(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::~QQuickTextureFactory", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectDestroyQQuickTextureFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::~QQuickTextureFactory")
	}
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactory() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactory(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactoryDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickTextureFactory_TimerEvent
func callbackQQuickTextureFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::timerEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::timerEvent")
	}
}

func (ptr *QQuickTextureFactory) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextureFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickTextureFactory_ChildEvent
func callbackQQuickTextureFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::childEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::childEvent")
	}
}

func (ptr *QQuickTextureFactory) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextureFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickTextureFactory_ConnectNotify
func callbackQQuickTextureFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextureFactory) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::connectNotify", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::connectNotify")
	}
}

func (ptr *QQuickTextureFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextureFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextureFactory_CustomEvent
func callbackQQuickTextureFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::customEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::customEvent")
	}
}

func (ptr *QQuickTextureFactory) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickTextureFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickTextureFactory_DeleteLater
func callbackQQuickTextureFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickTextureFactory) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::deleteLater", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::deleteLater")
	}
}

func (ptr *QQuickTextureFactory) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickTextureFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickTextureFactory_DisconnectNotify
func callbackQQuickTextureFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextureFactory) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::disconnectNotify", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::disconnectNotify")
	}
}

func (ptr *QQuickTextureFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextureFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextureFactory_Event
func callbackQQuickTextureFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextureFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickTextureFactory) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::event", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::event")
	}
}

func (ptr *QQuickTextureFactory) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickTextureFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickTextureFactory_EventFilter
func callbackQQuickTextureFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickTextureFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickTextureFactory) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::eventFilter", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::eventFilter")
	}
}

func (ptr *QQuickTextureFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickTextureFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickTextureFactory_MetaObject
func callbackQQuickTextureFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickTextureFactory::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickTextureFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickTextureFactory) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::metaObject", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickTextureFactory::metaObject")
	}
}

func (ptr *QQuickTextureFactory) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextureFactory_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextureFactory) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextureFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickView struct {
	QQuickWindow
}

type QQuickView_ITF interface {
	QQuickWindow_ITF
	QQuickView_PTR() *QQuickView
}

func (p *QQuickView) QQuickView_PTR() *QQuickView {
	return p
}

func (p *QQuickView) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func (p *QQuickView) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickWindow_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickView(ptr QQuickView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickView_PTR().Pointer()
	}
	return nil
}

func NewQQuickViewFromPointer(ptr unsafe.Pointer) *QQuickView {
	var n = new(QQuickView)
	n.SetPointer(ptr)
	return n
}
func (ptr *QQuickView) ResizeMode() QQuickView__ResizeMode {
	if ptr.Pointer() != nil {
		return QQuickView__ResizeMode(C.QQuickView_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) SetResizeMode(vre QQuickView__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetResizeMode(ptr.Pointer(), C.longlong(vre))
	}
}

func (ptr *QQuickView) Status() QQuickView__Status {
	if ptr.Pointer() != nil {
		return QQuickView__Status(C.QQuickView_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQuickView2(engine qml.QQmlEngine_ITF, parent gui.QWindow_ITF) *QQuickView {
	var tmpValue = NewQQuickViewFromPointer(C.QQuickView_NewQQuickView2(qml.PointerFromQQmlEngine(engine), gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickView(parent gui.QWindow_ITF) *QQuickView {
	var tmpValue = NewQQuickViewFromPointer(C.QQuickView_NewQQuickView(gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickView3(source core.QUrl_ITF, parent gui.QWindow_ITF) *QQuickView {
	var tmpValue = NewQQuickViewFromPointer(C.QQuickView_NewQQuickView3(core.PointerFromQUrl(source), gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickView) Engine() *qml.QQmlEngine {
	if ptr.Pointer() != nil {
		var tmpValue = qml.NewQQmlEngineFromPointer(C.QQuickView_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) Errors() []*qml.QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQuick_PackedList) []*qml.QQmlError {
			var out = make([]*qml.QQmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQQuickViewFromPointer(l.data).errors_atList(i)
			}
			return out
		}(C.QQuickView_Errors(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) InitialSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickView_InitialSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) KeyPressEvent(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MousePressEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) RootContext() *qml.QQmlContext {
	if ptr.Pointer() != nil {
		var tmpValue = qml.NewQQmlContextFromPointer(C.QQuickView_RootContext(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) RootObject() *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickView_RootObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickView_SetSource
func callbackQQuickView_SetSource(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQuickView) ConnectSetSource(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setSource", f)
	}
}

func (ptr *QQuickView) DisconnectSetSource(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setSource")
	}
}

func (ptr *QQuickView) SetSource(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickView) Source() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQuickView_Source(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQuickView_StatusChanged
func callbackQQuickView_StatusChanged(ptr unsafe.Pointer, status C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::statusChanged"); signal != nil {
		signal.(func(QQuickView__Status))(QQuickView__Status(status))
	}

}

func (ptr *QQuickView) ConnectStatusChanged(f func(status QQuickView__Status)) {
	if ptr.Pointer() != nil {
		C.QQuickView_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::statusChanged", f)
	}
}

func (ptr *QQuickView) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::statusChanged")
	}
}

func (ptr *QQuickView) StatusChanged(status QQuickView__Status) {
	if ptr.Pointer() != nil {
		C.QQuickView_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

//export callbackQQuickView_DestroyQQuickView
func callbackQQuickView_DestroyQQuickView(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::~QQuickView"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).DestroyQQuickViewDefault()
	}
}

func (ptr *QQuickView) ConnectDestroyQQuickView(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::~QQuickView", f)
	}
}

func (ptr *QQuickView) DisconnectDestroyQQuickView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::~QQuickView")
	}
}

func (ptr *QQuickView) DestroyQQuickView() {
	if ptr.Pointer() != nil {
		C.QQuickView_DestroyQQuickView(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickView) DestroyQQuickViewDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_DestroyQQuickViewDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickView) errors_atList(i int) *qml.QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = qml.NewQQmlErrorFromPointer(C.QQuickView_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*qml.QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

//export callbackQQuickView_Event
func callbackQQuickView_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickView) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::event", f)
	}
}

func (ptr *QQuickView) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::event")
	}
}

func (ptr *QQuickView) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickView) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickView_ExposeEvent
func callbackQQuickView_ExposeEvent(ptr unsafe.Pointer, vqe unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(vqe))
	} else {
		NewQQuickViewFromPointer(ptr).ExposeEventDefault(gui.NewQExposeEventFromPointer(vqe))
	}
}

func (ptr *QQuickView) ConnectExposeEvent(f func(vqe *gui.QExposeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::exposeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectExposeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::exposeEvent")
	}
}

func (ptr *QQuickView) ExposeEvent(vqe gui.QExposeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ExposeEvent(ptr.Pointer(), gui.PointerFromQExposeEvent(vqe))
	}
}

func (ptr *QQuickView) ExposeEventDefault(vqe gui.QExposeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ExposeEventDefault(ptr.Pointer(), gui.PointerFromQExposeEvent(vqe))
	}
}

//export callbackQQuickView_FocusInEvent
func callbackQQuickView_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::focusInEvent", f)
	}
}

func (ptr *QQuickView) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::focusInEvent")
	}
}

func (ptr *QQuickView) FocusInEvent(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQQuickView_FocusOutEvent
func callbackQQuickView_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::focusOutEvent", f)
	}
}

func (ptr *QQuickView) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::focusOutEvent")
	}
}

func (ptr *QQuickView) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQQuickView_HideEvent
func callbackQQuickView_HideEvent(ptr unsafe.Pointer, vqh unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(vqh))
	} else {
		NewQQuickViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(vqh))
	}
}

func (ptr *QQuickView) ConnectHideEvent(f func(vqh *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::hideEvent", f)
	}
}

func (ptr *QQuickView) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::hideEvent")
	}
}

func (ptr *QQuickView) HideEvent(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QQuickView) HideEventDefault(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

//export callbackQQuickView_MouseDoubleClickEvent
func callbackQQuickView_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::mouseDoubleClickEvent")
	}
}

func (ptr *QQuickView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickView_ReleaseResources
func callbackQQuickView_ReleaseResources(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickView) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::releaseResources", f)
	}
}

func (ptr *QQuickView) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::releaseResources")
	}
}

func (ptr *QQuickView) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickView_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickView) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_ReleaseResourcesDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ResizeEvent
func callbackQQuickView_ResizeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ConnectResizeEvent(f func(ev *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::resizeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::resizeEvent")
	}
}

func (ptr *QQuickView) ResizeEvent(ev gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

func (ptr *QQuickView) ResizeEventDefault(ev gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

//export callbackQQuickView_ShowEvent
func callbackQQuickView_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(vqs))
	} else {
		NewQQuickViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *QQuickView) ConnectShowEvent(f func(vqs *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showEvent", f)
	}
}

func (ptr *QQuickView) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showEvent")
	}
}

func (ptr *QQuickView) ShowEvent(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

func (ptr *QQuickView) ShowEventDefault(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackQQuickView_Update
func callbackQQuickView_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickView) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::update", f)
	}
}

func (ptr *QQuickView) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::update")
	}
}

func (ptr *QQuickView) Update() {
	if ptr.Pointer() != nil {
		C.QQuickView_Update(ptr.Pointer())
	}
}

func (ptr *QQuickView) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_WheelEvent
func callbackQQuickView_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickView) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::wheelEvent", f)
	}
}

func (ptr *QQuickView) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::wheelEvent")
	}
}

func (ptr *QQuickView) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQQuickView_SetHeight
func callbackQQuickView_SetHeight(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setHeight"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickViewFromPointer(ptr).SetHeightDefault(int(int32(arg)))
	}
}

func (ptr *QQuickView) ConnectSetHeight(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setHeight", f)
	}
}

func (ptr *QQuickView) DisconnectSetHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setHeight")
	}
}

func (ptr *QQuickView) SetHeight(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetHeight(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickView) SetHeightDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetHeightDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickView_SetMaximumHeight
func callbackQQuickView_SetMaximumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setMaximumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewQQuickViewFromPointer(ptr).SetMaximumHeightDefault(int(int32(h)))
	}
}

func (ptr *QQuickView) ConnectSetMaximumHeight(f func(h int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMaximumHeight", f)
	}
}

func (ptr *QQuickView) DisconnectSetMaximumHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMaximumHeight")
	}
}

func (ptr *QQuickView) SetMaximumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumHeight(ptr.Pointer(), C.int(int32(h)))
	}
}

func (ptr *QQuickView) SetMaximumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackQQuickView_SetMaximumWidth
func callbackQQuickView_SetMaximumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setMaximumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewQQuickViewFromPointer(ptr).SetMaximumWidthDefault(int(int32(w)))
	}
}

func (ptr *QQuickView) ConnectSetMaximumWidth(f func(w int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMaximumWidth", f)
	}
}

func (ptr *QQuickView) DisconnectSetMaximumWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMaximumWidth")
	}
}

func (ptr *QQuickView) SetMaximumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumWidth(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QQuickView) SetMaximumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackQQuickView_SetMinimumHeight
func callbackQQuickView_SetMinimumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setMinimumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewQQuickViewFromPointer(ptr).SetMinimumHeightDefault(int(int32(h)))
	}
}

func (ptr *QQuickView) ConnectSetMinimumHeight(f func(h int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMinimumHeight", f)
	}
}

func (ptr *QQuickView) DisconnectSetMinimumHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMinimumHeight")
	}
}

func (ptr *QQuickView) SetMinimumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumHeight(ptr.Pointer(), C.int(int32(h)))
	}
}

func (ptr *QQuickView) SetMinimumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackQQuickView_SetMinimumWidth
func callbackQQuickView_SetMinimumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setMinimumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewQQuickViewFromPointer(ptr).SetMinimumWidthDefault(int(int32(w)))
	}
}

func (ptr *QQuickView) ConnectSetMinimumWidth(f func(w int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMinimumWidth", f)
	}
}

func (ptr *QQuickView) DisconnectSetMinimumWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setMinimumWidth")
	}
}

func (ptr *QQuickView) SetMinimumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumWidth(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QQuickView) SetMinimumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackQQuickView_SetTitle
func callbackQQuickView_SetTitle(ptr unsafe.Pointer, vqs C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQQuickViewFromPointer(ptr).SetTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QQuickView) ConnectSetTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setTitle", f)
	}
}

func (ptr *QQuickView) DisconnectSetTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setTitle")
	}
}

func (ptr *QQuickView) SetTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QQuickView_SetTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QQuickView) SetTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QQuickView_SetTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQQuickView_SetVisible
func callbackQQuickView_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQQuickViewFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QQuickView) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setVisible", f)
	}
}

func (ptr *QQuickView) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setVisible")
	}
}

func (ptr *QQuickView) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QQuickView) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQQuickView_SetWidth
func callbackQQuickView_SetWidth(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setWidth"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickViewFromPointer(ptr).SetWidthDefault(int(int32(arg)))
	}
}

func (ptr *QQuickView) ConnectSetWidth(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setWidth", f)
	}
}

func (ptr *QQuickView) DisconnectSetWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setWidth")
	}
}

func (ptr *QQuickView) SetWidth(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetWidth(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickView) SetWidthDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetWidthDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickView_SetX
func callbackQQuickView_SetX(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setX"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickViewFromPointer(ptr).SetXDefault(int(int32(arg)))
	}
}

func (ptr *QQuickView) ConnectSetX(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setX", f)
	}
}

func (ptr *QQuickView) DisconnectSetX() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setX")
	}
}

func (ptr *QQuickView) SetX(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetX(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickView) SetXDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetXDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickView_SetY
func callbackQQuickView_SetY(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::setY"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickViewFromPointer(ptr).SetYDefault(int(int32(arg)))
	}
}

func (ptr *QQuickView) ConnectSetY(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setY", f)
	}
}

func (ptr *QQuickView) DisconnectSetY() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::setY")
	}
}

func (ptr *QQuickView) SetY(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetY(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickView) SetYDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetYDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickView_Alert
func callbackQQuickView_Alert(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::alert"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQQuickViewFromPointer(ptr).AlertDefault(int(int32(msec)))
	}
}

func (ptr *QQuickView) ConnectAlert(f func(msec int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::alert", f)
	}
}

func (ptr *QQuickView) DisconnectAlert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::alert")
	}
}

func (ptr *QQuickView) Alert(msec int) {
	if ptr.Pointer() != nil {
		C.QQuickView_Alert(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QQuickView) AlertDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QQuickView_AlertDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

//export callbackQQuickView_Close
func callbackQQuickView_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickViewFromPointer(ptr).CloseDefault())))
}

func (ptr *QQuickView) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::close", f)
	}
}

func (ptr *QQuickView) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::close")
	}
}

func (ptr *QQuickView) Close() bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickView) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickView_FocusObject
func callbackQQuickView_FocusObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::focusObject"); signal != nil {
		return core.PointerFromQObject(signal.(func() *core.QObject)())
	}

	return core.PointerFromQObject(NewQQuickViewFromPointer(ptr).FocusObjectDefault())
}

func (ptr *QQuickView) ConnectFocusObject(f func() *core.QObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::focusObject", f)
	}
}

func (ptr *QQuickView) DisconnectFocusObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::focusObject")
	}
}

func (ptr *QQuickView) FocusObject() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQuickView_FocusObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) FocusObjectDefault() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQuickView_FocusObjectDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickView_Format
func callbackQQuickView_Format(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::format"); signal != nil {
		return gui.PointerFromQSurfaceFormat(signal.(func() *gui.QSurfaceFormat)())
	}

	return gui.PointerFromQSurfaceFormat(NewQQuickViewFromPointer(ptr).FormatDefault())
}

func (ptr *QQuickView) ConnectFormat(f func() *gui.QSurfaceFormat) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::format", f)
	}
}

func (ptr *QQuickView) DisconnectFormat() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::format")
	}
}

func (ptr *QQuickView) Format() *gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQSurfaceFormatFromPointer(C.QQuickView_Format(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) FormatDefault() *gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQSurfaceFormatFromPointer(C.QQuickView_FormatDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

//export callbackQQuickView_Hide
func callbackQQuickView_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickView) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::hide", f)
	}
}

func (ptr *QQuickView) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::hide")
	}
}

func (ptr *QQuickView) Hide() {
	if ptr.Pointer() != nil {
		C.QQuickView_Hide(ptr.Pointer())
	}
}

func (ptr *QQuickView) HideDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_Lower
func callbackQQuickView_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickView) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::lower", f)
	}
}

func (ptr *QQuickView) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::lower")
	}
}

func (ptr *QQuickView) Lower() {
	if ptr.Pointer() != nil {
		C.QQuickView_Lower(ptr.Pointer())
	}
}

func (ptr *QQuickView) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_MoveEvent
func callbackQQuickView_MoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ConnectMoveEvent(f func(ev *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::moveEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::moveEvent")
	}
}

func (ptr *QQuickView) MoveEvent(ev gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

func (ptr *QQuickView) MoveEventDefault(ev gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

//export callbackQQuickView_NativeEvent
func callbackQQuickView_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickViewFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QQuickView) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::nativeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::nativeEvent")
	}
}

func (ptr *QQuickView) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QQuickView) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQQuickView_Raise
func callbackQQuickView_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickView) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::raise", f)
	}
}

func (ptr *QQuickView) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::raise")
	}
}

func (ptr *QQuickView) Raise() {
	if ptr.Pointer() != nil {
		C.QQuickView_Raise(ptr.Pointer())
	}
}

func (ptr *QQuickView) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_RequestActivate
func callbackQQuickView_RequestActivate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::requestActivate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).RequestActivateDefault()
	}
}

func (ptr *QQuickView) ConnectRequestActivate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::requestActivate", f)
	}
}

func (ptr *QQuickView) DisconnectRequestActivate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::requestActivate")
	}
}

func (ptr *QQuickView) RequestActivate() {
	if ptr.Pointer() != nil {
		C.QQuickView_RequestActivate(ptr.Pointer())
	}
}

func (ptr *QQuickView) RequestActivateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_RequestActivateDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_RequestUpdate
func callbackQQuickView_RequestUpdate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::requestUpdate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).RequestUpdateDefault()
	}
}

func (ptr *QQuickView) ConnectRequestUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::requestUpdate", f)
	}
}

func (ptr *QQuickView) DisconnectRequestUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::requestUpdate")
	}
}

func (ptr *QQuickView) RequestUpdate() {
	if ptr.Pointer() != nil {
		C.QQuickView_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QQuickView) RequestUpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_RequestUpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_Show
func callbackQQuickView_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickView) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::show", f)
	}
}

func (ptr *QQuickView) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::show")
	}
}

func (ptr *QQuickView) Show() {
	if ptr.Pointer() != nil {
		C.QQuickView_Show(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowFullScreen
func callbackQQuickView_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickView) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showFullScreen", f)
	}
}

func (ptr *QQuickView) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showFullScreen")
	}
}

func (ptr *QQuickView) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowMaximized
func callbackQQuickView_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickView) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showMaximized", f)
	}
}

func (ptr *QQuickView) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showMaximized")
	}
}

func (ptr *QQuickView) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowMinimized
func callbackQQuickView_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickView) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showMinimized", f)
	}
}

func (ptr *QQuickView) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showMinimized")
	}
}

func (ptr *QQuickView) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowNormal
func callbackQQuickView_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickView) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showNormal", f)
	}
}

func (ptr *QQuickView) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::showNormal")
	}
}

func (ptr *QQuickView) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_Size
func callbackQQuickView_Size(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::size"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickViewFromPointer(ptr).SizeDefault())
}

func (ptr *QQuickView) ConnectSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::size", f)
	}
}

func (ptr *QQuickView) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::size")
	}
}

func (ptr *QQuickView) Size() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickView_Size(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickView) SizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickView_SizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickView_SurfaceType
func callbackQQuickView_SurfaceType(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::surfaceType"); signal != nil {
		return C.longlong(signal.(func() gui.QSurface__SurfaceType)())
	}

	return C.longlong(NewQQuickViewFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *QQuickView) ConnectSurfaceType(f func() gui.QSurface__SurfaceType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::surfaceType", f)
	}
}

func (ptr *QQuickView) DisconnectSurfaceType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::surfaceType")
	}
}

func (ptr *QQuickView) SurfaceType() gui.QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickView_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) SurfaceTypeDefault() gui.QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickView_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickView_TabletEvent
func callbackQQuickView_TabletEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ConnectTabletEvent(f func(ev *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::tabletEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::tabletEvent")
	}
}

func (ptr *QQuickView) TabletEvent(ev gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

func (ptr *QQuickView) TabletEventDefault(ev gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

//export callbackQQuickView_TouchEvent
func callbackQQuickView_TouchEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ConnectTouchEvent(f func(ev *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::touchEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::touchEvent")
	}
}

func (ptr *QQuickView) TouchEvent(ev gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

func (ptr *QQuickView) TouchEventDefault(ev gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

//export callbackQQuickView_TimerEvent
func callbackQQuickView_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::timerEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::timerEvent")
	}
}

func (ptr *QQuickView) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickView) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickView_ChildEvent
func callbackQQuickView_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::childEvent", f)
	}
}

func (ptr *QQuickView) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::childEvent")
	}
}

func (ptr *QQuickView) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickView) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickView_ConnectNotify
func callbackQQuickView_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickView) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::connectNotify", f)
	}
}

func (ptr *QQuickView) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::connectNotify")
	}
}

func (ptr *QQuickView) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickView_CustomEvent
func callbackQQuickView_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickView) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::customEvent", f)
	}
}

func (ptr *QQuickView) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::customEvent")
	}
}

func (ptr *QQuickView) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickView) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickView_DeleteLater
func callbackQQuickView_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickView) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::deleteLater", f)
	}
}

func (ptr *QQuickView) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::deleteLater")
	}
}

func (ptr *QQuickView) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickView_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickView) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickView_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickView_DisconnectNotify
func callbackQQuickView_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickView) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::disconnectNotify", f)
	}
}

func (ptr *QQuickView) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::disconnectNotify")
	}
}

func (ptr *QQuickView) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickView_EventFilter
func callbackQQuickView_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickView) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::eventFilter", f)
	}
}

func (ptr *QQuickView) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::eventFilter")
	}
}

func (ptr *QQuickView) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickView_MetaObject
func callbackQQuickView_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickView::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickView) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::metaObject", f)
	}
}

func (ptr *QQuickView) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickView::metaObject")
	}
}

func (ptr *QQuickView) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickView_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidget_ITF interface {
	widgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func (p *QQuickWidget) QQuickWidget_PTR() *QQuickWidget {
	return p
}

func (p *QQuickWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QQuickWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickWidget(ptr QQuickWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWidget_PTR().Pointer()
	}
	return nil
}

func NewQQuickWidgetFromPointer(ptr unsafe.Pointer) *QQuickWidget {
	var n = new(QQuickWidget)
	n.SetPointer(ptr)
	return n
}
func (ptr *QQuickWidget) ResizeMode() QQuickWidget__ResizeMode {
	if ptr.Pointer() != nil {
		return QQuickWidget__ResizeMode(C.QQuickWidget_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWidget) SetResizeMode(vre QQuickWidget__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetResizeMode(ptr.Pointer(), C.longlong(vre))
	}
}

func (ptr *QQuickWidget) Status() QQuickWidget__Status {
	if ptr.Pointer() != nil {
		return QQuickWidget__Status(C.QQuickWidget_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQuickWidget2(engine qml.QQmlEngine_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	var tmpValue = NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget2(qml.PointerFromQQmlEngine(engine), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickWidget(parent widgets.QWidget_ITF) *QQuickWidget {
	var tmpValue = NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQuickWidget3(source core.QUrl_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	var tmpValue = NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget3(core.PointerFromQUrl(source), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQuickWidget_DragEnterEvent
func callbackQQuickWidget_DragEnterEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dragEnterEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dragEnterEvent")
	}
}

func (ptr *QQuickWidget) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QQuickWidget) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

//export callbackQQuickWidget_DragLeaveEvent
func callbackQQuickWidget_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dragLeaveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dragLeaveEvent")
	}
}

func (ptr *QQuickWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QQuickWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackQQuickWidget_DragMoveEvent
func callbackQQuickWidget_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dragMoveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dragMoveEvent")
	}
}

func (ptr *QQuickWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QQuickWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackQQuickWidget_DropEvent
func callbackQQuickWidget_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dropEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::dropEvent")
	}
}

func (ptr *QQuickWidget) DropEvent(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QQuickWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QQuickWidget) Engine() *qml.QQmlEngine {
	if ptr.Pointer() != nil {
		var tmpValue = qml.NewQQmlEngineFromPointer(C.QQuickWidget_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) Errors() []*qml.QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQuick_PackedList) []*qml.QQmlError {
			var out = make([]*qml.QQmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQQuickWidgetFromPointer(l.data).errors_atList(i)
			}
			return out
		}(C.QQuickWidget_Errors(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWidget_Event
func callbackQQuickWidget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickWidget) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::event", f)
	}
}

func (ptr *QQuickWidget) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::event")
	}
}

func (ptr *QQuickWidget) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWidget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickWidget_FocusInEvent
func callbackQQuickWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::focusInEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::focusInEvent")
	}
}

func (ptr *QQuickWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQQuickWidget_FocusOutEvent
func callbackQQuickWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::focusOutEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::focusOutEvent")
	}
}

func (ptr *QQuickWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) Format() *gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQSurfaceFormatFromPointer(C.QQuickWidget_Format(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) GrabFramebuffer() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QQuickWidget_GrabFramebuffer(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_HideEvent
func callbackQQuickWidget_HideEvent(ptr unsafe.Pointer, vqh unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(vqh))
	} else {
		NewQQuickWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(vqh))
	}
}

func (ptr *QQuickWidget) ConnectHideEvent(f func(vqh *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::hideEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::hideEvent")
	}
}

func (ptr *QQuickWidget) HideEvent(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QQuickWidget) HideEventDefault(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QQuickWidget) InitialSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWidget_InitialSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_KeyPressEvent
func callbackQQuickWidget_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::keyPressEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::keyPressEvent")
	}
}

func (ptr *QQuickWidget) KeyPressEvent(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWidget_KeyReleaseEvent
func callbackQQuickWidget_KeyReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::keyReleaseEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::keyReleaseEvent")
	}
}

func (ptr *QQuickWidget) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWidget_MouseDoubleClickEvent
func callbackQQuickWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mouseDoubleClickEvent")
	}
}

func (ptr *QQuickWidget) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQQuickWidget_MouseMoveEvent
func callbackQQuickWidget_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mouseMoveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mouseMoveEvent")
	}
}

func (ptr *QQuickWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQQuickWidget_MousePressEvent
func callbackQQuickWidget_MousePressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mousePressEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mousePressEvent")
	}
}

func (ptr *QQuickWidget) MousePressEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQQuickWidget_MouseReleaseEvent
func callbackQQuickWidget_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mouseReleaseEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::mouseReleaseEvent")
	}
}

func (ptr *QQuickWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) QuickWindow() *QQuickWindow {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickWindowFromPointer(C.QQuickWidget_QuickWindow(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) RootContext() *qml.QQmlContext {
	if ptr.Pointer() != nil {
		var tmpValue = qml.NewQQmlContextFromPointer(C.QQuickWidget_RootContext(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) RootObject() *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickWidget_RootObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_SceneGraphError
func callbackQQuickWidget_SceneGraphError(ptr unsafe.Pointer, error C.longlong, message C.struct_QtQuick_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), cGoUnpackString(message))
	}

}

func (ptr *QQuickWidget) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::sceneGraphError", f)
	}
}

func (ptr *QQuickWidget) DisconnectSceneGraphError() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::sceneGraphError")
	}
}

func (ptr *QQuickWidget) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	if ptr.Pointer() != nil {
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		C.QQuickWidget_SceneGraphError(ptr.Pointer(), C.longlong(error), messageC)
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

//export callbackQQuickWidget_SetSource
func callbackQQuickWidget_SetSource(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQuickWidget) ConnectSetSource(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setSource", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetSource(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setSource")
	}
}

func (ptr *QQuickWidget) SetSource(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQuickWidget_ShowEvent
func callbackQQuickWidget_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(vqs))
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *QQuickWidget) ConnectShowEvent(f func(vqs *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showEvent")
	}
}

func (ptr *QQuickWidget) ShowEvent(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

func (ptr *QQuickWidget) ShowEventDefault(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackQQuickWidget_StatusChanged
func callbackQQuickWidget_StatusChanged(ptr unsafe.Pointer, status C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::statusChanged"); signal != nil {
		signal.(func(QQuickWidget__Status))(QQuickWidget__Status(status))
	}

}

func (ptr *QQuickWidget) ConnectStatusChanged(f func(status QQuickWidget__Status)) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::statusChanged", f)
	}
}

func (ptr *QQuickWidget) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::statusChanged")
	}
}

func (ptr *QQuickWidget) StatusChanged(status QQuickWidget__Status) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

func (ptr *QQuickWidget) Source() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQuickWidget_Source(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_WheelEvent
func callbackQQuickWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::wheelEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::wheelEvent")
	}
}

func (ptr *QQuickWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QQuickWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQQuickWidget_DestroyQQuickWidget
func callbackQQuickWidget_DestroyQQuickWidget(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::~QQuickWidget"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).DestroyQQuickWidgetDefault()
	}
}

func (ptr *QQuickWidget) ConnectDestroyQQuickWidget(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::~QQuickWidget", f)
	}
}

func (ptr *QQuickWidget) DisconnectDestroyQQuickWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::~QQuickWidget")
	}
}

func (ptr *QQuickWidget) DestroyQQuickWidget() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DestroyQQuickWidget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWidget) DestroyQQuickWidgetDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DestroyQQuickWidgetDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWidget) errors_atList(i int) *qml.QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = qml.NewQQmlErrorFromPointer(C.QQuickWidget_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*qml.QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_ActionEvent
func callbackQQuickWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::actionEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectActionEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::actionEvent")
	}
}

func (ptr *QQuickWidget) ActionEvent(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QQuickWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQQuickWidget_EnterEvent
func callbackQQuickWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::enterEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::enterEvent")
	}
}

func (ptr *QQuickWidget) EnterEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_LeaveEvent
func callbackQQuickWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::leaveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::leaveEvent")
	}
}

func (ptr *QQuickWidget) LeaveEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_Metric
func callbackQQuickWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQQuickWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QQuickWidget) ConnectMetric(f func(m gui.QPaintDevice__PaintDeviceMetric) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::metric", f)
	}
}

func (ptr *QQuickWidget) DisconnectMetric() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::metric")
	}
}

func (ptr *QQuickWidget) Metric(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWidget_Metric(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

func (ptr *QQuickWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQQuickWidget_MinimumSizeHint
func callbackQQuickWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QQuickWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::minimumSizeHint", f)
	}
}

func (ptr *QQuickWidget) DisconnectMinimumSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::minimumSizeHint")
	}
}

func (ptr *QQuickWidget) MinimumSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_MoveEvent
func callbackQQuickWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::moveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::moveEvent")
	}
}

func (ptr *QQuickWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QQuickWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQQuickWidget_PaintEngine
func callbackQQuickWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQQuickWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QQuickWidget) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::paintEngine", f)
	}
}

func (ptr *QQuickWidget) DisconnectPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::paintEngine")
	}
}

func (ptr *QQuickWidget) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QQuickWidget_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QQuickWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWidget_PaintEvent
func callbackQQuickWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::paintEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::paintEvent")
	}
}

func (ptr *QQuickWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QQuickWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQQuickWidget_SetEnabled
func callbackQQuickWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setEnabled", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setEnabled")
	}
}

func (ptr *QQuickWidget) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQQuickWidget_SetStyleSheet
func callbackQQuickWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQQuickWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QQuickWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setStyleSheet", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setStyleSheet")
	}
}

func (ptr *QQuickWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QQuickWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QQuickWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QQuickWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQQuickWidget_SetVisible
func callbackQQuickWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setVisible", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setVisible")
	}
}

func (ptr *QQuickWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QQuickWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQQuickWidget_SetWindowModified
func callbackQQuickWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetWindowModified(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setWindowModified", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetWindowModified() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setWindowModified")
	}
}

func (ptr *QQuickWidget) SetWindowModified(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QQuickWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQQuickWidget_SetWindowTitle
func callbackQQuickWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQQuickWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QQuickWidget) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setWindowTitle", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setWindowTitle")
	}
}

func (ptr *QQuickWidget) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QQuickWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QQuickWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QQuickWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQQuickWidget_SizeHint
func callbackQQuickWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QQuickWidget) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::sizeHint", f)
	}
}

func (ptr *QQuickWidget) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::sizeHint")
	}
}

func (ptr *QQuickWidget) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_ChangeEvent
func callbackQQuickWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::changeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectChangeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::changeEvent")
	}
}

func (ptr *QQuickWidget) ChangeEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_Close
func callbackQQuickWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QQuickWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::close", f)
	}
}

func (ptr *QQuickWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::close")
	}
}

func (ptr *QQuickWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWidget_CloseEvent
func callbackQQuickWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::closeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::closeEvent")
	}
}

func (ptr *QQuickWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QQuickWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQQuickWidget_ContextMenuEvent
func callbackQQuickWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::contextMenuEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::contextMenuEvent")
	}
}

func (ptr *QQuickWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QQuickWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQQuickWidget_FocusNextPrevChild
func callbackQQuickWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QQuickWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::focusNextPrevChild", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusNextPrevChild() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::focusNextPrevChild")
	}
}

func (ptr *QQuickWidget) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QQuickWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQQuickWidget_HasHeightForWidth
func callbackQQuickWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QQuickWidget) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::hasHeightForWidth", f)
	}
}

func (ptr *QQuickWidget) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::hasHeightForWidth")
	}
}

func (ptr *QQuickWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWidget_HeightForWidth
func callbackQQuickWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQQuickWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QQuickWidget) ConnectHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::heightForWidth", f)
	}
}

func (ptr *QQuickWidget) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::heightForWidth")
	}
}

func (ptr *QQuickWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QQuickWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQuickWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQQuickWidget_Hide
func callbackQQuickWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickWidget) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::hide", f)
	}
}

func (ptr *QQuickWidget) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::hide")
	}
}

func (ptr *QQuickWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_InputMethodEvent
func callbackQQuickWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::inputMethodEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::inputMethodEvent")
	}
}

func (ptr *QQuickWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQQuickWidget_InputMethodQuery
func callbackQQuickWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::inputMethodQuery", f)
	}
}

func (ptr *QQuickWidget) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::inputMethodQuery")
	}
}

func (ptr *QQuickWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQuickWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWidget_Lower
func callbackQQuickWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickWidget) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::lower", f)
	}
}

func (ptr *QQuickWidget) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::lower")
	}
}

func (ptr *QQuickWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_NativeEvent
func callbackQQuickWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QQuickWidget) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::nativeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::nativeEvent")
	}
}

func (ptr *QQuickWidget) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QQuickWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQQuickWidget_Raise
func callbackQQuickWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickWidget) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::raise", f)
	}
}

func (ptr *QQuickWidget) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::raise")
	}
}

func (ptr *QQuickWidget) Raise() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_Repaint
func callbackQQuickWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QQuickWidget) ConnectRepaint(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::repaint", f)
	}
}

func (ptr *QQuickWidget) DisconnectRepaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::repaint")
	}
}

func (ptr *QQuickWidget) Repaint() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ResizeEvent
func callbackQQuickWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::resizeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::resizeEvent")
	}
}

func (ptr *QQuickWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QQuickWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQQuickWidget_SetDisabled
func callbackQQuickWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetDisabled(f func(disable bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setDisabled", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setDisabled")
	}
}

func (ptr *QQuickWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QQuickWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQQuickWidget_SetFocus2
func callbackQQuickWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QQuickWidget) ConnectSetFocus2(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setFocus2", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetFocus2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setFocus2")
	}
}

func (ptr *QQuickWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQQuickWidget_SetHidden
func callbackQQuickWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetHidden(f func(hidden bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setHidden", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::setHidden")
	}
}

func (ptr *QQuickWidget) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QQuickWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQQuickWidget_Show
func callbackQQuickWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickWidget) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::show", f)
	}
}

func (ptr *QQuickWidget) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::show")
	}
}

func (ptr *QQuickWidget) Show() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_Show(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowFullScreen
func callbackQQuickWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showFullScreen", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showFullScreen")
	}
}

func (ptr *QQuickWidget) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowMaximized
func callbackQQuickWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showMaximized", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showMaximized")
	}
}

func (ptr *QQuickWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowMinimized
func callbackQQuickWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showMinimized", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showMinimized")
	}
}

func (ptr *QQuickWidget) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowNormal
func callbackQQuickWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showNormal", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::showNormal")
	}
}

func (ptr *QQuickWidget) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_TabletEvent
func callbackQQuickWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::tabletEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::tabletEvent")
	}
}

func (ptr *QQuickWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QQuickWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQQuickWidget_Update
func callbackQQuickWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickWidget) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::update", f)
	}
}

func (ptr *QQuickWidget) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::update")
	}
}

func (ptr *QQuickWidget) Update() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_UpdateMicroFocus
func callbackQQuickWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QQuickWidget) ConnectUpdateMicroFocus(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::updateMicroFocus", f)
	}
}

func (ptr *QQuickWidget) DisconnectUpdateMicroFocus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::updateMicroFocus")
	}
}

func (ptr *QQuickWidget) UpdateMicroFocus() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_TimerEvent
func callbackQQuickWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::timerEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::timerEvent")
	}
}

func (ptr *QQuickWidget) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickWidget_ChildEvent
func callbackQQuickWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::childEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::childEvent")
	}
}

func (ptr *QQuickWidget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWidget_ConnectNotify
func callbackQQuickWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::connectNotify", f)
	}
}

func (ptr *QQuickWidget) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::connectNotify")
	}
}

func (ptr *QQuickWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWidget_CustomEvent
func callbackQQuickWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::customEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::customEvent")
	}
}

func (ptr *QQuickWidget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWidget_DeleteLater
func callbackQQuickWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWidget) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::deleteLater", f)
	}
}

func (ptr *QQuickWidget) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::deleteLater")
	}
}

func (ptr *QQuickWidget) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWidget_DisconnectNotify
func callbackQQuickWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::disconnectNotify", f)
	}
}

func (ptr *QQuickWidget) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::disconnectNotify")
	}
}

func (ptr *QQuickWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWidget_EventFilter
func callbackQQuickWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::eventFilter", f)
	}
}

func (ptr *QQuickWidget) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::eventFilter")
	}
}

func (ptr *QQuickWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickWidget_MetaObject
func callbackQQuickWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWidget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::metaObject", f)
	}
}

func (ptr *QQuickWidget) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWidget::metaObject")
	}
}

func (ptr *QQuickWidget) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickWindow struct {
	gui.QWindow
}

type QQuickWindow_ITF interface {
	gui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func (p *QQuickWindow) QQuickWindow_PTR() *QQuickWindow {
	return p
}

func (p *QQuickWindow) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWindow_PTR().Pointer()
	}
	return nil
}

func (p *QQuickWindow) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWindow_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickWindow(ptr QQuickWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func NewQQuickWindowFromPointer(ptr unsafe.Pointer) *QQuickWindow {
	var n = new(QQuickWindow)
	n.SetPointer(ptr)
	return n
}
func (ptr *QQuickWindow) ActiveFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickWindow_ActiveFocusItem(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QQuickWindow_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickWindow_ContentItem(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func NewQQuickWindow(parent gui.QWindow_ITF) *QQuickWindow {
	var tmpValue = NewQQuickWindowFromPointer(C.QQuickWindow_NewQQuickWindow(gui.PointerFromQWindow(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QQuickWindow_AccessibleRoot(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWindow_ActiveFocusItemChanged
func callbackQQuickWindow_ActiveFocusItemChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::activeFocusItemChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectActiveFocusItemChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::activeFocusItemChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectActiveFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::activeFocusItemChanged")
	}
}

func (ptr *QQuickWindow) ActiveFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ActiveFocusItemChanged(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterAnimating
func callbackQQuickWindow_AfterAnimating(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::afterAnimating"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterAnimating(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::afterAnimating", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterAnimating(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::afterAnimating")
	}
}

func (ptr *QQuickWindow) AfterAnimating() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterAnimating(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterRendering
func callbackQQuickWindow_AfterRendering(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::afterRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterRendering(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::afterRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterRendering(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::afterRendering")
	}
}

func (ptr *QQuickWindow) AfterRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterRendering(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterSynchronizing
func callbackQQuickWindow_AfterSynchronizing(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::afterSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterSynchronizing(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::afterSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::afterSynchronizing")
	}
}

func (ptr *QQuickWindow) AfterSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterSynchronizing(ptr.Pointer())
	}
}

//export callbackQQuickWindow_BeforeRendering
func callbackQQuickWindow_BeforeRendering(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::beforeRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeRendering(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::beforeRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeRendering(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::beforeRendering")
	}
}

func (ptr *QQuickWindow) BeforeRendering() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_BeforeRendering(ptr.Pointer())
	}
}

//export callbackQQuickWindow_BeforeSynchronizing
func callbackQQuickWindow_BeforeSynchronizing(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::beforeSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeSynchronizing(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::beforeSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::beforeSynchronizing")
	}
}

func (ptr *QQuickWindow) BeforeSynchronizing() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_BeforeSynchronizing(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_ClearBeforeRendering(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWindow_ColorChanged
func callbackQQuickWindow_ColorChanged(ptr unsafe.Pointer, vqc unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(vqc))
	}

}

func (ptr *QQuickWindow) ConnectColorChanged(f func(vqc *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::colorChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::colorChanged")
	}
}

func (ptr *QQuickWindow) ColorChanged(vqc gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(vqc))
	}
}

func (ptr *QQuickWindow) CreateTextureFromId(id uint, size core.QSize_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromId(ptr.Pointer(), C.uint(uint32(id)), core.PointerFromQSize(size), C.longlong(options)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImage_ITF) *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage2(ptr.Pointer(), gui.PointerFromQImage(image)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImage_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.longlong(options)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickWindow_EffectiveDevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWindow_Event
func callbackQQuickWindow_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQuickWindow) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::event", f)
	}
}

func (ptr *QQuickWindow) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::event")
	}
}

func (ptr *QQuickWindow) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWindow) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickWindow_ExposeEvent
func callbackQQuickWindow_ExposeEvent(ptr unsafe.Pointer, vqe unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(vqe))
	} else {
		NewQQuickWindowFromPointer(ptr).ExposeEventDefault(gui.NewQExposeEventFromPointer(vqe))
	}
}

func (ptr *QQuickWindow) ConnectExposeEvent(f func(vqe *gui.QExposeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::exposeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectExposeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::exposeEvent")
	}
}

func (ptr *QQuickWindow) ExposeEvent(vqe gui.QExposeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ExposeEvent(ptr.Pointer(), gui.PointerFromQExposeEvent(vqe))
	}
}

func (ptr *QQuickWindow) ExposeEventDefault(vqe gui.QExposeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ExposeEventDefault(ptr.Pointer(), gui.PointerFromQExposeEvent(vqe))
	}
}

//export callbackQQuickWindow_FocusInEvent
func callbackQQuickWindow_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::focusInEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::focusInEvent")
	}
}

func (ptr *QQuickWindow) FocusInEvent(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickWindow) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQQuickWindow_FocusOutEvent
func callbackQQuickWindow_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::focusOutEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::focusOutEvent")
	}
}

func (ptr *QQuickWindow) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickWindow) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQQuickWindow_FrameSwapped
func callbackQQuickWindow_FrameSwapped(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::frameSwapped"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectFrameSwapped(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::frameSwapped", f)
	}
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectFrameSwapped(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::frameSwapped")
	}
}

func (ptr *QQuickWindow) FrameSwapped() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_FrameSwapped(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) GrabWindow() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QQuickWindow_GrabWindow(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func QQuickWindow_HasDefaultAlphaBuffer() bool {
	return C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer() != 0
}

func (ptr *QQuickWindow) HasDefaultAlphaBuffer() bool {
	return C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer() != 0
}

//export callbackQQuickWindow_HideEvent
func callbackQQuickWindow_HideEvent(ptr unsafe.Pointer, vqh unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(vqh))
	} else {
		NewQQuickWindowFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(vqh))
	}
}

func (ptr *QQuickWindow) ConnectHideEvent(f func(vqh *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::hideEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::hideEvent")
	}
}

func (ptr *QQuickWindow) HideEvent(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QQuickWindow) HideEventDefault(vqh gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QQuickWindow) IncubationController() *qml.QQmlIncubationController {
	if ptr.Pointer() != nil {
		return qml.NewQQmlIncubationControllerFromPointer(C.QQuickWindow_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentOpenGLContext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentSceneGraph(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsSceneGraphInitialized(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWindow_KeyPressEvent
func callbackQQuickWindow_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWindowFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWindow) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::keyPressEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::keyPressEvent")
	}
}

func (ptr *QQuickWindow) KeyPressEvent(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWindow_KeyReleaseEvent
func callbackQQuickWindow_KeyReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWindowFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWindow) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::keyReleaseEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::keyReleaseEvent")
	}
}

func (ptr *QQuickWindow) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

//export callbackQQuickWindow_MouseDoubleClickEvent
func callbackQQuickWindow_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mouseDoubleClickEvent")
	}
}

func (ptr *QQuickWindow) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQuickItemFromPointer(C.QQuickWindow_MouseGrabberItem(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_MouseMoveEvent
func callbackQQuickWindow_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mouseMoveEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mouseMoveEvent")
	}
}

func (ptr *QQuickWindow) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickWindow_MousePressEvent
func callbackQQuickWindow_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mousePressEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mousePressEvent")
	}
}

func (ptr *QQuickWindow) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQQuickWindow_MouseReleaseEvent
func callbackQQuickWindow_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mouseReleaseEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::mouseReleaseEvent")
	}
}

func (ptr *QQuickWindow) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQOpenGLContextFromPointer(C.QQuickWindow_OpenglContext(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_OpenglContextCreated
func callbackQQuickWindow_OpenglContextCreated(ptr unsafe.Pointer, context unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::openglContextCreated"); signal != nil {
		signal.(func(*gui.QOpenGLContext))(gui.NewQOpenGLContextFromPointer(context))
	}

}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context *gui.QOpenGLContext)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectOpenglContextCreated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::openglContextCreated", f)
	}
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectOpenglContextCreated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::openglContextCreated")
	}
}

func (ptr *QQuickWindow) OpenglContextCreated(context gui.QOpenGLContext_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_OpenglContextCreated(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

//export callbackQQuickWindow_ReleaseResources
func callbackQQuickWindow_ReleaseResources(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::releaseResources"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectReleaseResources(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::releaseResources", f)
	}
}

func (ptr *QQuickWindow) DisconnectReleaseResources() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::releaseResources")
	}
}

func (ptr *QQuickWindow) ReleaseResources() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {
	if ptr.Pointer() != nil {
		return gui.NewQOpenGLFramebufferObjectFromPointer(C.QQuickWindow_RenderTarget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) RenderTargetId() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QQuickWindow_RenderTargetId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickWindow) RenderTargetSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWindow_RenderTargetSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) ResetOpenGLState() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ResetOpenGLState(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ResizeEvent
func callbackQQuickWindow_ResizeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ConnectResizeEvent(f func(ev *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::resizeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::resizeEvent")
	}
}

func (ptr *QQuickWindow) ResizeEvent(ev gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

func (ptr *QQuickWindow) ResizeEventDefault(ev gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

//export callbackQQuickWindow_SceneGraphAboutToStop
func callbackQQuickWindow_SceneGraphAboutToStop(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::sceneGraphAboutToStop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphAboutToStop(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphAboutToStop", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphAboutToStop(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphAboutToStop")
	}
}

func (ptr *QQuickWindow) SceneGraphAboutToStop() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphAboutToStop(ptr.Pointer())
	}
}

//export callbackQQuickWindow_SceneGraphError
func callbackQQuickWindow_SceneGraphError(ptr unsafe.Pointer, error C.longlong, message C.struct_QtQuick_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), cGoUnpackString(message))
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphError", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphError")
	}
}

func (ptr *QQuickWindow) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	if ptr.Pointer() != nil {
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		C.QQuickWindow_SceneGraphError(ptr.Pointer(), C.longlong(error), messageC)
	}
}

//export callbackQQuickWindow_SceneGraphInitialized
func callbackQQuickWindow_SceneGraphInitialized(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::sceneGraphInitialized"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInitialized(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphInitialized", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInitialized(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphInitialized")
	}
}

func (ptr *QQuickWindow) SceneGraphInitialized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphInitialized(ptr.Pointer())
	}
}

//export callbackQQuickWindow_SceneGraphInvalidated
func callbackQQuickWindow_SceneGraphInvalidated(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::sceneGraphInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInvalidated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphInvalidated", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInvalidated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::sceneGraphInvalidated")
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

func (ptr *QQuickWindow) SendEvent(item QQuickItem_ITF, e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_SendEvent(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetClearBeforeRendering(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
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

//export callbackQQuickWindow_ShowEvent
func callbackQQuickWindow_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(vqs))
	} else {
		NewQQuickWindowFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *QQuickWindow) ConnectShowEvent(f func(vqs *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showEvent")
	}
}

func (ptr *QQuickWindow) ShowEvent(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

func (ptr *QQuickWindow) ShowEventDefault(vqs gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackQQuickWindow_Update
func callbackQQuickWindow_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::update"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::update", f)
	}
}

func (ptr *QQuickWindow) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::update")
	}
}

func (ptr *QQuickWindow) Update() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(ptr.Pointer())
	}
}

//export callbackQQuickWindow_WheelEvent
func callbackQQuickWindow_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::wheelEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::wheelEvent")
	}
}

func (ptr *QQuickWindow) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickWindow) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQQuickWindow_DestroyQQuickWindow
func callbackQQuickWindow_DestroyQQuickWindow(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::~QQuickWindow"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).DestroyQQuickWindowDefault()
	}
}

func (ptr *QQuickWindow) ConnectDestroyQQuickWindow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::~QQuickWindow", f)
	}
}

func (ptr *QQuickWindow) DisconnectDestroyQQuickWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::~QQuickWindow")
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindow(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindowDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindowDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWindow_SetHeight
func callbackQQuickWindow_SetHeight(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setHeight"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetHeightDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) ConnectSetHeight(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setHeight", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setHeight")
	}
}

func (ptr *QQuickWindow) SetHeight(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetHeight(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetHeightDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetHeightDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_SetMaximumHeight
func callbackQQuickWindow_SetMaximumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setMaximumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMaximumHeightDefault(int(int32(h)))
	}
}

func (ptr *QQuickWindow) ConnectSetMaximumHeight(f func(h int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMaximumHeight", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMaximumHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMaximumHeight")
	}
}

func (ptr *QQuickWindow) SetMaximumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumHeight(ptr.Pointer(), C.int(int32(h)))
	}
}

func (ptr *QQuickWindow) SetMaximumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackQQuickWindow_SetMaximumWidth
func callbackQQuickWindow_SetMaximumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setMaximumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMaximumWidthDefault(int(int32(w)))
	}
}

func (ptr *QQuickWindow) ConnectSetMaximumWidth(f func(w int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMaximumWidth", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMaximumWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMaximumWidth")
	}
}

func (ptr *QQuickWindow) SetMaximumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumWidth(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QQuickWindow) SetMaximumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackQQuickWindow_SetMinimumHeight
func callbackQQuickWindow_SetMinimumHeight(ptr unsafe.Pointer, h C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setMinimumHeight"); signal != nil {
		signal.(func(int))(int(int32(h)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMinimumHeightDefault(int(int32(h)))
	}
}

func (ptr *QQuickWindow) ConnectSetMinimumHeight(f func(h int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMinimumHeight", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMinimumHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMinimumHeight")
	}
}

func (ptr *QQuickWindow) SetMinimumHeight(h int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumHeight(ptr.Pointer(), C.int(int32(h)))
	}
}

func (ptr *QQuickWindow) SetMinimumHeightDefault(h int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumHeightDefault(ptr.Pointer(), C.int(int32(h)))
	}
}

//export callbackQQuickWindow_SetMinimumWidth
func callbackQQuickWindow_SetMinimumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setMinimumWidth"); signal != nil {
		signal.(func(int))(int(int32(w)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMinimumWidthDefault(int(int32(w)))
	}
}

func (ptr *QQuickWindow) ConnectSetMinimumWidth(f func(w int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMinimumWidth", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMinimumWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setMinimumWidth")
	}
}

func (ptr *QQuickWindow) SetMinimumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumWidth(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QQuickWindow) SetMinimumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

//export callbackQQuickWindow_SetTitle
func callbackQQuickWindow_SetTitle(ptr unsafe.Pointer, vqs C.struct_QtQuick_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQQuickWindowFromPointer(ptr).SetTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QQuickWindow) ConnectSetTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setTitle", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setTitle")
	}
}

func (ptr *QQuickWindow) SetTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QQuickWindow_SetTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QQuickWindow) SetTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QQuickWindow_SetTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQQuickWindow_SetVisible
func callbackQQuickWindow_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQQuickWindowFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QQuickWindow) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setVisible", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setVisible")
	}
}

func (ptr *QQuickWindow) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QQuickWindow) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQQuickWindow_SetWidth
func callbackQQuickWindow_SetWidth(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setWidth"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetWidthDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) ConnectSetWidth(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setWidth", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setWidth")
	}
}

func (ptr *QQuickWindow) SetWidth(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetWidth(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetWidthDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetWidthDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_SetX
func callbackQQuickWindow_SetX(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setX"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetXDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) ConnectSetX(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setX", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetX() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setX")
	}
}

func (ptr *QQuickWindow) SetX(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetX(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetXDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetXDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_SetY
func callbackQQuickWindow_SetY(ptr unsafe.Pointer, arg C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::setY"); signal != nil {
		signal.(func(int))(int(int32(arg)))
	} else {
		NewQQuickWindowFromPointer(ptr).SetYDefault(int(int32(arg)))
	}
}

func (ptr *QQuickWindow) ConnectSetY(f func(arg int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setY", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetY() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::setY")
	}
}

func (ptr *QQuickWindow) SetY(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetY(ptr.Pointer(), C.int(int32(arg)))
	}
}

func (ptr *QQuickWindow) SetYDefault(arg int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_SetYDefault(ptr.Pointer(), C.int(int32(arg)))
	}
}

//export callbackQQuickWindow_Alert
func callbackQQuickWindow_Alert(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::alert"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQQuickWindowFromPointer(ptr).AlertDefault(int(int32(msec)))
	}
}

func (ptr *QQuickWindow) ConnectAlert(f func(msec int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::alert", f)
	}
}

func (ptr *QQuickWindow) DisconnectAlert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::alert")
	}
}

func (ptr *QQuickWindow) Alert(msec int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Alert(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QQuickWindow) AlertDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_AlertDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

//export callbackQQuickWindow_Close
func callbackQQuickWindow_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *QQuickWindow) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::close", f)
	}
}

func (ptr *QQuickWindow) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::close")
	}
}

func (ptr *QQuickWindow) Close() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWindow_FocusObject
func callbackQQuickWindow_FocusObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::focusObject"); signal != nil {
		return core.PointerFromQObject(signal.(func() *core.QObject)())
	}

	return core.PointerFromQObject(NewQQuickWindowFromPointer(ptr).FocusObjectDefault())
}

func (ptr *QQuickWindow) ConnectFocusObject(f func() *core.QObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::focusObject", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::focusObject")
	}
}

func (ptr *QQuickWindow) FocusObject() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQuickWindow_FocusObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) FocusObjectDefault() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQuickWindow_FocusObjectDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_Format
func callbackQQuickWindow_Format(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::format"); signal != nil {
		return gui.PointerFromQSurfaceFormat(signal.(func() *gui.QSurfaceFormat)())
	}

	return gui.PointerFromQSurfaceFormat(NewQQuickWindowFromPointer(ptr).FormatDefault())
}

func (ptr *QQuickWindow) ConnectFormat(f func() *gui.QSurfaceFormat) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::format", f)
	}
}

func (ptr *QQuickWindow) DisconnectFormat() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::format")
	}
}

func (ptr *QQuickWindow) Format() *gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQSurfaceFormatFromPointer(C.QQuickWindow_Format(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) FormatDefault() *gui.QSurfaceFormat {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQSurfaceFormatFromPointer(C.QQuickWindow_FormatDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_Hide
func callbackQQuickWindow_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickWindow) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::hide", f)
	}
}

func (ptr *QQuickWindow) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::hide")
	}
}

func (ptr *QQuickWindow) Hide() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Hide(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_Lower
func callbackQQuickWindow_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickWindow) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::lower", f)
	}
}

func (ptr *QQuickWindow) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::lower")
	}
}

func (ptr *QQuickWindow) Lower() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Lower(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_MoveEvent
func callbackQQuickWindow_MoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ConnectMoveEvent(f func(ev *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::moveEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::moveEvent")
	}
}

func (ptr *QQuickWindow) MoveEvent(ev gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

func (ptr *QQuickWindow) MoveEventDefault(ev gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

//export callbackQQuickWindow_NativeEvent
func callbackQQuickWindow_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QQuickWindow) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::nativeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::nativeEvent")
	}
}

func (ptr *QQuickWindow) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QQuickWindow) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQQuickWindow_Raise
func callbackQQuickWindow_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickWindow) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::raise", f)
	}
}

func (ptr *QQuickWindow) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::raise")
	}
}

func (ptr *QQuickWindow) Raise() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Raise(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_RequestActivate
func callbackQQuickWindow_RequestActivate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::requestActivate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RequestActivateDefault()
	}
}

func (ptr *QQuickWindow) ConnectRequestActivate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::requestActivate", f)
	}
}

func (ptr *QQuickWindow) DisconnectRequestActivate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::requestActivate")
	}
}

func (ptr *QQuickWindow) RequestActivate() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestActivate(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RequestActivateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestActivateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_RequestUpdate
func callbackQQuickWindow_RequestUpdate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::requestUpdate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RequestUpdateDefault()
	}
}

func (ptr *QQuickWindow) ConnectRequestUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::requestUpdate", f)
	}
}

func (ptr *QQuickWindow) DisconnectRequestUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::requestUpdate")
	}
}

func (ptr *QQuickWindow) RequestUpdate() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RequestUpdateDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestUpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_Show
func callbackQQuickWindow_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickWindow) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::show", f)
	}
}

func (ptr *QQuickWindow) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::show")
	}
}

func (ptr *QQuickWindow) Show() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_Show(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowFullScreen
func callbackQQuickWindow_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showFullScreen", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showFullScreen")
	}
}

func (ptr *QQuickWindow) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowMaximized
func callbackQQuickWindow_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showMaximized", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showMaximized")
	}
}

func (ptr *QQuickWindow) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowMinimized
func callbackQQuickWindow_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showMinimized", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showMinimized")
	}
}

func (ptr *QQuickWindow) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowNormal
func callbackQQuickWindow_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showNormal", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::showNormal")
	}
}

func (ptr *QQuickWindow) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_Size
func callbackQQuickWindow_Size(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::size"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWindowFromPointer(ptr).SizeDefault())
}

func (ptr *QQuickWindow) ConnectSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::size", f)
	}
}

func (ptr *QQuickWindow) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::size")
	}
}

func (ptr *QQuickWindow) Size() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWindow_Size(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QQuickWindow) SizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QQuickWindow_SizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQQuickWindow_SurfaceType
func callbackQQuickWindow_SurfaceType(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::surfaceType"); signal != nil {
		return C.longlong(signal.(func() gui.QSurface__SurfaceType)())
	}

	return C.longlong(NewQQuickWindowFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *QQuickWindow) ConnectSurfaceType(f func() gui.QSurface__SurfaceType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::surfaceType", f)
	}
}

func (ptr *QQuickWindow) DisconnectSurfaceType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::surfaceType")
	}
}

func (ptr *QQuickWindow) SurfaceType() gui.QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickWindow_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWindow) SurfaceTypeDefault() gui.QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickWindow_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWindow_TabletEvent
func callbackQQuickWindow_TabletEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ConnectTabletEvent(f func(ev *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::tabletEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::tabletEvent")
	}
}

func (ptr *QQuickWindow) TabletEvent(ev gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

func (ptr *QQuickWindow) TabletEventDefault(ev gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

//export callbackQQuickWindow_TouchEvent
func callbackQQuickWindow_TouchEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ConnectTouchEvent(f func(ev *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::touchEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::touchEvent")
	}
}

func (ptr *QQuickWindow) TouchEvent(ev gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

func (ptr *QQuickWindow) TouchEventDefault(ev gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

//export callbackQQuickWindow_TimerEvent
func callbackQQuickWindow_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::timerEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::timerEvent")
	}
}

func (ptr *QQuickWindow) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickWindow_ChildEvent
func callbackQQuickWindow_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::childEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::childEvent")
	}
}

func (ptr *QQuickWindow) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickWindow_ConnectNotify
func callbackQQuickWindow_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWindowFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWindow) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::connectNotify", f)
	}
}

func (ptr *QQuickWindow) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::connectNotify")
	}
}

func (ptr *QQuickWindow) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWindow) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWindow_CustomEvent
func callbackQQuickWindow_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::customEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::customEvent")
	}
}

func (ptr *QQuickWindow) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWindow) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickWindow_DeleteLater
func callbackQQuickWindow_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWindow) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::deleteLater", f)
	}
}

func (ptr *QQuickWindow) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::deleteLater")
	}
}

func (ptr *QQuickWindow) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWindow_DisconnectNotify
func callbackQQuickWindow_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWindowFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWindow) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::disconnectNotify", f)
	}
}

func (ptr *QQuickWindow) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::disconnectNotify")
	}
}

func (ptr *QQuickWindow) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWindow) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWindow_EventFilter
func callbackQQuickWindow_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQuickWindow) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::eventFilter", f)
	}
}

func (ptr *QQuickWindow) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::eventFilter")
	}
}

func (ptr *QQuickWindow) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickWindow) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickWindow_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickWindow_MetaObject
func callbackQQuickWindow_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQuickWindow::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWindowFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWindow) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::metaObject", f)
	}
}

func (ptr *QQuickWindow) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQuickWindow::metaObject")
	}
}

func (ptr *QQuickWindow) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWindow_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWindow_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//go:generate stringer -type=QSGAbstractRenderer__ClearModeBit
//QSGAbstractRenderer::ClearModeBit
type QSGAbstractRenderer__ClearModeBit int64

const (
	QSGAbstractRenderer__ClearColorBuffer   QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0001)
	QSGAbstractRenderer__ClearDepthBuffer   QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0002)
	QSGAbstractRenderer__ClearStencilBuffer QSGAbstractRenderer__ClearModeBit = QSGAbstractRenderer__ClearModeBit(0x0004)
)

type QSGAbstractRenderer struct {
	core.QObject
}

type QSGAbstractRenderer_ITF interface {
	core.QObject_ITF
	QSGAbstractRenderer_PTR() *QSGAbstractRenderer
}

func (p *QSGAbstractRenderer) QSGAbstractRenderer_PTR() *QSGAbstractRenderer {
	return p
}

func (p *QSGAbstractRenderer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGAbstractRenderer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGAbstractRenderer(ptr QSGAbstractRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGAbstractRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSGAbstractRendererFromPointer(ptr unsafe.Pointer) *QSGAbstractRenderer {
	var n = new(QSGAbstractRenderer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGAbstractRenderer) DestroyQSGAbstractRenderer() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QSGAbstractRenderer) ClearColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QSGAbstractRenderer_ClearColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
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

func (ptr *QSGAbstractRenderer) DeviceRect() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QSGAbstractRenderer_DeviceRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ProjectionMatrix() *gui.QMatrix4x4 {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQMatrix4x4FromPointer(C.QSGAbstractRenderer_ProjectionMatrix(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QMatrix4x4).DestroyQMatrix4x4)
		return tmpValue
	}
	return nil
}

//export callbackQSGAbstractRenderer_SceneGraphChanged
func callbackQSGAbstractRenderer_SceneGraphChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::sceneGraphChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGAbstractRenderer) ConnectSceneGraphChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectSceneGraphChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::sceneGraphChanged", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectSceneGraphChanged() {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectSceneGraphChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::sceneGraphChanged")
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

func (ptr *QSGAbstractRenderer) ViewportRect() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QSGAbstractRenderer_ViewportRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQSGAbstractRenderer_TimerEvent
func callbackQSGAbstractRenderer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::timerEvent", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::timerEvent")
	}
}

func (ptr *QSGAbstractRenderer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSGAbstractRenderer_ChildEvent
func callbackQSGAbstractRenderer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::childEvent", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::childEvent")
	}
}

func (ptr *QSGAbstractRenderer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGAbstractRenderer_ConnectNotify
func callbackQSGAbstractRenderer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGAbstractRenderer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::connectNotify", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::connectNotify")
	}
}

func (ptr *QSGAbstractRenderer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGAbstractRenderer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGAbstractRenderer_CustomEvent
func callbackQSGAbstractRenderer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::customEvent", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::customEvent")
	}
}

func (ptr *QSGAbstractRenderer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGAbstractRenderer_DeleteLater
func callbackQSGAbstractRenderer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGAbstractRendererFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGAbstractRenderer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::deleteLater", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::deleteLater")
	}
}

func (ptr *QSGAbstractRenderer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGAbstractRenderer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSGAbstractRenderer_DisconnectNotify
func callbackQSGAbstractRenderer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGAbstractRenderer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::disconnectNotify", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::disconnectNotify")
	}
}

func (ptr *QSGAbstractRenderer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGAbstractRenderer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGAbstractRenderer_Event
func callbackQSGAbstractRenderer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGAbstractRendererFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGAbstractRenderer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::event", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::event")
	}
}

func (ptr *QSGAbstractRenderer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGAbstractRenderer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGAbstractRenderer_EventFilter
func callbackQSGAbstractRenderer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGAbstractRendererFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGAbstractRenderer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::eventFilter", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::eventFilter")
	}
}

func (ptr *QSGAbstractRenderer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGAbstractRenderer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGAbstractRenderer_MetaObject
func callbackQSGAbstractRenderer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGAbstractRenderer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGAbstractRendererFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGAbstractRenderer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::metaObject", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGAbstractRenderer::metaObject")
	}
}

func (ptr *QSGAbstractRenderer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGAbstractRenderer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGAbstractRenderer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGBasicGeometryNode struct {
	QSGNode
}

type QSGBasicGeometryNode_ITF interface {
	QSGNode_ITF
	QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode
}

func (p *QSGBasicGeometryNode) QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode {
	return p
}

func (p *QSGBasicGeometryNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGBasicGeometryNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGBasicGeometryNode(ptr QSGBasicGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGBasicGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGBasicGeometryNode {
	var n = new(QSGBasicGeometryNode)
	n.SetPointer(ptr)
	return n
}
func (ptr *QSGBasicGeometryNode) Geometry2() *QSGGeometry {
	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) Geometry() *QSGGeometry {
	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometry_ITF) {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_SetGeometry(ptr.Pointer(), PointerFromQSGGeometry(geometry))
	}
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNode() {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGBasicGeometryNode_IsSubtreeBlocked
func callbackQSGBasicGeometryNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGBasicGeometryNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGBasicGeometryNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGBasicGeometryNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGBasicGeometryNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGBasicGeometryNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGBasicGeometryNode::isSubtreeBlocked")
	}
}

func (ptr *QSGBasicGeometryNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGBasicGeometryNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGBasicGeometryNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGBasicGeometryNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGBasicGeometryNode_Preprocess
func callbackQSGBasicGeometryNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGBasicGeometryNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGBasicGeometryNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGBasicGeometryNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGBasicGeometryNode::preprocess", f)
	}
}

func (ptr *QSGBasicGeometryNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGBasicGeometryNode::preprocess")
	}
}

func (ptr *QSGBasicGeometryNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGBasicGeometryNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_PreprocessDefault(ptr.Pointer())
	}
}

//go:generate stringer -type=QSGBatchRenderer__BatchCompatibility
//QSGBatchRenderer::BatchCompatibility
type QSGBatchRenderer__BatchCompatibility int64

const (
	QSGBatchRenderer__BatchBreaksOnCompare QSGBatchRenderer__BatchCompatibility = QSGBatchRenderer__BatchCompatibility(0)
	QSGBatchRenderer__BatchIsCompatible    QSGBatchRenderer__BatchCompatibility = QSGBatchRenderer__BatchCompatibility(1)
)

type QSGBatchRenderer struct {
	ptr unsafe.Pointer
}

type QSGBatchRenderer_ITF interface {
	QSGBatchRenderer_PTR() *QSGBatchRenderer
}

func (p *QSGBatchRenderer) QSGBatchRenderer_PTR() *QSGBatchRenderer {
	return p
}

func (p *QSGBatchRenderer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGBatchRenderer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSGBatchRenderer(ptr QSGBatchRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBatchRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSGBatchRendererFromPointer(ptr unsafe.Pointer) *QSGBatchRenderer {
	var n = new(QSGBatchRenderer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGBatchRenderer) DestroyQSGBatchRenderer() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

type QSGClipNode struct {
	QSGBasicGeometryNode
}

type QSGClipNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGClipNode_PTR() *QSGClipNode
}

func (p *QSGClipNode) QSGClipNode_PTR() *QSGClipNode {
	return p
}

func (p *QSGClipNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGClipNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGBasicGeometryNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGClipNode(ptr QSGClipNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGClipNode_PTR().Pointer()
	}
	return nil
}

func NewQSGClipNodeFromPointer(ptr unsafe.Pointer) *QSGClipNode {
	var n = new(QSGClipNode)
	n.SetPointer(ptr)
	return n
}
func NewQSGClipNode() *QSGClipNode {
	var tmpValue = NewQSGClipNodeFromPointer(C.QSGClipNode_NewQSGClipNode())
	runtime.SetFinalizer(tmpValue, (*QSGClipNode).DestroyQSGClipNode)
	return tmpValue
}

func (ptr *QSGClipNode) ClipRect() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGClipNode_ClipRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGClipNode) IsRectangular() bool {
	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsRectangular(ptr.Pointer()) != 0
	}
	return false
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

func (ptr *QSGClipNode) DestroyQSGClipNode() {
	if ptr.Pointer() != nil {
		C.QSGClipNode_DestroyQSGClipNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGClipNode_IsSubtreeBlocked
func callbackQSGClipNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGClipNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGClipNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGClipNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGClipNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGClipNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGClipNode::isSubtreeBlocked")
	}
}

func (ptr *QSGClipNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGClipNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGClipNode_Preprocess
func callbackQSGClipNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGClipNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGClipNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGClipNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGClipNode::preprocess", f)
	}
}

func (ptr *QSGClipNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGClipNode::preprocess")
	}
}

func (ptr *QSGClipNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGClipNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGClipNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGClipNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGDynamicTexture struct {
	QSGTexture
}

type QSGDynamicTexture_ITF interface {
	QSGTexture_ITF
	QSGDynamicTexture_PTR() *QSGDynamicTexture
}

func (p *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture {
	return p
}

func (p *QSGDynamicTexture) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGTexture_PTR().Pointer()
	}
	return nil
}

func (p *QSGDynamicTexture) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGTexture_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGDynamicTexture(ptr QSGDynamicTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGDynamicTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGDynamicTextureFromPointer(ptr unsafe.Pointer) *QSGDynamicTexture {
	var n = new(QSGDynamicTexture)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGDynamicTexture) DestroyQSGDynamicTexture() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQSGDynamicTexture_UpdateTexture
func callbackQSGDynamicTexture_UpdateTexture(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::updateTexture"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGDynamicTexture) ConnectUpdateTexture(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::updateTexture", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectUpdateTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::updateTexture")
	}
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_UpdateTexture(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_Bind
func callbackQSGDynamicTexture_Bind(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::bind"); signal != nil {
		signal.(func())()
	} else {

	}
}

func (ptr *QSGDynamicTexture) ConnectBind(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::bind", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectBind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::bind")
	}
}

func (ptr *QSGDynamicTexture) Bind() {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_Bind(ptr.Pointer())
	}
}

//export callbackQSGDynamicTexture_HasAlphaChannel
func callbackQSGDynamicTexture_HasAlphaChannel(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::hasAlphaChannel"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGDynamicTexture) ConnectHasAlphaChannel(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::hasAlphaChannel", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectHasAlphaChannel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::hasAlphaChannel")
	}
}

func (ptr *QSGDynamicTexture) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_HasMipmaps
func callbackQSGDynamicTexture_HasMipmaps(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::hasMipmaps"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGDynamicTexture) ConnectHasMipmaps(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::hasMipmaps", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectHasMipmaps() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::hasMipmaps")
	}
}

func (ptr *QSGDynamicTexture) HasMipmaps() bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_HasMipmaps(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_IsAtlasTexture
func callbackQSGDynamicTexture_IsAtlasTexture(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::isAtlasTexture"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).IsAtlasTextureDefault())))
}

func (ptr *QSGDynamicTexture) ConnectIsAtlasTexture(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::isAtlasTexture", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectIsAtlasTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::isAtlasTexture")
	}
}

func (ptr *QSGDynamicTexture) IsAtlasTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_IsAtlasTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) IsAtlasTextureDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_IsAtlasTextureDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_NormalizedTextureSubRect
func callbackQSGDynamicTexture_NormalizedTextureSubRect(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::normalizedTextureSubRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQSGDynamicTextureFromPointer(ptr).NormalizedTextureSubRectDefault())
}

func (ptr *QSGDynamicTexture) ConnectNormalizedTextureSubRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::normalizedTextureSubRect", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectNormalizedTextureSubRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::normalizedTextureSubRect")
	}
}

func (ptr *QSGDynamicTexture) NormalizedTextureSubRect() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGDynamicTexture_NormalizedTextureSubRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGDynamicTexture) NormalizedTextureSubRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGDynamicTexture_NormalizedTextureSubRectDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQSGDynamicTexture_RemovedFromAtlas
func callbackQSGDynamicTexture_RemovedFromAtlas(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::removedFromAtlas"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGDynamicTextureFromPointer(ptr).RemovedFromAtlasDefault())
}

func (ptr *QSGDynamicTexture) ConnectRemovedFromAtlas(f func() *QSGTexture) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::removedFromAtlas", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectRemovedFromAtlas() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::removedFromAtlas")
	}
}

func (ptr *QSGDynamicTexture) RemovedFromAtlas() *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGDynamicTexture_RemovedFromAtlas(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGDynamicTexture) RemovedFromAtlasDefault() *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGDynamicTexture_RemovedFromAtlasDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSGDynamicTexture_TextureId
func callbackQSGDynamicTexture_TextureId(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::textureId"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QSGDynamicTexture) ConnectTextureId(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::textureId", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectTextureId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::textureId")
	}
}

func (ptr *QSGDynamicTexture) TextureId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGDynamicTexture_TextureId(ptr.Pointer())))
	}
	return 0
}

//export callbackQSGDynamicTexture_TextureSize
func callbackQSGDynamicTexture_TextureSize(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QSGDynamicTexture) ConnectTextureSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::textureSize", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectTextureSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::textureSize")
	}
}

func (ptr *QSGDynamicTexture) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSGDynamicTexture_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSGDynamicTexture_TimerEvent
func callbackQSGDynamicTexture_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::timerEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::timerEvent")
	}
}

func (ptr *QSGDynamicTexture) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGDynamicTexture) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSGDynamicTexture_ChildEvent
func callbackQSGDynamicTexture_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::childEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::childEvent")
	}
}

func (ptr *QSGDynamicTexture) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGDynamicTexture) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGDynamicTexture_ConnectNotify
func callbackQSGDynamicTexture_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGDynamicTexture) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::connectNotify", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::connectNotify")
	}
}

func (ptr *QSGDynamicTexture) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGDynamicTexture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGDynamicTexture_CustomEvent
func callbackQSGDynamicTexture_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::customEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::customEvent")
	}
}

func (ptr *QSGDynamicTexture) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGDynamicTexture) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGDynamicTexture_DeleteLater
func callbackQSGDynamicTexture_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGDynamicTextureFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGDynamicTexture) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::deleteLater", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::deleteLater")
	}
}

func (ptr *QSGDynamicTexture) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGDynamicTexture) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSGDynamicTexture_DisconnectNotify
func callbackQSGDynamicTexture_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGDynamicTexture) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::disconnectNotify", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::disconnectNotify")
	}
}

func (ptr *QSGDynamicTexture) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGDynamicTexture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGDynamicTexture_Event
func callbackQSGDynamicTexture_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGDynamicTexture) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::event", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::event")
	}
}

func (ptr *QSGDynamicTexture) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_EventFilter
func callbackQSGDynamicTexture_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGDynamicTexture) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::eventFilter", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::eventFilter")
	}
}

func (ptr *QSGDynamicTexture) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_MetaObject
func callbackQSGDynamicTexture_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGDynamicTexture::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGDynamicTextureFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGDynamicTexture) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::metaObject", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGDynamicTexture::metaObject")
	}
}

func (ptr *QSGDynamicTexture) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGDynamicTexture_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGDynamicTexture) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGDynamicTexture_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QSGEngine struct {
	core.QObject
}

type QSGEngine_ITF interface {
	core.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func (p *QSGEngine) QSGEngine_PTR() *QSGEngine {
	return p
}

func (p *QSGEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGEngine(ptr QSGEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGEngine_PTR().Pointer()
	}
	return nil
}

func NewQSGEngineFromPointer(ptr unsafe.Pointer) *QSGEngine {
	var n = new(QSGEngine)
	n.SetPointer(ptr)
	return n
}
func NewQSGEngine(parent core.QObject_ITF) *QSGEngine {
	var tmpValue = NewQSGEngineFromPointer(C.QSGEngine_NewQSGEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSGEngine) CreateRenderer() *QSGAbstractRenderer {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGAbstractRendererFromPointer(C.QSGEngine_CreateRenderer(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromId(id uint, size core.QSize_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGEngine_CreateTextureFromId(ptr.Pointer(), C.uint(uint32(id)), core.PointerFromQSize(size), C.longlong(options)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromImage(image gui.QImage_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGEngine_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.longlong(options)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
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

func (ptr *QSGEngine) DestroyQSGEngine() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DestroyQSGEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSGEngine_TimerEvent
func callbackQSGEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::timerEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::timerEvent")
	}
}

func (ptr *QSGEngine) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSGEngine_ChildEvent
func callbackQSGEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::childEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::childEvent")
	}
}

func (ptr *QSGEngine) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGEngine_ConnectNotify
func callbackQSGEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::connectNotify", f)
	}
}

func (ptr *QSGEngine) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::connectNotify")
	}
}

func (ptr *QSGEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGEngine_CustomEvent
func callbackQSGEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::customEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::customEvent")
	}
}

func (ptr *QSGEngine) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGEngine_DeleteLater
func callbackQSGEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGEngine) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::deleteLater", f)
	}
}

func (ptr *QSGEngine) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::deleteLater")
	}
}

func (ptr *QSGEngine) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSGEngine_DisconnectNotify
func callbackQSGEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::disconnectNotify", f)
	}
}

func (ptr *QSGEngine) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::disconnectNotify")
	}
}

func (ptr *QSGEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGEngine_Event
func callbackQSGEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::event", f)
	}
}

func (ptr *QSGEngine) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::event")
	}
}

func (ptr *QSGEngine) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGEngine_EventFilter
func callbackQSGEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::eventFilter", f)
	}
}

func (ptr *QSGEngine) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::eventFilter")
	}
}

func (ptr *QSGEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGEngine_MetaObject
func callbackQSGEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::metaObject", f)
	}
}

func (ptr *QSGEngine) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGEngine::metaObject")
	}
}

func (ptr *QSGEngine) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGFlatColorMaterial struct {
	QSGMaterial
}

type QSGFlatColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial
}

func (p *QSGFlatColorMaterial) QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial {
	return p
}

func (p *QSGFlatColorMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGFlatColorMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGFlatColorMaterial(ptr QSGFlatColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGFlatColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGFlatColorMaterialFromPointer(ptr unsafe.Pointer) *QSGFlatColorMaterial {
	var n = new(QSGFlatColorMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGFlatColorMaterial) DestroyQSGFlatColorMaterial() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQSGFlatColorMaterial() *QSGFlatColorMaterial {
	var tmpValue = NewQSGFlatColorMaterialFromPointer(C.QSGFlatColorMaterial_NewQSGFlatColorMaterial())
	runtime.SetFinalizer(tmpValue, (*QSGFlatColorMaterial).DestroyQSGFlatColorMaterial)
	return tmpValue
}

func (ptr *QSGFlatColorMaterial) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGFlatColorMaterial_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QSGFlatColorMaterial_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQSGFlatColorMaterial_Compare
func callbackQSGFlatColorMaterial_Compare(ptr unsafe.Pointer, other unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGFlatColorMaterial::compare"); signal != nil {
		return C.int(int32(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other))))
	}

	return C.int(int32(NewQSGFlatColorMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other))))
}

func (ptr *QSGFlatColorMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGFlatColorMaterial::compare", f)
	}
}

func (ptr *QSGFlatColorMaterial) DisconnectCompare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGFlatColorMaterial::compare")
	}
}

func (ptr *QSGFlatColorMaterial) Compare(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGFlatColorMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

func (ptr *QSGFlatColorMaterial) CompareDefault(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGFlatColorMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

//export callbackQSGFlatColorMaterial_CreateShader
func callbackQSGFlatColorMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGFlatColorMaterial::createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGFlatColorMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGFlatColorMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGFlatColorMaterial::createShader", f)
	}
}

func (ptr *QSGFlatColorMaterial) DisconnectCreateShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGFlatColorMaterial::createShader")
	}
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGFlatColorMaterial::type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGFlatColorMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGFlatColorMaterial) ConnectType(f func() *QSGMaterialType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGFlatColorMaterial::type", f)
	}
}

func (ptr *QSGFlatColorMaterial) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGFlatColorMaterial::type")
	}
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

//go:generate stringer -type=QSGGeometry__DataPattern
//QSGGeometry::DataPattern
type QSGGeometry__DataPattern int64

const (
	QSGGeometry__AlwaysUploadPattern QSGGeometry__DataPattern = QSGGeometry__DataPattern(0)
	QSGGeometry__StreamPattern       QSGGeometry__DataPattern = QSGGeometry__DataPattern(1)
	QSGGeometry__DynamicPattern      QSGGeometry__DataPattern = QSGGeometry__DataPattern(2)
	QSGGeometry__StaticPattern       QSGGeometry__DataPattern = QSGGeometry__DataPattern(3)
)

type QSGGeometry struct {
	ptr unsafe.Pointer
}

type QSGGeometry_ITF interface {
	QSGGeometry_PTR() *QSGGeometry
}

func (p *QSGGeometry) QSGGeometry_PTR() *QSGGeometry {
	return p
}

func (p *QSGGeometry) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGGeometry) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSGGeometry(ptr QSGGeometry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometry_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryFromPointer(ptr unsafe.Pointer) *QSGGeometry {
	var n = new(QSGGeometry)
	n.SetPointer(ptr)
	return n
}
func (ptr *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_Allocate(ptr.Pointer(), C.int(int32(vertexCount)), C.int(int32(indexCount)))
	}
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

func (ptr *QSGGeometry) IndexData() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexData2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexDataAsUInt() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGGeometry_IndexDataAsUInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexDataAsUInt2() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QSGGeometry_IndexDataAsUInt2(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexDataAsUShort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSGGeometry_IndexDataAsUShort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexDataAsUShort2() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSGGeometry_IndexDataAsUShort2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_IndexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexType() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_IndexType(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) LineWidth() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QSGGeometry_LineWidth(ptr.Pointer()))
	}
	return 0
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

func (ptr *QSGGeometry) VertexCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGGeometry_VertexCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) VertexData() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexData2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_VertexDataPattern(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGGeometry_DestroyQSGGeometry
func callbackQSGGeometry_DestroyQSGGeometry(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGGeometry::~QSGGeometry"); signal != nil {
		signal.(func())()
	} else {
		NewQSGGeometryFromPointer(ptr).DestroyQSGGeometryDefault()
	}
}

func (ptr *QSGGeometry) ConnectDestroyQSGGeometry(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGGeometry::~QSGGeometry", f)
	}
}

func (ptr *QSGGeometry) DisconnectDestroyQSGGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGGeometry::~QSGGeometry")
	}
}

func (ptr *QSGGeometry) DestroyQSGGeometry() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometry(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGGeometry) DestroyQSGGeometryDefault() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometryDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QSGGeometryNode struct {
	QSGBasicGeometryNode
}

type QSGGeometryNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGGeometryNode_PTR() *QSGGeometryNode
}

func (p *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode {
	return p
}

func (p *QSGGeometryNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGGeometryNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGBasicGeometryNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGGeometryNode(ptr QSGGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGGeometryNode {
	var n = new(QSGGeometryNode)
	n.SetPointer(ptr)
	return n
}
func NewQSGGeometryNode() *QSGGeometryNode {
	var tmpValue = NewQSGGeometryNodeFromPointer(C.QSGGeometryNode_NewQSGGeometryNode())
	runtime.SetFinalizer(tmpValue, (*QSGGeometryNode).DestroyQSGGeometryNode)
	return tmpValue
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

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGGeometryNode_IsSubtreeBlocked
func callbackQSGGeometryNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGGeometryNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGGeometryNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGGeometryNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGGeometryNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGGeometryNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGGeometryNode::isSubtreeBlocked")
	}
}

func (ptr *QSGGeometryNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGGeometryNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGGeometryNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGGeometryNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGGeometryNode_Preprocess
func callbackQSGGeometryNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGGeometryNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGGeometryNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGGeometryNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGGeometryNode::preprocess", f)
	}
}

func (ptr *QSGGeometryNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGGeometryNode::preprocess")
	}
}

func (ptr *QSGGeometryNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGGeometryNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGGeometryNode_PreprocessDefault(ptr.Pointer())
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

type QSGMaterial struct {
	ptr unsafe.Pointer
}

type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (p *QSGMaterial) QSGMaterial_PTR() *QSGMaterial {
	return p
}

func (p *QSGMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSGMaterial(ptr QSGMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialFromPointer(ptr unsafe.Pointer) *QSGMaterial {
	var n = new(QSGMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterial) DestroyQSGMaterial() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQSGMaterial_Compare
func callbackQSGMaterial_Compare(ptr unsafe.Pointer, other unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterial::compare"); signal != nil {
		return C.int(int32(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other))))
	}

	return C.int(int32(NewQSGMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other))))
}

func (ptr *QSGMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterial::compare", f)
	}
}

func (ptr *QSGMaterial) DisconnectCompare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterial::compare")
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

//export callbackQSGMaterial_CreateShader
func callbackQSGMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterial::createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(nil)
}

func (ptr *QSGMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterial::createShader", f)
	}
}

func (ptr *QSGMaterial) DisconnectCreateShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterial::createShader")
	}
}

func (ptr *QSGMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGMaterial) Flags() QSGMaterial__Flag {
	if ptr.Pointer() != nil {
		return QSGMaterial__Flag(C.QSGMaterial_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGMaterial) SetFlag(flags QSGMaterial__Flag, on bool) {
	if ptr.Pointer() != nil {
		C.QSGMaterial_SetFlag(ptr.Pointer(), C.longlong(flags), C.char(int8(qt.GoBoolToInt(on))))
	}
}

//export callbackQSGMaterial_Type
func callbackQSGMaterial_Type(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterial::type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(nil)
}

func (ptr *QSGMaterial) ConnectType(f func() *QSGMaterialType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterial::type", f)
	}
}

func (ptr *QSGMaterial) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterial::type")
	}
}

func (ptr *QSGMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGMaterial_Type(ptr.Pointer()))
	}
	return nil
}

type QSGMaterialShader struct {
	ptr unsafe.Pointer
}

type QSGMaterialShader_ITF interface {
	QSGMaterialShader_PTR() *QSGMaterialShader
}

func (p *QSGMaterialShader) QSGMaterialShader_PTR() *QSGMaterialShader {
	return p
}

func (p *QSGMaterialShader) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGMaterialShader) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSGMaterialShader(ptr QSGMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGMaterialShader {
	var n = new(QSGMaterialShader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterialShader) DestroyQSGMaterialShader() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQSGMaterialShader_FragmentShader
func callbackQSGMaterialShader_FragmentShader(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterialShader::fragmentShader"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSGMaterialShaderFromPointer(ptr).FragmentShaderDefault())
}

func (ptr *QSGMaterialShader) ConnectFragmentShader(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::fragmentShader", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectFragmentShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::fragmentShader")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterialShader::vertexShader"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSGMaterialShaderFromPointer(ptr).VertexShaderDefault())
}

func (ptr *QSGMaterialShader) ConnectVertexShader(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::vertexShader", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectVertexShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::vertexShader")
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

//export callbackQSGMaterialShader_Activate
func callbackQSGMaterialShader_Activate(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterialShader::activate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).ActivateDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectActivate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::activate", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectActivate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::activate")
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

//export callbackQSGMaterialShader_Deactivate
func callbackQSGMaterialShader_Deactivate(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterialShader::deactivate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).DeactivateDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectDeactivate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::deactivate", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectDeactivate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::deactivate")
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

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQOpenGLShaderProgramFromPointer(C.QSGMaterialShader_Program(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGMaterialShader) DisconnectUpdateState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::updateState")
	}
}

//export callbackQSGMaterialShader_Compile
func callbackQSGMaterialShader_Compile(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterialShader::compile"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).CompileDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectCompile(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::compile", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectCompile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::compile")
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

//export callbackQSGMaterialShader_Initialize
func callbackQSGMaterialShader_Initialize(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGMaterialShader::initialize"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).InitializeDefault()
	}
}

func (ptr *QSGMaterialShader) ConnectInitialize(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::initialize", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectInitialize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGMaterialShader::initialize")
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
		var sourceFileC = C.CString(sourceFile)
		defer C.free(unsafe.Pointer(sourceFileC))
		C.QSGMaterialShader_SetShaderSourceFile(ptr.Pointer(), C.longlong(ty), sourceFileC)
	}
}

func (ptr *QSGMaterialShader) SetShaderSourceFiles(ty gui.QOpenGLShader__ShaderTypeBit, sourceFiles []string) {
	if ptr.Pointer() != nil {
		var sourceFilesC = C.CString(strings.Join(sourceFiles, "|"))
		defer C.free(unsafe.Pointer(sourceFilesC))
		C.QSGMaterialShader_SetShaderSourceFiles(ptr.Pointer(), C.longlong(ty), sourceFilesC)
	}
}

type QSGMaterialType struct {
	ptr unsafe.Pointer
}

type QSGMaterialType_ITF interface {
	QSGMaterialType_PTR() *QSGMaterialType
}

func (p *QSGMaterialType) QSGMaterialType_PTR() *QSGMaterialType {
	return p
}

func (p *QSGMaterialType) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGMaterialType) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSGMaterialType(ptr QSGMaterialType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialType_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialTypeFromPointer(ptr unsafe.Pointer) *QSGMaterialType {
	var n = new(QSGMaterialType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterialType) DestroyQSGMaterialType() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
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
)

type QSGNode struct {
	ptr unsafe.Pointer
}

type QSGNode_ITF interface {
	QSGNode_PTR() *QSGNode
}

func (p *QSGNode) QSGNode_PTR() *QSGNode {
	return p
}

func (p *QSGNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSGNode(ptr QSGNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func NewQSGNodeFromPointer(ptr unsafe.Pointer) *QSGNode {
	var n = new(QSGNode)
	n.SetPointer(ptr)
	return n
}
func (ptr *QSGNode) ChildAtIndex(i int) *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_ChildAtIndex(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QSGNode) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGNode_ChildCount(ptr.Pointer())))
	}
	return 0
}

func NewQSGNode() *QSGNode {
	return NewQSGNodeFromPointer(C.QSGNode_NewQSGNode())
}

func (ptr *QSGNode) AppendChildNode(node QSGNode_ITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_AppendChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) FirstChild() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_FirstChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) Flags() QSGNode__Flag {
	if ptr.Pointer() != nil {
		return QSGNode__Flag(C.QSGNode_Flags(ptr.Pointer()))
	}
	return 0
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

//export callbackQSGNode_IsSubtreeBlocked
func callbackQSGNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGNode::isSubtreeBlocked")
	}
}

func (ptr *QSGNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGNode) LastChild() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_LastChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) MarkDirty(bits QSGNode__DirtyStateBit) {
	if ptr.Pointer() != nil {
		C.QSGNode_MarkDirty(ptr.Pointer(), C.longlong(bits))
	}
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

func (ptr *QSGNode) PrependChildNode(node QSGNode_ITF) {
	if ptr.Pointer() != nil {
		C.QSGNode_PrependChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

//export callbackQSGNode_Preprocess
func callbackQSGNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGNode::preprocess", f)
	}
}

func (ptr *QSGNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGNode::preprocess")
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

func (ptr *QSGNode) PreviousSibling() *QSGNode {
	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_PreviousSibling(ptr.Pointer()))
	}
	return nil
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

func (ptr *QSGNode) SetFlag(f QSGNode__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QSGNode_SetFlag(ptr.Pointer(), C.longlong(f), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QSGNode) SetFlags(f QSGNode__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QSGNode_SetFlags(ptr.Pointer(), C.longlong(f), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QSGNode) Type() QSGNode__NodeType {
	if ptr.Pointer() != nil {
		return QSGNode__NodeType(C.QSGNode_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGNode_DestroyQSGNode
func callbackQSGNode_DestroyQSGNode(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGNode::~QSGNode"); signal != nil {
		signal.(func())()
	} else {
		NewQSGNodeFromPointer(ptr).DestroyQSGNodeDefault()
	}
}

func (ptr *QSGNode) ConnectDestroyQSGNode(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGNode::~QSGNode", f)
	}
}

func (ptr *QSGNode) DisconnectDestroyQSGNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGNode::~QSGNode")
	}
}

func (ptr *QSGNode) DestroyQSGNode() {
	if ptr.Pointer() != nil {
		C.QSGNode_DestroyQSGNode(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGNode) DestroyQSGNodeDefault() {
	if ptr.Pointer() != nil {
		C.QSGNode_DestroyQSGNodeDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QSGOpacityNode struct {
	QSGNode
}

type QSGOpacityNode_ITF interface {
	QSGNode_ITF
	QSGOpacityNode_PTR() *QSGOpacityNode
}

func (p *QSGOpacityNode) QSGOpacityNode_PTR() *QSGOpacityNode {
	return p
}

func (p *QSGOpacityNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGOpacityNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGOpacityNode(ptr QSGOpacityNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpacityNode_PTR().Pointer()
	}
	return nil
}

func NewQSGOpacityNodeFromPointer(ptr unsafe.Pointer) *QSGOpacityNode {
	var n = new(QSGOpacityNode)
	n.SetPointer(ptr)
	return n
}
func NewQSGOpacityNode() *QSGOpacityNode {
	var tmpValue = NewQSGOpacityNodeFromPointer(C.QSGOpacityNode_NewQSGOpacityNode())
	runtime.SetFinalizer(tmpValue, (*QSGOpacityNode).DestroyQSGOpacityNode)
	return tmpValue
}

func (ptr *QSGOpacityNode) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QSGOpacityNode_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpacityNode) SetOpacity(opacity float64) {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNode() {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_DestroyQSGOpacityNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGOpacityNode_IsSubtreeBlocked
func callbackQSGOpacityNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGOpacityNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGOpacityNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGOpacityNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpacityNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGOpacityNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpacityNode::isSubtreeBlocked")
	}
}

func (ptr *QSGOpacityNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGOpacityNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGOpacityNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGOpacityNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGOpacityNode_Preprocess
func callbackQSGOpacityNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGOpacityNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGOpacityNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGOpacityNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpacityNode::preprocess", f)
	}
}

func (ptr *QSGOpacityNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpacityNode::preprocess")
	}
}

func (ptr *QSGOpacityNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGOpacityNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGOpacityNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGOpaqueTextureMaterial struct {
	QSGMaterial
}

type QSGOpaqueTextureMaterial_ITF interface {
	QSGMaterial_ITF
	QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial
}

func (p *QSGOpaqueTextureMaterial) QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial {
	return p
}

func (p *QSGOpaqueTextureMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGOpaqueTextureMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGOpaqueTextureMaterial(ptr QSGOpaqueTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGOpaqueTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGOpaqueTextureMaterial {
	var n = new(QSGOpaqueTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGOpaqueTextureMaterial) DestroyQSGOpaqueTextureMaterial() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQSGOpaqueTextureMaterial() *QSGOpaqueTextureMaterial {
	var tmpValue = NewQSGOpaqueTextureMaterialFromPointer(C.QSGOpaqueTextureMaterial_NewQSGOpaqueTextureMaterial())
	runtime.SetFinalizer(tmpValue, (*QSGOpaqueTextureMaterial).DestroyQSGOpaqueTextureMaterial)
	return tmpValue
}

func (ptr *QSGOpaqueTextureMaterial) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) HorizontalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_MipmapFiltering(ptr.Pointer()))
	}
	return 0
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
		var tmpValue = NewQSGTextureFromPointer(C.QSGOpaqueTextureMaterial_Texture(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) VerticalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) M_texture() *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGOpaqueTextureMaterial_M_texture(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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

//export callbackQSGOpaqueTextureMaterial_Compare
func callbackQSGOpaqueTextureMaterial_Compare(ptr unsafe.Pointer, other unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGOpaqueTextureMaterial::compare"); signal != nil {
		return C.int(int32(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other))))
	}

	return C.int(int32(NewQSGOpaqueTextureMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other))))
}

func (ptr *QSGOpaqueTextureMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpaqueTextureMaterial::compare", f)
	}
}

func (ptr *QSGOpaqueTextureMaterial) DisconnectCompare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpaqueTextureMaterial::compare")
	}
}

func (ptr *QSGOpaqueTextureMaterial) Compare(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGOpaqueTextureMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) CompareDefault(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGOpaqueTextureMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

//export callbackQSGOpaqueTextureMaterial_CreateShader
func callbackQSGOpaqueTextureMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGOpaqueTextureMaterial::createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGOpaqueTextureMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGOpaqueTextureMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpaqueTextureMaterial::createShader", f)
	}
}

func (ptr *QSGOpaqueTextureMaterial) DisconnectCreateShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpaqueTextureMaterial::createShader")
	}
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGOpaqueTextureMaterial::type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGOpaqueTextureMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGOpaqueTextureMaterial) ConnectType(f func() *QSGMaterialType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpaqueTextureMaterial::type", f)
	}
}

func (ptr *QSGOpaqueTextureMaterial) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGOpaqueTextureMaterial::type")
	}
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

type QSGSimpleMaterial struct {
	QSGMaterial
}

type QSGSimpleMaterial_ITF interface {
	QSGMaterial_ITF
	QSGSimpleMaterial_PTR() *QSGSimpleMaterial
}

func (p *QSGSimpleMaterial) QSGSimpleMaterial_PTR() *QSGSimpleMaterial {
	return p
}

func (p *QSGSimpleMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGSimpleMaterial(ptr QSGSimpleMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterial {
	var n = new(QSGSimpleMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterial) DestroyQSGSimpleMaterial() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

type QSGSimpleMaterialShader struct {
	QSGMaterialShader
}

type QSGSimpleMaterialShader_ITF interface {
	QSGMaterialShader_ITF
	QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader
}

func (p *QSGSimpleMaterialShader) QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader {
	return p
}

func (p *QSGSimpleMaterialShader) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleMaterialShader) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterialShader_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGSimpleMaterialShader(ptr QSGSimpleMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterialShader {
	var n = new(QSGSimpleMaterialShader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterialShader) DestroyQSGSimpleMaterialShader() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

type QSGSimpleRectNode struct {
	QSGGeometryNode
}

type QSGSimpleRectNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleRectNode_PTR() *QSGSimpleRectNode
}

func (p *QSGSimpleRectNode) QSGSimpleRectNode_PTR() *QSGSimpleRectNode {
	return p
}

func (p *QSGSimpleRectNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleRectNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGGeometryNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGSimpleRectNode(ptr QSGSimpleRectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleRectNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleRectNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleRectNode {
	var n = new(QSGSimpleRectNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleRectNode) DestroyQSGSimpleRectNode() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQSGSimpleRectNode2() *QSGSimpleRectNode {
	var tmpValue = NewQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode2())
	runtime.SetFinalizer(tmpValue, (*QSGSimpleRectNode).DestroyQSGSimpleRectNode)
	return tmpValue
}

func NewQSGSimpleRectNode(rect core.QRectF_ITF, color gui.QColor_ITF) *QSGSimpleRectNode {
	var tmpValue = NewQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode(core.PointerFromQRectF(rect), gui.PointerFromQColor(color)))
	runtime.SetFinalizer(tmpValue, (*QSGSimpleRectNode).DestroyQSGSimpleRectNode)
	return tmpValue
}

func (ptr *QSGSimpleRectNode) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QSGSimpleRectNode_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QSGSimpleRectNode) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGSimpleRectNode_Rect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
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

//export callbackQSGSimpleRectNode_IsSubtreeBlocked
func callbackQSGSimpleRectNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGSimpleRectNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGSimpleRectNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGSimpleRectNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleRectNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGSimpleRectNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleRectNode::isSubtreeBlocked")
	}
}

func (ptr *QSGSimpleRectNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGSimpleRectNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGSimpleRectNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGSimpleRectNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGSimpleRectNode_Preprocess
func callbackQSGSimpleRectNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGSimpleRectNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleRectNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGSimpleRectNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleRectNode::preprocess", f)
	}
}

func (ptr *QSGSimpleRectNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleRectNode::preprocess")
	}
}

func (ptr *QSGSimpleRectNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGSimpleRectNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_PreprocessDefault(ptr.Pointer())
	}
}

//go:generate stringer -type=QSGSimpleTextureNode__TextureCoordinatesTransformFlag
//QSGSimpleTextureNode::TextureCoordinatesTransformFlag
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag int64

const (
	QSGSimpleTextureNode__NoTransform        QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x00)
	QSGSimpleTextureNode__MirrorHorizontally QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x01)
	QSGSimpleTextureNode__MirrorVertically   QSGSimpleTextureNode__TextureCoordinatesTransformFlag = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x02)
)

type QSGSimpleTextureNode struct {
	QSGGeometryNode
}

type QSGSimpleTextureNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode
}

func (p *QSGSimpleTextureNode) QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode {
	return p
}

func (p *QSGSimpleTextureNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleTextureNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGGeometryNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGSimpleTextureNode(ptr QSGSimpleTextureNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleTextureNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleTextureNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleTextureNode {
	var n = new(QSGSimpleTextureNode)
	n.SetPointer(ptr)
	return n
}
func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {
	var tmpValue = NewQSGSimpleTextureNodeFromPointer(C.QSGSimpleTextureNode_NewQSGSimpleTextureNode())
	runtime.SetFinalizer(tmpValue, (*QSGSimpleTextureNode).DestroyQSGSimpleTextureNode)
	return tmpValue
}

func (ptr *QSGSimpleTextureNode) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGSimpleTextureNode_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) OwnsTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_OwnsTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGSimpleTextureNode) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGSimpleTextureNode_Rect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
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

func (ptr *QSGSimpleTextureNode) SourceRect() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGSimpleTextureNode_SourceRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGSimpleTextureNode_Texture(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
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

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNode() {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGSimpleTextureNode_IsSubtreeBlocked
func callbackQSGSimpleTextureNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGSimpleTextureNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGSimpleTextureNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGSimpleTextureNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleTextureNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGSimpleTextureNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleTextureNode::isSubtreeBlocked")
	}
}

func (ptr *QSGSimpleTextureNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGSimpleTextureNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGSimpleTextureNode_Preprocess
func callbackQSGSimpleTextureNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGSimpleTextureNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleTextureNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGSimpleTextureNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleTextureNode::preprocess", f)
	}
}

func (ptr *QSGSimpleTextureNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGSimpleTextureNode::preprocess")
	}
}

func (ptr *QSGSimpleTextureNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGSimpleTextureNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_PreprocessDefault(ptr.Pointer())
	}
}

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
	QSGTexture__Repeat      QSGTexture__WrapMode = QSGTexture__WrapMode(0)
	QSGTexture__ClampToEdge QSGTexture__WrapMode = QSGTexture__WrapMode(1)
)

type QSGTexture struct {
	core.QObject
}

type QSGTexture_ITF interface {
	core.QObject_ITF
	QSGTexture_PTR() *QSGTexture
}

func (p *QSGTexture) QSGTexture_PTR() *QSGTexture {
	return p
}

func (p *QSGTexture) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGTexture) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGTexture(ptr QSGTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureFromPointer(ptr unsafe.Pointer) *QSGTexture {
	var n = new(QSGTexture)
	n.SetPointer(ptr)
	return n
}
func NewQSGTexture() *QSGTexture {
	var tmpValue = NewQSGTextureFromPointer(C.QSGTexture_NewQSGTexture())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSGTexture_Bind
func callbackQSGTexture_Bind(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::bind"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGTexture) ConnectBind(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::bind", f)
	}
}

func (ptr *QSGTexture) DisconnectBind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::bind")
	}
}

func (ptr *QSGTexture) Bind() {
	if ptr.Pointer() != nil {
		C.QSGTexture_Bind(ptr.Pointer())
	}
}

func (ptr *QSGTexture) ConvertToNormalizedSourceRect(rect core.QRectF_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGTexture_ConvertToNormalizedSourceRect(ptr.Pointer(), core.PointerFromQRectF(rect)))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_Filtering(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGTexture_HasAlphaChannel
func callbackQSGTexture_HasAlphaChannel(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::hasAlphaChannel"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGTexture) ConnectHasAlphaChannel(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::hasAlphaChannel", f)
	}
}

func (ptr *QSGTexture) DisconnectHasAlphaChannel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::hasAlphaChannel")
	}
}

func (ptr *QSGTexture) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGTexture_HasMipmaps
func callbackQSGTexture_HasMipmaps(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::hasMipmaps"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSGTexture) ConnectHasMipmaps(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::hasMipmaps", f)
	}
}

func (ptr *QSGTexture) DisconnectHasMipmaps() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::hasMipmaps")
	}
}

func (ptr *QSGTexture) HasMipmaps() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_HasMipmaps(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) HorizontalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGTexture_IsAtlasTexture
func callbackQSGTexture_IsAtlasTexture(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::isAtlasTexture"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).IsAtlasTextureDefault())))
}

func (ptr *QSGTexture) ConnectIsAtlasTexture(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::isAtlasTexture", f)
	}
}

func (ptr *QSGTexture) DisconnectIsAtlasTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::isAtlasTexture")
	}
}

func (ptr *QSGTexture) IsAtlasTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_IsAtlasTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) IsAtlasTextureDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_IsAtlasTextureDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGTexture_NormalizedTextureSubRect
func callbackQSGTexture_NormalizedTextureSubRect(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::normalizedTextureSubRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQSGTextureFromPointer(ptr).NormalizedTextureSubRectDefault())
}

func (ptr *QSGTexture) ConnectNormalizedTextureSubRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::normalizedTextureSubRect", f)
	}
}

func (ptr *QSGTexture) DisconnectNormalizedTextureSubRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::normalizedTextureSubRect")
	}
}

func (ptr *QSGTexture) NormalizedTextureSubRect() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGTexture_NormalizedTextureSubRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) NormalizedTextureSubRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSGTexture_NormalizedTextureSubRectDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQSGTexture_RemovedFromAtlas
func callbackQSGTexture_RemovedFromAtlas(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::removedFromAtlas"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGTextureFromPointer(ptr).RemovedFromAtlasDefault())
}

func (ptr *QSGTexture) ConnectRemovedFromAtlas(f func() *QSGTexture) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::removedFromAtlas", f)
	}
}

func (ptr *QSGTexture) DisconnectRemovedFromAtlas() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::removedFromAtlas")
	}
}

func (ptr *QSGTexture) RemovedFromAtlas() *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlas(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) RemovedFromAtlasDefault() *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlasDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
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

//export callbackQSGTexture_TextureId
func callbackQSGTexture_TextureId(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::textureId"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QSGTexture) ConnectTextureId(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::textureId", f)
	}
}

func (ptr *QSGTexture) DisconnectTextureId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::textureId")
	}
}

func (ptr *QSGTexture) TextureId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGTexture_TextureId(ptr.Pointer())))
	}
	return 0
}

//export callbackQSGTexture_TextureSize
func callbackQSGTexture_TextureSize(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QSGTexture) ConnectTextureSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::textureSize", f)
	}
}

func (ptr *QSGTexture) DisconnectTextureSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::textureSize")
	}
}

func (ptr *QSGTexture) TextureSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSGTexture_TextureSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSGTexture) UpdateBindOptions(force bool) {
	if ptr.Pointer() != nil {
		C.QSGTexture_UpdateBindOptions(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(force))))
	}
}

func (ptr *QSGTexture) VerticalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) DestroyQSGTexture() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DestroyQSGTexture(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTexture_TimerEvent
func callbackQSGTexture_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGTexture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::timerEvent", f)
	}
}

func (ptr *QSGTexture) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::timerEvent")
	}
}

func (ptr *QSGTexture) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTexture) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSGTexture_ChildEvent
func callbackQSGTexture_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGTexture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::childEvent", f)
	}
}

func (ptr *QSGTexture) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::childEvent")
	}
}

func (ptr *QSGTexture) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTexture) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGTexture_ConnectNotify
func callbackQSGTexture_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTexture) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::connectNotify", f)
	}
}

func (ptr *QSGTexture) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::connectNotify")
	}
}

func (ptr *QSGTexture) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTexture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTexture_CustomEvent
func callbackQSGTexture_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGTexture) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::customEvent", f)
	}
}

func (ptr *QSGTexture) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::customEvent")
	}
}

func (ptr *QSGTexture) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGTexture) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGTexture_DeleteLater
func callbackQSGTexture_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTextureFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGTexture) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::deleteLater", f)
	}
}

func (ptr *QSGTexture) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::deleteLater")
	}
}

func (ptr *QSGTexture) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGTexture) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTexture_DisconnectNotify
func callbackQSGTexture_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTexture) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::disconnectNotify", f)
	}
}

func (ptr *QSGTexture) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::disconnectNotify")
	}
}

func (ptr *QSGTexture) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTexture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTexture_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTexture_Event
func callbackQSGTexture_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGTexture) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::event", f)
	}
}

func (ptr *QSGTexture) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::event")
	}
}

func (ptr *QSGTexture) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGTexture) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGTexture_EventFilter
func callbackQSGTexture_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGTexture) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::eventFilter", f)
	}
}

func (ptr *QSGTexture) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::eventFilter")
	}
}

func (ptr *QSGTexture) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGTexture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGTexture_MetaObject
func callbackQSGTexture_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTexture::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGTextureFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGTexture) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::metaObject", f)
	}
}

func (ptr *QSGTexture) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTexture::metaObject")
	}
}

func (ptr *QSGTexture) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTexture_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTexture_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGTextureMaterial struct {
	QSGOpaqueTextureMaterial
}

type QSGTextureMaterial_ITF interface {
	QSGOpaqueTextureMaterial_ITF
	QSGTextureMaterial_PTR() *QSGTextureMaterial
}

func (p *QSGTextureMaterial) QSGTextureMaterial_PTR() *QSGTextureMaterial {
	return p
}

func (p *QSGTextureMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGTextureMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGOpaqueTextureMaterial_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGTextureMaterial(ptr QSGTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGTextureMaterial {
	var n = new(QSGTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGTextureMaterial) DestroyQSGTextureMaterial() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

//export callbackQSGTextureMaterial_Compare
func callbackQSGTextureMaterial_Compare(ptr unsafe.Pointer, other unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureMaterial::compare"); signal != nil {
		return C.int(int32(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other))))
	}

	return C.int(int32(NewQSGTextureMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other))))
}

func (ptr *QSGTextureMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureMaterial::compare", f)
	}
}

func (ptr *QSGTextureMaterial) DisconnectCompare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureMaterial::compare")
	}
}

func (ptr *QSGTextureMaterial) Compare(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGTextureMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

func (ptr *QSGTextureMaterial) CompareDefault(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGTextureMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

//export callbackQSGTextureMaterial_CreateShader
func callbackQSGTextureMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureMaterial::createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGTextureMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGTextureMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureMaterial::createShader", f)
	}
}

func (ptr *QSGTextureMaterial) DisconnectCreateShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureMaterial::createShader")
	}
}

func (ptr *QSGTextureMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGTextureMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureMaterial) CreateShaderDefault() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGTextureMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGTextureMaterial_Type
func callbackQSGTextureMaterial_Type(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureMaterial::type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGTextureMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGTextureMaterial) ConnectType(f func() *QSGMaterialType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureMaterial::type", f)
	}
}

func (ptr *QSGTextureMaterial) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureMaterial::type")
	}
}

func (ptr *QSGTextureMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGTextureMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureMaterial) TypeDefault() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGTextureMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProvider_ITF interface {
	core.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func (p *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider {
	return p
}

func (p *QSGTextureProvider) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGTextureProvider) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGTextureProvider(ptr QSGTextureProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureProvider_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureProviderFromPointer(ptr unsafe.Pointer) *QSGTextureProvider {
	var n = new(QSGTextureProvider)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGTextureProvider) DestroyQSGTextureProvider() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQSGTextureProvider_Texture
func callbackQSGTextureProvider_Texture(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::texture"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGTexture())
}

func (ptr *QSGTextureProvider) ConnectTexture(f func() *QSGTexture) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::texture", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTexture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::texture")
	}
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSGTextureFromPointer(C.QSGTextureProvider_Texture(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSGTextureProvider_TextureChanged
func callbackQSGTextureProvider_TextureChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::textureChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGTextureProvider) ConnectTextureChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectTextureChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::textureChanged", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTextureChanged() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectTextureChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::textureChanged")
	}
}

func (ptr *QSGTextureProvider) TextureChanged() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TextureChanged(ptr.Pointer())
	}
}

//export callbackQSGTextureProvider_TimerEvent
func callbackQSGTextureProvider_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::timerEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::timerEvent")
	}
}

func (ptr *QSGTextureProvider) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTextureProvider) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSGTextureProvider_ChildEvent
func callbackQSGTextureProvider_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::childEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::childEvent")
	}
}

func (ptr *QSGTextureProvider) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTextureProvider) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSGTextureProvider_ConnectNotify
func callbackQSGTextureProvider_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTextureProvider) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::connectNotify", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::connectNotify")
	}
}

func (ptr *QSGTextureProvider) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTextureProvider) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTextureProvider_CustomEvent
func callbackQSGTextureProvider_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::customEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::customEvent")
	}
}

func (ptr *QSGTextureProvider) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGTextureProvider) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSGTextureProvider_DeleteLater
func callbackQSGTextureProvider_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTextureProviderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGTextureProvider) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::deleteLater", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::deleteLater")
	}
}

func (ptr *QSGTextureProvider) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGTextureProvider) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTextureProvider_DisconnectNotify
func callbackQSGTextureProvider_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureProviderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTextureProvider) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::disconnectNotify", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::disconnectNotify")
	}
}

func (ptr *QSGTextureProvider) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTextureProvider) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTextureProvider_Event
func callbackQSGTextureProvider_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureProviderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSGTextureProvider) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::event", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::event")
	}
}

func (ptr *QSGTextureProvider) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGTextureProvider) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGTextureProvider_EventFilter
func callbackQSGTextureProvider_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTextureProviderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSGTextureProvider) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::eventFilter", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::eventFilter")
	}
}

func (ptr *QSGTextureProvider) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGTextureProvider) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGTextureProvider_MetaObject
func callbackQSGTextureProvider_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTextureProvider::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGTextureProviderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGTextureProvider) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::metaObject", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTextureProvider::metaObject")
	}
}

func (ptr *QSGTextureProvider) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTextureProvider_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureProvider) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTextureProvider_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGTransformNode struct {
	QSGNode
}

type QSGTransformNode_ITF interface {
	QSGNode_ITF
	QSGTransformNode_PTR() *QSGTransformNode
}

func (p *QSGTransformNode) QSGTransformNode_PTR() *QSGTransformNode {
	return p
}

func (p *QSGTransformNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGTransformNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGNode_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGTransformNode(ptr QSGTransformNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTransformNode_PTR().Pointer()
	}
	return nil
}

func NewQSGTransformNodeFromPointer(ptr unsafe.Pointer) *QSGTransformNode {
	var n = new(QSGTransformNode)
	n.SetPointer(ptr)
	return n
}
func NewQSGTransformNode() *QSGTransformNode {
	var tmpValue = NewQSGTransformNodeFromPointer(C.QSGTransformNode_NewQSGTransformNode())
	runtime.SetFinalizer(tmpValue, (*QSGTransformNode).DestroyQSGTransformNode)
	return tmpValue
}

func (ptr *QSGTransformNode) Matrix() *gui.QMatrix4x4 {
	if ptr.Pointer() != nil {
		return gui.NewQMatrix4x4FromPointer(C.QSGTransformNode_Matrix(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTransformNode) SetMatrix(matrix gui.QMatrix4x4_ITF) {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_SetMatrix(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QSGTransformNode) DestroyQSGTransformNode() {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_DestroyQSGTransformNode(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTransformNode_IsSubtreeBlocked
func callbackQSGTransformNode_IsSubtreeBlocked(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTransformNode::isSubtreeBlocked"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSGTransformNodeFromPointer(ptr).IsSubtreeBlockedDefault())))
}

func (ptr *QSGTransformNode) ConnectIsSubtreeBlocked(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTransformNode::isSubtreeBlocked", f)
	}
}

func (ptr *QSGTransformNode) DisconnectIsSubtreeBlocked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTransformNode::isSubtreeBlocked")
	}
}

func (ptr *QSGTransformNode) IsSubtreeBlocked() bool {
	if ptr.Pointer() != nil {
		return C.QSGTransformNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTransformNode) IsSubtreeBlockedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSGTransformNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGTransformNode_Preprocess
func callbackQSGTransformNode_Preprocess(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGTransformNode::preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTransformNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGTransformNode) ConnectPreprocess(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTransformNode::preprocess", f)
	}
}

func (ptr *QSGTransformNode) DisconnectPreprocess() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGTransformNode::preprocess")
	}
}

func (ptr *QSGTransformNode) Preprocess() {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGTransformNode) PreprocessDefault() {
	if ptr.Pointer() != nil {
		C.QSGTransformNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGVertexColorMaterial struct {
	QSGMaterial
}

type QSGVertexColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial
}

func (p *QSGVertexColorMaterial) QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial {
	return p
}

func (p *QSGVertexColorMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGVertexColorMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
}

func PointerFromQSGVertexColorMaterial(ptr QSGVertexColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGVertexColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGVertexColorMaterialFromPointer(ptr unsafe.Pointer) *QSGVertexColorMaterial {
	var n = new(QSGVertexColorMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGVertexColorMaterial) DestroyQSGVertexColorMaterial() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQSGVertexColorMaterial() *QSGVertexColorMaterial {
	var tmpValue = NewQSGVertexColorMaterialFromPointer(C.QSGVertexColorMaterial_NewQSGVertexColorMaterial())
	runtime.SetFinalizer(tmpValue, (*QSGVertexColorMaterial).DestroyQSGVertexColorMaterial)
	return tmpValue
}

//export callbackQSGVertexColorMaterial_Compare
func callbackQSGVertexColorMaterial_Compare(ptr unsafe.Pointer, other unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGVertexColorMaterial::compare"); signal != nil {
		return C.int(int32(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other))))
	}

	return C.int(int32(NewQSGVertexColorMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other))))
}

func (ptr *QSGVertexColorMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGVertexColorMaterial::compare", f)
	}
}

func (ptr *QSGVertexColorMaterial) DisconnectCompare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGVertexColorMaterial::compare")
	}
}

func (ptr *QSGVertexColorMaterial) Compare(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGVertexColorMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

func (ptr *QSGVertexColorMaterial) CompareDefault(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSGVertexColorMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other))))
	}
	return 0
}

//export callbackQSGVertexColorMaterial_CreateShader
func callbackQSGVertexColorMaterial_CreateShader(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGVertexColorMaterial::createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGVertexColorMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGVertexColorMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGVertexColorMaterial::createShader", f)
	}
}

func (ptr *QSGVertexColorMaterial) DisconnectCreateShader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGVertexColorMaterial::createShader")
	}
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSGVertexColorMaterial::type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGVertexColorMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGVertexColorMaterial) ConnectType(f func() *QSGMaterialType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSGVertexColorMaterial::type", f)
	}
}

func (ptr *QSGVertexColorMaterial) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSGVertexColorMaterial::type")
	}
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
