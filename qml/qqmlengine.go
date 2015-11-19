package qml

//#include "qqmlengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"strings"
	"unsafe"
)

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func PointerFromQQmlEngine(ptr QQmlEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlEngineFromPointer(ptr unsafe.Pointer) *QQmlEngine {
	var n = new(QQmlEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return ptr
}

//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership = QQmlEngine__ObjectOwnership(1)
)

func (ptr *QQmlEngine) OfflineStoragePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlEngine_OfflineStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOfflineStoragePath(ptr.Pointer(), C.CString(dir))
	}
}

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	return NewQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImageProvider(ptr.Pointer(), C.CString(providerId), PointerFromQQmlImageProviderBase(provider))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImportPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_AddPluginPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) ClearComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(ptr.Pointer())
	}
}

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {
	return NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	if ptr.Pointer() != nil {
		return NewQQmlImageProviderBaseFromPointer(C.QQmlEngine_ImageProvider(ptr.Pointer(), C.CString(providerId)))
	}
	return nil
}

func (ptr *QQmlEngine) ImportPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_ImportPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	if ptr.Pointer() != nil {
		return NewQQmlIncubationControllerFromPointer(C.QQmlEngine_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QQmlEngine_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	if ptr.Pointer() != nil {
		return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlEngine_NetworkAccessManagerFactory(ptr.Pointer()))
	}
	return nil
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlEngine_OutputWarningsToStandardError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlEngine) PluginPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_PluginPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectQuit(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "quit", f)
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "quit")
	}
}

//export callbackQQmlEngineQuit
func callbackQQmlEngineQuit(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "quit").(func())()
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_RemoveImageProvider(ptr.Pointer(), C.CString(providerId))
	}
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlEngine_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) SetBaseUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetImportPathList(ptr.Pointer(), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(ptr.Pointer(), PointerFromQQmlIncubationController(controller))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(ptr.Pointer(), PointerFromQQmlNetworkAccessManagerFactory(factory))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.int(ownership))
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetPluginPathList(ptr.Pointer(), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
