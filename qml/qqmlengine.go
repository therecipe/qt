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

type QQmlEngineITF interface {
	QJSEngineITF
	QQmlEnginePTR() *QQmlEngine
}

func PointerFromQQmlEngine(ptr QQmlEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEnginePTR().Pointer()
	}
	return nil
}

func QQmlEngineFromPointer(ptr unsafe.Pointer) *QQmlEngine {
	var n = new(QQmlEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlEngine) QQmlEnginePTR() *QQmlEngine {
	return ptr
}

//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int

var (
	QQmlEngine__CppOwnership        = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership = QQmlEngine__ObjectOwnership(1)
)

func (ptr *QQmlEngine) OfflineStoragePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlEngine_OfflineStoragePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOfflineStoragePath(C.QtObjectPtr(ptr.Pointer()), C.CString(dir))
	}
}

func NewQQmlEngine(parent core.QObjectITF) *QQmlEngine {
	return QQmlEngineFromPointer(unsafe.Pointer(C.QQmlEngine_NewQQmlEngine(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBaseITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImageProvider(C.QtObjectPtr(ptr.Pointer()), C.CString(providerId), C.QtObjectPtr(PointerFromQQmlImageProviderBase(provider)))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImportPath(C.QtObjectPtr(ptr.Pointer()), C.CString(path))
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_AddPluginPath(C.QtObjectPtr(ptr.Pointer()), C.CString(path))
	}
}

func (ptr *QQmlEngine) BaseUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlEngine_BaseUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlEngine) ClearComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QQmlEngine_ContextForObject(object core.QObjectITF) *QQmlContext {
	return QQmlContextFromPointer(unsafe.Pointer(C.QQmlEngine_QQmlEngine_ContextForObject(C.QtObjectPtr(core.PointerFromQObject(object)))))
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	if ptr.Pointer() != nil {
		return QQmlImageProviderBaseFromPointer(unsafe.Pointer(C.QQmlEngine_ImageProvider(C.QtObjectPtr(ptr.Pointer()), C.CString(providerId))))
	}
	return nil
}

func (ptr *QQmlEngine) ImportPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_ImportPathList(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	if ptr.Pointer() != nil {
		return QQmlIncubationControllerFromPointer(unsafe.Pointer(C.QQmlEngine_IncubationController(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return network.QNetworkAccessManagerFromPointer(unsafe.Pointer(C.QQmlEngine_NetworkAccessManager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	if ptr.Pointer() != nil {
		return QQmlNetworkAccessManagerFactoryFromPointer(unsafe.Pointer(C.QQmlEngine_NetworkAccessManagerFactory(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func QQmlEngine_ObjectOwnership(object core.QObjectITF) QQmlEngine__ObjectOwnership {
	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(C.QtObjectPtr(core.PointerFromQObject(object))))
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlEngine_OutputWarningsToStandardError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlEngine) PluginPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_PluginPathList(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectQuit(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "quit", f)
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "quit")
	}
}

//export callbackQQmlEngineQuit
func callbackQQmlEngineQuit(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "quit").(func())()
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_RemoveImageProvider(C.QtObjectPtr(ptr.Pointer()), C.CString(providerId))
	}
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	if ptr.Pointer() != nil {
		return QQmlContextFromPointer(unsafe.Pointer(C.QQmlEngine_RootContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlEngine) SetBaseUrl(url string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObjectITF, context QQmlContextITF) {
	C.QQmlEngine_QQmlEngine_SetContextForObject(C.QtObjectPtr(core.PointerFromQObject(object)), C.QtObjectPtr(PointerFromQQmlContext(context)))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetImportPathList(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationControllerITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQmlIncubationController(controller)))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactoryITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQmlNetworkAccessManagerFactory(factory)))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObjectITF, ownership QQmlEngine__ObjectOwnership) {
	C.QQmlEngine_QQmlEngine_SetObjectOwnership(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(ownership))
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetPluginPathList(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
