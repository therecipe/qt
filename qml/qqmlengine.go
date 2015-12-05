package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
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
	for len(n.ObjectName()) < len("QQmlEngine_") {
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::offlineStoragePath")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlEngine_OfflineStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setOfflineStoragePath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOfflineStoragePath(ptr.Pointer(), C.CString(dir))
	}
}

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::QQmlEngine")
		}
	}()

	return NewQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::addImageProvider")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImageProvider(ptr.Pointer(), C.CString(providerId), PointerFromQQmlImageProviderBase(provider))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::addImportPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddImportPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::addPluginPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_AddPluginPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QQmlEngine) ClearComponentCache() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::clearComponentCache")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(ptr.Pointer())
	}
}

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::contextForObject")
		}
	}()

	return NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::imageProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQmlImageProviderBaseFromPointer(C.QQmlEngine_ImageProvider(ptr.Pointer(), C.CString(providerId)))
	}
	return nil
}

func (ptr *QQmlEngine) ImportPathList() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::importPathList")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_ImportPathList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::incubationController")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQmlIncubationControllerFromPointer(C.QQmlEngine_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::networkAccessManager")
		}
	}()

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QQmlEngine_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::networkAccessManagerFactory")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlEngine_NetworkAccessManagerFactory(ptr.Pointer()))
	}
	return nil
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::objectOwnership")
		}
	}()

	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::outputWarningsToStandardError")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlEngine_OutputWarningsToStandardError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlEngine) PluginPathList() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::pluginPathList")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_PluginPathList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::quit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectQuit(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "quit", f)
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::quit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "quit")
	}
}

//export callbackQQmlEngineQuit
func callbackQQmlEngineQuit(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::quit")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "quit").(func())()
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::removeImageProvider")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_RemoveImageProvider(ptr.Pointer(), C.CString(providerId))
	}
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::rootContext")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlEngine_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) SetBaseUrl(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setBaseUrl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setContextForObject")
		}
	}()

	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setImportPathList")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetImportPathList(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setIncubationController")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(ptr.Pointer(), PointerFromQQmlIncubationController(controller))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setNetworkAccessManagerFactory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(ptr.Pointer(), PointerFromQQmlNetworkAccessManagerFactory(factory))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setObjectOwnership")
		}
	}()

	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.int(ownership))
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setOutputWarningsToStandardError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::setPluginPathList")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetPluginPathList(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::trimComponentCache")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlEngine::~QQmlEngine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
